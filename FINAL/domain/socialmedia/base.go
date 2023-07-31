package socialmedia

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup, db *sql.DB) {
	repo := newSocialmediaRepo(db)
	svc := newSocialmediaService(repo, repo)
	handler := newSocialmediaHandler(svc)

	socialmedia := router.Group("/socialmedia")
	{
		socialmedia.POST("", handler.createSocialmedia)
		socialmedia.GET("", handler.getSocialmedia)
		socialmedia.GET("list", handler.getSocialmediaList)
		socialmedia.PUT("", handler.editSocialmedia)
		socialmedia.DELETE("", handler.removeSocialmedia)
	}
}
