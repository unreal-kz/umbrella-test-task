package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{s: svc}
}

func (e *Endpoint) Status(c echo.Context) error {
	d := e.s.DaysLeft()
	s := fmt.Sprintf("Days left: %d", d)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
