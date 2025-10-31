namespace go model

struct BaseResp{
    1: i64 code,
    2: string msg,
}

struct UserInfo{
    1: string username,  //姓名
    2: string userId,   // 学号
    4: string Major // 专业
    5: string college, //学院
    6: string grade,  // 年级
    7: string email //邮箱
    8: string role //角色
    9: required string created_at
    10: required string updated_at
    11: required string deleted_at
}