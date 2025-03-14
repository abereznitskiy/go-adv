package auth

type UserCreateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,len=11"`
}

type UserVerifyRequest struct {
	SessionId string `json:"sessionId"`
	Code      string `json:"code"`
}

type LoginResponse struct {
	SessionId string `json:"sessionId"`
}

type VerifyResponse struct {
	Token string `json:"token"`
}
