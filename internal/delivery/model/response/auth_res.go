package response

type LoginResponse struct {
	Id        string `json:"id"`
	Access    string `json:"access_token"`
	AccessTTL int64  `json:"expired,omitempty"`
	Refresh   string `json:"refresh_token"`
}
