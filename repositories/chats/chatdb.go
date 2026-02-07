package chats

import mods "github.com/Ilyasich/hackaton/models"

type ChatConxRep map[int64]mods.Context

func (rep *ChatConxRep) AddChat(Id int64) bool {
	if *rep == nil {
		*rep = make(map[int64]mods.Context)
	}
	(*rep)[Id] = mods.EMPTY
	return true
}

func (rep *ChatConxRep) ChangeChatCont(Id int64, cont mods.Context) bool {
	if *rep == nil {
		return false
	}
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

func (rep *ChatConxRep) GetChatCont(Id int64) (mods.Context, bool) {
	if *rep == nil {
		return mods.EMPTY, false
	}
	c, ok := (*rep)[Id]
	return c, ok
}
