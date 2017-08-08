package prpc

import (
	"bytes"
	"errors"
	"suzuki/prpc"
)

type COM_ServerToClient_ErrorMessage struct {
	id int //0
}
type COM_ServerToClient_LoginOK struct {
	info COM_AccountInfo //0
}
type COM_ServerToClient_CreatePlayerOK struct {
	player COM_Player //0
}
type COM_ServerToClient_JoinBattleOk struct {
	Camp int32 //0
}
type COM_ServerToClient_SetBattleUnitOK struct {
	instId int64 //0
}
type COM_ServerToClient_BattleReport struct {
	report COM_BattleReport //0
}
type COM_ServerToClient_BattleExit struct {
	result COM_BattleResult //0
}
type COM_ServerToClientStub struct {
	Sender prpc.StubSender
}
type COM_ServerToClientProxy interface {
	ErrorMessage(id int) error                  // 0
	LoginOK(info COM_AccountInfo) error         // 1
	CreatePlayerOK(player COM_Player) error     // 2
	JoinBattleOk(Camp int32) error              // 3
	SetupBattleOK() error                       // 4
	SetBattleUnitOK(instId int64) error         // 5
	BattleReport(report COM_BattleReport) error // 6
	BattleExit(result COM_BattleResult) error   // 7
}

func (this *COM_ServerToClient_ErrorMessage) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.id != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize id
	{
		if this.id != 0 {
			err := prpc.Write(buffer, this.id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ServerToClient_ErrorMessage) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize id
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_LoginOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //info
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize info
	{
		err := this.info.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_LoginOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize info
	if mask.ReadBit() {
		err := this.info.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_CreatePlayerOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //player
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize player
	{
		err := this.player.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_CreatePlayerOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize player
	if mask.ReadBit() {
		err := this.player.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_JoinBattleOk) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.Camp != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize Camp
	{
		if this.Camp != 0 {
			err := prpc.Write(buffer, this.Camp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ServerToClient_JoinBattleOk) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize Camp
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.Camp)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(this.instId != 0)
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize instId
	{
		if this.instId != 0 {
			err := prpc.Write(buffer, this.instId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (this *COM_ServerToClient_SetBattleUnitOK) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize instId
	if mask.ReadBit() {
		err := prpc.Read(buffer, &this.instId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_BattleReport) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //report
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize report
	{
		err := this.report.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_BattleReport) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize report
	if mask.ReadBit() {
		err := this.report.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_BattleExit) Serialize(buffer *bytes.Buffer) error {
	//field mask
	mask := prpc.NewMask1(1)
	mask.WriteBit(true) //result
	{
		err := prpc.Write(buffer, mask.Bytes())
		if err != nil {
			return err
		}
	}
	// serialize result
	{
		err := this.result.Serialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClient_BattleExit) Deserialize(buffer *bytes.Buffer) error {
	//field mask
	mask, err := prpc.NewMask0(buffer, 1)
	if err != nil {
		return err
	}
	// deserialize result
	if mask.ReadBit() {
		err := this.result.Deserialize(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}
func (this *COM_ServerToClientStub) ErrorMessage(id int) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(0))
	if err != nil {
		return err
	}
	_0 := COM_ServerToClient_ErrorMessage{}
	_0.id = id
	err = _0.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) LoginOK(info COM_AccountInfo) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(1))
	if err != nil {
		return err
	}
	_1 := COM_ServerToClient_LoginOK{}
	_1.info = info
	err = _1.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) CreatePlayerOK(player COM_Player) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(2))
	if err != nil {
		return err
	}
	_2 := COM_ServerToClient_CreatePlayerOK{}
	_2.player = player
	err = _2.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) JoinBattleOk(Camp int32) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(3))
	if err != nil {
		return err
	}
	_3 := COM_ServerToClient_JoinBattleOk{}
	_3.Camp = Camp
	err = _3.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) SetupBattleOK() error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(4))
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) SetBattleUnitOK(instId int64) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(5))
	if err != nil {
		return err
	}
	_5 := COM_ServerToClient_SetBattleUnitOK{}
	_5.instId = instId
	err = _5.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) BattleReport(report COM_BattleReport) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(6))
	if err != nil {
		return err
	}
	_6 := COM_ServerToClient_BattleReport{}
	_6.report = report
	err = _6.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func (this *COM_ServerToClientStub) BattleExit(result COM_BattleResult) error {
	buffer := this.Sender.MethodBegin()
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	err := prpc.Write(buffer, uint16(7))
	if err != nil {
		return err
	}
	_7 := COM_ServerToClient_BattleExit{}
	_7.result = result
	err = _7.Serialize(buffer)
	if err != nil {
		return err
	}
	return this.Sender.MethodEnd()
}
func Bridging_COM_ServerToClient_ErrorMessage(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_0 := COM_ServerToClient_ErrorMessage{}
	err := _0.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.ErrorMessage(_0.id)
}
func Bridging_COM_ServerToClient_LoginOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_1 := COM_ServerToClient_LoginOK{}
	err := _1.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.LoginOK(_1.info)
}
func Bridging_COM_ServerToClient_CreatePlayerOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_2 := COM_ServerToClient_CreatePlayerOK{}
	err := _2.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.CreatePlayerOK(_2.player)
}
func Bridging_COM_ServerToClient_JoinBattleOk(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_3 := COM_ServerToClient_JoinBattleOk{}
	err := _3.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.JoinBattleOk(_3.Camp)
}
func Bridging_COM_ServerToClient_SetupBattleOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	return p.SetupBattleOK()
}
func Bridging_COM_ServerToClient_SetBattleUnitOK(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_5 := COM_ServerToClient_SetBattleUnitOK{}
	err := _5.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.SetBattleUnitOK(_5.instId)
}
func Bridging_COM_ServerToClient_BattleReport(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_6 := COM_ServerToClient_BattleReport{}
	err := _6.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.BattleReport(_6.report)
}
func Bridging_COM_ServerToClient_BattleExit(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	_7 := COM_ServerToClient_BattleExit{}
	err := _7.Deserialize(buffer)
	if err != nil {
		return err
	}
	return p.BattleExit(_7.result)
}
func COM_ServerToClientDispatch(buffer *bytes.Buffer, p COM_ServerToClientProxy) error {
	if buffer == nil {
		return errors.New(prpc.NoneBufferError)
	}
	if p == nil {
		return errors.New(prpc.NoneProxyError)
	}
	pid := uint16(0XFFFF)
	err := prpc.Read(buffer, &pid)
	if err != nil {
		return err
	}
	switch pid {
	case 0:
		return Bridging_COM_ServerToClient_ErrorMessage(buffer, p)
	case 1:
		return Bridging_COM_ServerToClient_LoginOK(buffer, p)
	case 2:
		return Bridging_COM_ServerToClient_CreatePlayerOK(buffer, p)
	case 3:
		return Bridging_COM_ServerToClient_JoinBattleOk(buffer, p)
	case 4:
		return Bridging_COM_ServerToClient_SetupBattleOK(buffer, p)
	case 5:
		return Bridging_COM_ServerToClient_SetBattleUnitOK(buffer, p)
	case 6:
		return Bridging_COM_ServerToClient_BattleReport(buffer, p)
	case 7:
		return Bridging_COM_ServerToClient_BattleExit(buffer, p)
	default:
		return errors.New(prpc.NoneDispatchMatchError)
	}
}
