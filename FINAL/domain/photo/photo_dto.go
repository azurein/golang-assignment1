package photo

type PhotoReq struct {
	PhotoId  int    `json:"photoId"`
	UserId   int    `json:"userId"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
}

func (c PhotoReq) PhotoReqIntoPhoto() Photo {
	return Photo{
		PhotoId:  c.PhotoId,
		UserId:   c.UserId,
		Title:    c.Title,
		Caption:  c.Caption,
		PhotoUrl: c.PhotoUrl,
	}
}

type PhotoResp struct {
	PhotoId   int    `json:"photoId"`
	UserId    int    `json:"userId"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photoUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c Photo) PhotoIntoPhotoResp() PhotoResp {
	return PhotoResp{
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		Title:     c.Title,
		Caption:   c.Caption,
		PhotoUrl:  c.PhotoUrl,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
