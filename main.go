package main

import (
	"log"

	"github.com/joho/godotenv"

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

	//carga las rutas
	handler.RouterManager()
}
