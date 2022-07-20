package usecase

import (
	domain "Office-Booking/domain/gedung"
	"Office-Booking/domain/gedung/request"
	"errors"
)

type gedungUsecase struct {
	GedungRepo domain.GedungRepository
}

// Create implements domain.GedungUsecase
func (u *gedungUsecase) Create(request request.PostRequest) (*domain.Gedung, error) {
	if request.Name == "" {
		return nil, errors.New("nama gedung belum diisi")
	}
	if request.Location == "" {
		return nil, errors.New("lokasi gedung belum diisi")
	}
	if request.Price == "" {
		return nil, errors.New("price belum diisi")
	}
	if request.Latitude == "" {
		return nil, errors.New("latitude belum diisi")
	}
	if request.Longitude == "" {
		return nil, errors.New("longitude belum diisi")
	}
	if request.Description == "" {
		return nil, errors.New("description belum diisi")
	}
	if request.IDBooking == 0 {
		return nil, errors.New("id_booking empty")
	}
	gedung := &domain.Gedung{
		Name:        request.Name,
		Location:    request.Location,
		Price:       request.Price,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		Description: request.Description,
		IDBooking:   request.IDBooking,
	}

	postGedung, err := u.GedungRepo.Create(gedung)
	if err != nil {
		return nil, err
	}

	return postGedung, nil
}

// Delete implements domain.GedungUsecase
func (u *gedungUsecase) Delete(id int) (*domain.Gedung, error) {
	gedung, err := u.GedungRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return gedung, err
}

// GetAll implements domain.GedungUsecase
func (u *gedungUsecase) GetAll() ([]domain.Gedung, error) {
	// var gedungs []Gedung
	// err := db.Model(&Gedung{}).Preload("Review").Find(&gedungs).Error
	// return gedungs, err
	foundGedung, err := u.GedungRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundGedung, nil
}

// GetByID implements domain.GedungUsecase
func (u *gedungUsecase) GetByID(id int) (*domain.Gedung, error) {
	gedung, err := u.GedungRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return gedung, err
}

// GetByPrice implements domain.GedungUsecase
func (u *gedungUsecase) GetByPrice(price string) (*domain.Gedung, error) {
	gedung, err := u.GedungRepo.GetByPrice(price)
	if err != nil {
		return nil, err
	}

	return gedung, err
}

func (u *gedungUsecase) GetByLocation(location string) (*domain.Gedung, error) {
	gedung, err := u.GedungRepo.GetByLocation(location)
	if err != nil {
		return nil, err
	}

	return gedung, err
}

// Update implements domain.GedungUsecase
func (u *gedungUsecase) Update(id int) (*domain.Gedung, error) {
	gedung, err := u.GedungRepo.Update(id)
	if err != nil {
		return nil, err
	}

	return gedung, err
}

func NewGedungUseCase(ur domain.GedungRepository) domain.GedungUsecase {
	return &gedungUsecase{GedungRepo: ur}
}
