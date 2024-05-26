package memory

import mods "github.com/Ilyasich/hackaton/models"

type Repository []mods.User

func (rep *Repository) AddUser(user mods.User) bool {
	*rep = append(*rep, user)
	return true
}

func (rep *Repository) UserExists(tgacc mods.TelegramID) bool {
	for _, el := range *rep {
		if el.Tgacc == tgacc {
			return true
		}
	}
	return false
}

func (rep *Repository) TonAccExists(tonacc mods.AccountID) bool {
	for _, el := range *rep {
		if el.Tonacc == tonacc {
			return true
		}
	}
	return false
}

func (rep *Repository) ChangeUsersTonAcc(tgacc mods.TelegramID, newtonacc mods.AccountID) bool {
	for i, el := range *rep {
		if el.Tgacc == tgacc {
			(*rep)[i].Tonacc = newtonacc
			return true
		}
	}
	return false
}

func (rep *Repository) WalkByUsers(operation func(*mods.User)) {
	for i := range *rep {
		operation(&(*rep)[i])
	}
}

func (rep *Repository) SetBanned(tgacc mods.TelegramID, state bool) bool {
	for i, el := range *rep {
		if el.Tgacc == tgacc {
			(*rep)[i].IsBanned = state
			return true
		}
	}
	return false
}

func (rep *Repository) IsBanned(tgacc mods.TelegramID) bool {
	for _, el := range *rep {
		if el.Tgacc == tgacc {
			return el.IsBanned
		}
	}
	return false
}
