package main

import (
	"time"

	"github.com/labstack/echo/v4"
)

func LogTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}

		elapsed := time.Since(start)

		c.Logger().Printf("time elapsed: %s", elapsed)
		return nil
	}
}
