package repository

import (
	domain "Office-Booking/domain/booking"
	"fmt"

	"gorm.io/gorm"
)

type bookingRepository struct {
	Conn *gorm.DB
}

// Create implements domain.BookingRepository
func (u *bookingRepository) Create(booking *domain.Booking) (*domain.Booking, error) {
	if err := u.Conn.Create(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

// GetAll implements domain.BookingRepository
func (u *bookingRepository) GetAll() ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := u.Conn.Preload("User").Preload("Jenis").Preload("Gedung").Find(&bookings)
	if err.Error != nil {
		return []domain.Booking{}, err.Error
	}
	fmt.Println(bookings)
	return bookings, nil
}

// GetByID implements domain.BookingRepository
func (u *bookingRepository) GetByID(id int) (*domain.Booking, error) {
	booking := &domain.Booking{ID: id}
	if err := u.Conn.Preload("User").Preload("Jenis").Preload("Gedung").First(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

// Update implements domain.BookingRepository
func (u *bookingRepository) Update(id int) (*domain.Booking, error) {
	booking := &domain.Booking{ID: id}
	if err := u.Conn.Updates(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

// Delete implements domain.BookingRepository
func (u *bookingRepository) Delete(id int) (*domain.Booking, error) {
	booking := &domain.Booking{ID: id}
	if err := u.Conn.Delete(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func NewBookingRepository(Conn *gorm.DB) domain.BookingRepository {
	return &bookingRepository{Conn: Conn}
}
