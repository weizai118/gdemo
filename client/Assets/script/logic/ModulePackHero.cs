﻿

using System;

public class ModulePackHero : IModule
{
    private ModuleNetGameServ mGameServer = null;

    public override bool init()
    {
        return true;
    }

    public override bool prestart()
    {
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(int delta)
    {
        
    }
}