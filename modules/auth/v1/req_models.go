package gortc_auth_v1

type LogInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
