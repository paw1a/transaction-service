package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/paw1a/transaction-service/internal/repository"
	"github.com/paw1a/transaction-service/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	clientService      service.Clients
	transactionService service.Transactions
}

func NewHandler(conn *sqlx.DB) *Handler {
	clientRepo := repository.NewClientRepo(conn)
	transactionRepo := repository.NewTransactionRepo(conn)

	clientService := service.NewClientService(clientRepo)
	transactionService, err := service.NewTransactionService(transactionRepo, clientService)
	if err != nil {
		log.Fatalf("failed to init transaction service: %v", err)
	}

	return &Handler{
		clientService:      clientService,
		transactionService: transactionService,
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
		clientsApi := api.Group("/clients")
		{
			clientsApi.GET("/", h.clientFindAll)
			clientsApi.GET("/:id", h.clientFindById)
			clientsApi.POST("/", h.clientCreate)
			clientsApi.DELETE("/:id", h.clientDelete)
		}

		transactionApi := api.Group("/transactions")
		{
			transactionApi.GET("/", h.transactionFindAll)
			transactionApi.GET("/:id", h.transactionFindById)
			transactionApi.POST("/", h.transactionCreate)
		}
	}
}
