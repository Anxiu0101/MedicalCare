package main

import (
	"MedicalCare/cache"
	"MedicalCare/conf"
	"MedicalCare/model"
	"MedicalCare/pkg/logging"
	"MedicalCare/router"
	"MedicalCare/service"
	"fmt"
	"net/http"
)

func init() {
	conf.Setup()
	model.Setup()
	cache.Setup()
	logging.Setup()
}

func main() {
	router := router.InitRouter()

	go service.Manager.Listen()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeout,
		WriteTimeout:   conf.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
