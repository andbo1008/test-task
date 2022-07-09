package handler

import (
	"net/http"
	"rest-api/film"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateFilm(c *gin.Context) {
	var f film.Film
	logrus.Print(f)

	if err := c.ShouldBind(&f); err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Print(f)
	id, err := h.service.CreateFilm(f)
	if err != nil {
		logrus.Fatal(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"id":     id,
	})
}
func (h *Handler) GetAll(c *gin.Context) {
	allFilms, err := h.service.GetAll()
	if err != nil {
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   allFilms,
	})

}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	film, err := h.service.GetById(id)
	if err != nil {
		logrus.Fatal(c, http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success",
		"id":          film.Id,
		"title":       film.Title,
		"releaseDate": film.ReleaseDate,
	})
}
