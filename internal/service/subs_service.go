package service

import (
	"context"
	"strconv"
	"testovoe/internal/handler/dto"
	"testovoe/internal/utils"
	"time"
)

func (s *service) CreateSubscription(sub dto.CreateSubscription) error {
	ctx := context.Background()
	now := time.Now()

	newSub := dto.Subscription{
		UserId:      sub.UserId,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		StartDate:   now,
		EndDate:     now.Add(30 * (24 * time.Hour)),
	}

	return s.repo.CreateSubscription(ctx, newSub)
}

func (s *service) SubscriptionsList() ([]dto.Subscription, error) {
	ctx := context.Background()

	return s.repo.SubscriptionsList(ctx)
}

func (s *service) SubscriptionsSum(startDate, endDate, userId, serviceName string) (int, error) {
	ctx := context.Background()

	start, err := utils.FromStringToTime(startDate)
	if err != nil {
		return 0, err
	}
	end, err := utils.FromStringToTime(endDate)
	if err != nil {
		return 0, err
	}

	return s.repo.SubscriptionsSum(ctx, start, end, userId, serviceName)
}

func (s *service) SubscriptionByUserId(userId string) ([]dto.Subscription, error) {
	ctx := context.Background()

	return s.repo.SubscriptionByUserId(ctx, userId)
}

func (s *service) UpdateSubscription(idStr string, sub dto.UpdateSubscription) (*dto.Subscription, error) {
	ctx := context.Background()

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	start, err := utils.FromStringToTime(sub.StartDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.FromStringToTime(sub.EndDate)
	if err != nil {
		return nil, err
	}

	newSub := dto.Subscription{
		Id:          id,
		UserId:      sub.UserId,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		StartDate:   start,
		EndDate:     end,
	}

	if err := s.repo.UpdateSubscription(ctx, &newSub); err != nil {
		return nil, err
	}

	return &newSub, nil
}

func (s *service) DeleteSubscription(id string) error {
	ctx := context.Background()

	return s.repo.DeleteSubscription(ctx, id)
}
