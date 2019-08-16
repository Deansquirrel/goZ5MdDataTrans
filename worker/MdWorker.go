package worker

import (
	"github.com/Deansquirrel/goServiceSupportHelper"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ5MdDataTrans/object"
	"github.com/Deansquirrel/goZ5MdDataTrans/repository"
	"math"
	"time"
)

var mdHpXsSlHzDate time.Time

var zxKc map[int]float64
var xsSl map[int]float64

const (
	minKcDifference = 0.000001
)

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
	endDate := tClose.Add(time.Hour * 24)
	list, err := repMd.GetMdYyInfo(goToolCommon.GetDateStr(lastUpdate), goToolCommon.GetDateStr(endDate))
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	lastYyr := goToolMSSqlHelper.GetDefaultOprTime()
	for _, d := range list {
		err = repOnline.UpdateMdYyInfo(d)
		if err != nil {
			_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
			return
		}
		if goToolCommon.GetDateStr(d.FYyr) > goToolCommon.GetDateStr(lastYyr) {
			lastYyr = d.FYyr
		}
	}
	err = repOnline.UpdateMdYyInfoLastUpdate(lastYyr)
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
}

func (r *mdWorker) UpdateZxKc(id string) {
	repMd := repository.NewRepMd()
	kcList, err := repMd.GetZxKc()
	if err != nil {
		_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
		return
	}
	for _, kc := range kcList {
		zkc, ok := zxKc[kc.FHpId]
		if ok {
			if math.Dim(math.Max(zkc, kc.FSl), math.Min(zkc, kc.FSl)) >= minKcDifference {
				err = r.updateZxKc(kc)
				if err != nil {
					_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
					return
				}
			}
		} else {
			err = r.updateZxKc(kc)
			if err != nil {
				_ = goServiceSupportHelper.JobErrRecord(id, err.Error())
				return
			}
		}
	}
}

func (r *mdWorker) updateZxKc(d *object.ZxKc) error {
	repOnline, err := repository.NewRepOnline()
	if err != nil {
		return err
	}
	err = repOnline.UpdateZxKc(d)
	if err != nil {
		return err
	} else {
		zxKc[d.FHpId] = d.FSl
	}
	return nil
}

func (r *mdWorker) UpdateMdHpXsSlHz(id string) {
	//TODO
}
