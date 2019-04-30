package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {

	////================ set config using primitive code
	// var configuration model.Configuration
	// file, err := os.Open("/home/ebdesk/goproject/GoLivyRestAPI/config/config.json")
	// if err != nil {
	// 	panic(err)
	// }
	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&configuration)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(configuration.Port)

	//========= set config using viper library
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/config/")
	errs := viper.ReadInConfig()
	if errs != nil {
		log.Fatal(errs)
	}
	conf := viper.GetString(key)
	// fmt.Println(viper.Get("port"))
	return conf
}
