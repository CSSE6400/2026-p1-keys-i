package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TODOResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
	DeadlineAt string `json:"deadline_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func RegisterTODO(rtr *gin.RouterGroup) {
	todo := rtr.Group("todos")
	{
		// GET all
		todo.GET("", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, TODOResponse{
					ID:          1,
					Title:       "Watch CSSE6400 Lecture",
					Description: "Watch the CSSE6400 lecture on ECHO360 for week 1",
					Completed:   true,
					DeadlineAt:  "2026-02-27T18:00:00",
					CreatedAt:   "2026-02-20T14:00:00",
					UpdatedAt:   "2026-02-20T14:00:00",
			})
		})
		// GET :id
		// create a todo entry
		// update a todo entry
		// delete a todo entry
	}
}
