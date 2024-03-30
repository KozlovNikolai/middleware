package endpoint

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	srv Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		srv: s,
	}
}

func (e Endpoint) Status(ctx echo.Context) error {
	dur := e.srv.DaysLeft()
	s := fmt.Sprintf("До 1го Апреля 2025 года осталось %d дней.\n", dur)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
