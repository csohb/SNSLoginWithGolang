package endpoint

import (
	"SNSLoginWithGolang/oauth2.0/domain"
	"SNSLoginWithGolang/oauth2.0/rest_apis/context"
	"SNSLoginWithGolang/oauth2.0/service"
	"fmt"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
}
type LoginResponse struct {
}
type LoginEndpoint struct {
	*context.APIContext
	req    LoginRequest
	logger *logrus.Entry
}

func (ep *LoginEndpoint) Logger() *logrus.Entry {
	return ep.logger
}
func (ep *LoginEndpoint) GetRequestData() interface{} {
	return &ep.req
}
func (ep *LoginEndpoint) Handler() error {

	// sns token 정보 가지고 오기

	// sns Service 생성
	snsinfo := &domain.SnsInfo{
		SnsID:        "",
		HostURL:      "",
		ClientID:     "",
		RedirectUrl:  "",
		State:        "",
		Scope:        "",
		ResponseType: "",
		AdminKey:     "",
		RestAPIKey:   "",
		ClientSecret: "",
	}
	sns, err := service.NewSnsService(snsinfo)
	if err != nil {
		logrus.WithError(err).Error("new sns service error")
		return err
	}

	// redirect url 생성
	err = sns.RedirectSnsAuthorization(ep.Context)
	if err != nil {
		logrus.WithError(err).Error("redirect to authorization failed.")
		return err
	}

	// accesstoken 받기
	token, err := sns.RequestToken(ep.Context)
	if err != nil {
		logrus.WithError(err).Error("request token get err")
		return err
	}

	// token 정보 저장
	fmt.Println(token)

	// member 정보 받아오기

	// session 생성

	return nil
}
func (ep LoginEndpoint) IsRequiredAuth() bool {
	return true
}
func (ep LoginEndpoint) ApiName() string {
	return "Login"
}
func NewLoginEndpoint(c *context.APIContext) *LoginEndpoint {
	ret := &LoginEndpoint{
		APIContext: c,
		logger:     c.Log.WithField("api", "Login"),
	}
	return ret
}
