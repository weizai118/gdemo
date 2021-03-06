﻿

using System;

public class ModuleShop : IModule
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

    public void requestBuyDiamond(int idx)
    {
        netproto.CS_ShopAddDiamond msg = new netproto.CS_ShopAddDiamond
        {
            idx = idx,
        };
        this.mGameServer.send(msg);
    }

    public void requestBuyGold()
    {
        netproto.CS_ShopBuyGold msg = new netproto.CS_ShopBuyGold
        {
        };
        this.mGameServer.send(msg);
    }

    [MsgRoute]
    public void response(netproto.SC_ShopAddDiamond msg)
    {
        SysLog.debug("buy diamond result " + msg.code);
    }

    [MsgRoute]
    public void response(netproto.SC_ShopBuyGold msg)
    {
        SysLog.debug("buy gold result " + msg.code);
    }
}