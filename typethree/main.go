package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DbConf DBConfig 
	UserConf UserConfig 
}

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
	fmt.Println("--------reading with single struct---------")

	fmt.Println(config)

	var userConf UserConfig

	err = viper.Unmarshal(&userConf)

	if err != nil {
		fmt.Println("Error reading config", err)
	}

	var dbConf DBConfig
	err = viper.Unmarshal(&dbConf)

	if err != nil {
		fmt.Println("Error reading config", err)
	}

	configtwo := Config{
		UserConf: userConf,
		DbConf:   dbConf,
	}

	fmt.Println("--------reading with multiple structs and combining them---------")
	fmt.Println(configtwo)

}


