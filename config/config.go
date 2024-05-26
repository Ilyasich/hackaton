package config

import (
	"os"
	"strconv"
)

type RestConfig struct {
	TonHost string
}

type TgConfig struct {
	InvLink  string
	BotToken string
}

type WorkerConfig struct {
	delay int
}

type Config struct {
	R  RestConfig
	Tg TgConfig
	Wc WorkerConfig
}

func Init() Config {
	var tg TgConfig
	var r RestConfig
	var w WorkerConfig
	var c Config
	initStringWithDefVal("invlink", "", &tg.InvLink)
	initStringWithDefVal("token", "", &tg.BotToken)
	initStringWithDefVal("tonaddr", "https://tonapi.io/v2/accounts/", &r.TonHost)
	initIntWithDefVal("checktime", 120, &w.delay)
	c.R = r
	c.Tg = tg
	c.Wc = w
	return c
}

func initStringWithDefVal(envname string, defval string, variable *string) {
	if _, ok := os.LookupEnv(envname); !ok {
		*variable = defval
		return
	}
	*variable = os.Getenv(envname)
}

func initIntWithDefVal(envname string, defval int, variable *int) {
	if _, ok := os.LookupEnv(envname); !ok {
		*variable = defval
		return
	}
	*variable, _ = strconv.Atoi(os.Getenv(envname))
}
