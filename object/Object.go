package object

import "time"

type MdYyInfo struct {
	FMdId    int
	FYyr     time.Time
	FTc      int
	FSr      float64
	FOprTime time.Time
}

type ZxKc struct {
	FMdId    int
	FHpId    int
	FSl      float64
	FOprTime time.Time
}

type MdHpXsSlHz struct {
	FYyDate  time.Time
	FMdId    int
	FHpId    int
	FXsQty   float64
	FOprTime time.Time
}
