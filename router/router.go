package router

import (
	"rest-api/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	H handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	return &Router{H: *h}
}
func (r *Router) RouterStart() {
	router := gin.Default()

	router.POST("/film", r.H.CreateFilm)
	router.GET("/film", r.H.GetAll)
	router.GET("/film/:id", r.H.GetById)

	router.Run(":28000")
}
