package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbConf DBConfig `mapstructure:"DBCON"`
	DbConf1 DBConfig `mapstructure:"DBCONF"`
}

type DBConfig struct {
	Db_name string `mapstructure:"DBNAME"`
	Db_pass string `mapstructure:"DBPASS"`
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
}


