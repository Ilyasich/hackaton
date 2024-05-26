package services

import (
	"hackaton/models"
)

func (ser *Service) SetBanned(user *models.User) {
	if bal, _ := ser.rest.GetBalance(user.Tonacc); bal > 0 {
		ser.rep.SetBanned(user.Tgacc, true)
		return
	} else {
		ser.rep.SetBanned(user.Tgacc, false)
	}
}
