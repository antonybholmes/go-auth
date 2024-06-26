package userdb

import (
	"net/mail"

	"github.com/antonybholmes/go-auth"
)

// pretend its a global const
var userdb *auth.UserDb

func InitDB(file string) error {
	var err error

	userdb, err = auth.NewUserDB(file)

	return err

}

func CreateStandardUser(user *auth.SignupReq) (*auth.AuthUser, error) {
	return userdb.CreateStandardUser(user)
}

func FindUserById(id string) (*auth.AuthUser, error) {
	return userdb.FindUserById(id)
}

func FindUserByEmail(email *mail.Address) (*auth.AuthUser, error) {
	return userdb.FindUserByEmail(email)
}

func FindUserByUsername(username string) (*auth.AuthUser, error) {
	return userdb.FindUserByUsername(username)
}

func FindUserByUuid(uuid string) (*auth.AuthUser, error) {
	return userdb.FindUserByUuid(uuid)
}

func UserRoles(user *auth.AuthUser) (*[]auth.Role, error) {
	return userdb.UserRoles(user)
}

func RoleList(user *auth.AuthUser) (*[]string, error) {

	roles, err := userdb.UserRoles(user)

	if err != nil {
		return nil, err
	}

	ret := make([]string, len(*roles))

	for ri, role := range *roles {
		ret[ri] = role.Name
	}

	return &ret, nil

	//ret := strings.Join(tokens, ",")

	//return ret, nil

}

func UserPermissions(user *auth.AuthUser) (*[]auth.Permission, error) {
	return userdb.UserPermissions(user)
}

func PermissionList(user *auth.AuthUser) (*[]string, error) {

	permissions, err := userdb.UserPermissions(user)

	if err != nil {
		return nil, err
	}

	ret := make([]string, len(*permissions))

	for pi, permission := range *permissions {
		ret[pi] = permission.Name
	}

	return &ret, nil

}

func PublicUserRolePermissions(user *auth.AuthUser) (*[]auth.PublicRole, error) {
	return userdb.PublicUserRolePermissions(user)
}

func PublicUserRolePermissionsList(user *auth.AuthUser) (*auth.RoleMap, error) {

	roles, err := userdb.PublicUserRolePermissions(user)

	if err != nil {
		return nil, err
	}

	ret := make(auth.RoleMap)

	for _, role := range *roles {
		//for _, permission := range role.Permissions {
		//	tokens = append(tokens, fmt.Sprintf("%s::%s", role.Name, permission))
		//}

		_, ok := ret[role.Name]

		if !ok {
			ret[role.Name] = make([]string, 0, 10)
		}

		ret[role.Name] = append(ret[role.Name], role.Permissions...)
	}

	return &ret, nil

}

func SetIsVerified(user string) error {
	return userdb.SetIsVerified(user)
}

func SetPassword(uuid string, password string) error {
	return userdb.SetPassword(uuid, password)
}

func SetUsername(uuid string, username string) error {
	return userdb.SetUsername(uuid, username)
}

func SetName(uuid string, firstName string, lastName string) error {
	return userdb.SetName(uuid, firstName, lastName)
}

func SetUserInfo(uuid string, username string, firstName string, lastName string) error {
	return userdb.SetUserInfo(uuid, username, firstName, lastName)
}

func SetEmail(uuid string, email string) error {
	return userdb.SetEmail(uuid, email)
}

func SetEmailAddress(uuid string, address *mail.Address) error {
	return userdb.SetEmailAddress(uuid, address)
}
