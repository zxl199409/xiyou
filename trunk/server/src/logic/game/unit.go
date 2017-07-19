package game

import (
	"sync/atomic"
	"logic/prpc"
)

var genInstId int64 = 1

type GameUnit struct {
	UnitId	    int32
	InstId      int64
	InstName 	string
	IProperties []int32
	CProperties []float32
}

func CreateUnitFromTable(id int32) *GameUnit {
	t := GetUnitRecordById(id)
	if t == nil {
		return nil
	}
	u := GameUnit{}
	u.UnitId = t.Id
	u.InstId = atomic.AddInt64(&genInstId, 1)
	u.IProperties = append(u.IProperties, t.IProp...)
	u.CProperties = append(u.CProperties, t.CProp...)
	//fmt.Println("CreateUnitFromTable, I", t.IProp, "it", it1, it2)
	return &u
}

func(this* GameUnit)GetUnitCOM()prpc.COM_Unit{
	u := prpc.COM_Unit{}
	u.UnitId = this.UnitId
	u.InstId = this.InstId
	u.IProperties = append(u.IProperties, this.IProperties...)
	u.CProperties = append(u.CProperties, this.CProperties...)
	return u
}
