package chats

import mods "github.com/Ilyasich/hackaton/models"

type ChatConxRep map[int64]mods.Context

func (rep *ChatConxRep) AddChat(Id int64) bool {
	(*rep)[Id] = mods.EMPTY
	return true
}

func (rep *ChatConxRep) ChangeChatCont(Id int64, cont mods.Context) bool {
	(*rep)[Id] = cont
	return true
}

func (rep *ChatConxRep) DeleteChat(Id int64) bool {
	delete(*rep, Id)
	return true
}

func (rep *ChatConxRep) ChatExists(Id int64) bool {
	_, ok := (*rep)[Id]
	return ok
}
