package review

import "Office-Booking/domain/review/request"

type Review struct {
	ID          int     `json:"id" gorm:"PrimaryKey"`
	Img         string  `json:"img"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
	IDGedung    int     `json:"id_gedung"`
}

type Reviews []Review

type ReviewRepository interface {
	Create(review *Review) (*Review, error)
	GetByID(id int) (*Review, error)
	GetAll() (*Reviews, error)
	Delete(id int) (*Review, error)
}

type ReviewUsecase interface {
	Create(request request.ReviewPost) (*Review, error)
	GetByID(id int) (*Review, error)
	GetAll() (*Reviews, error)
	Delete(id int) (*Review, error)
}
