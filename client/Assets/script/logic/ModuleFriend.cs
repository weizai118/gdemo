﻿
public class ModuleFriend : IModule
{
    private ModuleNetGameServ mGameServer = null;
    public override bool init()
    {
        return true;
    }

    public override void uninit()
    {

    }

    public override bool prestart()
    {
        return true;
    }

    public override void update(int delta)
    {

    }
}