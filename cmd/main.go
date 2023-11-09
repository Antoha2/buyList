package main

import (
	"fmt"
)

func main() {

	Run()
}

func Run() {

	fmt.Println("!!!!!!!!!!")
	cfg := cfg.GetConfig()
	// gormDB, err := initDb(cfg)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }

	// TgBotRepository := repository.NewRepository(gormDB)

	// TgBotGeokoder := geokoder.NewGeokoder()
	// TgBotMeteoYandex := yandeX.NewYandex(cfg)
	// TgBotMeteoGismeteo := gismeteo.NewGismeteo()

	// TgBotService := service.NewService(TgBotRepository, TgBotMeteoYandex, TgBotMeteoGismeteo, TgBotGeokoder)
	// TgBotTransport := trans.NewWeb(TgBotService, cfg)
	// HTTPTransport := http.NewHTTP(TgBotService, cfg)

	// go TgBotTransport.StartBot()
	// go HTTPTransport.StartHTTP()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// <-quit

	// HTTPTransport.Stop()
}
