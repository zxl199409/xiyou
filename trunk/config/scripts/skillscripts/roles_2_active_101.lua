sys.log("主角 主动技能 2 开始")

--主角2技能  对敌方全体造成法术攻击

function SK_101_Action(battleid, casterid)
	
	
	local skillid = 101	-- 技能id

	local level = 1

	local attackNum = 0
	
	local  t = Player.GetTargets(battleid,casterid,attackNum)  --获取目标

	for i,v in ipairs(t) do
	
		Battle.TargetOn(battleid) -- 清空数据
	
	    local  truedamage  = Player.GetMagicDamage(battleid,casterid,v)       --伤害 公式（）
		
		sys.log("主角对目标造成的法术伤害    "..truedamage)
		
		local damage = ClacDamageByAllBuff(battleid,casterid,v,truedamage)
		sys.log("主角对目标造成的最终法术伤害    "..damage)
		--判断伤害
		if damage <= 0 then 
		
			damage = 1
		
		end
		
		local crit = Battle.GetCrit(skillid)   --是否暴击
		
		Battle.Attack(battleid,casterid,v,damage,crit)   --调用服务器 （伤害）(战斗者，释放者，承受者，伤害，暴击）
		
	
		Battle.TargetOver(battleid)  -- 赋给下个目标
		
		
	end

end


sys.log("主角 主动技能 2 结束")