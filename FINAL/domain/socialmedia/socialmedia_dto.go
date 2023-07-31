package socialmedia

type SocialmediaReq struct {
	SocialmediaId  int    `json:"socialmediaId"`
	UserId         int    `json:"userId"`
	Name           string `json:"name"`
	SocialmediaUrl string `json:"socialmediaUrl"`
}

func (c SocialmediaReq) SocialmediaReqIntoSocialmedia() Socialmedia {
	return Socialmedia{
		SocialmediaId:  c.SocialmediaId,
		UserId:         c.UserId,
		Name:           c.Name,
		SocialmediaUrl: c.SocialmediaUrl,
	}
}

type SocialmediaResp struct {
	SocialmediaId  int    `json:"socialmediaId"`
	UserId         int    `json:"userId"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	SocialmediaUrl string `json:"socialmediaUrl"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

func (c Socialmedia) SocialmediaIntoSocialmediaResp() SocialmediaResp {
	return SocialmediaResp{
		SocialmediaId:  c.SocialmediaId,
		UserId:         c.UserId,
		Name:           c.Name,
		SocialmediaUrl: c.SocialmediaUrl,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}
