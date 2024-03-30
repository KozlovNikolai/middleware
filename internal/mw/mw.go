package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleValue = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		s := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(roleValue, s) {
			log.Println("red button user detected")
		}
		next(ctx)
		return nil
	}
}
