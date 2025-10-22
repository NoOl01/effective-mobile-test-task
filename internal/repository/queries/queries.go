package queries

// noinspection SqlResolve,SqlNoDataSourceInspection
const (
	QueryCreateSubscription = `
		INSERT INTO subs (user_id, service_name, price, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
	`

	QueryGetSubscription = `
		SELECT * FROM subs WHERE user_id = $1
	`

	QueryListSubscriptions = `
		SELECT * FROM subs
	`

	QueryUpdateSubscription = `
		UPDATE subs
		SET user_id = $2, service_name = $3, price = $4, start_date = $5, end_date = $6
		WHERE id = $1
		RETURNING id, user_id, service_name, price, start_date, end_date
	`

	QueryDeleteSubscription = `
		DELETE FROM subs
		WHERE id = $1
	`

	QueryCalculateSubscriptionCost = `
		SELECT COALESCE(SUM(price), 0) as total
		FROM subs
		WHERE start_date <= $2 AND (end_date IS NULL OR end_date >= $1)
	`

	QueryCheckExistSub = `
		SELECT EXISTS(SELECT 1 FROM subs WHERE user_id = $1 AND service_name = $2)
	`
)
