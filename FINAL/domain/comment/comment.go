package comment

type Comment struct {
	CommentId int    `db:"comment_id"`
	PhotoId   int    `db:"photo_id"`
	UserId    int    `db:"user_id"`
	Message   string `db:"message"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	Deleted   bool   `db:"deleted"`
}
