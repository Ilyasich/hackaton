package models

type AccountID string

type TelegramID int64

type User struct {
	Tgacc    TelegramID
	Tonacc   AccountID
	IsBanned bool
}
