package Meituan

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func Meituan_orm() (bool, error) {
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
	for i := 0; i < len(AllMeituan); i++ {
		utime := int64(AllMeituan[i].PushTime)
		uloc := time.FixedZone("CST", 8*3600)
		t := time.Unix(utime/1000, 0).In(uloc)
		now := t.Format("2006-01-02 15:04:05")
		if len(AllMeituan[i].WorkPlace) < 1 {
			continue
		}
		job := &global.Hello{
			Company:       "美团",
			Title:         AllMeituan[i].Title,
			Job_category:  AllMeituan[i].Job_category,
			Job_type_name: "实习生",
			Job_detail:    AllMeituan[i].Job_Detail + AllMeituan[i].Job_Obj,
			Job_Location:  global.Work{AllMeituan[i].WorkPlace[0].City},
			PushTime:      now,
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
