package mysql

import (
	"context"
	"judgeMore/pkg/constants"
	"judgeMore/pkg/errno"
)

func QueryRecognizedEvent(ctx context.Context) ([]*RecognizedEvent, int64, error) {
	var reconize_event []*RecognizedEvent
	var total int64
	err := db.WithContext(ctx).
		Table(constants.TableReconizedEvent).
		Find(&reconize_event).
		Count(&total).
		Error
	if err != nil {
		return nil, -1, errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql:failed to query reconized_event"+err.Error())
	}
	return reconize_event, total, nil
}
