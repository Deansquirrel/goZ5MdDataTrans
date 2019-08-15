package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/Deansquirrel/goZ5MdDataTrans/object"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlUpdateMdYyInfo = "" +
		"INSERT INTO [mdyyinfo]([mdid],[yyr],[tc],[sr],[oprtime]) " +
		"VALUES (?,?,?,?,?)"
	sqlGetMdYyInfoLastUpdate = "" +
		"select yyr " +
		"from mdyyinfolastupdate " +
		"where mdid = ?"
	sqlUpdateMdYyInfoLastUpdate = "" +
		"IF EXISTS (SELECT * FROM mdyyinfolastupdate WHERE MDID=?) " +
		"	BEGIN " +
		"		UPDATE mdyyinfolastupdate " +
		"		SET yyr = ?,oprtime=GETDATE() " +
		"		WHERE mdid = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO mdyyinfolastupdate(mdid,yyr,oprtime) " +
		"		VALUES (?,?,GETDATE()) " +
		"	END"
)

type repOnline struct {
	dbConfig *goToolMSSql.MSSqlConfig
}

func NewRepOnline() (*repOnline, error) {
	dbConfig, err := NewCommon().GetOnLineDbConfig()
	if err != nil {
		return nil, err
	}
	return &repOnline{
		dbConfig: dbConfig,
	}, nil
}

func (r *repOnline) GetMdYyInfoLastUpdate() (time.Time, error) {
	repMd := NewRepMd()
	mdId, err := repMd.GetMdId()
	if err != nil {
		return time.Now(), err
	}
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetMdYyInfoLastUpdate, mdId)
	if err != nil {
		errMsg := fmt.Sprintf("GetMdYyInfoLastUpdate err: %s", err.Error())
		log.Error(errMsg)
		return time.Now(), errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rTime := time.Now()
	flag := false
	for rows.Next() {
		err = rows.Scan(&rTime)
		if err != nil {
			errMsg := fmt.Sprintf("GetMdYyInfoLastUpdate read data err: %s", err.Error())
			log.Error(errMsg)
			return time.Now(), errors.New(errMsg)
		}
		flag = true
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetMdYyInfoLastUpdate read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return time.Now(), errors.New(errMsg)
	}
	if flag {
		return rTime, nil
	} else {
		return goToolMSSqlHelper.GetDefaultOprTime(), nil
	}
}

func (r *repOnline) UpdateMdYyInfoLastUpdate(t time.Time) error {
	repMd := NewRepMd()
	mdId, err := repMd.GetMdId()
	if err != nil {
		return err
	}
	err = goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdYyInfoLastUpdate,
		mdId,
		t, mdId,
		mdId, t)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateMdYyInfoLastUpdate err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repOnline) UpdateMdYyInfo(d *object.MdYyInfo) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdYyInfo,
		d.FMdId, d.FYyr, d.FTc, d.FSr, d.FOprTime)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateMdYyInfo err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repOnline) UpdateZxKc(d *object.ZxKc) error {
	//TODO
	return nil
}

func (r *repOnline) UpdateMdHpXsSlHz(d []*object.MdHpXsSlHz) error {
	//TODO
	return nil
}

func (r *repOnline) GetMdYyInfo() ([]*object.MdYyInfo, error) {
	//TODO
	return nil, nil
}

func (r *repOnline) GetZxKc() ([]*object.ZxKc, error) {
	//TODO
	return nil, nil
}

func (r *repOnline) GetMdHpXsSlHz() ([]*object.MdHpXsSlHz, error) {
	//TODO
	return nil, nil
}
