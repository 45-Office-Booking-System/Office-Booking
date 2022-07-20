package booking

import (
	"Office-Booking/domain/booking/request"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID           int            `json:"id" gorm:"PrimaryKey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
	Status       string         `json:"status"`
	BookingCode  string         `json:"bookingcode"`
	TotalBooking string         `json:"totalbooking"`
	OrderDate    string         `json:"orderdate"`
	CheckIn      string         `json:"checkin"`
	CheckOut     string         `json:"checkout"`
	Name         string         `json:"fullname"`
	Phone        string         `json:"phone"`
	User         []User         `gorm:"Foreignkey:IDBooking;" json:"user"`
	Gedung       []Gedung       `gorm:"Foreignkey:IDBooking;" json:"gedung"`
	Jenis        []Jenis        `gorm:"Foreignkey:IDBooking;" json:"jenis"`
}
type User struct {
	ID        int
	Email     string
	Name      string
	Fullname  string
	Alamat    string
	Phone     string
	IDBooking string
}
type Gedung struct {
	ID          int
	Name        string
	Price       string
	Location    string
	Latitude    string
	Longitude   string
	Description string
	IDBooking   string
}
type Jenis struct {
	ID        int
	Jenis     string
	IDBooking string
}

type BookingRepository interface {
	Create(booking *Booking) (*Booking, error)
	GetAll() ([]Booking, error)
	GetByID(id int) (*Booking, error)
	Update(id int) (*Booking, error)
	Delete(id int) (*Booking, error)
}

type BookingUsecase interface {
	Create(request request.BookingRequest) (*Booking, error)
	GetAll() ([]Booking, error)
	GetByID(id int) (*Booking, error)
	Update(id int) (*Booking, error)
	Delete(id int) (*Booking, error)
}
