package domain

import (
	"Office-Booking/domain/nearby/request"
	"time"

	"gorm.io/gorm"
)

type Nearby struct {
	ID             int            `json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
	NameFacilities string         `json:"namefacilities"`
	Jenis          string         `json:"jenis"`
	Jarak          string         `json:"jarak"`
	Latitude       string         `json:"latitude"`
	Longtitude     string         `json:"longtitude"`
	IDGedung       int            `json:"id_gedung"`
}

type Nearbys []Nearby

type NearbyRepository interface {
	Create(nearby *Nearby) (*Nearby, error)
	ReadByID(id int) (*Nearby, error)
	ReadAll() (*Nearbys, error)
	Delete(id int) (*Nearby, error)
	Updates(id int) (*Nearby, error)
}

type NearbyUsecase interface {
	Create(request request.NearbyCreateRequest) (*Nearby, error)
	ReadByID(id int) (*Nearby, error)
	ReadAll() (*Nearbys, error)
	Delete(id int) (*Nearby, error)
	Updates(id int) (*Nearby, error)
}
