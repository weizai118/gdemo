﻿

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerFight : Layer
{
    public LayerFight() : base("ui", "fight")
    {
    }

    protected override bool init()
    {

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }
}
