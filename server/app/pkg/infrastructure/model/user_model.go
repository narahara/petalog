package model

import "common/pkg/domain/entity"

type User struct {
	UserId string `gorm:"column:user_id;primaryKey"`
}

func (a *User) ConvertToEntity() *entity.User {
	return &entity.User{
		UserId: a.UserId,
	}
}
