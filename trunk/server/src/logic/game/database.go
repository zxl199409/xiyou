package game

import (
	"bytes"
	"database/sql"
	"jimny/logs"

	_ "github.com/go-sql-driver/mysql"

	"logic/prpc"
	_ "jimny/sqlite3"

	"sync/atomic"
	"encoding/json"
	"fmt"
)

var (
	MaxPlayerInstId 		int64 	=  1
	MaxUnitInstId 			int64 	= 1
	MaxGuildId	  			int32 	= 1
	MaxGuildAssistantId	  	int32 	= 1
	MaxBattleID	  			int64 	= 1
)

func GenPlayerInstId() int64{
	return atomic.AddInt64(&MaxPlayerInstId,1)
}

func GenUnitInstId() int64 {
	return atomic.AddInt64(&MaxUnitInstId,1)
}

func GenGuildInstId() int32 {
	return atomic.AddInt32(&MaxGuildId,1)
}

func GenGuildAssistantInstId() int32 {
	return atomic.AddInt32(&MaxGuildAssistantId,1)
}

func GenBattleId() int64 {
	return atomic.AddInt64(&MaxBattleID,1)
}

func InitDB() {
	c, e := ConnectDB()
	if e != nil {
		logs.Debug(e.Error())
		return
	}
	defer c.Close()
	r, e := c.Query("SELECT MAX(`PlayerId`) AS MaxID FROM `Player`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxPlayerInstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}

	r, e = c.Query("SELECT MAX(`UnitId`) AS MaxID FROM `Unit`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxUnitInstId)
		if e != nil {
			logs.Debug(e.Error())
			//return
		}
	}

	r, e = c.Query("SELECT MAX(`GuildId`) AS MaxID FROM `Guild`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxGuildId)
		if e != nil {
			logs.Debug(e.Error())
			//return
		}
	}

	r, e = c.Query("SELECT MAX(`AssistantId`) AS MaxID FROM `GuildAssistant`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxGuildAssistantId)
		if e != nil {
			logs.Debug(e.Error())
			//return
		}
	}

	r, e = c.Query("SELECT MAX(`BattleID`) AS MaxID FROM `BattleReport`")

	if e != nil {
		logs.Debug(e.Error())
		return
	}


	if r.Next() {
		e = r.Scan(&MaxBattleID)
		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}

	logs.Infof("MAX PLAYER ID %d MAX UNIT ID %d, MAX GUILD ID %d, MAX GUILD ASSISTAN ID %d, ",MaxPlayerInstId,MaxUnitInstId,MaxGuildId,MaxGuildAssistantId)
}



func ConnectDB() (*sql.DB, error) {
	//dsn := beego.AppConfig.String("dbuser") + ":" + beego.AppConfig.String("dbpass") + "@tcp(" + beego.AppConfig.String("dbhost") + ":" + beego.AppConfig.String("dbport") + ")/" + beego.AppConfig.String("dbname")
	dsn := GetEnvString("V_MySqlData")
	return sql.Open("mysql", dsn)
}

func QueryPlayer(username string) <- chan *prpc.SGE_DBPlayer {
	logs.Debug("Query player")
	rChan := make(chan *prpc.SGE_DBPlayer)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("QueryPlayer panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `Player` WHERE `Username` = ?", username)

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		if r.Next() {
			a := int64(0)
			b := []byte{}
			c := ""
			d := int64(0)
			e = r.Scan(&a, &c, &b, &d)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := &prpc.SGE_DBPlayer{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p.PlayerId = a


			p.COM_Player.Employees = <- QueryUnit(a)

			rChan <- p

			close(rChan)
			return
		}

		rChan <- nil
		close(rChan)
		return
	}()

	return rChan
}


func QueryPlayerById(InstId int64) <- chan *prpc.SGE_DBPlayer {
	logs.Debug("Query player")
	rChan := make(chan *prpc.SGE_DBPlayer)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("QueryPlayerById panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `Player` WHERE `InstId` = ?", InstId)

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		if r.Next() {
			a := int64(0)
			b := []byte{}
			c := ""
			d := int64(0)
			e = r.Scan(&a, &c, &b, &d)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := &prpc.SGE_DBPlayer{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p.PlayerId = a


			p.COM_Player.Employees = <- QueryUnit(a)

			rChan <- p

			close(rChan)
			return
		}

		rChan <- nil
		close(rChan)
		return
	}()

	return rChan
}

func InsertPlayer(p prpc.SGE_DBPlayer) <- chan int64 {

	rChan := make (chan int64)

	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertPlayer panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b :=  bytes.NewBuffer(nil)

		p.Serialize(b)

		r , e := c.Exec("INSERT INTO `Player`(`PlayerId`, `Username`, `InstId`, `BinData`)VALUES(?,?,?,?)", p.PlayerId , p.Username, p.COM_Player.InstId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <-  (i + 1)
		close(rChan)
	}()
	return rChan
}

func QueryUnit(ownerId int64) <- chan []prpc.COM_Unit {

	rChan := make(chan []prpc.COM_Unit)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("QueryUnit panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `Unit` WHERE `OwnerId` = ?",ownerId )

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		arr := []prpc.COM_Unit{}

		for r.Next() {
			a := int64(0)
			b := []byte{}
			c := int64(0)
			e = r.Scan(&a, &c, &b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := prpc.COM_Unit{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p.InstId = a


			arr = append(arr, p)

		}

		rChan <- arr
		close(rChan)

	}()

	return rChan
}

func InsertUnit(ownerId int64, p prpc.COM_Unit) <- chan int64 {
	rChan := make (chan int64)
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertUnit panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b := bytes.NewBuffer(nil)

		p.Serialize(b)

		r , e := c.Exec("INSERT INTO `Unit`(`UnitId`, `OwnerId`, `BinData`)VALUES(?,?,?)", p.InstId, ownerId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <- (i + 1)
		close(rChan)
	}()
	return  rChan
}

func UpdatePlayer(p prpc.SGE_DBPlayer) {

	//logs.Debug(p.UnitGroup)
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdatePlayer panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		for _, u := range p.Employees{
			UpdateUnit(u)
		}

		p.Employees = nil

		e = p.Serialize(&b)

		if e != nil {
			logs.Debug(e.Error())
			return
		}

		logs.Debug("GamePlayerSave", p.Friends)
		_, e = c.Exec("UPDATE `Player` SET `BinData` = ? WHERE `PlayerId` = ?", b.Bytes(), p.PlayerId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func UpdateUnit(p prpc.COM_Unit) {
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateUnit panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		p.Serialize(&b)

		_, e = c.Exec("UPDATE `Unit` SET `BinData` = ? WHERE `UnitId` = ?", b.Bytes(), p.InstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func QueryAllTopList()  <- chan []prpc.COM_TopUnit {		//取出来整张表的数据
	rChan := make(chan []prpc.COM_TopUnit)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("QueryAllTopList panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `TopList`" )

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		arr := []prpc.COM_TopUnit{}

		for r.Next() {
			a := int64(0)
			b := []byte{}
			e = r.Scan(&a, &b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				logs.Debug("e.Error() 1 ", e.Error())
				continue
			}

			p := prpc.COM_TopUnit{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				logs.Debug("e.Error() ", e.Error(), "  ", a)
				continue
			}

			arr = append(arr, p)

		}

		rChan <- arr
		close(rChan)

	}()

	return rChan
}

func UpdateTopList(InstId int64, t prpc.SGE_DBTopUnit) {
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateTopList panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		t.Serialize(&b)

		_, e = c.Exec("UPDATE `TopList` SET `BinData` = ? WHERE `InstId` = ?", b.Bytes(), InstId)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}


func InsertTopList (InstId int64, t prpc.SGE_DBTopUnit) <- chan int64 {

	rChan := make (chan int64)

	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertTopList panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b :=  bytes.NewBuffer(nil)

		t.Serialize(b)

		r , e := c.Exec("INSERT INTO `TopList`(`InstId`, `BinData`)VALUES(?,?)", InstId, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <-  (i + 1)
		close(rChan)
	}()
	return rChan
}

//帮派
func InsertGuild(pGuild prpc.COM_Guild,member prpc.COM_GuildMember) <- chan bool {
	rChan := make (chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertGuild panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			return
		}
		defer c.Close()

		stmt1, e := c.Prepare("INSERT INTO `Guild`(`GuildId`, `GuildName`,`Master`,`MasterName`,`GuildVal`,`CreatTime`,`RequestList`,`RequestFlag`,`Require`,`Contribution`)" +
			"VALUES(?,?,?,?,?,?,?,?,?,?)")
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			return
		}
		stmt2, e := c.Prepare("INSERT INTO `GuildMember`(`GuildId`, `RoleId`,`RoleName`,`Rolelevel`,`Job`,`TianTiVal`,`UnitId`)VALUES(?,?,?,?,?,?,?)")
		if  e != nil {
			logs.Debug(e.Error())
			rChan <- false
			return
		}

		buffs, erro  :=  json.Marshal(pGuild.RequestList)

		if erro != nil {
			close(rChan)
			return
		}

		var isFlag int = 0

		if pGuild.IsRatify {
			isFlag = 1
		}
		
		_, e = stmt1.Exec( pGuild.GuildId, pGuild.GuildName, pGuild.Master, pGuild.MasterName, pGuild.GuildVal, pGuild.CreateTime, buffs,isFlag,pGuild.Require,pGuild.Contribution )
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			return
		}

		stmt1.Close()

		_, e = stmt2.Exec(member.GuildId,member.RoleId,member.RoleName,member.Level,member.Job,member.TianTiVal,member.UnitId)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			return
		}

		stmt2.Close()

		rChan <-  true
		close(rChan)
	}()
	return rChan
}

func DeleteDBGuild(guildId int32) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("DeleteDBGuild panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("DELETE FROM `Guild` WHERE `GuildId`= ?  ")
		stmt.Exec(guildId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func FetchGuild() <- chan []prpc.COM_Guild {
	rChan := make(chan []prpc.COM_Guild)

	guildCatch := []prpc.COM_Guild{}

	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("FetchGuild panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM Guild")
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		for r.Next() {
			guild := prpc.COM_Guild{}
			qlist := []byte{}
			isFlag := 0
			r.Scan(&guild.GuildId,&guild.GuildName,&guild.Master,&guild.MasterName,&guild.GuildVal,&guild.CreateTime,&qlist,&isFlag,&guild.Require,&guild.Contribution)
			
			e = json.Unmarshal(qlist,&guild.RequestList)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			if isFlag==1 {
				guild.IsRatify = true
			}else {
				guild.IsRatify = false
			}
			
			guildCatch = append(guildCatch, guild)
		}

		rChan <- guildCatch
		close(rChan)
	}()

	return rChan
}

func FetchGuildMember() <-chan []prpc.COM_GuildMember {
	rChan := make(chan []prpc.COM_GuildMember)
	guildMemberCatch := []prpc.COM_GuildMember{}
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("FetchGuildMember panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM GuildMember")
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		for r.Next() {
			m := prpc.COM_GuildMember{}
			e = r.Scan(&m.GuildId,&m.RoleId,&m.RoleName,&m.Level,&m.Job,&m.TianTiVal,&m.UnitId)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
			guildMemberCatch = append(guildMemberCatch,m)
		}

		rChan <- guildMemberCatch
		close(rChan)
	}()
	return rChan
}

func UpdateGuildRequestList(guildId int32,qlist []prpc.COM_GuildRequestData) <- chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateGuildRequestList panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		buffs, erro  :=  json.Marshal(qlist)

		if erro != nil {
			logs.Debug(erro.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt, e := c.Prepare("UPDATE `Guild` SET `RequestList`=? WHERE `GuildId`=?")
		stmt.Exec(buffs,guildId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildVal(guildId int32,val int32) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildVal panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `Guild` SET `GuildVal`=? WHERE `GuildId`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(val,guildId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildRatify(guildId int32,isRatify bool,require int32) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildRatify panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `Guild` SET `RequestFlag` = ? , `Require` = ?  WHERE `GuildId` = ?")
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		if isRatify {
			stmt.Exec(1,require,guildId)
		}else {
			stmt.Exec(0,require,guildId)
		}
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildContribution(guildId int32,contribution int32) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildContribution panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `Guild` SET `Contribution`=?   WHERE `GuildId`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(contribution,guildId)

		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func ResetDBGuildContribution() <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("ResetDBGuildContribution panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `Guild` SET `Contribution`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(0)

		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func InsertGuildMember(member prpc.COM_GuildMember) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertGuildMember panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("INSERT INTO `GuildMember`(`GuildId`, `RoleId`,`RoleName`,`Rolelevel`,`Job`,`TianTiVal`,`UnitId`)VALUES(?,?,?,?,?,?,?)")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(member.GuildId,member.RoleId,member.RoleName,member.Level,member.Job,member.TianTiVal,member.UnitId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func DeleteDBGuildMember(roleId int64) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("DeleteDBGuildMember panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("DELETE FROM `GuildMember` WHERE `RoleId`= ?  ")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(roleId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildMemberVal(player int64,val int32) <- chan bool  {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildMemberVal panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `GuildMember` SET `TianTiVal`=? WHERE `RoleId`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(val,player)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildMemberJob(player int64,job int) <- chan bool  {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildMemberJob panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `GuildMember` SET `Job`=? WHERE `RoleId`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(job,player)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func UpdateDBGuildMemberLevel(player int64,level int32) <- chan bool  {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateDBGuildMemberLevel panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `GuildMember` SET `Rolelevel`=? WHERE `RoleId`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(level,player)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

//////////////////捐赠表/////////////////

func InsertGuildAssistant(data prpc.SGE_DBGuildAssistant) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertGuildAssistant panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()
		buffs , e := json.Marshal(data.Donator)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		stmt, e := c.Prepare("INSERT INTO `GuildAssistant`(`AssistantId`, `RoleName`,`GuildId`,`AssistantItem`,`CrtCount`,`MaxCount`,`CatchNum`,`Donator`)VALUES(?,?,?,?,?,?,?,?)")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(data.Id,data.RoleName,data.GuildId,data.ItemId,data.CrtCount,data.MaxCount,data.CatchNum,buffs)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func DelGuildAssistant(assistantId int32) <-chan bool {
	rChan := make( chan bool)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("DelGuildAssistant panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("DELETE FROM `GuildAssistant` WHERE `AssistantId`= ?  ")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(assistantId)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}

func FindGuildAssistantById(assistantId int32) <-chan *prpc.SGE_DBGuildAssistant {
	rChan := make(chan *prpc.SGE_DBGuildAssistant)
	data := &prpc.SGE_DBGuildAssistant{}
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("FindGuildAssistantById panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM `GuildAssistant` WHERE `AssistantId` = ?",assistantId)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		if r.Next() {
			buffs := []byte{}
			r.Scan(&data.Id,&data.RoleName,&data.GuildId,&data.ItemId,&data.CrtCount,&data.MaxCount,&data.CatchNum,&buffs)
			e = json.Unmarshal(buffs,&data.Donator)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
		}

		rChan <- data
		close(rChan)
	}()
	return rChan
}

func FindGuildAssistantByPlayerName(name string) <-chan *prpc.SGE_DBGuildAssistant {
	rChan := make(chan *prpc.SGE_DBGuildAssistant)
	data := &prpc.SGE_DBGuildAssistant{}
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("FindGuildAssistantByPlayerName panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM `GuildAssistant` WHERE `RoleName` = ?",name)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		if r.Next() {
			buffs := []byte{}
			r.Scan(&data.Id,&data.RoleName,&data.GuildId,&data.ItemId,&data.CrtCount,&data.MaxCount,&data.CatchNum,&buffs)
			e = json.Unmarshal(buffs,&data.Donator)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
		}

		rChan <- data
		close(rChan)
	}()
	return rChan
}

func FindGuildAssistantByGuildId(guildId int32) <-chan []prpc.SGE_DBGuildAssistant {
	rChan := make(chan []prpc.SGE_DBGuildAssistant)
	data := []prpc.SGE_DBGuildAssistant{}
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("FindGuildAssistantByGuildId panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM `GuildAssistant` WHERE `GuildId` = ?",guildId)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		for r.Next() {
			buffs := []byte{}
			m := prpc.SGE_DBGuildAssistant{}
			r.Scan(&m.Id,&m.RoleName,&m.GuildId,&m.ItemId,&m.CrtCount,&m.MaxCount,&m.CatchNum,&buffs)
			e = json.Unmarshal(buffs,&m.Donator)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
			data = append(data,m)
		}

		rChan <- data
		close(rChan)
	}()
	return rChan
}

func UpdateGuildAssistant(data prpc.SGE_DBGuildAssistant) <-chan bool {
	rChan := make( chan bool )
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateGuildAssistant panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()
		buffs , e := json.Marshal(data.Donator)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		stmt, e := c.Prepare("UPDATE `GuildAssistant` SET `CrtCount`=? , `CatchNum`=? , `Donator`=? WHERE `AssistantId`=?")
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		stmt.Exec(data.CrtCount,data.CatchNum,buffs,data.Id)
		stmt.Close()

		rChan <- true
		close(rChan)
	}()
	return rChan
}


////////////////////////////
////
////////////////////////////

func InsertBattleReport(battleid int64, report prpc.SGE_DBBattleReport) <- chan int64 {
	rChan := make (chan int64)
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertBattleReport panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b := bytes.NewBuffer(nil)

		report.Serialize(b)

		r , e := c.Exec("INSERT INTO `BattleReport`(`BattleID`,`Report`)VALUES(?,?)", battleid, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <- (i + 1)
		close(rChan)
	}()
	return  rChan
}

func QueryBattleReport(Battleid int64) <- chan *prpc.SGE_DBBattleReport {
	rChan := make(chan *prpc.SGE_DBBattleReport)
	go func() {
		defer func() {
		if r := recover(); r != nil {
			logs.Error("QueryBattleReport panic %s", fmt.Sprint(r))
			}
		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT * FROM `BattleReport` WHERE `BattleID` = ?", Battleid)

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		if r.Next() {
			a := int64(0)
			b := []byte{}

			e = r.Scan(&a, &b)
				if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := &prpc.SGE_DBBattleReport{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
				if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}


			rChan <- p

			close(rChan)
			return
		}

		rChan <- nil
		close(rChan)
		return
	}()

	return rChan
}

////////////////////////////
////
////////////////////////////


func InsertCheckPointRecordDetail(checkpointid int32, detail prpc.SGE_BattleRecord_Detail) <- chan int64 {
	rChan := make (chan int64)
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertCheckPointRecordDetail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}
		defer c.Close()
		b := bytes.NewBuffer(nil)

		detail.Serialize(b)

		r , e := c.Exec("INSERT INTO `CheckPointBattleRecord`(`CheckPointId`,`Data`)VALUES(?,?)", checkpointid, b.Bytes())

		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		i, e := r.LastInsertId()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- 0
			return
		}

		rChan <- (i + 1)
		close(rChan)
	}()
	return  rChan
}

func UpdateCheckPointRecordDetail(checkpointid int32, detail prpc.SGE_BattleRecord_Detail) {
	go func () {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateCheckPointRecordDetail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			return
		}
		defer c.Close()
		b := bytes.Buffer{}

		detail.Serialize(&b)

		_, e = c.Exec("UPDATE `CheckPointBattleRecord` SET `Data` = ? WHERE `CheckPointId` = ?", b.Bytes(), checkpointid)

		if e != nil {
			logs.Debug(e.Error())
			return
		}
	}()
}

func QueryCheckPointRecordDetail(checkpointid int32) <- chan *prpc.SGE_BattleRecord_Detail {
	rChan := make(chan *prpc.SGE_BattleRecord_Detail)
	go func() {

		defer func() {
			if r := recover(); r != nil {
				logs.Error("QueryPlayerById panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()
		r, e := c.Query("SELECT `Data` FROM `CheckPointBattleRecord` WHERE `CheckPointId` = ?", checkpointid)

		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		if r.Next() {
			b := []byte{}

			e = r.Scan(&b)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}

			p := &prpc.SGE_BattleRecord_Detail{}

			bb := bytes.NewBuffer(b)
			e = p.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}


			rChan <- p

			close(rChan)
			return
		}

		rChan <- nil
		close(rChan)
		return
	}()

	return rChan
}


//Mail
func InsertMail(mail prpc.COM_Mail) <-chan bool {
	rChan := make( chan bool )

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertMail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("INSERT INTO `Mail`( `RecvName` , `SendTime` , `ItemNum` ,`Hero`,`Copper`,`Gold`, `BinData` ) VALUES(?,?,?,?,?,?,?)")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		b :=  bytes.NewBuffer(nil)
		mail.Serialize(b)

		stmt.Exec(mail.RecvPlayerName,mail.MailTimestamp,len(mail.Items),mail.Hero,mail.Copper,mail.Gold,b.Bytes())
		stmt.Close()
		rChan <- true
		close(rChan)
	}()

	return rChan
}

func EraseMail(mailId int32) <-chan bool {
	rChan := make( chan bool )

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("EraseMail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("DELETE FROM `Mail` WHERE `MailGuid`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		stmt.Exec(mailId)
		stmt.Close()
		rChan <- true
		close(rChan)
	}()

	return rChan
}

func UpdateMail(mail prpc.COM_Mail) <-chan bool {
	rChan := make( chan bool )
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("UpdateMail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}
		defer c.Close()

		stmt, e := c.Prepare("UPDATE `Mail` SET `BinData`=?,`ItemNum`=?,`Hero`=?,`Copper`=?,`Gold`=? WHERE `MailGuid`=?")

		if e != nil {
			logs.Debug(e.Error())
			rChan <- false
			close(rChan)
			return
		}

		b :=  bytes.NewBuffer(nil)
		mail.Serialize(b)

		stmt.Exec(b.Bytes(),len(mail.Items),mail.Hero,mail.Copper,mail.Gold,mail.MailId)
		stmt.Close()
		rChan <- true
		close(rChan)
	}()
	return rChan
}

func FatchDBMail(recv string, fatchid int32) <- chan []prpc.COM_Mail {
	rChan := make( chan []prpc.COM_Mail)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logs.Error("InsertMail panic %s", fmt.Sprint(r))
			}

		}()

		c, e := ConnectDB()
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}
		defer c.Close()

		r, e := c.Query("SELECT * FROM `Mail` WHERE `RecvName`= ? and `MailGuid` > ? order by `MailGuid` asc limit 100;",recv,fatchid)
		if e != nil {
			logs.Debug(e.Error())
			rChan <- nil
			close(rChan)
			return
		}

		arr := []prpc.COM_Mail{}

		for r.Next() {
			mailId 		:= int32(0)
			recvName	:= ""
			sendTime	:= int64(0)
			itemNum 	:= int32(0)
			hero 		:= int32(0)
			copper 		:= int32(0)
			gold 		:= int32(0)
			data 		:= []byte{}

			e = r.Scan(&mailId,&recvName,&sendTime,&itemNum,&hero,&copper,&gold,&data)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
			m := prpc.COM_Mail{}
			bb := bytes.NewBuffer(data)
			e = m.Deserialize(bb)
			if e != nil {
				logs.Debug(e.Error())
				rChan <- nil
				close(rChan)
				return
			}
			m.MailId = mailId

			arr = append(arr,m)
		}

		rChan <- arr
		close(rChan)
	}()

	return rChan
}