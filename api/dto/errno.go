package dto

const (
	NoErr = iota
	ErrEmailExist
	ErrCaptcha
	ErrPassword
	ErrUserNotFound

	ErrShortLinkExist
	ErrNoShortLink
	ErrPrivilege

	ErrShortLinkActive
	ErrShortLinkTime
)
