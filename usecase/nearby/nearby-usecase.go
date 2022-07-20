package usecase

import (
	domain "Office-Booking/domain/nearby"
	"Office-Booking/domain/nearby/request"
	"errors"
)

type nearbyUsecase struct {
	NearbyRepo domain.NearbyRepository
}

func NewNearbyUsecase(ur domain.NearbyRepository) domain.NearbyUsecase {
	return &nearbyUsecase{NearbyRepo: ur}
}

func (u *nearbyUsecase) Create(request request.NearbyCreateRequest) (*domain.Nearby, error) {
	if request.NameFacilities == "" {
		return nil, errors.New("name facilities empty")
	}
	if request.Jenis == "" {
		return nil, errors.New("jenis empty")
	}
	if request.Jarak == "" {
		return nil, errors.New("jarak empty")
	}
	if request.Latitude == "" {
		return nil, errors.New("latitude empty")
	}
	if request.Longtitude == "" {
		return nil, errors.New("longtitude empty")
	}
	if request.IDGedung == 0 {
		return nil, errors.New("id_gedung empty")
	}
	nearby := &domain.Nearby{
		NameFacilities: request.NameFacilities,
		Jenis:          request.Jenis,
		Jarak:          request.Jarak,
		Latitude:       request.Latitude,
		Longtitude:     request.Longtitude,
		IDGedung:       request.IDGedung,
	}

	createdNearby, err := u.NearbyRepo.Create(nearby)
	if err != nil {
		return nil, err
	}

	return createdNearby, nil
}

func (u *nearbyUsecase) ReadByID(id int) (*domain.Nearby, error) {
	nearby, err := u.NearbyRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return nearby, err
}

func (u *nearbyUsecase) ReadAll() (*domain.Nearbys, error) {
	foundNearbys, err := u.NearbyRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundNearbys, nil
}

func (u *nearbyUsecase) Delete(id int) (*domain.Nearby, error) {
	nearby, err := u.NearbyRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return nearby, err
}

func (u *nearbyUsecase) Updates(id int) (*domain.Nearby, error) {
	nearby, err := u.NearbyRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return nearby, err
}
