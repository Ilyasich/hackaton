package transport

import (
	"encoding/json"

	"github.com/Ilyasich/hackaton/config"
	"github.com/Ilyasich/hackaton/models"

	"github.com/go-resty/resty/v2"
)

type Rest struct {
	cl   *resty.Client
	conf config.Config
}

// New returns a Rest client (value) that wraps a pointer to resty.Client.
func New(con config.Config) Rest {
	return Rest{
		cl:   resty.New(),
		conf: con,
	}
}

func (r Rest) GetBalance(tonId models.AccountID) (float64, bool) {
	url := r.conf.R.TonHost + string(tonId)
	resp, err := r.cl.R().Get(url)
	if err != nil {
		return 0, false
	}
	response := make(map[string]interface{})
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return 0, false
	}
	// balance may come as number or string depending on API; handle both safely
	if b, ok := response["balance"].(float64); ok {
		return b, true
	}
	if bs, ok := response["balance"].(string); ok {
		// try to parse string as float
		var val float64
		if err := json.Unmarshal([]byte(bs), &val); err == nil {
			return val, true
		}
	}
	return 0, false
}
