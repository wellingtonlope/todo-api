package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/application/myerrors"
)

func handlerError(c echo.Context, err *myerrors.Error) error {
	switch err.Type {
	case myerrors.DOMAIN, myerrors.REGISTER_ALREADY_EXISTS:
		return c.JSON(http.StatusBadRequest, err)
	case myerrors.REGISTER_NOT_FOUND:
		return c.JSON(http.StatusNotFound, err)
	case myerrors.REPOSITORY, myerrors.UNIDENTIFIED:
		return c.JSON(http.StatusInternalServerError, err)
	default:
		return c.JSON(http.StatusInternalServerError, err)
	}
}
