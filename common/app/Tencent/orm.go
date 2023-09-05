package Tencent

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func Tencent_orm() (bool, error) {
	time1 := time.Now().Format("2006-01-02 15:04:05")
	// 转换为 MongoDB 支持的格式
	collection, err := mongoDB.MongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = collection.Database().Client().Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	for i := 0; i < len(AllInformation); i++ {
		job := &global.Hello{
			Company:       "腾讯",
			Title:         AllInformation[i].Title,
			Job_category:  "2023校园招聘",
			Job_type_name: AllInformation[i].Job_type_name,
			Job_detail:    AllInformation[i].Job_Detail + AllInformation[i].Job_Obj,
			Job_Location:  AllInformation[i].WorkPlace,
			PushTime:      "",
			Fetch_time:    time1,
		}
		filter := bson.M{"Title": job.Title}
		replace := bson.M{"$set": job}
		_, err := collection.UpdateOne(context.Background(), filter, replace, options.Update().SetUpsert(true))
		if err != nil {
			log.Fatal(err)
		}
	}
	return true, nil
}
