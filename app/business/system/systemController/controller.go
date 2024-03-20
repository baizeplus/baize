package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUser, NewRole, NewPost, NewMenu, NewLogin, NewDictType, NewConfig, NewFile,
	NewProfile, NewDictData, NewDept, wire.Struct(new(System), "*"))

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
}
