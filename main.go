package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"mysql/resource"
)


func main() {
	//初始化数据库
	configFile := flag.String("conf", "config/a.yaml", "path of config file")
	flag.Parse()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper read config is failed, err is %v configFile is %s ", err, configFile)
	}
	log.Println("logger init ")
	dbConf := viper.GetStringMapString("database")
	resource.InitDB(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])



}