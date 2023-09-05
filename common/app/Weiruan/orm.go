package Weiruan

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func Weiruan_orm() (bool, error) {
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
	for i := 0; i < len(Allhref); i++ {
		job := &global.Hello{
			Company:       "微软",
			Title:         Title[i],
			Job_category:  Job_category[i],
			Job_type_name: "实习生",
			Job_detail:    Job_detail[i],
			Job_Location:  WorkLocation,
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
