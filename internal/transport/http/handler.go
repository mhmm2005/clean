package http

import (
	"clean/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
)

type serviceInterface interface {
	GetHealth() string
	GetUserDataByName(string) *models.User
}

type Handler struct {
	Service serviceInterface
	Router  *gin.Engine
}

func NewHandler(service serviceInterface) *Handler {
	h := &Handler{
		Router:  gin.Default(),
		Service: service,
	}

	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200,
			gin.H{
				"message": h.Service.GetHealth(),
			})
	})

	h.Router.GET("/user/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.JSON(200,
			gin.H{
				"user": h.Service.GetUserDataByName(username),
			})
	})

}

func (h *Handler) Serve() error {
	err := h.Router.Run("0.0.0.0:6363")
	if err != nil {
		return errors.New("can not run server")
	}
	return nil
}
