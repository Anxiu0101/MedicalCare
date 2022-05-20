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

// @title MedicalCare API
// @version 1.0
// @description This is MedicalCare Project api docs.
// @contact.name Anxiu0101
// @contact.email anxiu.fyc@foxmail.com
// @contact.url https://github.com/Anxiu0101/MedicalCare/blob/master/README.md
// @host localhost:8000
// @BasePath /api/v1
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
