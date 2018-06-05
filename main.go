package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"microservice/upload"
	"net/http"
)

func main() {
	initConfig()
	startWebServer()
}

func startWebServer() {
	fmt.Println("start server")
	router := mux.NewRouter()
	router.HandleFunc("/tool/upload/image", upload.SaveImageURL)
	http.Handle("/", router)
	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("reading config failed: %s \n", err))
	} else {
		fmt.Println("Config Loaded")
	}
}
func initLogger() {

}
