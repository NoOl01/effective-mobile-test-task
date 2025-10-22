package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"testovoe/internal/errs"
	"testovoe/internal/handler/dto"
	"testovoe/internal/repository/queries"
	"time"
)

func (r *repository) CreateSubscription(ctx context.Context, sub dto.Subscription) error {
	var exists bool

	if err := r.pool.QueryRow(ctx, queries.QueryCheckExistSub, sub.UserId, sub.ServiceName).Scan(&exists); err != nil {
		log.Printf("failed to check record: %v\n", err)
		return err
	}

	if exists {
		log.Printf("subscription already exists for user_id: '%s', service_name: '%s'\n", sub.UserId, sub.ServiceName)
		return errs.ErrRecordAlreadyExist
	}

	if _, err := r.pool.Exec(ctx, queries.QueryCreateSubscription, sub.UserId, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate); err != nil {
		log.Printf("failed to create record: %v\n", err)
		return err
	}

	return nil
}

func (r *repository) SubscriptionsList(ctx context.Context) ([]dto.Subscription, error) {
	var result []dto.Subscription

	rows, err := r.pool.Query(ctx, queries.QueryListSubscriptions)
	if err != nil {
		log.Printf("failed to get subcsriptions: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub dto.Subscription
		err := rows.Scan(
			&sub.Id,
			&sub.UserId,
			&sub.ServiceName,
			&sub.Price,
			&sub.StartDate,
			&sub.EndDate,
		)
		if err != nil {
			log.Printf("error with scan: %v\n", err)
			return nil, err
		}
		result = append(result, sub)
	}

	return result, nil
}

func (r *repository) SubscriptionsSum(ctx context.Context, startDate, endDate time.Time, userId uuid.UUID, serviceName string) (int, error) {
	query := queries.QueryCalculateSubscriptionCost
	args := []interface{}{startDate, endDate}
	paramCount := 2

	if userId != uuid.Nil {
		paramCount++
		query += fmt.Sprintf(" AND user_id = $%d", paramCount)
		args = append(args, userId)
	}
	if serviceName != "" {
		paramCount++
		query += fmt.Sprintf(" AND service_name = $%d", paramCount)
		args = append(args, serviceName)
	}

	var result int
	if err := r.pool.QueryRow(ctx, query, args...).Scan(&result); err != nil {
		log.Printf("failed to calculate sum: %v", err)
		return 0, err
	}

	return result, nil
}

func (r *repository) SubscriptionByUserId(ctx context.Context, userId uuid.UUID) ([]dto.Subscription, error) {
	var result []dto.Subscription

	rows, err := r.pool.Query(ctx, queries.QueryGetSubscription, userId)
	if err != nil {
		log.Printf("failed to find subs by user id: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sub dto.Subscription
		if err := rows.Scan(
			&sub.Id,
			&sub.UserId,
			&sub.ServiceName,
			&sub.Price,
			&sub.StartDate,
			&sub.EndDate,
		); err != nil {
			fmt.Printf("error with scan: %v\n", err)
			return nil, err
		}
		result = append(result, sub)
	}

	if len(result) == 0 {
		return nil, errs.ErrRecordsNotFound
	}

	return result, nil
}

func (r *repository) UpdateSubscription(ctx context.Context, sub *dto.Subscription) error {
	if err := r.pool.QueryRow(ctx, queries.QueryUpdateSubscription,
		sub.Id,
		sub.UserId,
		sub.ServiceName,
		sub.Price,
		sub.StartDate,
		sub.EndDate,
	).Scan(
		&sub.Id,
		&sub.UserId,
		&sub.ServiceName,
		&sub.Price,
		&sub.StartDate,
		&sub.EndDate,
	); err != nil {
		log.Printf("failed to update sub: %v\n", err)
		return err
	}

	return nil
}

func (r *repository) DeleteSubscription(ctx context.Context, Id string) error {
	result, err := r.pool.Exec(ctx, queries.QueryDeleteSubscription, Id)
	if err != nil {
		log.Printf("failed to delete sub: %v\n", err)
		return err
	}

	if result.RowsAffected() == 0 {
		log.Println("record not found")
		return errors.New("record not found")
	}

	return nil
}
