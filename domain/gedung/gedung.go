package gedung

import (
	"Office-Booking/domain/gedung/request"
	"time"

	"gorm.io/gorm"
)

type Gedung struct {
	ID          int            `json:"id" gorm:"PrimaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Name        string         `json:"name"`
	Location    string         `json:"location"`
	Price       string         `json:"price"`
	Latitude    string         `json:"latitude"`
	Longitude   string         `json:"longitude"`
	Description string         `json:"description"`
	IDBooking   int            `json:"id_booking"`
	Review      []Review       `gorm:"Foreignkey:IDGedung;" json:"review"`
	Nearby      []Nearby       `gorm:"Foreignkey:IDGedung;" json:"nearby"`
	Jenis       []Jenis        `gorm:"Foreignkey:IDGedung;" json:"jenis"`
}

type Review struct {
	ID          int
	Rating      float64
	Description string
	IDGedung    string
}
type Nearby struct {
	ID             int
	NameFacilities string
	Jenis          string
	Jarak          string
	Latitude       string
	Longtitude     string
	IDGedung       string
}
type Jenis struct {
	ID       int
	Jenis    string
	IDGedung string
}

type GedungRepository interface {
	Create(gedung *Gedung) (*Gedung, error)
	GetAll() ([]Gedung, error)
	GetByID(id int) (*Gedung, error)
	GetByPrice(price string) (*Gedung, error)
	GetByLocation(location string) (*Gedung, error)
	Update(id int) (*Gedung, error)
	Delete(id int) (*Gedung, error)
}

type GedungUsecase interface {
	Create(request request.PostRequest) (*Gedung, error)
	GetAll() ([]Gedung, error)
	GetByID(id int) (*Gedung, error)
	GetByPrice(price string) (*Gedung, error)
	GetByLocation(location string) (*Gedung, error)
	Update(id int) (*Gedung, error)
	Delete(id int) (*Gedung, error)
}
