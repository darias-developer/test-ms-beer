package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"

	"github.com/darias-developer/test-ms-beer/config"
	"github.com/darias-developer/test-ms-beer/handler"
	"github.com/darias-developer/test-ms-beer/util"
)

func main() {

	//carga variables desde archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//inicializa la configuracion de los logs
	util.LoggerInit()

	util.LogInfo.Println("init main")

	//verifica conexion a la db
	util.LogInfo.Println("init CheckConnection")
	config.CheckConnection()

	c := cron.New()

	c.AddFunc("@every 60m", func() {
		util.LogInfo.Println("reconect CheckConnection")
		config.CheckConnection()
	})

	c.Start()
	//time.Sleep(time.Second * 5)

	//carga las rutas
	handler.RouterManager()
}
