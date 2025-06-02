package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		val = strings.ToLower(val)

		if val == "admin" {
			log.Println("red button user deetected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
