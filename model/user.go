package model

import "time"

type User struct {
	Email       string    `bson:"_id"`
	Password    string    `bson:"password"`
	DisplayName string    `bson:"display_name"`
	WatchId     string    `bson:"watch_id"`
	Roles       []string  `bson:"roles"`
	BirthDate time.Time `bson:"birth_date"`
	Begin       time.Time `bson:"begin"`
	End         time.Time `bson:"end"`
	CreatedAt   time.Time `bson:"created_at"`
	CreatedBy   string    `bson:"created_by"`
	UpdatedAt   time.Time `bson:"updated_at"`
	UpdatedBy   string    `bson:"updated_by"`
}

type UserJSON struct {
	Email       string    `json:"email"`
	DisplayName string    `json:"displayName"`
	WatchId     string    `json:"watchId"`
	Roles       []string  `json:"roles"`
	Begin       time.Time `json:"begin"`
	End         time.Time `json:"end"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedBy   string    `json:"updatedBy"`
}

type CreateUserRequestJSON struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	DisplayName string `json:"displayName" validate:"required,min=4"`
	WatchId     string `json:"watchId" validate:"required"`
	CreatedBy string `json:"-"`
}

type UpdateUserRequestJSON struct {
	Email       string `json:"-"`
	DisplayName string `json:"displayName" validate:"required,min=4"`
	WatchId     string `json:"watchId" validate:"required"`
	Begin 	time.Time `json:"begin"`
	End     time.Time `json:"end"`
}

type UserLoginRequestJSON struct {
	Email string    `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type OldUserLoginRequestJSON struct {
	Email string    `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type MeJSON struct {
	Email       string   `json:"email"`
	DisplayName string   `json:"displayName"`
	WatchId     string   `json:"watchId"`
	Roles       []string `json:"roles"`
	BirthDate time.Time `json:"birthDate"`
}

type OldUser struct {
	Username    string    `bson:"_id" json:"username"`
	Password    string    `bson:"password" json:"-"`
	DisplayName string    `bson:"display_name" json:"display_name"`
	WatchId     string    `bson:"watch_id" json:"watch_id"`
	PhoneNumber string    `bson:"phone_number" json:"phone_number"`
	Email       string    `bson:"email" json:"email"`
	RoleId      string    `bson:"role_id" json:"-"`
	Created     time.Time `bson:"created" json:"created"`
	Modified    time.Time `bson:"modified" json:"modified"`
}