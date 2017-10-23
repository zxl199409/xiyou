sys.log("虾兵 SK_315_Action  开始")

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
-- 铁剪蟹将技能 对敌人造成物理强度的伤害。

-- 法术强度视作buff  Battle.buff

function SK_315_Action(battleid, casterid)

	Battle.TargetOn(battleid)
	local skillid = 315		-- 技能id
	
	local  t = Player.GetTarget(battleid,casterid)  --获取目标 
	local  true_damage = Player.GetUnitDamage(battleid,casterid,t)
	sys.log("铁剪蟹将技能 的伤害"..t.."造成物理伤害"..true_damage)
	local damage = ClacDamageByAllBuff(battleid,casterid,t,true_damage)
	sys.log("铁剪蟹将技能 的伤害"..t.."造成最终物理伤害"..damage)
	--判断伤害
	if damage <= 0 then 
		damage = 0
	end
	local crit = Battle.GetCrit(skillid)   --是否暴击
	Battle.Attack(battleid,casterid,t,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
	sys.log("铁剪蟹将 被动技能连击   伤害    "   ..  damage)

	--被动技能连击 
	Battle.Attack(battleid,casterid,t,damage,crit)
	Battle.TargetOver(battleid)
		
	sys.log("铁剪蟹将对id为"..t.."的目标造成"..damage.."点伤害")
	
	return  true
	 
end
sys.log("虾兵 SK_315_Action  结束")