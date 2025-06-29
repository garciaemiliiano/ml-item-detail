package users

import "item-detail-api/src/repositories/basemodel"

type UserDAO struct {
	basemodel.BaseModel
	Username string `gorm:"column:username"`
}

func (UserDAO) TableName() string {
	return "users"
}
