package service

import (
	"SNSLoginWithGolang/oauth2.0/domain"
	"SNSLoginWithGolang/oauth2.0/sns_const"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SnsService struct {
	SnsHandler
	sns    *domain.SnsInfo
	logger *logrus.Entry
	// todo : session
}

func (s *SnsService) RedirectSnsAuthorization(c echo.Context) error {
	url := s.GetUserAuthorizationURL(c)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *SnsService) RequestToken(c echo.Context) (SnsToken, error) {
	return s.RequestToken(c)
}

func (s *SnsService) NewMember(snsID string, user *SnsUserInfo) (*domain.Member, error) {
	member := domain.Member{
		ID:         snsID + user.SnsUserID,
		SNSID:      snsID,
		Name:       user.Name,
		Gender:     user.Gender,
		CTN:        user.CTN,
		Email:      user.Email,
		NickName:   user.Nickname,
		ProfileImg: user.ProfileImg,
		CreatedAt:  time.Time{},
	}
	return &member, nil
}

// todo : create Member Info

func getSnsHandler(sns *domain.SnsInfo) SnsHandler {
	switch sns.SnsID {
	case sns_const.KakaoSNSID:
		// todo : new Kakao Service
		return nil
	case sns_const.GoogleSNSID:
		// todo : new Google Service
		return nil
	case sns_const.NaverSNSID:
		// todo : new Naver Service
		return nil
	default:
		return nil
	}
}

func NewSnsService(sns *domain.SnsInfo) (*SnsService, error) {
	ret := &SnsService{
		SnsHandler: getSnsHandler(sns),
		sns:        sns,
		logger: logrus.WithFields(logrus.Fields{
			"sns": sns.SnsID,
		}),
	}
	return ret, nil
}
