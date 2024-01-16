package model

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey, type=varchar(50)"`
	Email     string    `json:"email" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
}

func (user User) GetName() string {
	return "User"
}
