package solid

type SocialMedia int

const (
	Facebook SocialMedia = iota
	Instagram
	WhatsApp
)

type Post struct {
	From     SocialMedia
	PostText string
}
