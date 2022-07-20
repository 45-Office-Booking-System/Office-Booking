package domain

import (
	"Office-Booking/domain/users/request"
	"Office-Booking/domain/users/response"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int            `json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `json:"deletedAt"`
	Email           string         `json:"email"`
	Name            string         `json:"name"`
	Fullname        string         `json:"fullname"`
	Alamat          string         `json:"alamat"`
	Phone           string         `json:"phone"`
	Password        string         `json:"password"`
	NewPassword     string         `json:"newpassword"`
	ConfirmPassword string         `json:"Confirmpassword"`
	IDBooking       int            `json:"id_booking"`
}

type Users []User

type UserRepository interface {
	Create(user *User) (*User, error)
	CreateBooking(user *User) (*User, error)
	ReadByID(id int) (*User, error)
	ReadByName(name string) (*User, error)
	ReadAll() (*Users, error)
	CheckLogin(user *User) (*User, bool, error)
	RegisterUser(user *User) (*User, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
	UpdatesAdmin(id int) (*User, error)
}

type UserUsecase interface {
	Create(request request.UserCreateRequest) (*User, error)
	CreateBooking(request request.UserBookingReq) (*User, error)
	ReadByID(id int) (*User, error)
	ReadByName(name string) (*User, error)
	ReadAll() (*Users, error)
	Login(request request.LoginRequest) (*response.SuccessLogin, error)
	RegisterUser(request request.RegisterRequest) (*User, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
}
