package storage

import "time"

type UserPaymentHistoryAction string

const (
	UserPaymentHistoryActionPay  UserPaymentHistoryAction = "pay"
	UserPaymentHistoryActionSkip UserPaymentHistoryAction = "skip"
)

type User struct {
	Id        int
	Username  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
