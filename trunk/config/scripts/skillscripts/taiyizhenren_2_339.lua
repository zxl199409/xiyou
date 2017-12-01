sys.log(" 太乙真人 SK_339_Action 开始")

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
--太乙真人2号技能 重塑真身：从你的手牌中随机上阵一张卡牌（没有手牌或者场上没有位置则不动一回合）

function SK_339_Action(battleid,casterid)
	Battle.TargetOn(battleid)
	
	local skillid = 339
	
	Battle.InToBattleOnFighting(battleid,casterid)

	Battle.TargetOver(battleid)
end
sys.log(" 太乙真人 SK_339_Action 结束")