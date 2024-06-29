package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/internal/app/usecase"
)

var mapErrorTypeStatus = map[usecase.ErrorType]int{
	usecase.ErrorTypeInternalError: http.StatusInternalServerError,
	usecase.ErrorTypeBadRequest:    http.StatusBadRequest,
	usecase.ErrorTypeNotFound:      http.StatusNotFound,
}

type errorMessage struct {
	Message string `json:"message"`
}

func Error(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		errUC, ok := err.(usecase.Error)
		if !ok {
			return c.JSON(http.StatusInternalServerError, errorMessage{
				Message: "internal server error",
			})
		}
		status, ok := mapErrorTypeStatus[errUC.Type]
		if !ok {
			return c.JSON(http.StatusInternalServerError, errorMessage{
				Message: "internal server error",
			})
		}
		return c.JSON(status, errorMessage{
			Message: errUC.Message,
		})
	}
}
