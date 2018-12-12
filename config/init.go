package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	projectName := "ss"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Failed to read config file in project: ", err)
	}
}

func GetMysqlConnectingString() string {
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", user, password, host, db, charset)
}
