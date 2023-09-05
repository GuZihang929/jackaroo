package huawei

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func huaweiOrmS(list []List, list1 []List1) (bool, error) {
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
			Company:       "华为",
			Title:         list[i].Jobname + list1[i].DISPLAYNAME,
			Job_detail:    "Duty:\n" + list1[i].MAINBUSINESS + "\nRequire:\n" + list1[i].DEMAND,
			Job_category:  list[i].JobFamilyName,
			Job_type_name: "校园招聘",
			Job_Location:  global.Work{list1[i].LOCDESCS},
			PushTime:      list[i].CreationDate,
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
