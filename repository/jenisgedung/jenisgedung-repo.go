package repository

import (
	domain "Office-Booking/domain/jenisgedung"

	"gorm.io/gorm"
)

type jenisgedungRepository struct {
	Conn *gorm.DB
}

func NewJenisgedungRepository(Conn *gorm.DB) domain.JenisgedungRepository {
	return &jenisgedungRepository{Conn: Conn}
}

func (u *jenisgedungRepository) Create(jenisgedung *domain.Jenisgedung) (*domain.Jenisgedung, error) {
	if err := u.Conn.Create(&jenisgedung).Error; err != nil {
		return nil, err
	}

	return jenisgedung, nil
}

func (u *jenisgedungRepository) ReadByID(id int) (*domain.Jenisgedung, error) {
	jenisgedung := &domain.Jenisgedung{ID: id}
	if err := u.Conn.First(&jenisgedung).Error; err != nil {
		return nil, err
	}

	return jenisgedung, nil
}

func (u *jenisgedungRepository) ReadAll() (*domain.Jenisgedungs, error) {
	jenisgedungs := &domain.Jenisgedungs{}
	u.Conn.Find(&jenisgedungs)

	return jenisgedungs, nil
}

func (u *jenisgedungRepository) Delete(id int) (*domain.Jenisgedung, error) {
	jenisgedung := &domain.Jenisgedung{ID: id}
	if err := u.Conn.Delete(&jenisgedung).Error; err != nil {
		return nil, err
	}
	return jenisgedung, nil
}

func (u *jenisgedungRepository) Updates(id int) (*domain.Jenisgedung, error) {
	jenisgedung := &domain.Jenisgedung{ID: id}
	if err := u.Conn.Updates(&jenisgedung).Error; err != nil {
		return nil, err
	}

	return jenisgedung, nil
}
