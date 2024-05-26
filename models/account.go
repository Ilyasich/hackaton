package models

type AccountID string

type TelegramID int32

type User struct {
	Tgacc    TelegramID
	Tonacc   AccountID
	IsBanned bool
}
