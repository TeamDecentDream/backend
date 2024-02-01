package models

import "time"

type Member struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Address     string      `json:"address"`
	Authorities []Authority `json:"authorities"`
	RegDate     time.Time   `json:"reg_date"`
}

type Authority struct {
	MemberId int
	Role     string
}

type Credential struct {
	Provider string `json:"provider"`
	Code     string `json:"code"`
}

type KakaoOAuthResponse struct {
	AccessToken           string  `json:"access_token"`
	ExpiresIn             int     `json:"expires_in"`
	IDToken               string  `json:"id_token"`
	RefreshToken          string  `json:"refresh_token"`
	RefreshTokenExpiresIn float64 `json:"refresh_token_expires_in"`
	Scope                 string  `json:"scope"`
	TokenType             string  `json:"token_type"`
	Error                 string  `json:"error"`
}

type KakaoMemberResponse struct {
	ConnectedAt  string  `json:"connected_at"`
	ID           float64 `json:"id"`
	KakaoAccount struct {
		Email               string `json:"email"`
		EmailNeedsAgreement bool   `json:"email_needs_agreement"`
		HasEmail            bool   `json:"has_email"`
		IsEmailValid        bool   `json:"is_email_valid"`
		IsEmailVerified     bool   `json:"is_email_verified"`
		Profile             struct {
			IsDefaultImage    bool   `json:"is_default_image"`
			Nickname          string `json:"nickname"`
			ProfileImageURL   string `json:"profile_image_url"`
			ThumbnailImageURL string `json:"thumbnail_image_url"`
		} `json:"profile"`
		ProfileImageNeedsAgreement    bool `json:"profile_image_needs_agreement"`
		ProfileNicknameNeedsAgreement bool `json:"profile_nickname_needs_agreement"`
	} `json:"kakao_account"`
	Properties struct {
		Nickname       string `json:"nickname"`
		ProfileImage   string `json:"profile_image"`
		ThumbnailImage string `json:"thumbnail_image"`
	} `json:"properties"`
}
