package gortc_auth_v1

type logInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type verificationReq struct {
	Code string `json:"code"`
}

type forgetPasswordSendOTPReq struct {
	email string `json:"email",validate:"required,email"`
}
