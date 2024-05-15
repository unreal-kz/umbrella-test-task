package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo"
)

const (
	role = "admin"
)

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(val, role) {
			log.Println("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
