package domain

import "time"

type Member struct {
	ID         string
	SNSID      string
	Name       string
	Gender     string
	CTN        string
	Email      string
	NickName   string
	ProfileImg string
	CreatedAt  time.Time
}
