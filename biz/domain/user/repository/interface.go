package repository

import (
	"context"
	"judgeMore/biz/service/model"
)

type UserDB interface {
	IsUserExist(ctx context.Context, user *model.User) (bool, error)
	CreateUser(ctx context.Context, user *model.User) (string, error)
	GetUserInfoByRoleId(ctx context.Context, role_id string) (*model.User, error)
	UpdateInfoByRoleId(ctx context.Context, role_id string, element ...string) (*model.User, error)
	ActivateUser(ctx context.Context, uid string) error
}
type UserCache interface {
	GetCodeCache(ctx context.Context, key string) (code string, err error)
	PutCodeToCache(ctx context.Context, key string) (code string, err error)
	IsKeyExist(ctx context.Context, key string) bool
}
