package worker

import (
	"time"

	"github.com/Ilyasich/hackaton/services"
)

type Worker struct {
	checktime time.Duration
	service   services.Service
}

func New(service services.Service, checktime time.Duration) *Worker {
	return &Worker{
		service:   service,
		checktime: checktime,
	}
}

func (w *Worker) RunCheck() {
	for {
		select {
		case <-time.After(w.checktime):
		}
		w.service.SetBans()
	}
}
