package profiledto

import "waysbean/models"

type ProfileResponse struct {
	ID           int                 `json:"id" gorm:"primary_key:auto_increment"`
	Phone        string              `json:"phone" gorm:"type: varchar(255)"`
	Name         string              `json:"name" gorm:"type: varchar(255)"`
	Gender       string              `json:"gender" gorm:"type: varchar(255)"`
	PostCode     string              `json:"post_code" gorm:"type: varchar(255)"`
	Address      string              `json:"address" gorm:"type: text"`
	ImageProfile string              `json:"image_profile" gorm:"type: varchar(255)"`
	UserID       int                 `json:"user_id"`
	User         models.UserResponse `json:"user"`
}
