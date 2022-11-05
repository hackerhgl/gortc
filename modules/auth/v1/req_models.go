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
	Email string `json:"email" validate:"required,email"`
}

type forgetPasswordVerifyOTPReq struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

type forgetPasswordVerifyResetReq struct {
	Email         string `json:"email" validate:"required,email"`
	Code          string `json:"code" validate:"required"`
	Password      string `json:"password" validate:"required"`
	ResetPassword string `json:"resetPassword" validate:"required"`
}
