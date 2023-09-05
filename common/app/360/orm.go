package _60

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xiangxiang/jackaroo/common/mongoDB"
	"xiangxiang/jackaroo/global"
)

func _60Orm(list []List) (bool, error) {
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
			Company:       "360",
			Title:         list[i].JobAdName,
			Job_category:  list[i].Category,
			Job_detail:    "Duty:\n" + list[i].Duty + "\nRequire:\n" + list[i].Require,
			Job_type_name: list[i].ClassificationOne,
			Job_Location:  list[i].LocNames,
			PushTime:      time1,
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
