package game

import (
	"time"
	"logic/conf"
	"logic/log"
)

const (
	timer		= 100			//计时器更新间隔
)

type (
	OncePlayer struct {
		PlayerInstId	int64
		TianTiVal		int32
		MatchingTime	float64
	}
	RobotTableData struct {
		RobotId			int32
		RobotScoreL		int32			//积分区间
		RobotScoreH		int32
		RobotBattleId	int32
		RobotIntegral	int32
	}
)

var (
	TianTiStore		[]*OncePlayer
	Robot	=map[int32]*RobotTableData{}
)

func LoadRobotTable(filename string) error {
	csv, err := conf.NewCSVFile(filename)
	if err != nil {
		return err
	}

	for r := 0; r < csv.Length(); r++ {
		c := RobotTableData{}

		c.RobotId 		= int32(csv.GetInt(r,"ID"))
		c.RobotScoreL	= csv.GetInt32(r,"ScoreL")
		c.RobotScoreH	= csv.GetInt32(r,"ScoreH")
		c.RobotIntegral	= csv.GetInt32(r,"Integral")
		c.RobotBattleId	= csv.GetInt32(r,"BalltID")

		Robot[c.RobotId] = &c
	}
	return nil
}

func GetRobotData(tiantiV int32) *RobotTableData {
	for _,r := range Robot{
		if r.RobotScoreL <= tiantiV && r.RobotScoreH >= tiantiV {
			return r
		}
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func InitTianTi()  {
	go func() {
		t1 := time.NewTimer(time.Millisecond * timer)

		for {
			select {
			case <-t1.C:
				//fmt.Println("50ms timer")
				Tick(timer)
				t1.Reset(time.Millisecond * timer)
			}
		}
	}()
}

func Tick(dt float64)  {
	var TempSecond	float64
	TempSecond = dt/1000
	for i:=0;i<len(TianTiStore) ;i++  {
		if TianTiStore[i]==nil {
			continue
		}
		CheckMatching(TianTiStore[i],TempSecond)
	}

	PlayerTick(TempSecond)
}

func CheckMatching(oncePlayer *OncePlayer, dt float64)  {
	oncePlayer.MatchingTime += dt
	for _,t := range TianTiStore{
		if oncePlayer.MatchingTime > 20{
			robot := GetRobotData(oncePlayer.TianTiVal)
			if robot==nil {
				log.Info("Can Not Find Robot PlayerId=",oncePlayer.PlayerInstId,"TiantiV=",oncePlayer.TianTiVal);
				continue
			}
			myself := FindPlayerByInstId(oncePlayer.PlayerInstId)
			RemoveMatching(oncePlayer.PlayerInstId)
			if CreatePvR(myself,robot.RobotBattleId) == nil {
				log.Info("Tianti CreatePvR Loser PlayerId=",oncePlayer.PlayerInstId,"RobotBattleId=",robot.RobotBattleId);
			}
			return
		}

		tempV := (int32(oncePlayer.MatchingTime/10) +1)*50

		if oncePlayer.TianTiVal >= (t.TianTiVal - tempV) && oncePlayer.TianTiVal <= (t.TianTiVal + tempV) {
			if t.PlayerInstId == oncePlayer.PlayerInstId {
				continue
			}
			//fmt.Println("Matching InstId=",oncePlayer.PlayerInstId,"MyTiantiVal",oncePlayer.TianTiVal,"tempV=",tempV,"[",(t.TianTiVal - tempV),(t.TianTiVal + tempV),"]","MatchingTime",
			//	oncePlayer.MatchingTime)
			myself := FindPlayerByInstId(oncePlayer.PlayerInstId)
			rival  := FindPlayerByInstId(t.PlayerInstId)
			RemoveMatching(oncePlayer.PlayerInstId)
			RemoveMatching(t.PlayerInstId)
			if CreatePvP(myself,rival) != nil {
				log.Info("Matching Succeed")
			}else {
				log.Info("Tianti CreatePvP Loser",oncePlayer.PlayerInstId,t.PlayerInstId)
			}

		}
	}
}

func StartMatching(player *GamePlayer,groupId int32)  {
	if player==nil {
		return
	}

	if player.GetUnitGroupById(groupId) == nil {
		log.Info("Can Not Find UnitGroup GroupId=",groupId)
		return
	}

	player.BattleUnitGroup = groupId

	tmp := OncePlayer{}
	tmp.PlayerInstId = player.MyUnit.InstId
	tmp.TianTiVal	 = player.TianTiVal
	TianTiStore = append(TianTiStore,&tmp)

	log.Info("StartMatching OK InstId=",tmp.PlayerInstId,"TianTiVal=",tmp.TianTiVal)
}

func StopMatching(player *GamePlayer)  {
	if player==nil {
		return
	}
	if RemoveMatching(player.MyUnit.InstId) {
		player.BattleUnitGroup = 0
	}
}

func RemoveMatching(instId int64) bool {
	for i:=0;i<len(TianTiStore) ;i++  {
		if instId == TianTiStore[i].PlayerInstId {
			TianTiStore = append(TianTiStore[:i], TianTiStore[i+1:]...)
			log.Info("RemoveMatching...",instId)
			return true
		}
	}
	return false
}

func CaleTianTiVal(player1 *GamePlayer,player2 *GamePlayer,winCamp int) int32 {
	if player1 == nil || player2 == nil {
		return 0
	}
	coef := int32((player1.TianTiVal - player2.TianTiVal)/5)

	if player1.BattleCamp == winCamp {
		player1.TianTiVal += (30-coef*2)
	}else {
		if player1.TianTiVal > 400 && player1.TianTiVal <= 1000 {
			player1.TianTiVal = player1.TianTiVal - (15-coef)
		}else if player1.TianTiVal > 1000 {
			player1.TianTiVal = player1.TianTiVal - (30-coef*2)
		}
	}
	if player1.session != nil {
		player1.session.UpdateTiantiVal(player1.TianTiVal)
	}

	tableId := GetTianTiIdByVal(player1.TianTiVal)
	ttData := GetTianTiTableDataById(tableId)
	if ttData == nil {
		log.Info("Can Not Find TianTiTableData By TableId=",tableId)
		return 0
	}

	var dropId int32 = 0;

	if player1.BattleCamp == winCamp {
		dropId = ttData.WinDrop
		log.Info("Tianti Battle Over CaleVal Winer Player[",player1.MyUnit.InstId,"]","TianTiVal[",player1.TianTiVal,"]","DropId=",ttData.WinDrop)
	}else {
		dropId = ttData.LoseDop
		log.Info("Tianti Battle Over CaleVal Loser Player[",player1.MyUnit.InstId,"]","TianTiVal[",player1.TianTiVal,"]","DropId=",ttData.LoseDop)
	}

	return dropId
}

func CaleTiantiPVR(player *GamePlayer,winCamp int) int32 {
	if player==nil {
		return 0
	}

	robot := GetRobotData(player.TianTiVal)

	coef := int32((player.TianTiVal - robot.RobotIntegral)/5)

	if player.BattleCamp == winCamp {
		player.TianTiVal += (30-coef*2)
	}else {
		if player.TianTiVal > 400 && player.TianTiVal <= 1000 {
			player.TianTiVal = player.TianTiVal - (15-coef)
		}else if player.TianTiVal > 1000 {
			player.TianTiVal = player.TianTiVal - (30-coef*2)
		}
	}

	if player.session != nil {
		player.session.UpdateTiantiVal(player.TianTiVal)
	}

	tableId := GetTianTiIdByVal(player.TianTiVal)
	ttData := GetTianTiTableDataById(tableId)
	if ttData == nil {
		log.Info("Can Not Find TianTiTableData By TableId=",tableId)
		return 0
	}

	var dropId int32 = 0;

	if player.BattleCamp == winCamp {
		dropId = ttData.WinDrop
		log.Info("Tianti PVR Battle Over CaleVal Winer Player[",player.MyUnit.InstId,"]","TianTiVal[",player.TianTiVal,"]","DropId=",ttData.WinDrop)
	}else {
		dropId = ttData.LoseDop
		log.Info("Tianti PVR Battle Over CaleVal Loser Player[",player.MyUnit.InstId,"]","TianTiVal[",player.TianTiVal,"]","DropId=",ttData.LoseDop)
	}
	return dropId
}