package user

import (
	"errors"
	"travel-server/global"
	"travel-server/model"

	"gorm.io/gorm"
)

// 创建用户
func Create(req CreateReq) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		user := model.User{
			Username: req.Username,
			Password: req.Password,
		}
		if err := tx.Where("username = ?", req.Username).First(&user).Error; err == nil {
			return errors.New("用户已存在！")
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
}
