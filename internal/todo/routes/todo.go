package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TODOResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	DeadlineAt  string `json:"deadline_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type errorResp struct {
	Error string `json:"error"`
}

var baseTodo = TODOResponse{
	ID:          1,
	Title:       "Watch CSSE6400 Lecture",
	Description: "Watch the CSSE6400 lecture on ECHO360 for week 1",
	Completed:   true,
	DeadlineAt:  "2026-02-27T18:00:00",
	CreatedAt:   "2026-02-20T14:00:00",
	UpdatedAt:   "2026-02-20T14:00:00",
}

var baseTodoList = []TODOResponse{baseTodo}

func RegisterTODO(rtr *gin.RouterGroup) {
	todo := rtr.Group("/todos")
	{
		// GET all
		todo.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, baseTodoList)
		})

		// GET :id
		todo.GET("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResp{Error: "ID must be a valid integer"})
				return
			}

			resp := baseTodo
			resp.ID = id
			ctx.JSON(http.StatusOK, resp)
		})

		// POST create
		todo.POST("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, baseTodo)
		})

		// PUT update
		todo.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, baseTodo)
		})

		// DELETE
		todo.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, baseTodo)
		})
	}
}
