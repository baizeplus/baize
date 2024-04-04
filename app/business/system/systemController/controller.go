package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUser, NewRole, NewPost, NewMenu, NewLogin, NewDictType, NewConfig, NewFile, NewNotice,
	NewProfile, NewDictData, NewDept, NewSse, wire.Struct(new(System), "*"))

type System struct {
	Login    *Login
	User     *User
	Dept     *Dept
	DictType *DictType
	DictData *DictData
	Menu     *Menu
	Role     *Role
	Post     *Post
	Profile  *Profile
	Config   *Config
	File     *File
	Sse      *Sse
	Notice   *Notice
}
