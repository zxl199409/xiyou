sys.log("skill 1 start")

-- 技能释放 传入战斗ID和释放者的ID
-- 通过释放者和battleid取得对应的目标 单体或者多个
-- 循环/直接使用接口操控战斗 类似 战斗.攻击(战斗id, 释放者id, 承受者ID, 伤害数值, 是否暴击)
-- 
-- 
-- 所需接口
--  取得对应属性
--  计算伤害数值
--  计算是否暴击
--  攻击
-- 
-- 

function SK_1_Action(battleid, casterid)
	local skillid = 1		-- 技能id
	local skillAttack = 10	-- 技能攻击
	
	local t = Player.GetTarget(battleid, casterid)	-- 获取到的目标,可以为单体也可以为复数,根据不同需求选择
	sys.log("GetTarget  ".. t)
	
	local caster_attack = Player.GetUnitProperty(battleid, casterid, 6)	-- 获取到攻击者的属性
	local defender_def = Player.GetUnitProperty(battleid, t, 3)
	
	local damage = caster_attack - defender_def
	
	if damage <= 0 then 
		damage = 1
	end
	
	Battle.Attack(battleid, t, damage, true)
	
	-- 只给游戏返回 对谁造成了多少伤害
	-- 并不参与计算
	return true
end

sys.log("skill1 end")