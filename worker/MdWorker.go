package worker

import (
	"github.com/Deansquirrel/goServiceSupportHelper"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ5MdDataTrans/repository"
	"time"
)

var mdHpXsSlHzDate time.Time

var zxKc map[int]float64
var xsSl map[int]float64

func init() {
	mdHpXsSlHzDate = goToolMSSqlHelper.GetDefaultOprTime()
	zxKc = make(map[int]float64)
	xsSl = make(map[int]float64)
}

type mdWorker struct {
}

func NewMdWorker() *mdWorker {
	return &mdWorker{}
}

func (r *mdWorker) UpdateMdYyInfo(id string) {
	repOnline, err := repository.NewRepOnline()
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	lastUpdate, err := repOnline.GetMdYyInfoLastUpdate()
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	repMd := repository.NewRepMd()
	tClose, err := repMd.GetLastMdYyDate()
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	if goToolCommon.GetDateStr(lastUpdate) < "2000-01-01" {
		lastUpdate = tClose
	}
	endDate := tClose.Add(time.Hour * 24)
	list, err := repMd.GetMdYyInfo(goToolCommon.GetDateStr(lastUpdate), goToolCommon.GetDateStr(endDate))
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	for _, d := range list {
		err = repOnline.UpdateMdYyInfo(d)
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
			return
		}
	}
	err = repOnline.UpdateMdYyInfoLastUpdate(tClose)
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
}

func (r *mdWorker) UpdateZxKc(id string) {
	//TODO
}

func (r *mdWorker) UpdateMdHpXsSlHz(id string) {
	//TODO
}
