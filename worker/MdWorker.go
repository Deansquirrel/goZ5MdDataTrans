package worker

import (
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

func (r *mdWorker) Test(id string) {
	log.Debug(fmt.Sprintf("Test %s", id))
}
