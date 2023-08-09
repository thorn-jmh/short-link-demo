package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/ent"
	"testing"
)

func setup() {
	viper.SetConfigFile("../../config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	testDB, err := ent.Open("mysql", viper.GetString("Test.TestDB"))
	if err != nil {
		panic(err)
	}
	dao.DB = testDB
	if err := dao.DB.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func teardown() {
	dao.InitDB()
	gin.SetMode(gin.ReleaseMode)
}
