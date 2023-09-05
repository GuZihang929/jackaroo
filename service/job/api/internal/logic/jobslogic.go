package logic

import (
	"context"
	"fmt"
	"job/rpc/rpc"
	"strconv"

	"job/api/internal/svc"
	"job/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobsLogic {
	return &JobsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobsLogic) Jobs(req *types.JobsRequest) (resp *types.JobsResponse, err error) {

	sql := GetSql(req)

	//get, _ := l.svcCtx.Redis.Hgetall("id:" + strconv.FormatInt(req.NextId, 10) + "_limit:" + strconv.FormatInt(req.PageSize, 10))
	//if len(get) != 0 {
	//
	//	job := []types.Job{}
	//	fmt.Println(len(get))
	//	for i := 0; i < len(get); i++ {
	//		fmt.Println(i)
	//		atoi := strconv.Itoa(i)
	//		hget, _ := l.svcCtx.Redis.Hget("id:"+strconv.FormatInt(req.NextId, 10)+"_limit:"+strconv.FormatInt(req.PageSize, 10), atoi)
	//		job = append(job, types.Job{})
	//
	//		err = json.Unmarshal([]byte(hget), &job[i])
	//		if err != nil {
	//			fmt.Println(err)
	//			return nil, err
	//		}
	//	}
	//	return &types.JobsResponse{
	//		Code: 200,
	//		Msg:  "",
	//		Date: job,
	//	}, err
	//}
	res, err := l.svcCtx.SJobRpc.Login(l.ctx, &rpc.SerchJobRequest{Sql: sql})
	fmt.Println("----========----", err)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}

	resp = &types.JobsResponse{
		Code: res.Code,
		Msg:  res.Msg,
		Date: []types.Job{},
	}
	fmt.Println("-==================", err)
	for _, date := range res.Date {
		//marshal, _ := json.Marshal(date)
		//_ = l.svcCtx.Redis.Hset("id:"+strconv.FormatInt(req.NextId, 10)+"_limit:"+strconv.FormatInt(req.PageSize, 10),
		//	strconv.Itoa(i), string(marshal))

		fmt.Println(date.Company)
		resp.Date = append(resp.Date, types.Job{
			Id:          date.Id,
			CId:         date.CId,
			Company:     date.Company,
			Title:       date.Title,
			JobCategory: date.JobCategory,
			JobTypeName: date.JobTypeName,
			JobDetail:   date.JobDetail,
			JobLocation: date.JobLocation,
			Push_time:   date.PushTime,
			Fetch_time:  date.FetchTime,
		})
		fmt.Println("-========-=-=-=-=------------------------")
	}

	return
}

func GetSql(req *types.JobsRequest) string {
	res := "select * from job where job_category = '" + req.JobCategory + "'"
	if req.Name != "" {
		res += " and company = '" + req.Name + "'"

	}

	if req.JobTypeName != "" {
		res += " and job_type_name = '" + req.JobTypeName + "'"
	}

	if req.Title != "" {
		res += " and title like '" + req.Title + "%'"
	}
	if req.JobLocation != "" {
		res += " and job_location = '%" + req.JobLocation + "%'"
	}

	res += " and uuid > '" + strconv.FormatInt(req.NextId, 10) + "'"

	res += " limit " + strconv.FormatInt(req.PageSize, 10)
	return res
}
