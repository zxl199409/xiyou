﻿using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CreateRoleSceneTouch : MonoBehaviour {

    Actor[] _Actors;

    Vector3 _MalePos, _FemalePos, _ToPos;
    float _MaleRotY, _FemaleRotY;

    string _DefaultAnim;
    string _SelectAnim;

	// Use this for initialization
	void Start () {
        UIManager.Show("xuanren");

        string[] posStr = Define.GetStr("CreateMalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _MalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));
        _MaleRotY = Define.GetFloat("CreateMaleRotY");

        posStr = Define.GetStr("CreateFemalePos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _FemalePos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));
        _FemaleRotY = Define.GetFloat("CreateFemaleRot");

        posStr = Define.GetStr("CreateSelectPos").Split(new char[]{','}, StringSplitOptions.RemoveEmptyEntries);
        _ToPos = new Vector3(float.Parse(posStr[0]), float.Parse(posStr[1]), float.Parse(posStr[2]));

        _DefaultAnim = Define.GetStr("CreateDefaultClip");
        _SelectAnim = Define.GetStr("CreateSelectClip");

        _Actors = new Actor[2];
        EntityData eData = EntityData.GetData(Define.MALE_ID);
        DisplayData dData = DisplayData.GetData(eData._DisplayId);
        GameObject actorObj = AssetLoader.LoadAsset(dData._AssetPath);
        actorObj.transform.Rotate(Vector3.up, _MaleRotY);
        _Actors [0] = new Actor(actorObj, _MalePos, 0, "", "", null, dData._Id);
        _Actors [0].Play(_DefaultAnim);

        eData = EntityData.GetData(Define.FEMALE_ID);
        dData = DisplayData.GetData(eData._DisplayId);
        actorObj = AssetLoader.LoadAsset(dData._AssetPath);
        actorObj.transform.Rotate(Vector3.up, _FemaleRotY);
        _Actors [1] = new Actor(actorObj, _FemalePos, 0, "", "", null, dData._Id);
        _Actors [1].Play(_DefaultAnim);
	}

    public void SelectMale()
    {
        if (_Actors [0] == null)
            return;

        _Actors [0].Play(_SelectAnim);
        _Actors [0].MoveTo(_ToPos, delegate {
            _Actors [0].Play(Define.ANIMATION_PLAYER_ACTION_IDLE);
            _Actors [0]._ActorObj.transform.localRotation = Quaternion.identity;
        });

        if (_Actors [1] == null)
            return;

        if (Vector3.Distance(_Actors [1]._ActorObj.transform.position, _FemalePos) <= 1f)
            return;
        
        _Actors [1].Play(_SelectAnim);
        _Actors [1].MoveTo(_FemalePos, delegate {
            _Actors [1].Play(_DefaultAnim);
            _Actors [1]._ActorObj.transform.localRotation = Quaternion.identity;
            _Actors [1]._ActorObj.transform.Rotate(Vector3.up, _FemaleRotY);
        });
    }

    public void SelectFemale()
    {
        if (_Actors [1] == null)
            return;

        _Actors [1].Play(_SelectAnim);
        _Actors [1].MoveTo(_ToPos, delegate {
            _Actors [1].Play(Define.ANIMATION_PLAYER_ACTION_IDLE);
            _Actors [1]._ActorObj.transform.localRotation = Quaternion.identity;
        });

        if (_Actors [0] == null)
            return;

        if (Vector3.Distance(_Actors [0]._ActorObj.transform.position, _MalePos) <= 1f)
            return;

        _Actors [0].Play(_SelectAnim);
        _Actors [0].MoveTo(_MalePos, delegate {
            _Actors [0].Play(_DefaultAnim);
            _Actors [0]._ActorObj.transform.localRotation = Quaternion.identity;
            _Actors [0]._ActorObj.transform.Rotate(Vector3.up, _MaleRotY);
        });
    }
	
	// Update is called once per frame
	void Update () {
		
	}
}
