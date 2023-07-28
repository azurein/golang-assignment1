package order

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup, db *sql.DB) {
	repo := newOrderRepo(db)
	svc := newOrderService(repo, repo)
	handler := newOrderHandler(svc)

	order := router.Group("/order")
	{
		order.POST("create", handler.createOrder)
		order.POST("list", handler.getOrderList)
		order.POST("edit", handler.editOrder)
		order.POST("remove", handler.removeOrder)
	}
}
