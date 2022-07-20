package usecase

import (
	"Office-Booking/delivery/http/helper"
	domain "Office-Booking/domain/users"
	"Office-Booking/domain/users/request"
	"Office-Booking/domain/users/response"
	"errors"
)

type userUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{UserRepo: ur}
}

func (u *userUsecase) Login(request request.LoginRequest) (*response.SuccessLogin, error) {
	if request.Email == "" || request.Password == "" {
		return nil, errors.New("email or password empty")
	}
	user := &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	resUser, _, err := u.UserRepo.CheckLogin(user)
	if err != nil {
		return nil, errors.New("email or password wrong")
	}

	jwt := helper.NewGoJWT()

	token, err := jwt.CreateTokenJWT(int(resUser.ID), resUser.Email, resUser.Name, resUser.Phone)
	if err != nil {
		return nil, err
	}

	resLogin := &response.SuccessLogin{ID: int(resUser.ID), Email: resUser.Email, Token: token}

	return resLogin, nil

}

func (u *userUsecase) RegisterUser(request request.RegisterRequest) (*domain.User, error) {
	if request.Email == "" {
		return nil, errors.New("email belum di isi")
	}
	if request.Password == "" {
		return nil, errors.New("password belum di isi")
	}
	if request.Name == "" {
		return nil, errors.New("username belum di isi")
	}
	if request.Fullname == "" {
		return nil, errors.New("nama lengkap belum di isi")
	}
	if request.Alamat == "" {
		return nil, errors.New("alamat belum di isi")
	}
	if request.Phone == "" {
		return nil, errors.New("no hp belum di isi")
	}

	user := &domain.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Fullname: request.Fullname,
		Alamat:   request.Alamat,
		Phone:    request.Phone,
	}

	createdUser, err := u.UserRepo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *userUsecase) Create(request request.UserCreateRequest) (*domain.User, error) {
	if request.Email == "" {
		return nil, errors.New("email empty")
	}
	user := &domain.User{
		Email:    request.Email,
		Password: request.Password,
		Name:     request.Name,
		Fullname: request.Fullname,
		Alamat:   request.Alamat,
		Phone:    request.Phone,
	}

	createdUser, err := u.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *userUsecase) ReadByID(id int) (*domain.User, error) {
	user, err := u.UserRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUsecase) ReadByName(name string) (*domain.User, error) {
	user, err := u.UserRepo.ReadByName(name)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUsecase) ReadAll() (*domain.Users, error) {
	foundUsers, err := u.UserRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundUsers, nil
}

func (u *userUsecase) Delete(id int) (*domain.User, error) {
	user, err := u.UserRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUsecase) Updates(id int) (*domain.User, error) {
	user, err := u.UserRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

// UpdatesAdmin implements domain.UserUsecase
func (u *userUsecase) UpdatesAdmin(id int) (*domain.User, error) {
	user, err := u.UserRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

// CreateBooking implements domain.UserUsecase
func (u *userUsecase) CreateBooking(request request.UserBookingReq) (*domain.User, error) {
	if request.Name == "" {
		return nil, errors.New("name belum di isi")
	}
	if request.Email == "" {
		return nil, errors.New("email belum di isi")
	}
	if request.IDBooking == 0 {
		return nil, errors.New("IDBooking belum di isi")
	}
	user := &domain.User{
		Name:      request.Name,
		Email:     request.Email,
		IDBooking: request.IDBooking,
	}

	createdUserbooking, err := u.UserRepo.CreateBooking(user)
	if err != nil {
		return nil, err
	}

	return createdUserbooking, nil
}
