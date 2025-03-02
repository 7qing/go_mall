package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (user *User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetByID(db *gorm.DB, id uint) (user *User, err error) {
	err = db.Where("id = ?", id).First(&user).Error
	return user, err
}

func DeleteUser(db *gorm.DB, userId uint) error {
	// 使用 user_id 查找并删除用户
	err := db.Unscoped().Where("id = ?", userId).Delete(&User{}).Error
	return err
}
