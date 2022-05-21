package service

import "github.com/labstack/echo/v4"

type SnsToken struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	TokenType             string `json:"token_type"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	IdToken               string `json:"id_token"`
	OauthToken            string `json:"oauth_token"`
	OauthTokenSecret      string `json:"oauth_token_secret"`
	UserID                string `json:"user_id"`
	ScreenName            string `json:"screen_name"`
}

type SnsUserInfo struct {
	SnsUserID  string
	Name       string
	Country    string
	CTN        string
	Birthday   string
	Gender     string
	Email      string
	Nickname   string
	ProfileImg string
	AgeRange   string
	Age        string
	BirthYeat  string
}

type SnsHandler interface {
	GetUserAuthorizationURL(c echo.Context) string
	Authorization(c echo.Context) (SnsToken, error)
	GetUserInfoWithToken(token string) (SnsUserInfo, error)
	//Unlink(userID string) error
	//Logout(userID string) error
}
