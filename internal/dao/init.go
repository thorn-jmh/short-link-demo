package dao

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-svc-tpl/internal/dao/ent"

	_ "github.com/go-sql-driver/mysql"
)

var DB *ent.Client

// >>>>>>>>>>>> init >>>>>>>>>>>>

type DBCfg struct {
	DSN string
}

func InitDB() {
	var cfg DBCfg
	err := viper.Sub("Database").UnmarshalExact(&cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := ent.Open("mysql", cfg.DSN)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		logrus.Fatal(err)
	}

	DB = client
	if viper.GetString("App.RunLevel") == "debug" {
		DB = DB.Debug()
	}

}
