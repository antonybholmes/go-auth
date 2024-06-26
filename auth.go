package auth

import (
	"fmt"
	"strconv"

	"github.com/antonybholmes/go-sys"
	"github.com/google/uuid"
	"github.com/xyproto/randomstring"
	"golang.org/x/crypto/bcrypt"
)

type UrlReq struct {
	Url string `json:"url"`
}

type UrlCallbackReq struct {
	// the url that should form the email link in any emails that are sent
	CallbackUrl string `json:"callbackUrl"`
	// The url the callback url should redirect to once it completes
	Url string `json:"url"`
}

type User struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserName  string `db:"username"`
	Email     string `db:"email"`
}

type Permission struct {
	Uuid string `json:"uuid" db:"uuid"`
	Name string `json:"name" db:"name"`
}

type Role struct {
	Uuid        string       `json:"uuid" db:"uuid"`
	Name        string       `json:"name" db:"name"`
	Permissions []Permission `json:"permissions"`
}

type PublicRole struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// type PublicUser struct {
// 	Uuid      string `json:"uuid"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Username  string `json:"username"`
// 	Email     string `json:"email"`
// }

type AuthUser struct {
	Uuid           string `json:"uuid" db:"uuid"`
	FirstName      string `json:"firstName" db:"first_name"`
	LastName       string `json:"lastName" db:"last_name"`
	Username       string `json:"username" db:"username"`
	Email          string `json:"email" db:"email"`
	HashedPassword string `json:"-"`
	EmailVerified  bool   `json:"-"`
	CanSignIn      bool   `json:"-"`
	Updated        uint64 `json:"-"`
}

// func (user *AuthUser) Address() *mail.Address {
// 	return &mail.Address{Name: user.Name, Address: user.Email}
// }

func init() {
	randomstring.Seed()
}

func NewAuthUser(
	uuid string,
	firstName string,
	lastName string,
	userName string,
	email string,
	hashedPassword string,
	isVerified bool,
	canSignIn bool,
	updated uint64) *AuthUser {
	return &AuthUser{
		Uuid:           uuid,
		FirstName:      firstName,
		LastName:       lastName,
		Username:       userName,
		Email:          email,
		HashedPassword: hashedPassword,
		EmailVerified:  isVerified,
		CanSignIn:      canSignIn,
		Updated:        updated}
}

func (user *AuthUser) CheckPasswordsMatch(plainPwd string) error {
	return CheckPasswordsMatch(user.HashedPassword, plainPwd)
}

// Returns user details suitable for a web app to display
// func (user *AuthUser) ToPublicUser() *PublicUser {
// 	log.Debug().Msgf("here")
// 	return &PublicUser{Uuid: user.Uuid,
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Username:  user.Username,
// 		Email:     user.Email.Address}
// }

// Generate a one time code
func RandCode() string {
	return randomstring.CookieFriendlyString(32)
}

func Uuid() string {
	return uuid.New().String() // strings.ReplaceAll(u1.String(), "-", ""), nil
}

func HashPassword(password string) string {
	return string(sys.Must(bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)))
}

func CheckPasswordsMatch(hashedPassword string, plainPwd string) error {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice

	//log.Printf("comp %s %s\n", string(user.HashedPassword), string(plainPwd))

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPwd))

	if err != nil {
		return fmt.Errorf("passwords do not match")
	}

	return nil
}

func CreateOtp(user *AuthUser) string {
	return HashPassword(strconv.FormatUint(user.Updated, 10))

}

func CheckOtpValid(user *AuthUser, otp string) error {
	err := CheckPasswordsMatch(otp, strconv.FormatUint(user.Updated, 10))

	if err != nil {
		return fmt.Errorf("one time code has expired")
	}

	return nil
}
