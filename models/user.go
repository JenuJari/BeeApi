package models

import (
	"BeeApi/common"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(User), new(Profile))
	// UserList = make(map[string]*User)
	// u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	// UserList["user_11111"] = &u
}

type Profile struct {
	Id        int32
	BirthDate time.Time `orm:"type(date)"`
	Gender    string    `orm:"size(10)"`
	User      *User     `orm:"reverse(one)"` // Reverse relationship (optional)
}

type User struct {
	Id       int32
	Email    string   `orm:"size(500)"`
	Password string   `orm:"size(500)"`
	Profile  *Profile `orm:"rel(one)"` // OneToOne relation
}

/*
========================================================================================================================
    Oprations on User model
========================================================================================================================
*/

// RegisterUser Register a new user
func RegisterUser(u *User, p *Profile) int32 {
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	beego.Debug(p)

	o.Insert(p)
	u.Profile = p

	beego.Debug(p)

	hash, er := common.HashPassword(u.Password)

	if er != nil {
		panic("Error in hasing the passowrd.")
	}

	u.Password = hash

	o.Insert(u)
	beego.Debug(u)

	return u.Id
}

// AddUser add new User
func AddUser(u User) string {
	return "true"
}

// GetUser Get a User
func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

// GetAllUsers Get List of users
func GetAllUsers() map[string]*User {
	return UserList
}

// UpdateUser update user
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		// if uu.Username != "" {
		// 	u.Username = uu.Username
		// }
		// if uu.Password != "" {
		// 	u.Password = uu.Password
		// }
		// if uu.Profile.Age != 0 {
		// 	u.Profile.Age = uu.Profile.Age
		// }
		// if uu.Profile.Address != "" {
		// 	u.Profile.Address = uu.Profile.Address
		// }
		// if uu.Profile.Gender != "" {
		// 	u.Profile.Gender = uu.Profile.Gender
		// }
		// if uu.Profile.Email != "" {
		// 	u.Profile.Email = uu.Profile.Email
		// }
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

// Login check login cred for user
func Login(email, password string) bool {

	o := orm.NewOrm()
	user := User{Email: email}
	err := o.Read(&user, "Email")

	if err != nil {
		return false
	}

	return common.CheckPasswordHash(password, user.Password)
}

// DeleteUser  Delets a user.
func DeleteUser(uid string) {
	delete(UserList, uid)
}
