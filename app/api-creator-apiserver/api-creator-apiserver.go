package main

import (
	_apiRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/api/repository"
	_methodRepository "github.com/Hajime3778/api-creator-backend/pkg/admin/method/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/handler"
	_apiserverRepository "github.com/Hajime3778/api-creator-backend/pkg/apiserver/repository"
	"github.com/Hajime3778/api-creator-backend/pkg/apiserver/usecase"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/config"
	"github.com/Hajime3778/api-creator-backend/pkg/infrastructure/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// #region MongoDBテスト用のコード
	// logger.LoggingSetting("./log/")
	// apiserverConfig := config.NewConfig("../../apiserver.config.json")
	// db := database.NewDB(apiserverConfig)
	// mongoConn, ctx, cancel := db.NewMongoDBConnection()
	// defer cancel()
	//
	// apiserverRepository := _apiserverRepository.NewAPIServerRepository(ctx, mongoConn)
	//
	// collection := conn.Collection("test")
	// res, err := collection.InsertOne(ctx, bson.M{"name": "foo", "value": 123})
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// id := res.InsertedID
	//
	// log.Println(id)
	// #endregion

	// API Server側のMongoDB接続
	apiserverCfg := config.NewConfig("../../apiserver.config.json")
	db := database.NewDB(apiserverCfg)
	mongoConn, ctx, _ := db.NewMongoDBConnection()

	// 管理画面で設定されたMysql接続
	mysqlCfg := config.NewConfig("../../admin.config.json")
	mysqlDB := database.NewDB(mysqlCfg)
	mysqlConn := mysqlDB.NewMysqlConnection()

	apiRepository := _apiRepository.NewAPIRepository(mysqlConn)
	methodRepository := _methodRepository.NewMethodRepository(mysqlConn)
	apiserverRepository := _apiserverRepository.NewAPIServerRepository(ctx, mongoConn)

	engine := gin.Default()

	apiserverUsecase := usecase.NewAPIServerUsecase(apiRepository, methodRepository, apiserverRepository)
	handler.NewAPIServerHandler(engine, apiserverUsecase)

	engine.Run(":9000")
}
