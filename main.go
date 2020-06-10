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

	// result, _ := service.GetReviewer(6299)
	//{6299 1 笑笑老师 https://wx.hexiaoxiang.com/miniApp/avatar/ua.jpg 1373041051 2020-03-02 10:37:01 +0800 CST 2020-03-02 10:56:04 +0800 CST <nil>}

	//result , _:= service.SelectNameReviewer(6299)
	//{0 1 笑笑老师   <nil> <nil> <nil>}

	//_ = service.UpdateReviewer(6299)
	// UPDATE `reviewer` SET `cellphone` = '12', `updated_at` = '2020-06-09 15:11:49'  WHERE `reviewer`.`deleted_at` IS NULL AND ((id = 6299))

	// fmt.Println(result)

	// _ =  service.MapUpdateReviewer(6299, map[string]interface{}{"avatar_url":"www.baidu.com","staff_id":12,"cellphone":"123"})
	//UPDATE `reviewer` SET `avatar_url` = 'www.baidu.com', `cellphone` = '{', `staff_id` = 12, `updated_at` = '2020-06-09 15:36:23'  WHERE `reviewer`.`deleted_at` IS NULL AND ((id = 6299))

	//re ,_:= service.SelectNameReviewer(1)
	//[{29 木瓜老师} {30 贤斌老师} {12 笑笑老师} {1 人渣老师} {32 青老师} {3 小明老师}]


    //re, _ := service.OffSetReviewer()
    //[{6300 1 人渣老师 https://wx.hexiaoxiang.com/miniApp/avatar/ua.jpg 1383041051 2020-03-02 10:37:34 +0800 CST 2020-03-02 6:04 +0800 CST <nil>}]

	//fmt.Println(re)
    //fmt.Printf("%T\n",re)
    //fmt.Println(len(re))

}
