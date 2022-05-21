package context

import (
	"SNSLoginWithGolang/oauth2.0/service"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type APIContext struct {
	echo.Context
	Log *logrus.Entry
	service.SnsService
}

func (c *APIContext) GetContext() echo.Context {
	return c.Context
}

func (c *APIContext) GetRequestContext() context.Context {
	return c.Request().Context()
}
