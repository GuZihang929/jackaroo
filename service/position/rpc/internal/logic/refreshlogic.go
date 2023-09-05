package logic

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/service/company/rpc/types/company"
	"xiangxiang/jackaroo/service/position/model/model"
	"xiangxiang/jackaroo/service/position/rpc/internal/svc"
	"xiangxiang/jackaroo/service/position/rpc/types/position"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *position.RefreshRequest) (*position.RefreshResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.CompanyRpc.Detail(l.ctx, &company.DetailRequest{
		Id: in.Cid,
	})
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	collection, err := mongoDB.MongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = collection.Database().Client().Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	filter := bson.M{"company": res.Company}
	cursor, err := collection.Find(context.Background(), filter)
	defer cursor.Close(context.Background())

	type Pos struct {
		Title       string   `bson:"title"`
		JobCategory string   `bson:"job_category"`
		JobTypeName string   `bson:"job_type_name"`
		JobDetail   string   `bson:"job_detail"`
		JobLocation []string `bson:"job_location"`
	}

	//var positions []*model.Position

	for cursor.Next(context.Background()) {
		var p Pos
		if err := cursor.Decode(&p); err != nil {
			fmt.Printf("Error decoding document: %v\n", err)
		}

		currentTime := time.Now()
		currentTimeString := currentTime.Format("2006-01-02 15:04:05")

		posi := &model.Position{
			Cid:         res.Id,
			Company:     res.Company,
			Title:       p.Title,
			JobCategory: p.JobCategory,
			JobTypeName: p.JobTypeName,
			JobDetail:   p.JobDetail,
			JobLocation: strings.Join(p.JobLocation, "/"),
			PushTime:    currentTimeString,
			FetchTime:   currentTimeString,
		}

		//positions = append(positions, posi)
		// 查询语句，用于检查是否已存在相同的COMPANY和TITLE

		existsQuery := "SELECT id FROM position WHERE COMPANY = ? AND TITLE = ? LIMIT 1"

		// 插入语句
		insertQuery := "INSERT INTO position (CID, COMPANY, TITLE, JOB_CATEGORY, JOB_TYPE_NAME, JOB_DETAIL, JOB_LOCATION, PUSH_TIME, FETCH_TIME) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
		var id = 0

		//	TODO 查询语句有问题，结构输出也有问题rpc error: code = DeadlineExceeded desc = context deadline exceeded
		l.svcCtx.DB.QueryRow(&id, existsQuery, posi.Company, posi.Title)

		if id == 0 {
			_, err := l.svcCtx.DB.Exec(insertQuery, posi.Cid, posi.Company, posi.Title, posi.JobCategory, posi.JobTypeName, posi.JobDetail, posi.JobLocation, posi.PushTime, posi.FetchTime)
			if err != nil {
				return nil, err
			}
		}
	}
	refresh, err := l.svcCtx.CompanyRpc.Refresh(l.ctx, &company.RefreshRequest{
		Id: res.Id,
	})
	return &position.RefreshResponse{
		Id:      refresh.Id,
		Company: refresh.Company,
		Number:  refresh.Number,
		Brief:   refresh.Brief,
		Picture: refresh.Picture,
	}, nil
}
