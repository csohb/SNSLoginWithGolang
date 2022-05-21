package domain

import "time"

type SnsInfo struct {
	UpdatedAt    time.Time `json:"updated_at"`
	SnsID        string    `json:"sns_id"`
	HostURL      string    `json:"host_url"`
	ClientID     string    `json:"client_id"`
	RedirectUrl  string    `json:"redirect_url"`
	State        string    `json:"state"`
	Scope        string    `json:"scope"`
	ResponseType string    `json:"response_type"`
	AdminKey     string    `json:"admin_key"`
	RestAPIKey   string    `json:"rest_api_key"`
	ClientSecret string    `json:"client_secret"`
}
