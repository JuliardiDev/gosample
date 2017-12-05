package model

type UserInfo struct {
	Fullname string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	UserID   int64  `json:"user_id"`
	Language string `json:"lang,omitempty"`
}
