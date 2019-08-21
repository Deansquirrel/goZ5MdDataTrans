package global

import (
	"context"
	"github.com/Deansquirrel/goZ5MdDataTrans/object"
)

const (
	//goServiceSupportHelper Version "1.0.6 Build20190821"
	//PreVersion = "1.0.3 Build20190821"
	//TestVersion = "0.0.0 Build20190101"
	Version   = "0.0.0 Build20190101"
	Type      = "Z5MdDataTrans"
	SecretKey = "Z5MdDataTrans"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig
