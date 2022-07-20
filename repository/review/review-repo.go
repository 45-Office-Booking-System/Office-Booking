package review

import (
	domain "Office-Booking/domain/review"

	"gorm.io/gorm"
)

type reviewRepository struct {
	Conn *gorm.DB
}

// GetAll implements domain.ReviewRepository
func (u *reviewRepository) GetAll() (*domain.Reviews, error) {
	reviews := &domain.Reviews{}
	u.Conn.Find(&reviews)

	return reviews, nil
}

// GetByID implements domain.ReviewRepository
func (u *reviewRepository) GetByID(id int) (*domain.Review, error) {
	review := &domain.Review{ID: id}
	if err := u.Conn.First(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

// Create implements domain.ReviewRepository
func (u *reviewRepository) Create(review *domain.Review) (*domain.Review, error) {
	if err := u.Conn.Create(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

// Delete implements domain.ReviewRepository
func (u *reviewRepository) Delete(id int) (*domain.Review, error) {
	review := &domain.Review{ID: id}
	if err := u.Conn.Delete(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func NewReviewRepository(Conn *gorm.DB) domain.ReviewRepository {
	return &reviewRepository{Conn: Conn}
}
