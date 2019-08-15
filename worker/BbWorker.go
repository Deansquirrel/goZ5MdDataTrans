package worker

import (
	"fmt"
	log "github.com/Deansquirrel/goToolLog"
)

type bbWorker struct {
}

func NewBbWorker() *bbWorker {
	return &bbWorker{}
}

func (r *bbWorker) Test(id string) {
	log.Debug(fmt.Sprintf("Test %s", id))
}
