package user

type UserReq struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func (c UserReq) UserReqIntoUser() User {
	return User{
		UserId:   c.UserId,
		Username: c.Username,
		Email:    c.Email,
		Password: c.Password,
		Age:      c.Age,
	}
}

type UserResp struct {
	UserId    int    `json:"userId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c User) UserIntoUserResp() UserResp {
	return UserResp{
		UserId:    c.UserId,
		Username:  c.Username,
		Email:     c.Email,
		Age:       c.Age,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

type LoginResp struct {
	AccessToken string `json:"accessToken"`
}
