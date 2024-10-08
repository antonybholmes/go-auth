package auth

// type PasswordlessLogin struct {
// 	Username string `json:"username" `
// 	UrlCallbackReq
// }

// type EmailPasswordLoginReq struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	UrlCallbackReq
// }

// type UsernamePasswordLoginReq struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	UrlCallbackReq
// }

// type PasswordLoginReq struct {
// 	Password string `json:"password"`
// 	UrlCallbackReq
// }

// type PasswordResetReq struct {
// 	Password string `json:"password"`
// }

// type UsernameReq struct {
// 	Username string `json:"username"`
// }

type NewPasswordReq struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type LoginReq struct {
	UrlCallbackReq
	PublicId        string   `json:"publicId"`
	Username        string   `json:"username"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	FirstName       string   `json:"firstName"`
	LastName        string   `json:"lastName"`
	Roles           []string `json:"roles"`
	EmailIsVerified bool     `json:"emailIsVerified"`
	StaySignedIn    bool     `json:"staySignedIn"`
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
