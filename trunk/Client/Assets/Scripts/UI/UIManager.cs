﻿using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using LuaInterface;
using System;

public class UIManager {

    static Transform _UIRoot;

    static float _Timer;

    static public LuaState _Lua;

    static public void Init(Transform uiroot)
    {
        _UIRoot = uiroot;
        _Lua = new LuaState();
        LuaBinder.Bind(_Lua);
        _Lua.Start();
    }

    static public Transform UIRoot
    {
        get
        {
            return _UIRoot;
        }
    }

    static Dictionary<string, UIWindow> _Windows = new Dictionary<string, UIWindow>();
    static Dictionary<string, bool> _DirtyPool = new Dictionary<string, bool>();
    static List<string> _WantClearDirty = new List<string>();

    static public void Show(string uiName)
    {
        if (IsShow(uiName))
            return;

        if (!_Windows.ContainsKey(uiName))
            _Windows.Add(uiName, new UIWindow(uiName));

        _Windows [uiName].Show();
        _Windows [uiName].Start();
    }

    static public bool IsShow(string uiName)
    {
        bool isShow = _Windows.ContainsKey(uiName) && _Windows [uiName].IsShow;
        return isShow;
    }

    static public void Hide(string uiName)
    {
        if (!IsShow(uiName))
            return;

        //TODO 暂时及时销毁界面
        _Windows [uiName].Dispose();

        AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + uiName);
    }

    static public void HideAll()
    {
        //TODO 暂时及时销毁界面
        foreach(UIWindow window in _Windows.Values)
        {
            window.Dispose();
            AssetLoader.UnloadAsset(PathDefine.UI_ASSET_PATH + window.UIName);
        }
    }

    static public void Update()
    {
        _Timer += Time.deltaTime;

        foreach(UIWindow window in _Windows.Values)
        {
            window.Update();
            if (_Timer >= 1f)
                window.Tick();
        }

        if (_Timer >= 1f)
            _Timer = 0f;

        //重置界面Dirty属性
        for(int i=0; i < _WantClearDirty.Count; ++i)
        {
            if (!_DirtyPool.ContainsKey(_WantClearDirty[i]))
                continue;

            _DirtyPool [_WantClearDirty[i]] = false;
        }
        _WantClearDirty.Clear();
    }

    static public void SetDirty(string uiName)
    {
        if (!_DirtyPool.ContainsKey(uiName))
            return;

        _DirtyPool [uiName] = true;
    }

    static public bool IsDirty(string uiName)
    {
        if (!_DirtyPool.ContainsKey(uiName))
            return false;

        return _DirtyPool [uiName];
    }

    static public void ClearDirty(string uiName)
    {
        if (_WantClearDirty.Contains(uiName))
            return;
        
        _WantClearDirty.Add(uiName);
    }
}
