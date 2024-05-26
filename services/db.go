package services

import (
	mods "github.com/Ilyasich/hackaton/models"
	"github.com/Ilyasich/hackaton/transport"
)

type AccountsRepository interface {
	AddUser(mods.User) bool
	UserExists(mods.TelegramID) bool
	ChangeUsersTonAcc(mods.TelegramID, mods.AccountID) bool
	TonAccExists(mods.AccountID) bool
	WalkByUsers(func(*mods.User))
	SetBanned(mods.TelegramID, bool) bool
	IsBanned(mods.TelegramID) bool
}

type ChatsRepository interface {
	AddChat(int64) bool
	ChangeChatCont(int64, mods.Context) bool
	DeleteChat(int64) bool
	ChatExists(int64) bool
}

type Service struct {
	rep     AccountsRepository
	rest    transport.Rest
	chatrep ChatsRepository
}

func New(rep AccountsRepository, rest transport.Rest, chatrep ChatsRepository) Service {
	return Service{
		rep:     rep,
		rest:    rest,
		chatrep: chatrep,
	}
}

func (ser *Service) AddUser(tgac mods.TelegramID, tonac mods.AccountID) bool {
	if ser.rep.UserExists(tgac) || ser.rep.TonAccExists(tonac) {
		return false
	}
	_, ok := ser.rest.GetBalance(tonac)
	if !ok {
		return false
	}
	user := mods.User{tgac, tonac, false}
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

func (ser *Service) IsBanned(tgacc mods.TelegramID) bool {
	return ser.rep.IsBanned(tgacc)
}

func (ser *Service) UserExists(tgacc mods.TelegramID) bool {
	return ser.rep.UserExists(tgacc)
}

func (ser *Service) AddChat(ID int64) bool {
	if ser.chatrep.ChatExists(ID) {
		return false
	}
	return ser.chatrep.AddChat(ID)
}

func (ser *Service) ChangeChatCont(ID int64, cont mods.Context) bool {
	if ser.chatrep.ChatExists(ID) {
		return false
	}
	return ser.chatrep.ChangeChatCont(ID, cont)
}

func (ser *Service) ChatExists(ID int64) bool {
	return ser.chatrep.ChatExists(ID)
}
