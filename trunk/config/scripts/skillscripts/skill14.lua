sys.log("skill 14 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--	取得目标 （GetTarget（）  单   GetTargets（）  复）
--  取得对应属性 GetUnitProperty()
--  计算伤害数值 demage
--  计算是否暴击
--  攻击
-- 女娲2号技能 增加一个友方单位30%法术伤害的荆棘，持续3回合。

-- 物理强度视作buff Battle.buff

function SK_113_Action(battleid, casterid)

	Battle.TargetOn(battleid) -- 清空数据
	local skillid = 113		-- 技能id
	
	local t = Player.GetFriend(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	
	local  caster_attack = Player.GetUnitMtk(battleid,casterid)  --获取攻击者属性  法术
	
	Battle.Cure(battleid,t,0,0)
	
	local mag_damage = caster_attack*0.3
	
	Battle.AddBuff(battleid,casterid, t,107, mag_damage)  --给一个友方单位增加一个荆棘（暂时么有这个函数）   持续三回合
	
	Battle.TargetOver(battleid) -- 赋给下个目标
	
	sys.log("skill14")
	
	
	return true
end

sys.log("skill 14 end")