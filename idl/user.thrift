namespace go user
include "./model.thrift"

// register
struct RegisterRequest {
    1: required string username,
    2: required string password,
    3: required string email,
    4: required string Id,
}
struct RegisterResponse {
    1: required model.BaseResp base,
    2: required string UserId,
}

// login
struct LoginRequest{
    1: required string Id,
    2: required string password,
}
struct LoginResponse{
    1: required model.UserInfo data,
    2: required model.BaseResp base,
}
// logout
struct LogoutReq {
}

struct LogoutResp {

}
// QueryUserInfo
struct QueryUserInfoRequest {
    1 :required string Id,
}

struct QueryUserInfoResponse {
    1: required model.BaseResp base,
    2: required model.UserInfo data,
}

// UpdateUserInfo
struct UpdateUserInfoRequest{
    1: optional string college,
    2: optional string grade,
    3: optional string major,
}
struct UpdateUserInfoResponse{
    1: required model.BaseResp base,
    2: required model.UserInfo data,
}

// VerifyEmail
struct VerifyEmailRequest{
      1: required string email,
      2: required string code,
}
struct VerifyEmailResponse{
    1: required model.BaseResp base,
}
//
service UserService {
    RegisterResponse Register(1: RegisterRequest req)(api.post = "/api/auth/register"),
    LoginResponse Login(1: LoginRequest req)(api.post = "/api/auth/login"),
    LogoutResp Logout(1: LogoutReq req) (api.post="/api/auth/logout"),
    VerifyEmailResponse VerifyEmail(1: VerifyEmailRequest req)(api.post = "/api/auth/email"),
    QueryUserInfoResponse QueryUserInfo(1: QueryUserInfoRequest req)(api.get = '/api/users/'),
    UpdateUserInfoResponse UpdateUserInfo(1: UpdateUserInfoRequest req)(api.put ='/api/users/me'),
}
