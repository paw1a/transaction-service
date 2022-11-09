package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/transaction-service/internal/repository"
	"github.com/paw1a/transaction-service/internal/service"
	"net/http"
)

type Handler struct {
	clientService service.Clients
}

func NewHandler(conn *sqlx.DB) *Handler {
	clientRepo := repository.NewClientRepo(conn)
	return &Handler{
		clientService: service.NewClientService(clientRepo),
	}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(cors.Default())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/clients")
	}
}
