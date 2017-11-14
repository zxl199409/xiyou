package prpc
import(
  "bytes"
  "sync"
  "encoding/json"
)
type COM_BattleAction struct{
  sync.Mutex
  InstId int64  //0
  BuffList []COM_BattleBuffAction  //1
  SkillId int32  //2
  SkillBuff []COM_BattleBuff  //3
  TargetList []COM_BattleActionTarget  //4
  UnitList []COM_ChangeUnit  //5
}
func (this *COM_BattleAction)SetInstId(value int64) {
  this.Lock()
  defer this.Unlock()
  this.InstId = value
}
func (this *COM_BattleAction)GetInstId() int64 {
  this.Lock()
  defer this.Unlock()
  return this.InstId
}
func (this *COM_BattleAction)SetBuffList(value []COM_BattleBuffAction) {
  this.Lock()
  defer this.Unlock()
  this.BuffList = value
}
func (this *COM_BattleAction)GetBuffList() []COM_BattleBuffAction {
  this.Lock()
  defer this.Unlock()
  return this.BuffList
}
func (this *COM_BattleAction)SetSkillId(value int32) {
  this.Lock()
  defer this.Unlock()
  this.SkillId = value
}
func (this *COM_BattleAction)GetSkillId() int32 {
  this.Lock()
  defer this.Unlock()
  return this.SkillId
}
func (this *COM_BattleAction)SetSkillBuff(value []COM_BattleBuff) {
  this.Lock()
  defer this.Unlock()
  this.SkillBuff = value
}
func (this *COM_BattleAction)GetSkillBuff() []COM_BattleBuff {
  this.Lock()
  defer this.Unlock()
  return this.SkillBuff
}
func (this *COM_BattleAction)SetTargetList(value []COM_BattleActionTarget) {
  this.Lock()
  defer this.Unlock()
  this.TargetList = value
}
func (this *COM_BattleAction)GetTargetList() []COM_BattleActionTarget {
  this.Lock()
  defer this.Unlock()
  return this.TargetList
}
func (this *COM_BattleAction)SetUnitList(value []COM_ChangeUnit) {
  this.Lock()
  defer this.Unlock()
  this.UnitList = value
}
func (this *COM_BattleAction)GetUnitList() []COM_ChangeUnit {
  this.Lock()
  defer this.Unlock()
  return this.UnitList
}
func (this *COM_BattleAction)Serialize(buffer *bytes.Buffer) error {
  this.Lock()
  defer this.Unlock()
  //field mask
  mask := newMask1(1)
  mask.writeBit(this.InstId!=0)
  mask.writeBit(len(this.BuffList) != 0)
  mask.writeBit(this.SkillId!=0)
  mask.writeBit(len(this.SkillBuff) != 0)
  mask.writeBit(len(this.TargetList) != 0)
  mask.writeBit(len(this.UnitList) != 0)
  {
    err := write(buffer,mask.bytes())
    if err != nil {
      return err
    }
  }
  // serialize InstId
  {
    if(this.InstId!=0){
      err := write(buffer,this.InstId)
      if err != nil{
        return err
      }
    }
  }
  // serialize BuffList
  if len(this.BuffList) != 0{
    {
      err := write(buffer,uint(len(this.BuffList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.BuffList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize SkillId
  {
    if(this.SkillId!=0){
      err := write(buffer,this.SkillId)
      if err != nil{
        return err
      }
    }
  }
  // serialize SkillBuff
  if len(this.SkillBuff) != 0{
    {
      err := write(buffer,uint(len(this.SkillBuff)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.SkillBuff {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize TargetList
  if len(this.TargetList) != 0{
    {
      err := write(buffer,uint(len(this.TargetList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.TargetList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  // serialize UnitList
  if len(this.UnitList) != 0{
    {
      err := write(buffer,uint(len(this.UnitList)))
      if err != nil {
        return err
      }
    }
    for _, value := range this.UnitList {
      err := value.Serialize(buffer)
      if err != nil {
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleAction)Deserialize(buffer *bytes.Buffer) error{
  this.Lock()
  defer this.Unlock()
  //field mask
  mask, err:= newMask0(buffer,1);
  if err != nil{
    return err
  }
  // deserialize InstId
  if mask.readBit() {
    err := read(buffer,&this.InstId)
    if err != nil{
      return err
    }
  }
  // deserialize BuffList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.BuffList = make([]COM_BattleBuffAction,size)
    for i,_ := range this.BuffList{
      err := this.BuffList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize SkillId
  if mask.readBit() {
    err := read(buffer,&this.SkillId)
    if err != nil{
      return err
    }
  }
  // deserialize SkillBuff
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.SkillBuff = make([]COM_BattleBuff,size)
    for i,_ := range this.SkillBuff{
      err := this.SkillBuff[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize TargetList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.TargetList = make([]COM_BattleActionTarget,size)
    for i,_ := range this.TargetList{
      err := this.TargetList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  // deserialize UnitList
  if mask.readBit() {
    var size uint
    err := read(buffer,&size)
    if err != nil{
      return err
    }
    this.UnitList = make([]COM_ChangeUnit,size)
    for i,_ := range this.UnitList{
      err := this.UnitList[i].Deserialize(buffer)
      if err != nil{
        return err
      }
    }
  }
  return nil
}
func (this *COM_BattleAction)String() string{
  b, e := json.Marshal(this)
  if e != nil{
    return e.Error()
  }else{
    return string(b)
  }
}
