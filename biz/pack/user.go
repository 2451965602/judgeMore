package pack

import (
	resp "judgeMore/biz/model/model"
	"judgeMore/biz/service/model"
	"strconv"
)

func User(user *model.User) *resp.UserInfo {
	return &resp.UserInfo{
		Username:  user.UserName,
		UserId:    user.Uid,
		Major:     user.Major,
		College:   user.College,
		Grade:     user.Grade,
		Role:      user.Role,
		Email:     user.Email,
		CreatedAt: strconv.FormatInt(user.CreateAT, 10),
		UpdatedAt: strconv.FormatInt(user.UpdateAT, 10),
		DeletedAt: strconv.FormatInt(user.DeleteAT, 10),
	}
}
