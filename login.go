package auth

type EmailOnlyLoginReq struct {
	Email string `json:"email" `
	UrlCallbackReq
}

type EmailPasswordLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UrlCallbackReq
}

type UsernamePasswordLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UrlCallbackReq
}

type PasswordLoginReq struct {
	Password string `json:"password"`
	UrlCallbackReq
}

// type LoginUser struct {
// 	Email    string
// 	Password []byte
// }

// func NewLoginUser(email string, password string) *LoginUser {
// 	return &LoginUser{Email: email, Password: []byte(password)}
// }

// func LoginUserFromReq(req *LoginReq) *LoginUser {
// 	return NewLoginUser(req.Email, req.Password)
// }
