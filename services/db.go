package services

import (
	mods "github.com/Ilyasich/hackaton/tree/internal_dev/models"
	"github.com/Ilyasich/hackaton/tree/internal_dev/transport"
)

type AccountsRepository interface {
	AddUser(mods.User) bool
	UserExists(mods.TelegramID) bool
	ChangeUsersTonAcc(mods.TelegramID, mods.AccountID) bool
	TonAccExists(mods.AccountID) bool
	WalkByUsers(func(*mods.User))
	SetBanned(mods.TelegramID, bool) bool
}

type Service struct {
	rep  AccountsRepository
	rest transport.Rest
}

func New(rep AccountsRepository, rest transport.Rest) Service {
	return Service{
		rep:  rep,
		rest: rest,
	}
}

func (ser *Service) AddUser(tgacc mods.TelegramID, tonacc mods.AccountID) bool {
	if ser.rep.UserExists(tgacc) || ser.rep.TonAccExists(tonacc) {
		return false
	}
	_, ok := ser.rest.GetBalance(tonacc)
	if !ok {
		return false
	}
	user := mods.User{tgacc, tonacc, false}
	ser.SetBanned(&user)

	return ser.rep.AddUser(user)
}

func (ser *Service) ChangeUsersTonAcc(tgacc mods.TelegramID, newtonacc mods.AccountID) bool {
	if !ser.rep.UserExists(tgacc) {
		return false
	}
	if ser.rep.TonAccExists(newtonacc) {
		return false
	}
	return ser.rep.ChangeUsersTonAcc(tgacc, newtonacc)
}

func (ser *Service) WalkByUsers(operation func(*mods.User)) {
	ser.rep.WalkByUsers(operation)
}

func (ser *Service) SetBans() {
	ser.WalkByUsers(ser.SetBanned)
}
