//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Option: missing-value detection (*Specified/ShouldSerialize*/Reset*) enabled
    
// Generated from: logic.proto
namespace netproto
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_Tick")]
  public partial class CS_Tick : global::ProtoBuf.IExtensible
  {
    public CS_Tick() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_Tick")]
  public partial class SC_Tick : global::ProtoBuf.IExtensible
  {
    public SC_Tick() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_Login")]
  public partial class CS_Login : global::ProtoBuf.IExtensible
  {
    public CS_Login() {}
    
    private string _token;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"token", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string token
    {
      get { return _token; }
      set { _token = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool tokenSpecified
    {
      get { return this._token != null; }
      set { if (value == (this._token== null)) this._token = value ? this.token : (string)null; }
    }
    private bool ShouldSerializetoken() { return tokenSpecified; }
    private void Resettoken() { tokenSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_Login")]
  public partial class SC_Login : global::ProtoBuf.IExtensible
  {
    public SC_Login() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_ShopAddDiamond")]
  public partial class CS_ShopAddDiamond : global::ProtoBuf.IExtensible
  {
    public CS_ShopAddDiamond() {}
    
    private int? _idx;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"idx", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int? idx
    {
      get { return _idx; }
      set { _idx = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool idxSpecified
    {
      get { return this._idx != null; }
      set { if (value == (this._idx== null)) this._idx = value ? this.idx : (int?)null; }
    }
    private bool ShouldSerializeidx() { return idxSpecified; }
    private void Resetidx() { idxSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_ShopAddDiamond")]
  public partial class SC_ShopAddDiamond : global::ProtoBuf.IExtensible
  {
    public SC_ShopAddDiamond() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_ShopBuyGold")]
  public partial class CS_ShopBuyGold : global::ProtoBuf.IExtensible
  {
    public CS_ShopBuyGold() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_ShopBuyGold")]
  public partial class SC_ShopBuyGold : global::ProtoBuf.IExtensible
  {
    public SC_ShopBuyGold() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_HeroDraw")]
  public partial class CS_HeroDraw : global::ProtoBuf.IExtensible
  {
    public CS_HeroDraw() {}
    
    private int? _num;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"num", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int? num
    {
      get { return _num; }
      set { _num = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool numSpecified
    {
      get { return this._num != null; }
      set { if (value == (this._num== null)) this._num = value ? this.num : (int?)null; }
    }
    private bool ShouldSerializenum() { return numSpecified; }
    private void Resetnum() { numSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_HeroDraw")]
  public partial class SC_HeroDraw : global::ProtoBuf.IExtensible
  {
    public SC_HeroDraw() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_HeroUplevel")]
  public partial class CS_HeroUplevel : global::ProtoBuf.IExtensible
  {
    public CS_HeroUplevel() {}
    
    private int? _heroId;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"heroId", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int? heroId
    {
      get { return _heroId; }
      set { _heroId = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool heroIdSpecified
    {
      get { return this._heroId != null; }
      set { if (value == (this._heroId== null)) this._heroId = value ? this.heroId : (int?)null; }
    }
    private bool ShouldSerializeheroId() { return heroIdSpecified; }
    private void ResetheroId() { heroIdSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_HeroUpLevel")]
  public partial class SC_HeroUpLevel : global::ProtoBuf.IExtensible
  {
    public SC_HeroUpLevel() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_HeroEquip")]
  public partial class CS_HeroEquip : global::ProtoBuf.IExtensible
  {
    public CS_HeroEquip() {}
    
    private int? _heroId;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"heroId", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int? heroId
    {
      get { return _heroId; }
      set { _heroId = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool heroIdSpecified
    {
      get { return this._heroId != null; }
      set { if (value == (this._heroId== null)) this._heroId = value ? this.heroId : (int?)null; }
    }
    private bool ShouldSerializeheroId() { return heroIdSpecified; }
    private void ResetheroId() { heroIdSpecified = false; }
    
    private int? _equipId;
    [global::ProtoBuf.ProtoMember(2, IsRequired = false, Name=@"equipId", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int? equipId
    {
      get { return _equipId; }
      set { _equipId = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool equipIdSpecified
    {
      get { return this._equipId != null; }
      set { if (value == (this._equipId== null)) this._equipId = value ? this.equipId : (int?)null; }
    }
    private bool ShouldSerializeequipId() { return equipIdSpecified; }
    private void ResetequipId() { equipIdSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_HeroEquip")]
  public partial class SC_HeroEquip : global::ProtoBuf.IExtensible
  {
    public SC_HeroEquip() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_ChapterGetPrize")]
  public partial class CS_ChapterGetPrize : global::ProtoBuf.IExtensible
  {
    public CS_ChapterGetPrize() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_ChapterGetPrize")]
  public partial class SC_ChapterGetPrize : global::ProtoBuf.IExtensible
  {
    public SC_ChapterGetPrize() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_ChapterChallenge")]
  public partial class CS_ChapterChallenge : global::ProtoBuf.IExtensible
  {
    public CS_ChapterChallenge() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"SC_ChapterChallenge")]
  public partial class SC_ChapterChallenge : global::ProtoBuf.IExtensible
  {
    public SC_ChapterChallenge() {}
    
    private netproto.E_Code? _code;
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"code", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public netproto.E_Code? code
    {
      get { return _code; }
      set { _code = value; }
    }
    [global::System.Xml.Serialization.XmlIgnore]
    [global::System.ComponentModel.Browsable(false)]
    public bool codeSpecified
    {
      get { return this._code != null; }
      set { if (value == (this._code== null)) this._code = value ? this.code : (netproto.E_Code?)null; }
    }
    private bool ShouldSerializecode() { return codeSpecified; }
    private void Resetcode() { codeSpecified = false; }
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"CS_None")]
  public partial class CS_None : global::ProtoBuf.IExtensible
  {
    public CS_None() {}
    
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
    [global::ProtoBuf.ProtoContract(Name=@"E_Code")]
    public enum E_Code
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_ERROR", Value=0)]
      E_ERROR = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_OK", Value=1)]
      E_OK = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_RELOGIN", Value=2)]
      E_RELOGIN = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_NONE_EXIST", Value=3)]
      E_NONE_EXIST = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_UNKNOWN", Value=4)]
      E_UNKNOWN = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_SERVER_INTERNAL_ERROR", Value=5)]
      E_SERVER_INTERNAL_ERROR = 5,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_INVALID_PARAM", Value=6)]
      E_INVALID_PARAM = 6,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_INVALID_OPT", Value=7)]
      E_INVALID_OPT = 7,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_LIMIT_GOLD", Value=20)]
      E_LIMIT_GOLD = 20,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_LIMIT_DIAMOND", Value=21)]
      E_LIMIT_DIAMOND = 21,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_LIMIT_TIMES", Value=22)]
      E_LIMIT_TIMES = 22,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_LIMIT_CHAPTER_LEVEL", Value=23)]
      E_LIMIT_CHAPTER_LEVEL = 23,
            
      [global::ProtoBuf.ProtoEnum(Name=@"E_LIMIT_PLAYER_LEVEL", Value=24)]
      E_LIMIT_PLAYER_LEVEL = 24
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"EMsgIds")]
    public enum EMsgIds
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_None", Value=0)]
      ECS_None = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_Tick", Value=101)]
      ECS_Tick = 101,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_Tick", Value=102)]
      ESC_Tick = 102,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_Login", Value=103)]
      ECS_Login = 103,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_Login", Value=104)]
      ESC_Login = 104,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_ShopAddDiamond", Value=105)]
      ECS_ShopAddDiamond = 105,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_ShopAddDiamond", Value=106)]
      ESC_ShopAddDiamond = 106,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_ShopBuyGold", Value=107)]
      ECS_ShopBuyGold = 107,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_ShopBuyGold", Value=108)]
      ESC_ShopBuyGold = 108,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_HeroDraw", Value=109)]
      ECS_HeroDraw = 109,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_HeroDraw", Value=110)]
      ESC_HeroDraw = 110,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_HeroUplevel", Value=111)]
      ECS_HeroUplevel = 111,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_HeroUpLevel", Value=112)]
      ESC_HeroUpLevel = 112,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_HeroEquip", Value=113)]
      ECS_HeroEquip = 113,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_HeroEquip", Value=114)]
      ESC_HeroEquip = 114,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_ChapterGetPrize", Value=115)]
      ECS_ChapterGetPrize = 115,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_ChapterGetPrize", Value=116)]
      ESC_ChapterGetPrize = 116,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ECS_ChapterChallenge", Value=117)]
      ECS_ChapterChallenge = 117,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ESC_ChapterChallenge", Value=118)]
      ESC_ChapterChallenge = 118
    }
  
}