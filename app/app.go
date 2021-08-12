package app

import (
	"fmt"
	_ "github.com/InuNora/incrementer/docs"
	"github.com/InuNora/incrementer/service"
	"github.com/InuNora/incrementer/service/impl"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
)

// @title Swagger Incrementer
// @version 1.0
// @description Incrementer (Golang)

// @contact.name Julia Konopatova
// @contact.email j.pikareva@gmail.com

// @host localhost:1323
// @BasePath /

func StartApplication() {
	var incr service.IncrementerService
	incr, err := impl.NewIncrementer(0, impl.MaxUint)
	if err != nil {
		fmt.Errorf("Error: %s \n", err.Error())
	}
	fmt.Printf("Start Value: %v \n", incr.GetNumber())
	incr.IncrementNumber()
	fmt.Printf("Current Value: %v \n", incr.GetNumber())
	_ = incr.SetMaximumValue(3)
	for i := 0; i < 3; i++ {
		incr.IncrementNumber()
		fmt.Printf("Current Value: %v \n", incr.GetNumber())
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))

}
