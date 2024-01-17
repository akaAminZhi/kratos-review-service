package data

import (
	"context"

	"review-service/internal/biz"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.WithContext(ctx).ReviewInfo.Save(review)
	return review, err
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, orderId int64) (*model.ReviewInfo, error) {
	return r.data.query.WithContext(ctx).ReviewInfo.Where(r.data.query.ReviewInfo.OrderID).First()
}
