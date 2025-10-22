package service

import (
	"testovoe/internal/handler/dto"
	"testovoe/internal/repository"
)

type Service interface {
	CreateSubscription(sub dto.CreateSubscription) error
	SubscriptionsList() ([]dto.Subscription, error)
	SubscriptionsSum(startDate, endDate, userId, serviceName string) (int, error)
	SubscriptionByUserId(userId string) ([]dto.Subscription, error)
	UpdateSubscription(idStr string, sub dto.UpdateSubscription) (*dto.Subscription, error)
	DeleteSubscription(id string) error
}

type service struct {
	repo repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{repo: repository}
}
