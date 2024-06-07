package entity

type LoginType int

const (
	LoginTypeUsername LoginType = 1
	LoginTypePhone    LoginType = 2
	LoginTypeEmail    LoginType = 3
)

type LoginRequest struct {
	Type     LoginType `json:"type"`
	LoginKey string    `json:"login_key"`
	LoginPwd string    `json:"login_pwd"`
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type AddPolicyRequest struct {
	Sub string `json:"sub"`
	Dom string `json:"dom"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}
