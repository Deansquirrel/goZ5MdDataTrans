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
		"		SET yyr = ?,oprtime= ? " +
		"		WHERE mdid = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO mdyyinfolastupdate(mdid,yyr,oprtime) " +
		"		VALUES (?,?,?) " +
		"	END"
	sqlUpdateZxKc = "" +
		"INSERT INTO [zxkc]([mdid],[hpid],[sl],[oprtime]) " +
		"VALUES (?,?,?,?)"
	sqlUpdateMdHpXsSlHz = "" +
		"INSERT INTO [mdhpxsslhz]([yydate],[mdid],[hpid],[xsqty],[jlsj]) " +
		"VALUES (?,?,?,?,?)"
	sqlGetMdHpXsSlHzLastUpdate = "" +
		"select yyr " +
		"from mdhpxsslhzlastupdate " +
		"where mdid = ?"
	sqlUpdateMdHpXsSlHzLastUpdate = "" +
		"IF EXISTS (SELECT * FROM mdhpxsslhzlastupdate WHERE MDID=?) " +
		"	BEGIN " +
		"		UPDATE mdhpxsslhzlastupdate " +
		"		SET yyr = ?,oprtime= ? " +
		"		WHERE mdid = ? " +
		"	END " +
		"ELSE " +
		"	BEGIN " +
		"		INSERT INTO mdhpxsslhzlastupdate(mdid,yyr,oprtime) " +
		"		VALUES (?,?,?) " +
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
		return goToolMSSqlHelper.GetDefaultOprTime(), err
	}
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetMdYyInfoLastUpdate, mdId)
	if err != nil {
		errMsg := fmt.Sprintf("GetMdYyInfoLastUpdate err: %s", err.Error())
		log.Error(errMsg)
		return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
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
			return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
		}
		flag = true
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetMdYyInfoLastUpdate read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
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
		t, time.Now(), mdId,
		mdId, t, time.Now())
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
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateZxKc,
		d.FMdId, d.FHpId, d.FSl, d.FOprTime)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateZxKc err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repOnline) GetMdHpXsSlHzLastUpdate() (time.Time, error) {
	repMd := NewRepMd()
	mdId, err := repMd.GetMdId()
	if err != nil {
		return goToolMSSqlHelper.GetDefaultOprTime(), err
	}
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetMdHpXsSlHzLastUpdate, mdId)
	if err != nil {
		errMsg := fmt.Sprintf("GetMdHpXsSlHzLastUpdate err: %s", err.Error())
		log.Error(errMsg)
		return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rTime := time.Now()
	flag := false
	for rows.Next() {
		err = rows.Scan(&rTime)
		if err != nil {
			errMsg := fmt.Sprintf("GetMdHpXsSlHzLastUpdate read data err: %s", err.Error())
			log.Error(errMsg)
			return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
		}
		flag = true
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetMdHpXsSlHzLastUpdate read data err: %s", rows.Err().Error())
		log.Error(errMsg)
		return goToolMSSqlHelper.GetDefaultOprTime(), errors.New(errMsg)
	}
	if flag {
		return rTime, nil
	} else {
		return goToolMSSqlHelper.GetDefaultOprTime(), nil
	}
}

func (r *repOnline) UpdateMdHpXsSlHzLastUpdate(t time.Time) error {
	repMd := NewRepMd()
	mdId, err := repMd.GetMdId()
	if err != nil {
		return err
	}
	err = goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdHpXsSlHzLastUpdate,
		mdId,
		t, time.Now(), mdId,
		mdId, t, time.Now())
	if err != nil {
		errMsg := fmt.Sprintf("UpdateMdHpXsSlHzLastUpdate err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repOnline) UpdateMdHpXsSlHz(d *object.MdHpXsSlHz) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateMdHpXsSlHz,
		d.FYyDate, d.FMdId, d.FHpId, d.FXsQty, d.FOprTime)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateMdHpXsSlHz err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
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
