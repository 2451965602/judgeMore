namespace go score
include "./model.thrift"
struct QueryScoreByScoreIdRequest{
    1: required string score_id,
}
struct QueryScoreByScoreIdResponse{
     1: required model.BaseResp base,
     2: required model.ScoreRecord data,
}
struct QueryScoreByStuIdRequest{
     1: required string stu_id,
}
struct QueryScoreByStuIdResponse{
     1: required model.BaseResp base,
     2: required model.ScoreRecordList data,
}
struct QueryScoreByEventIdRequest{
      1: required string event_id,
}
struct QueryScoreByEventIdResponse{
     1: required model.BaseResp base,
     2: required model.ScoreRecord data,
}
// 为辅导员提供的积分修改接口
struct ReviseEventScoreRequest{
    1: required string result_id
    2: required double score
}
struct ReviseEventScoreResponse{
     1: required model.BaseResp base,
}
struct ScoreRankRequest{
    1: optional string college,
    2: optional string grade,
    3: optional string stu_name,
}
struct ScoreRankResponse{
         1: required model.BaseResp base,
         2: required model.StuScoreMessageList data,
}
service ScoreService {
    QueryScoreByScoreIdResponse QueryScoreByScoreId(1:QueryScoreByScoreIdRequest req)(api.get="/api/query/score/id"),
    QueryScoreByEventIdResponse QueryScoreByEventId(1:QueryScoreByEventIdRequest req)(api.get="/api/query/score/material"),
    QueryScoreByStuIdResponse QueryScoreByStuId(1:QueryScoreByStuIdRequest req)(api.get="/api/query/score/stu"),
    ReviseEventScoreResponse ReviseScore(1:ReviseEventScoreRequest req)(api.post = "/api/update/score/id"),
    ScoreRankResponse ScoreRank(1:ScoreRankRequest req)(api.get= "/api/score/query/rank")
}