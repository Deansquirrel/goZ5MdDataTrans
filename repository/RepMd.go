package repository

import (
	"github.com/Deansquirrel/goToolMSSql2000"
	"github.com/Deansquirrel/goToolMSSqlHelper"
)

import log "github.com/Deansquirrel/goToolLog"

type repMd struct {
	dbConfig *goToolMSSql2000.MSSqlConfig
}

func NewRepMd() *repMd {
	comm := NewCommon()
	return &repMd{
		dbConfig: goToolMSSqlHelper.ConvertDbConfigTo2000(comm.GetLocalDbConfig()),
	}
}

func (r *repMd) Test() {
	log.Debug("Test")
}
