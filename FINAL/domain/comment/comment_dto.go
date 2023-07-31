package comment

type CommentReq struct {
	CommentId int    `json:"commentId"`
	PhotoId   int    `json:"photoId"`
	UserId    int    `json:"userId"`
	Message   string `json:"message"`
}

func (c CommentReq) CommentReqIntoComment() Comment {
	return Comment{
		CommentId: c.CommentId,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		Message:   c.Message,
	}
}

type CommentResp struct {
	CommentId int    `json:"commentId"`
	PhotoId   int    `json:"photoId"`
	UserId    int    `json:"userId"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c Comment) CommentIntoCommentResp() CommentResp {
	return CommentResp{
		CommentId: c.CommentId,
		PhotoId:   c.PhotoId,
		UserId:    c.UserId,
		Message:   c.Message,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
