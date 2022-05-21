package endpoint

import (
	"SNSLoginWithGolang/oauth2.0/rest_apis/context"
	"github.com/sirupsen/logrus"
)

type AccessTokenRequest struct {
}
type AccessTokenResponse struct {
}
type AccessTokenEndpoint struct {
	*context.APIContext
	req    AccessTokenRequest
	logger *logrus.Entry
}

func (ep *AccessTokenEndpoint) Logger() *logrus.Entry {
	return ep.logger
}
func (ep *AccessTokenEndpoint) GetRequestData() interface{} {
	return &ep.req
}
func (ep *AccessTokenEndpoint) Handler() error {
	return nil
}
func (ep AccessTokenEndpoint) IsRequiredAuth() bool {
	return true
}
func (ep AccessTokenEndpoint) ApiName() string {
	return "AccessToken"
}
func NewAccessTokenEndpoint(c *context.APIContext) *AccessTokenEndpoint {
	ret := &AccessTokenEndpoint{
		APIContext: c,
		logger:     c.Log.WithField("api", "AccessToken"),
	}
	return ret
}
