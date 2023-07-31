package photo

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.RouterGroup, db *sql.DB) {
	repo := newPhotoRepo(db)
	svc := newPhotoService(repo, repo)
	handler := newPhotoHandler(svc)

	photo := router.Group("/photo")
	{
		photo.POST("", handler.createPhoto)
		photo.GET("", handler.getPhoto)
		photo.GET("list", handler.getPhotoList)
		photo.PUT("", handler.editPhoto)
		photo.DELETE("", handler.removePhoto)
	}
}
