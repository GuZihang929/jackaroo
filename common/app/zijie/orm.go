package zijie

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func zijieOrm(list []List) (bool, error) {
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
	for i := 0; i < len(list); i++ {
		job := &global.Hello{
			Company:       "字节跳动",
			Title:         list[i].Title,
			Job_detail:    "Duty:\n" + list[i].Description + "\nRequire:\n" + list[i].Requirement,
			Job_category:  list[i].JobCategory.Name,
			Job_type_name: list[i].JobSubject.Name.Name,
			Job_Location:  global.Work{strings.Trim(strings.Join(strings.Fields(fmt.Sprint(list[i].City)), ","), "[]")},
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
