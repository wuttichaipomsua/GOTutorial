package main

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/fibo", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(fn(5)))
	})

	e.GET("/fibo/:name", func(c echo.Context) error {
		name := c.Param("name")
		//var param = strconv.Atoi(name)
		param, err := strconv.Atoi(name)
   		if err != nil {
      		// handle error
   		}
		return c.String(http.StatusOK, strconv.Itoa(fn(param)))
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func fn(n int) (r int) {
	if n <= 1 {
		r = n
	} else {
		r = fn(n-1) + fn(n-2)
	}
	return r
}
