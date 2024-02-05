package main

import (
	"github.com/labstack/echo"
)

func LogTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Before the process

		if err := next(c); err != nil {
			c.Error(err)
		}

		// After the process
		// Tips: https://pkg.go.dev/time#Since

		c.Logger().Printf("time elapsed?")
		return nil
	}
}
