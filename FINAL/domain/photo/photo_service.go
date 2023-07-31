package photo

import (
	"context"
	"log"
)

type repository interface {
	writePhotoRepo
}

type writePhotoRepo interface {
	createPhoto(ctx context.Context, param Photo) (err error, photo Photo)
	getPhoto(ctx context.Context, photoId int) (err error, photo Photo, username string)
	getPhotoList(ctx context.Context) (err error, photoList []Photo)
	editPhoto(ctx context.Context, param Photo) (err error, photo Photo)
	removePhoto(ctx context.Context, photoId int, userId int) (err error)
}

type photoService struct {
	repo repository
	r    photoRepo
}

func newPhotoService(repo repository, r photoRepo) photoService {
	return photoService{
		repo: repo,
		r:    r,
	}
}

func (u photoService) createPhoto(ctx context.Context, req PhotoReq) (err error, resp PhotoResp) {
	var photo Photo

	err, photo = u.repo.createPhoto(ctx, req.PhotoReqIntoPhoto())
	if err != nil {
		log.Println(err)
		return
	}
	resp = photo.PhotoIntoPhotoResp()

	return err, resp
}

func (u photoService) getPhoto(ctx context.Context, req PhotoReq) (err error, resp PhotoResp) {
	err, photo, username := u.repo.getPhoto(ctx, req.PhotoId)
	resp = photo.PhotoIntoPhotoResp()
	resp.Username = username

	return err, resp
}

func (u photoService) getPhotoList(ctx context.Context) (err error, respList []PhotoResp) {
	err, photoList := u.repo.getPhotoList(ctx)
	_ = photoList

	for _, photo := range photoList {
		respList = append(respList, photo.PhotoIntoPhotoResp())
	}

	return err, respList
}

func (u photoService) editPhoto(ctx context.Context, req PhotoReq) (err error, resp PhotoResp) {
	var photo Photo

	err, photo = u.repo.editPhoto(ctx, req.PhotoReqIntoPhoto())
	if err != nil {
		log.Println(err)
		return
	}
	resp = photo.PhotoIntoPhotoResp()

	return err, resp
}

func (u photoService) removePhoto(ctx context.Context, req PhotoReq) (err error) {
	err = u.repo.removePhoto(ctx, req.PhotoId, req.UserId)

	return
}
