﻿

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerTemplate : Layer
{
    public LayerTemplate(string group, string name) : base(group, name)
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