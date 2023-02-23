package authdto

type LoginResponse struct {
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
	Role  string `json:"role" gorm:"type: varchar(255)"`
	Token string `json:"token" gorm:"type: varchar(255)"`
	Image string `json:"image"`
}

type RegisterResponse struct {
	Name  string `json:"name" gorm:"type: varchar(255)"`
	Email string `json:"email" gorm:"type: varchar(255)"`
	Role  string `json:"role" gorm:"type: varchar(55)"`
}

type CheckAuthResponse struct {
	Id      int    `gorm:"type: int" json:"id"`
	Name    string `gorm:"type: varchar(255)" json:"name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	Role    string `gorm:"type: varchar(55)"  json:"role"`
	Image   string `json:"image"`
	Phone   string `json:"phone"`
	PosCode string `json:"pos_code"`
	Address string `json:"address"`
}
