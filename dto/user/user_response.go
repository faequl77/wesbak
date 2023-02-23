package userdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type UserGetResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Image   string `json:"image"`
	Phone   string `json:"phone"`
	PosCode string `json:"pos_code"`
	Address string `json:"address"`
}

type UserUpdateResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
	Phone    string `json:"phone"`
	PosCode  string `json:"pos_code"`
	Address  string `json:"address"`
}

type UserDeleteResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}
