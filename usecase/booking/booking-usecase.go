package usecase

import (
	domain "Office-Booking/domain/booking"
	"Office-Booking/domain/booking/request"
	"errors"
)

type bookingUsecase struct {
	BookingRepo domain.BookingRepository
}

// Create implements domain.BookingUsecase
func (u *bookingUsecase) Create(request request.BookingRequest) (*domain.Booking, error) {
	if request.Status == "" {
		return nil, errors.New("status empty")
	}
	if request.BookingCode == "" {
		return nil, errors.New("bookingcode empty")
	}
	if request.TotalBooking == "" {
		return nil, errors.New("totalbooking empty")
	}
	if request.OrderDate == "" {
		return nil, errors.New("orderdate empty")
	}
	if request.CheckIn == "" {
		return nil, errors.New("checkin empty")
	}
	if request.CheckOut == "" {
		return nil, errors.New("checkout empty")
	}
	if request.Name == "" {
		return nil, errors.New("name empty")
	}
	if request.Phone == "" {
		return nil, errors.New("phone empty")
	}
	booking := &domain.Booking{
		Status:       request.Status,
		BookingCode:  request.BookingCode,
		TotalBooking: request.TotalBooking,
		OrderDate:    request.OrderDate,
		CheckIn:      request.CheckIn,
		CheckOut:     request.CheckOut,
		Name:         request.Name,
		Phone:        request.Phone,
	}

	postBooking, err := u.BookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	return postBooking, nil
}

// GetAll implements domain.BookingUsecase
func (u *bookingUsecase) GetAll() ([]domain.Booking, error) {
	foundBooking, err := u.BookingRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundBooking, nil
}

// GetByID implements domain.BookingUsecase
func (u *bookingUsecase) GetByID(id int) (*domain.Booking, error) {
	booking, err := u.BookingRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return booking, err
}

// Delete implements domain.BookingUsecase
func (u *bookingUsecase) Delete(id int) (*domain.Booking, error) {
	booking, err := u.BookingRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return booking, err
}

// Update implements domain.GedungUsecase
func (u *bookingUsecase) Update(id int) (*domain.Booking, error) {
	booking, err := u.BookingRepo.Update(id)
	if err != nil {
		return nil, err
	}

	return booking, err
}

func NewBookingUsecase(ur domain.BookingRepository) domain.BookingUsecase {
	return &bookingUsecase{BookingRepo: ur}
}
