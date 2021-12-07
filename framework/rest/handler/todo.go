package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-api/application/dto"
	"github.com/wellingtonlope/todo-api/application/usecase"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func initTodoHandler(e *echo.Echo, todoUsecase usecase.TodoUseCase) {
	h := TodoHandler{todoUseCase: todoUsecase}
	e.GET("/todos", h.GetAll)
	e.GET("/todos/:id", h.GetById)
	e.DELETE("/todos/:id", h.Delete)
	e.POST("/todos", h.Insert)
	e.PUT("/todos/:id", h.Update)
}

func (h *TodoHandler) GetAll(c echo.Context) error {
	todos, err := h.todoUseCase.GetAll()
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	todo, err := h.todoUseCase.GetById(id)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.todoUseCase.Delete(id)
	if err != nil {
		return handlerError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) Insert(c echo.Context) error {
	dto := dto.TodoNewDTO{}
	c.Bind(&dto)

	todo, err := h.todoUseCase.Insert(dto)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id := c.Param("id")
	dto := dto.TodoNewDTO{}
	c.Bind(&dto)

	todo, err := h.todoUseCase.Update(id, dto)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(http.StatusOK, todo)
}
