package controllers

import (
	//"os"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
	//"gopkg.in/mgo.v2"

	"github.com/fish-journal/models"
)

func respond(c echo.Context, err error, result interface{}) error {
	var msg string
	if err != nil {
		msg = fmt.Sprint(err)
	}

	statusCode := http.StatusOK
	/*
		switch err {
		case gorm.RecordNotFound:
			statusCode = http.StatusNotFound
		case gorm.InvalidSql, gorm.NoNewAttrs, gorm.NoValidTransaction, gorm.CantStartTransaction:
			statusCode = http.StatusInternalServerError
		}
	*/

	fmt.Println(result)

	return c.JSON(statusCode, map[string]interface{} {
		"result":  result,
		"error":   err != nil,
		"message": msg,
	})
}

func CreateEvent(my interface{}) echo.HandlerFunc {
	var err error
	return func(c echo.Context) error {
		if c.Param("type") == "catch" {
			err = nil
			name := c.Param("name")
			return respond(c, err, models.NewCatch(*models.NewSpecie(name), 15))
		}

		myType := reflect.ValueOf(my).Elem()

		x := reflect.New(myType.Type())
		x.Elem().Set(myType)


		if err := c.Bind(x.Interface()); err != nil {
			return respond(c, err, nil)
		}

		var err error
		err = nil

		return respond(c, err, x.Interface())
	}
}

/*
func GetAllEvents() echo.HandlerFunc {
	session, err := mgo.Dial(os.Getenv("127.0.0.1"))
	if err != nil {
		panic(err)
	}

	fmt.Println(session)

	return func(c *echo.Context) error {
		s := session.Clone()

		defer s.Close()
		//c.Next()
		return respond(c, err, models.NewCatch(models.NewSpecie("Yellow Fin Cutthroat"), 12))
	}
}
*/

/*
func AddEvent() *echo.HandlerFunc {

}
*/
