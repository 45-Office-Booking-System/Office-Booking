package domain

import (
	"Office-Booking/domain/jenisgedung/request"
	"time"

	"gorm.io/gorm"
)

type Jenisgedung struct {
	ID        int            `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Jenis     string         `json:"jenis"`
	IDGedung  int            `json:"id_gedung"`
}

type Jenisgedungs []Jenisgedung

type JenisgedungRepository interface {
	Create(jenisgedung *Jenisgedung) (*Jenisgedung, error)
	ReadByID(id int) (*Jenisgedung, error)
	ReadAll() (*Jenisgedungs, error)
	Delete(id int) (*Jenisgedung, error)
	Updates(id int) (*Jenisgedung, error)
}

type JenisgedungUsecase interface {
	Create(request request.JenisgedungCreateRequest) (*Jenisgedung, error)
	ReadByID(id int) (*Jenisgedung, error)
	ReadAll() (*Jenisgedungs, error)
	Delete(id int) (*Jenisgedung, error)
	Updates(id int) (*Jenisgedung, error)
}
