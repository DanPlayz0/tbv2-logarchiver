package v2

import (
	"fmt"

	"github.com/TicketsBot-cloud/gdl/objects/channel"
	"github.com/TicketsBot-cloud/gdl/objects/guild"
	"github.com/TicketsBot-cloud/gdl/objects/user"
)

type Entities struct {
	Users    map[uint64]User    `json:"users"`
	Channels map[uint64]Channel `json:"channels"`
	Roles    map[uint64]Role    `json:"roles"`
}

type User struct {
	Id       uint64 `json:"id,string"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Bot      bool   `json:"bot"`
}

func (u *User) AvatarUrl(size int) string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.webp?size=%d", u.Id, u.Avatar, size)
}

type Channel struct {
	Id   uint64              `json:"id"`
	Name string              `json:"name"`
	Type channel.ChannelType `json:"type"`
}

type Role struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Colour uint32 `json:"color"`
}

func UserFromGdl(entity user.User) User {
	return User{
		Id:       entity.Id,
		Username: entity.Username,
		Avatar:   entity.Avatar.String(),
		Bot:      entity.Bot,
	}
}

func ChannelFromGdl(entity channel.Channel) Channel {
	return Channel{
		Id:   entity.Id,
		Name: entity.Name,
		Type: entity.Type,
	}
}

func RoleFromGdl(entity guild.Role) Role {
	return Role{
		Id:     entity.Id,
		Name:   entity.Name,
		Colour: uint32(entity.Color),
	}
}
