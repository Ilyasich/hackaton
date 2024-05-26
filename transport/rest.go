package transport

import (
	"hackaton/config"
	"hackaton/models"
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Rest struct {
	cl   resty.Client
	conf config.Config
}

func New(con config.Config) *Rest {
	return &Rest{
		*resty.New(),
		con,
	}
}

func (r *Rest) GetBalance(tonId models.AccountID) (float64, bool) {
	url := (*r).conf.R.TonHost + string(tonId)
	acc, err := r.cl.R().Get(url)
	if err != nil {
		return 0, false
	}
	response := make(map[string]interface{})
	json.Unmarshal(acc.Body(), &response)
	return response["balance"].(float64), true
}
