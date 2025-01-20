package models

type User struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Username   string `json:"username" gorm:"null"`
	Password   string `json:"password" gorm:"null"`
	Email      string `json:"email" gorm:"null"`
	IsLoggedIn bool   `gorm:"default:false"`
}

func (User) TableName() string {
	return "user"
}
