package socialmedia

import (
	"context"
	"log"
)

type repository interface {
	writeSocialmediaRepo
}

type writeSocialmediaRepo interface {
	createSocialmedia(ctx context.Context, param Socialmedia) (err error, socialmedia Socialmedia)
	getSocialmedia(ctx context.Context, socialmediaId int) (err error, socialmedia Socialmedia, username string)
	getSocialmediaList(ctx context.Context) (err error, socialmediaList []Socialmedia)
	editSocialmedia(ctx context.Context, param Socialmedia) (err error, socialmedia Socialmedia)
	removeSocialmedia(ctx context.Context, socialmediaId int, userId int) (err error)
}

type socialmediaService struct {
	repo repository
	r    socialmediaRepo
}

func newSocialmediaService(repo repository, r socialmediaRepo) socialmediaService {
	return socialmediaService{
		repo: repo,
		r:    r,
	}
}

func (u socialmediaService) createSocialmedia(ctx context.Context, req SocialmediaReq) (err error, resp SocialmediaResp) {
	var socialmedia Socialmedia

	err, socialmedia = u.repo.createSocialmedia(ctx, req.SocialmediaReqIntoSocialmedia())
	if err != nil {
		log.Println(err)
		return
	}
	resp = socialmedia.SocialmediaIntoSocialmediaResp()

	return err, resp
}

func (u socialmediaService) getSocialmedia(ctx context.Context, req SocialmediaReq) (err error, resp SocialmediaResp) {
	err, socialmedia, username := u.repo.getSocialmedia(ctx, req.SocialmediaId)
	resp = socialmedia.SocialmediaIntoSocialmediaResp()
	resp.Username = username

	return err, resp
}

func (u socialmediaService) getSocialmediaList(ctx context.Context) (err error, respList []SocialmediaResp) {
	err, socialmediaList := u.repo.getSocialmediaList(ctx)
	_ = socialmediaList

	for _, socialmedia := range socialmediaList {
		respList = append(respList, socialmedia.SocialmediaIntoSocialmediaResp())
	}

	return err, respList
}

func (u socialmediaService) editSocialmedia(ctx context.Context, req SocialmediaReq) (err error, resp SocialmediaResp) {
	var socialmedia Socialmedia

	err, socialmedia = u.repo.editSocialmedia(ctx, req.SocialmediaReqIntoSocialmedia())
	if err != nil {
		log.Println(err)
		return
	}
	resp = socialmedia.SocialmediaIntoSocialmediaResp()

	return err, resp
}

func (u socialmediaService) removeSocialmedia(ctx context.Context, req SocialmediaReq) (err error) {
	err = u.repo.removeSocialmedia(ctx, req.SocialmediaId, req.UserId)

	return
}
