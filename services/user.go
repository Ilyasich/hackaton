package services

import (
	"github.com/Ilyasich/hackaton/models"
)

func (ser *Service) SetBanned(user *models.User) {
	// If user has balance (>0) they should NOT be banned. Otherwise they are banned.
	if bal, ok := ser.rest.GetBalance(user.Tonacc); ok && bal > 0 {
		ser.rep.SetBanned(user.Tgacc, false)
		return
	}
	ser.rep.SetBanned(user.Tgacc, true)
}
