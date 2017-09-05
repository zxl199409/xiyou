require "FairyGUI"

qiecuo = fgui.window_class(WindowBase)
local Window;

local allCardGroupList;
local cardGroupList;
local cardGroupUrl = "ui://qiecuo/paizuanniu_Button";
local cardItemUrl = "ui://qiecuo/touxiangkuang_Label";
local crtCardName;
local crtGroupIdx = 0;

function qiecuo:OnEntry()
	Window = qiecuo.New();
	Window:Show();
end

function qiecuo:GetWindow()
	return Window;
end

function qiecuo:OnInit()
 	self.contentPane = UIPackage.CreateObject("qiecuo", "qiecuo_com").asCom;
	self:Center();
	self.closeButton = self.contentPane:GetChild("n7").asButton;

	local rightPart = self.contentPane:GetChild("n5").asCom;
	local bg = rightPart:GetChild("n3");
	allCardGroupList = bg:GetChild("n5").asList;

	local cardGroup = rightPart:GetChild("n6");
	crtCardName = cardGroup:GetChild("n23");
	cardGroupList = cardGroup:GetChild("n27").asList;
	local setBattleBtn = cardGroup:GetChild("n29").asButton;
	setBattleBtn.visible = false;
	for i=1, 5 do
		local groupItem = allCardGroupList:AddItemFromPool(cardGroupUrl);
		groupItem.onClick:Add(qiecuo_OnSelectGroup);
		groupItem.data = i - 1;
	end

	allCardGroupList.selectedIndex = crtGroupIdx;

	local liftPanel = self.contentPane:GetChild("n6").asCom;
	local startBtn = liftPanel:GetChild("n35");
	startBtn.onClick:Add(qiecuo_OnStart);

	qiecuo_FlushData();
end



function qiecuo_RenderListItem(index, obj)

end

function qiecuo_OnDeleteGroup(context)
	local MessageBox = UIManager.ShowMessageBox();
end
 
function qiecuo_OnSetBattle(context)
	local MessageBox = UIManager.ShowMessageBox();
end





function qiecuo:OnUpdate()
	if UIManager.IsDirty("qiecuo") then
		qiecuo_FlushData();
		UIManager.ClearDirty("qiecuo");
	end
end

function qiecuo:OnTick()
	 
end

function qiecuo:isShow()
	return Window.isShowing;
end

function qiecuo:OnDispose()
	Window:Dispose();
end

function qiecuo:OnHide()
	Window:Hide();
end

function qiecuo_FlushData()

	cardGroupList:RemoveChildrenToPool(); 
	local groupCards = GamePlayer.GetGroupCards(crtGroupIdx);
	if groupCards == nil then
		return;
	end
	local displayData;
	local entityData;
	for i=1, groupCards.Count do
		displayData = GamePlayer.GetDisplayDataByIndexFromGroup(crtGroupIdx, i - 1);
		entityData = GamePlayer.GetEntityDataByIndexFromGroup(crtGroupIdx, i - 1);
		local itemBtn = cardGroupList:AddItemFromPool(cardItemUrl);
		itemBtn:GetChild("n5").asLoader.url = "ui://" .. displayData._HeadIcon;
		local fee = itemBtn:GetChild("n7");
		fee.text = entityData._Cost
		itemBtn.onClick:Add(qiecuo_OnCardInGroup);
		itemBtn.data = GamePlayer.GetInstIDFromGroup(crtGroupIdx, i - 1);
	end


	local groupItem;
	local groupName;
	for i=1, 5 do
			groupItem = allCardGroupList:GetChildAt(i-1);
			groupName = GamePlayer.GetGroupName(i - 1);
			if groupName == "" then
				groupName = "卡组" .. i;
			end
			groupItem:GetChild("n3").text = groupName;
			if crtGroupIdx == i - 1 then
				crtCardName.text = groupName;
			end
			if i - 1 == GamePlayer._CrtBattleGroupIdx then
				groupItem:GetChild("n4").visible = true;
			else
				groupItem:GetChild("n4").visible = false;
			end
		end
end



function qiecuo_OnSelectGroup(context)
	crtGroupIdx = context.sender.data;
	qiecuo_FlushData();
end


function qiecuo_OnStart(context)
	Proxy4Lua.StartMatching(crtGroupIdx);
	qiecuo_FlushData();
end

function qiecuo_OnCardInGroup(context)
	crtCardInstID = context.sender.data;

	UIParamHolder.Set("qiecuo1", crtCardInstID);
	UIParamHolder.Set("qiecuo2", true);
	UIManager.Show("xiangxiziliao");
end
