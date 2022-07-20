package usecase

import (
	domain "Office-Booking/domain/jenisgedung"
	"Office-Booking/domain/jenisgedung/request"
	"errors"
)

type jenisgedungUsecase struct {
	JenisgedungRepo domain.JenisgedungRepository
}

func NewJenisgedungUsecase(ur domain.JenisgedungRepository) domain.JenisgedungUsecase {
	return &jenisgedungUsecase{JenisgedungRepo: ur}
}

func (u *jenisgedungUsecase) Create(request request.JenisgedungCreateRequest) (*domain.Jenisgedung, error) {
	if request.Jenis == "" {
		return nil, errors.New("jenis gedung belum diisi")
	}
	if request.IDGedung == 0 {
		return nil, errors.New("id-gedung empty")
	}
	jenisgedung := &domain.Jenisgedung{
		Jenis:    request.Jenis,
		IDGedung: request.IDGedung,
	}

	createdJenisgedung, err := u.JenisgedungRepo.Create(jenisgedung)
	if err != nil {
		return nil, err
	}

	return createdJenisgedung, nil
}

func (u *jenisgedungUsecase) ReadByID(id int) (*domain.Jenisgedung, error) {
	jenisgedung, err := u.JenisgedungRepo.ReadByID(id)
	if err != nil {
		return nil, err
	}

	return jenisgedung, err
}

func (u *jenisgedungUsecase) ReadAll() (*domain.Jenisgedungs, error) {
	foundJenisgedungs, err := u.JenisgedungRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return foundJenisgedungs, nil
}

func (u *jenisgedungUsecase) Delete(id int) (*domain.Jenisgedung, error) {
	jenisgedung, err := u.JenisgedungRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return jenisgedung, err
}

func (u *jenisgedungUsecase) Updates(id int) (*domain.Jenisgedung, error) {
	jenisgedung, err := u.JenisgedungRepo.Updates(id)
	if err != nil {
		return nil, err
	}

	return jenisgedung, err
}
