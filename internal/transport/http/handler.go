package http

import (
	"clean/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type serviceInterface interface {
	GetHealth() string
	GetUserDataByName(string) *models.User
	GenerateJWT(string) string
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

	h.Router.GET("/health", JWTAuth(h.GetHealth))

	h.Router.GET("/user/:username", JWTAuth(h.GetUserDataByName))

	h.Router.POST("/token", h.GenerateJWT)

}

func (h *Handler) Serve() error {

	go func() {
		err := h.Router.Run("0.0.0.0:6363")
		if err != nil {
			fmt.Println("can not run server")
		}
	}()

	return nil
}

func (h *Handler) GetHealth(c *gin.Context) {
	c.JSON(200,
		gin.H{
			"health": h.Service.GetHealth(),
		})

}

func (h *Handler) GetUserDataByName(c *gin.Context) {
	username := c.Param("username")
	c.JSON(200,
		gin.H{
			"user": h.Service.GetUserDataByName(username),
		})
}

func (h *Handler) GenerateJWT(c *gin.Context) {
	username := c.PostForm("username")
	fmt.Println(username)
	c.JSON(200,
		gin.H{
			"token": h.Service.GenerateJWT(username),
		})
}
