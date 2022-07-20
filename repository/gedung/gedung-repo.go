package repository

import (
	domain "Office-Booking/domain/gedung"
	"fmt"

	"gorm.io/gorm"
)

type gedungRepository struct {
	Conn *gorm.DB
}

func NewGedungRepository(Conn *gorm.DB) domain.GedungRepository {
	return &gedungRepository{Conn: Conn}
}

// Create implements domain.GedungRepository
func (u *gedungRepository) Create(gedung *domain.Gedung) (*domain.Gedung, error) {
	if err := u.Conn.Create(&gedung).Error; err != nil {
		return nil, err
	}

	return gedung, nil
}

// Delete implements domain.GedungRepository
func (u *gedungRepository) Delete(id int) (*domain.Gedung, error) {
	gedung := &domain.Gedung{ID: id}
	if err := u.Conn.Delete(&gedung).Error; err != nil {
		return nil, err
	}
	return gedung, nil
}

// GetAll implements domain.GedungRepository
func (u *gedungRepository) GetAll() ([]domain.Gedung, error) {
	var gedungs []domain.Gedung
	err := u.Conn.Preload("Review").Preload("Jenis").Preload("Nearby").Find(&gedungs)
	if err.Error != nil {
		return []domain.Gedung{}, err.Error
	}
	fmt.Println(gedungs)
	return gedungs, nil
	// 	gedungs := &domain.Gedungs{}
	// 	u.Conn.Find(&gedungs)

	// 	return gedungs, nil
}

// GetByID implements domain.GedungRepository
func (u *gedungRepository) GetByID(id int) (*domain.Gedung, error) {
	gedung := &domain.Gedung{ID: id}
	if err := u.Conn.Preload("Review").Preload("Jenis").Preload("Nearby").First(&gedung).Error; err != nil {
		return nil, err
	}

	return gedung, nil
}

// GetByPrice implements domain.GedungRepository
func (u *gedungRepository) GetByPrice(price string) (*domain.Gedung, error) {
	gedung := &domain.Gedung{Price: price}
	if err := u.Conn.Where("price = ?", price).First(&gedung).Error; err != nil {
		return nil, err
	}

	return gedung, nil
}

func (u *gedungRepository) GetByLocation(location string) (*domain.Gedung, error) {
	gedung := &domain.Gedung{Location: location}
	if err := u.Conn.Where("location = ?", location).First(&gedung).Error; err != nil {
		return nil, err
	}

	return gedung, nil
}

// Update implements domain.GedungRepository
func (u *gedungRepository) Update(id int) (*domain.Gedung, error) {
	gedung := &domain.Gedung{ID: id}
	if err := u.Conn.Updates(&gedung).Error; err != nil {
		return nil, err
	}

	return gedung, nil
}
