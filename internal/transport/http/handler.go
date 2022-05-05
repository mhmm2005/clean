package http

import (
	"clean/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type serviceInterface interface {
	GetHealth() string
	GetUserDataByName(string) *models.User
	GetUsers() []*models.User
	GenerateJWT(string) string
	GetLogger() *zap.Logger
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

	h.Service.GetLogger().Info("mapHandler called")

	h.Router.GET("/health", JWTAuth(h.GetHealth))

	h.Router.GET("/user/:username", JWTAuth(h.GetUserDataByName))

	h.Router.GET("/users", JWTAuth(h.GetUsers))

	h.Router.POST("/token", JWTAuthAdmin(h.GenerateJWT))

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
	h.Service.GetLogger().Info("GetHealth handler")
	c.JSON(200,
		gin.H{
			"health": h.Service.GetHealth(),
		})

}

func (h *Handler) GetUserDataByName(c *gin.Context) {
	h.Service.GetLogger().Info("GetUserDataByName handler")
	username := c.Param("username")
	c.JSON(200,
		gin.H{
			"user": h.Service.GetUserDataByName(username),
		})
}

func (h *Handler) GetUsers(c *gin.Context) {
	h.Service.GetLogger().Info("GetUsers handler")
	c.JSON(200,
		gin.H{
			"users": h.Service.GetUsers(),
		})
}

func (h *Handler) GenerateJWT(c *gin.Context) {
	h.Service.GetLogger().Info("GenerateJWT handler")
	username := c.PostForm("username")
	fmt.Println(username)
	c.JSON(200,
		gin.H{
			"token": h.Service.GenerateJWT(username),
		})
}
