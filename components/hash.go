package components

import (
	"fmt"
	"strings"
)

// Hash is a discord profile picture hash
// https://discord.com/developers/docs/reference#image-formatting
type Hash string

// Animated returns wether the url is an animated gif
func (h Hash) Animated() bool {
	return strings.HasPrefix(string(h), "a_")
}

// AvatarPNG returns a png url of this hash
func (u User) AvatarPNG() string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.png", u.ID, u.Avatar)
}

// AvatarURL returns a url of this user, being animated if it can
func (u User) AvatarURL() string {
	if u.Avatar.Animated() {
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s.gif", u.ID, u.Avatar)
	}
	return u.AvatarPNG()
}
