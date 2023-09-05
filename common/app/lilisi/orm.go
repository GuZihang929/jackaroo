package lilisi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func lilisiOrm(list []List) (bool, error) {
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
		str := ""
		for j := 0; j < len(list[i].CityList); j++ {
			str += list[i].CityList[j].Local
		}
		job := &global.Hello{
			Company:       "莉莉丝",
			Title:         list[i].Title,
			Job_detail:    "Duty:\n" + list[i].Description + "\nRequire:\n" + list[i].Requirement,
			Job_category:  fmt.Sprintf("%v", list[i].JobCategory),
			Job_type_name: fmt.Sprintf("%v", list[i].RecruitType),
			Job_Location:  global.Work{str},
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
