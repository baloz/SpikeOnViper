package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig `mapstructure:",squash"`
	UserConfig `mapstructure:",squash"`
}

type DBConfig struct {
	Db_name string `mapstructure:"DBNAME"`
	Db_pass string `mapstructure:"DBPASS"`
}

type UserConfig struct {
	User_name string `mapstructure:"USERNAME"`
	User_pass string `mapstructure:"USERPASS"`
}

func main() {
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	

	var config Config

	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Println("Error reading config", err)
	}

	fmt.Println(config)
	fmt.Println(config.DBConfig.Db_name)
	fmt.Println(config.DBConfig.Db_pass)
	fmt.Println(config.UserConfig.User_name)
	fmt.Println(config.UserConfig.User_pass)

	printDBConfig(config.DBConfig)
	

}

func printDBConfig(dbConfig DBConfig){
	fmt.Println(dbConfig.Db_name)
	fmt.Println(dbConfig.Db_pass)

}
