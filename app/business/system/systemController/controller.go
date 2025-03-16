package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUser, NewRole, NewPost, NewLogin, NewDictType, NewConfig, NewFile, NewNotice, NewPermission, NewSelectBox,
	NewProfile, NewDictData, NewDept, NewSse, wire.Struct(new(System), "*"))

type System struct {
	Login      *Login
	User       *User
	Dept       *Dept
	DictType   *DictType
	DictData   *DictData
	Role       *Role
	Post       *Post
	Profile    *Profile
	Config     *Config
	File       *File
	Sse        *Sse
	Notice     *Notice
	Permission *Permission
	SelectBox  *SelectBox
}
