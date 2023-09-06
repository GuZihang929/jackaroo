package Alibaba

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

// 阿里巴巴 向数据库表中插入数据
func Alibaba_orm() (bool, error) {
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
	for i := 0; i < len(Allali); i++ {
		job := &global.Hello{
			Company:       "阿里巴巴",
			Title:         Allali[i].Title,
			Job_category:  Allali[i].Job_category,
			Job_type_name: Allali[i].Job_type_name,
			Job_detail:    Allali[i].Job_Detail + Allali[i].Job_Obj,
			Job_Location:  Allali[i].WorkLocation,
			PushTime:      Allali[i].PushTime,
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
