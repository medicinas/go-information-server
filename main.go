package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"medicinas/information-server/config"
	"medicinas/information-server/service"
)

var appName = "infomationserver"

func init() {
	profile := flag.String("profile", "dev", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "https://med-config-server.herokuapp.com", "Address to config server")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")
	flag.Parse()
	fmt.Println("Specified configBranch is " + *configBranch)
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func main() {
	fmt.Printf("Starting %v\n", appName)
	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))
	service := service.WebServerService{}
	service.Run(viper.GetString("server_port"))
}
