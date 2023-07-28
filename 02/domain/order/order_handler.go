package order

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	svc orderService
}

func newOrderHandler(svc orderService) orderHandler {
	return orderHandler{
		svc: svc,
	}
}

func writeMessageError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   err,
	})
}

func (u orderHandler) createOrder(ctx *gin.Context) {
	var req = OrderReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.createOrder(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create order success",
	})
}

func (u orderHandler) getOrderList(ctx *gin.Context) {
	err, resp := u.svc.getOrderList(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "get order list success",
	})
}

func (u orderHandler) editOrder(ctx *gin.Context) {
	var req = OrderReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err, resp := u.svc.editOrder(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data":    resp,
		"success": true,
		"message": "create order success",
	})
}

func (u orderHandler) removeOrder(ctx *gin.Context) {
	var req = OrderReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	err := u.svc.removeOrder(ctx.Request.Context(), req)
	if err != nil {
		log.Println(err)
		writeMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "remove order success",
	})
}
