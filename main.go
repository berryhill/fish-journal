// To Run:
// 	OWM_API_KEY=942a3a6482b1e600bf29d53e568cea09 go run main.go

package main

import (
	//"log"
	"fmt"

	"github.com/fish-journal/models"

	//owm "github.com/briandowns/openweathermap"
	"github.com/labstack/echo"
	"github.com/fish-journal/controllers"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	e.Get("/GetAllEvents", controllers.CreateEvent(&models.Event{}))
	e.GET("/CreateEvent/:type/:name/:length", controllers.CreateEvent(&models.Event{}))

/*
	w, err := owm.NewCurrent("F", "en") // fahrenheit (imperial) with English output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Phoenix")
	fmt.Println(w)

	w.CurrentByName("Denver")
	fmt.Println(w)
*/


	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}