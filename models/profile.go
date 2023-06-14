package models

type Profile struct {
	ID      int                 `json:"id" gorm:"primary_key:auto_increment"`
	Image   string              `json:"image" gorm:"type: varchar(255)"`
	Phone   string              `json:"phone" gorm:"type: varchar(255)"`
	Address string              `json:"address" gorm:"type: varchar(255)"`
	UserID  int                 `json:"user_id" gorm:"type:int"`
	User    UserProfileResponse `json:"user"`
}

type ProfileResponse struct {
	Image   string `json:"image"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	UserID  int    `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
