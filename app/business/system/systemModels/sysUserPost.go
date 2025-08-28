package systemModels

type SysUserPost struct {
	UserId string `db:"user_id"`
	PostId string `db:"post_id"`
}

func NewSysUserPost(userId string, postId string) *SysUserPost {
	return &SysUserPost{UserId: userId, PostId: postId}
}
