package repository

import (
	domain "Office-Booking/domain/nearby"

	"gorm.io/gorm"
)

type nearbyRepository struct {
	Conn *gorm.DB
}

func NewNearbyRepository(Conn *gorm.DB) domain.NearbyRepository {
	return &nearbyRepository{Conn: Conn}
}

func (u *nearbyRepository) Create(nearby *domain.Nearby) (*domain.Nearby, error) {
	if err := u.Conn.Create(&nearby).Error; err != nil {
		return nil, err
	}

	return nearby, nil
}

func (u *nearbyRepository) ReadByID(id int) (*domain.Nearby, error) {
	nearby := &domain.Nearby{ID: id}
	if err := u.Conn.First(&nearby).Error; err != nil {
		return nil, err
	}

	return nearby, nil
}

func (u *nearbyRepository) ReadAll() (*domain.Nearbys, error) {
	nearbys := &domain.Nearbys{}
	u.Conn.Find(&nearbys)

	return nearbys, nil
}

func (u *nearbyRepository) Delete(id int) (*domain.Nearby, error) {
	nearby := &domain.Nearby{ID: id}
	if err := u.Conn.Delete(&nearby).Error; err != nil {
		return nil, err
	}
	return nearby, nil
}

func (u *nearbyRepository) Updates(id int) (*domain.Nearby, error) {
	nearby := &domain.Nearby{ID: id}
	if err := u.Conn.Updates(&nearby).Error; err != nil {
		return nil, err
	}

	return nearby, nil
}
