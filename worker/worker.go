package worker

import (
	"github.com/Ilyasich/hackaton/tree/internal_dev/services"
	"time"
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
