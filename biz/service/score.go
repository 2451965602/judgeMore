package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"judgeMore/biz/dal/mysql"
	"judgeMore/biz/service/model"
	"judgeMore/pkg/errno"
	"judgeMore/pkg/utils"
)

type ScoreService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewScoreService(ctx context.Context, c *app.RequestContext) *ScoreService {
	return &ScoreService{
		ctx: ctx,
		c:   c,
	}
}

func (svc *ScoreService) QueryScoreRecordByScoreId(score_id string) (*model.ScoreRecord, error) {
	exist, err := mysql.IsScoreRecordExist(svc.ctx, score_id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceRecordNotExistCode, "Socre Result not exist")
	}
	recordInfo, err := mysql.QueryScoreRecordByScoreId(svc.ctx, score_id)
	if err != nil {
		return nil, err
	}
	return recordInfo, nil
}

func (svc *ScoreService) QueryScoreRecordByEventId(event_id string) (*model.ScoreRecord, error) {
	exist, err := mysql.IsScoreRecordExist_Event(svc.ctx, event_id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.NewErrNo(errno.ServiceRecordNotExistCode, "Socre Result not exist")
	}
	recordInfo, err := mysql.QueryScoreRecordByEventId(svc.ctx, event_id)
	if err != nil {
		return nil, err
	}
	return recordInfo, nil
}

func (svc *ScoreService) QueryScoreRecordByStuId(stu_id string) ([]*model.ScoreRecord, int64, error) {
	exist, err := mysql.IsUserExist(svc.ctx, &model.User{Uid: stu_id})
	if err != nil {
		return nil, -1, err
	}
	if !exist {
		return nil, -1, errno.NewErrNo(errno.ServiceUserExistCode, "user not exist")
	}
	recordInfoList, count, err := mysql.QueryScoreRecordByStuId(svc.ctx, stu_id)
	if err != nil {
		return nil, -1, err
	}
	return recordInfoList, count, nil
}

// 用于直接修改积分。
func (svc *ScoreService) ReviseScore(result_id string, score float64) error {
	exist, err := mysql.IsScoreRecordExist(svc.ctx, result_id)
	if err != nil {
		return err
	}
	if !exist {
		return errno.NewErrNo(errno.ServiceRecordNotExistCode, "Socre Result not exist")
	}
	info, err := mysql.QueryScoreRecordByScoreId(svc.ctx, result_id)
	if err != nil {
		return err
	}
	// 验证用户权限
	user_id := GetUserIDFromContext(svc.c)
	exist, err = mysql.IsAdminRelationExist(svc.ctx, user_id, info.UserId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.NewErrNo(errno.ServiceNoAuthToDo, "No permission to update the stu's score")
	}
	err = mysql.UpdatesScore(svc.ctx, result_id, score)
	if err != nil {
		return err
	}
	return nil
}

// TODO:代码写的可读性很不好，有兴趣可以优化一下
func (svc *ScoreService) ScoreRank(req *model.ScoreRankReq) ([]*model.StuScoreMessage, int64, error) {
	var users []*model.User
	var err error
	// 参数检验
	if req.StuName == "" && req.Grade == "" && req.College == "" {
		users, err = mysql.QueryAllStu(svc.ctx)
		if err != nil {
			return nil, 0, err
		}
	}
	if req.College != "" {
		collegeExist, err := IsCollegeExist(svc.ctx, req.College)
		if err != nil {
			return nil, 0, err
		}
		if !collegeExist {
			return nil, 0, errno.NewErrNo(errno.ServiceCollegeNotExistCode, "college not exist")
		}
	} else if req.Grade != "" {
		if !utils.IsGradeValid(req.Grade) {
			return nil, 0, errno.NewErrNo(errno.ServiceGradeNotExistCode, "invalid grade")
		}
	}
	// 按优先级：姓名 → 学院 → 年级 逐步过滤
	if req.StuName != "" {
		users, err = mysql.QueryUserByUserName(svc.ctx, req.StuName)
		if err != nil {
			return nil, 0, err
		}
		if len(users) == 0 {
			return nil, 0, nil
		}
	}

	if req.College != "" {
		if users != nil {
			// 在现有用户基础上按学院内存过滤
			filtered := make([]*model.User, 0)
			for _, user := range users {
				if user.College == req.College {
					filtered = append(filtered, user)
				}
			}
			users = filtered
		} else {
			// 直接查询学院
			users, err = mysql.QueryUserByCollege(svc.ctx, req.College)
			if err != nil {
				return nil, 0, err
			}
		}
		if len(users) == 0 {
			return nil, 0, nil
		}
	}

	if req.Grade != "" {
		if users != nil {
			// 在现有用户基础上按年级内存过滤
			filtered := make([]*model.User, 0)
			for _, user := range users {
				if user.Grade == req.Grade {
					filtered = append(filtered, user)
				}
			}
			users = filtered
		} else {
			// 直接查询年级
			users, err = mysql.QueryUserByUserGrade(svc.ctx, req.Grade)
			if err != nil {
				return nil, 0, err
			}
		}
		if len(users) == 0 {
			return nil, 0, nil
		}
	}
	// 以学生为单位查询积分
	var result []*model.StuScoreMessage
	var count int64
	for _, u := range users {
		temp := &model.StuScoreMessage{
			Uid:     u.Uid,
			Name:    u.UserName,
			Grade:   u.Grade,
			College: u.College,
		}
		var sum float64
		scoreList, _, err := mysql.QueryScoreRecordByStuId(svc.ctx, u.Uid)
		if err != nil {
			return nil, 0, err
		}
		for _, s := range scoreList {
			sum += s.FinalIntegral
		}
		temp.Score = sum
		result = append(result, temp)
		count++
	}
	return result, count, nil
}
