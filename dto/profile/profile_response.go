package profiledto

import "waysbean/models"

type ProfileResponse struct {
	ID      int                        `json:"id" gorm:"primary_key:auto_increment"`
	Image   string                     `json:"image" gorm:"type: varchar(255)"`
	Phone   string                     `json:"phone" gorm:"type: varchar(255)"`
	Address string                     `json:"address" gorm:"type: varchar(255)"`
	UserID  int                        `json:"user_id" gorm:"type:int"`
	User    models.UserProfileResponse `json:"user"`
}