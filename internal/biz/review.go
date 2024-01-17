package biz

import (
	"context"

	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
}

// GreeterUsecase is a Greeter usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (r *ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	r.log.Debugf("biz CreateReview %#v \n", review)
	// 1. validate
	// 2. generate review id
	// 3. query order and good info
	// 4. collect data and save to database
	return r.repo.SaveReview(ctx, review)
}
