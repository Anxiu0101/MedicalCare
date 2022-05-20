package model

import (
	"MedicalCare/conf"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	// Account Info
	UserName string `json:"username" gorm:"column:username;not null;uniqueIndex"`
	Password string `json:"password" gorm:"column:password"`
	Avatars  string `json:"avatars" gorm:"column:avatars"`

	// User Info
	Email  string `json:"email" gorm:"type:varchar(100);unique"`
	Gender int    `json:"gender" gorm:"size:3"`
	Age    int    `json:"age" gorm:"size:8"`
	Tel    string `json:"tel" gorm:"size:64"`

	State bool `json:"state" gorm:"default:true;column:state"`
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), conf.AppSetting.PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// LoadRelation 将用户好友关系载入缓存
func LoadRelation() error {
	var err error
	return err
}

// UploadRelation 将用户关系从缓存中更新
func UploadRelation() {

}

/* serialization */

// AccountInfo 账户资料结构体
type AccountInfo struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"username" form:"username" example:"Anxiu"`
	Avatars  string `json:"avatars" form:"avatars"`
}

// BuildAccountInfo 将 User 对象序列化为 AccountInfo
func BuildAccountInfo(user User) AccountInfo {
	return AccountInfo{
		ID:       user.ID,
		UserName: user.UserName,
		Avatars:  user.Avatars,
	}
}
