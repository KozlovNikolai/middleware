package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()
	s.GET("/status", handler)
	err := s.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server running. ")

}

func handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "test")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
