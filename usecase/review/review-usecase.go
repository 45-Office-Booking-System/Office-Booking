package review

import (
	domain "Office-Booking/domain/review"
	"Office-Booking/domain/review/request"
	"errors"
)

type reviewUsecase struct {
	ReviewRepo domain.ReviewRepository
}

// GetAll implements domain.ReviewUsecase
func (u *reviewUsecase) GetAll() (*domain.Reviews, error) {
	foundreview, err := u.ReviewRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundreview, nil
}

// GetByID implements domain.ReviewUsecase
func (u *reviewUsecase) GetByID(id int) (*domain.Review, error) {
	review, err := u.ReviewRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return review, err
}

// Create implements domain.ReviewRepository
func (u *reviewUsecase) Create(request request.ReviewPost) (*domain.Review, error) {
	if request.Rating == 0.0 {
		return nil, errors.New("rating belum isi")
	}
	if request.Description == "" {
		return nil, errors.New("description belum isi")
	}
	if request.IDGedung == 0 {
		return nil, errors.New("id_gedung empty")
	}
	review := &domain.Review{
		Rating:      request.Rating,
		Description: request.Description,
		IDGedung:    request.IDGedung,
		Img:         request.Img,
	}

	postReview, err := u.ReviewRepo.Create(review)
	if err != nil {
		return nil, err
	}

	return postReview, nil
}

// Delete implements domain.ReviewRepository
func (u *reviewUsecase) Delete(id int) (*domain.Review, error) {
	review, err := u.ReviewRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return review, err
}

func NewReviewUseCase(ur domain.ReviewRepository) domain.ReviewUsecase {
	return &reviewUsecase{ReviewRepo: ur}
}
