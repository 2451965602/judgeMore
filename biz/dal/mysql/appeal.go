package mysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"judgeMore/biz/service/model"
	"judgeMore/pkg/constants"
	"judgeMore/pkg/errno"
	"time"
)

// database层往往需要集成并暴露与业务有关的接口，而不在该层进行业务复杂逻辑的处理
func IsAppealExist(ctx context.Context, result_Id string) (bool, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("result_id = ?", result_Id).
		First(&appeal).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return true, nil
}

func IsAppealExistByAppealId(ctx context.Context, appeal_id string) (bool, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		First(&appeal).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //没找到了用户不存在
			return false, nil
		}
		return false, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return true, nil
}

// 该函数调用前检验存在性
func QueryAppealById(ctx context.Context, appeal_id string) (*model.Appeal, error) {
	var appeal *Appeal
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		First(&appeal).
		Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return buildAppeal(appeal), nil
}

func QueryAppealByUserId(ctx context.Context, userId string) ([]*model.Appeal, int64, error) {
	var appeal []*Appeal
	var count int64
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("user_id = ?", userId).
		Find(&appeal).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	return buildAppealList(appeal), count, nil
}

func CreateAppeal(ctx context.Context, a *model.Appeal) (string, error) {
	appeal := &Appeal{
		ResultId:       a.ResultId,
		UserId:         a.UserId,
		AppealReason:   a.AppealReason,
		AttachmentPath: a.AttachmentPath,
		AppealType:     a.AppealType,
		Status:         "pending",
	}
	var count int64
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("result_id = ?", a.ResultId).
		Count(&count).
		Error
	if err != nil {
		return "", errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to query appeal: %v", err)
	}
	appeal.AppealCount = count + 1
	// 这边可以对申诉次数进行控制，但后期有具体需求时则需要重构。控制逻辑应该放在service层
	err = db.WithContext(ctx).
		Table(constants.TableAppeal).
		Create(appeal).
		Error
	if err != nil {
		return "", errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to insert appeal: %v", err)
	}
	return appeal.AppealId, nil
}
func DeleteAppealById(ctx context.Context, appeal_id string) error {
	err := db.WithContext(ctx).
		Table(constants.TableAppeal).
		Where("appeal_id = ?", appeal_id).
		Delete(&Appeal{Status: "deleted"}).
		Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to delete appeal: %v", err)
	}
	return nil
}
func UpdateAppealInfo(ctx context.Context, appeal *model.Appeal) error {
	// 更新内容多 事务提交
	err := db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			return tx.Table(constants.TableAppeal).
				Where("appeal_id = ?", appeal.AppealId).
				Update("status", appeal.Status).
				Update("handled_by", appeal.UserId).
				Update("handled_result", appeal.HandleResult).
				Update("handled_at", time.Now()).
				Error
		})
	return err
}
func buildAppeal(data *Appeal) *model.Appeal {
	r := &model.Appeal{
		ResultId:       data.ResultId,
		AppealId:       data.AppealId,
		UserId:         data.UserId,
		AppealReason:   data.AppealReason,
		AttachmentPath: data.AttachmentPath,
		AppealCount:    data.AppealCount,
		AppealType:     data.AppealType,
		Status:         data.Status,
		HandleResult:   data.HandledResult,
		HandledBy:      data.HandledBy,
		UpdateAT:       data.UpdatedAt.Unix(),
		CreateAT:       data.CreatedAt.Unix(),
		DeleteAT:       0,
	}
	if data.HandledAt == nil {
		r.HandledAt = 0
	} else {
		r.HandledAt = data.HandledAt.Unix()
	}
	return r
}

func buildAppealList(data []*Appeal) []*model.Appeal {
	result := make([]*model.Appeal, 0)
	for i := range data {
		result = append(result, buildAppeal(data[i]))
	}
	return result
}
