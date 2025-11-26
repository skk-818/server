package request

type LoginReq struct {
	Username string `json:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type RegisterReq struct {
	Phone    string `json:"phone" validate:"required,len=11,numeric"`
	Password string `json:"password" validate:"required,min=6,max=128"`
	Nickname string `json:"nickname" validate:"required,min=2,max=50"`
}

type EmailLoginReq struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,len=6,numeric"`
}
