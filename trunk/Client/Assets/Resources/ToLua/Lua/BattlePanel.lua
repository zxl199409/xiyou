require "FairyGUI"

BattlePanel = fgui.window_class(WindowBase)
local Window;

local autoBtn;
local stateIcon;
local battStartIcon;

local battleStart = {};
battleStart.max = 1;
battleStart.count = 0;

local cards = {};
local cardsInGroup = {};
local cardsInGroupNum;

function BattlePanel:OnEntry()
	Window = BattlePanel.New();
	Window:Show();
end

function BattlePanel:OnInit()
	self.contentPane = UIPackage.CreateObject("BattlePanel", "BattlePanel_com").asCom;
	self:Center();

	battStartIcon = self.contentPane:GetChild("n7").asImage;

	local returnBtn = self.contentPane:GetChild("n8").asButton;
	returnBtn.onClick:Add(BattlePanel_OnReturnBtn);

	autoBtn = self.contentPane:GetChild("n12").asButton;
	autoBtn.onClick:Add(BattlePanel_OnAutoBtn);

	stateIcon = self.contentPane:GetChild("n16").asLoader;
	stateIcon.onClick:Add(BattlePanel_OnTurnOver);

	for i=1, 5 do
		cards[i] = {};
		cards[i]["card"] = self.contentPane:GetChild("n" .. (16 + i)).asCom;
		cards[i]["power"] = cards[i]["card"]:GetChild("power");
		cards[i]["cost"] = cards[i]["card"]:GetChild("cost");
	end

	for i=1, 10 do
		cardsInGroup[i] = self.contentPane:GetChild("n" .. (23 + i));
		if cardsInGroup[i] == nil then
			print(i .. " is nil");
		end
	end

	cardsInGroupNum = self.contentPane:GetChild("n43");

	BattlePanel_FlushData();
end

function BattlePanel:GetWindow()
	return Window;
end

function BattlePanel:OnUpdate()
	if UIManager.IsDirty("BattlePanel") then
		BattlePanel_FlushData();
		UIManager.ClearDirty("BattlePanel");
	end
end

function BattlePanel:OnTick()
	--2秒倒计时 注意return
	if battleStart == nil then
		return;
	end

	battleStart.count = battleStart.count + 1;
	if battleStart.count > battleStart.max then
		BattlePanel_OnHiddenBattleStart();
		battleStart = nil;
	end
end

function BattlePanel:isShow()
	return Window.isShowing;
end

function BattlePanel:OnDispose()
	Window:Dispose();
end

function BattlePanel:OnHide()
	Window:Hide();
end

function BattlePanel_FlushData()
	if Battle._CurrentState == Battle.BattleState.BS_Oper then
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_jieshuhuihe");
		stateIcon.touchable = true;
	else
		stateIcon.url = UIPackage.GetItemURL("BattlePanel", "battle_dengdaizhong");
		stateIcon.touchable = false;
	end

	local cardNum = Battle._LeftCardNum;
	for i=1, 5 do
		if i <= cardNum then
			cards[i]["card"].data = i;
			cards[i]["card"].onClick:Add(BattlePanel_OnCardClick);
			cards[i]["power"].text = i;
			cards[i]["cost"].text = i;
			cards[i]["card"].visible = true;
			if Battle._Turn == 1 and not Battle.IsSelfCard(i-1) then
				cards[i]["card"].enabled = false;
			else
				cards[i]["card"].enabled = true;
			end
		else
			cards[i]["card"].visible = false;
		end
	end

	local mainActor = Battle.GetActor(GamePlayer._InstID);
	if Battle._Turn == 1 and mainActor == NULL then
		if Battle._CurrentState == Battle.BattleState.BS_Oper then
			stateIcon.enabled = false;
		end
	else 
		stateIcon.enabled = true;
	end

	for i=1, 10 do
		if i <= Battle.CardsInGroupCount then
			cardsInGroup[i].visible = true;
		else
			cardsInGroup[i].visible = false;
		end
	end

	cardsInGroupNum.text = Battle.CardsInGroupCount;
	if Battle.CardsInGroupCount <= 0 then
		cardsInGroupNum.visible = false;
	else
		cardsInGroupNum.visible = true;
	end
end

function BattlePanel_OnCardClick(context)
	print(context.sender.data);
	Proxy4Lua.SelectCard4Ready(context.sender.data - 1);
end

function BattlePanel_OnReturnBtn()
	print("OnReturnBtn");
	SceneLoader.LoadScene("main");
end

function BattlePanel_OnTurnOver()
	print("OnTurnOver");
	Proxy4Lua.BattleSetup();
end

function BattlePanel_OnAutoBtn()
	print("OnAutoBtn");
	GamePlayer._IsAuto = not GamePlayer._IsAuto;
	if GamePlayer._IsAuto then
		autoBtn:GetController("icon").selectedIndex = 1;
	else
		autoBtn:GetController("icon").selectedIndex = 0;
	end
end

function BattlePanel_OnHiddenBattleStart()
	print("BattlePanel_OnHiddenBattleStart");
	battStartIcon.visible = false;
end