package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig `mapstructure:",squash"`
	UserConfig `mapstructure:",squash"`
}

// type DBConfig struct {
// 	Db_name string `mapstructure:"DBNAME"`
// 	Db_pass string `mapstructure:"DBPASS"`
// }

// type UserConfig struct {
// 	User_name string `mapstructure:"USERNAME"`
// 	User_pass string `mapstructure:"USERPASS"`
// }

type DBConfig struct {
	Dbname string 
	Dbpass string 
}

type UserConfig struct {
	Username string 
	Userpass string 
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

	// fmt.Println(config)
	fmt.Println(config.DBConfig.Dbname)
	fmt.Println(config.DBConfig.Dbpass)
	fmt.Println(config.UserConfig.Username)
	fmt.Println(config.UserConfig.Userpass)

	// printDBConfig(config.DBConfig)
	

}

// func printDBConfig(dbConfig DBConfig){
// 	fmt.Println(dbConfig.Db_name)
// 	fmt.Println(dbConfig.Db_pass)

// }
