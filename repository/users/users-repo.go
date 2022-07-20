package repository

import (
	domain "Office-Booking/domain/users"

	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &userRepository{Conn: Conn}
}

func (u *userRepository) CheckLogin(user *domain.User) (*domain.User, bool, error) {
	if err := u.Conn.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return nil, false, err
	}

	return user, true, nil
}

func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	if err := u.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) ReadByID(id int) (*domain.User, error) {
	user := &domain.User{ID: id}
	if err := u.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) ReadByName(name string) (*domain.User, error) {
	user := &domain.User{Name: name}
	if err := u.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) ReadAll() (*domain.Users, error) {
	users := &domain.Users{}
	u.Conn.Find(&users)

	return users, nil
}

func (u *userRepository) RegisterUser(user *domain.User) (*domain.User, error) {
	if err := u.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) Delete(id int) (*domain.User, error) {
	user := &domain.User{ID: id}
	if err := u.Conn.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Updates(id int) (*domain.User, error) {
	user := &domain.User{ID: id}
	if err := u.Conn.Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatesAdmin implements domain.UserRepository
func (u *userRepository) UpdatesAdmin(id int) (*domain.User, error) {
	user := &domain.User{ID: id}
	if err := u.Conn.Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// CreateBooking implements domain.UserRepository
func (u *userRepository) CreateBooking(user *domain.User) (*domain.User, error) {
	if err := u.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
