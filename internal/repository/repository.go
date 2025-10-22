package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"testovoe/internal/handler/dto"
	"time"
)

type Repository interface {
	CreateSubscription(ctx context.Context, sub dto.Subscription) error
	SubscriptionsList(ctx context.Context) ([]dto.Subscription, error)
	SubscriptionsSum(ctx context.Context, startDate, endDate time.Time, userId uuid.UUID, serviceName string) (int, error)
	SubscriptionByUserId(ctx context.Context, userId uuid.UUID) ([]dto.Subscription, error)
	UpdateSubscription(ctx context.Context, sub *dto.Subscription) error
	DeleteSubscription(ctx context.Context, id string) error
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &repository{
		pool: pool,
	}
}
