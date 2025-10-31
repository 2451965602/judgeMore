package model

type User struct {
	Uid      string
	UserName string
	Email    string
	Password string
	College  string
	Major    string
	Grade    string
	Status   int
	Role     string
	CreateAT int64
	UpdateAT int64
	DeleteAT int64
}

type EmailAuth struct {
	Code  string
	Email string
	Uid   string
	Time  int64 //时间戳
}
