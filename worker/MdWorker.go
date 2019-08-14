package worker

import log "github.com/Deansquirrel/goToolLog"

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

func (r *mdWorker) Test() {
	log.Debug("Test")
}
