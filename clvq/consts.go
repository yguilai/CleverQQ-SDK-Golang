package clvq

// 更新于: 2019年7月30日 15:23:27 燕归来

//常量设计来自 铃兰:1774902807
//前缀MT=MessageType=消息类型
//中间HE=HandleEvent=处理事件
//中间F=Friend=好友
//中间G=Group=群
//中间QQ=QQ (抚脸)
//const int中缩写定义：
//SB=SomeBody 某人
//OT=Online Type 在线状态
//SP=ShowPicture 修图
//EF=EFFECT 特效

///常用消息类型
const (
	MT_UNDEFINED    = -1 + iota //未定义:-1
	MT_ONLINETEMP               //在线状态临时会话: 0
	MT_FRIEND                   //好友:1
	MT_GROUP                    //群消息:2
	MT_DISGROUP                 //讨论组:3
	MT_GROUPTP                  //群临时:4
	MT_DISGROUPTP               //讨论组临时:5
	MT_TENPAY                   //财付通:6
	MT_FRIENDVERIFY             //好友验证回复会话:7 （Pro版可用）
)

///处理事件HandleEvent
const (
	MT_HE_AGREE  = 10 + (iota * 10) //同意:10
	MT_HE_REFUSE                    //拒绝:20
	MT_HE_IGNORE                    //忽略:30
)

///好友事件
const (
	MT_F_SINGLE               = 100 + iota //被单项添加为好友: 100
	MT_F_SBADDME                           //某人请求加为好友：101
	MT_F_AGREEADDME                        //被同意加为好友：102
	MT_F_IGNOREADDSB                       //被拒绝加为好友：103
	MT_F_DELETED                           //被删除好友:104
	MT_F_OFFLINE_FILE_RECEIVE              //好友离线文件接收:105
	MT_F_SIGNCHANGE                        //好友签名变更:106
)

///群事件
const (
	MT_G_SBQUITGROUP             = 201 + iota //某人退出群：201
	MT_G_ADMINkICKSB                          //某人被管理移除群：202
	MT_G_SBISSHUTUP                           //对象被禁言：203
	MT_G_SBREMOVESHUTUP                       //对象被解除禁言：204
	MT_G_OPENALLGROUPSHUTUP                   //开启全群禁言：205
	MT_G_CLOSEALLGROUPSHUTUP                  //关闭全群禁言：206
	MT_G_OPENANONYMOUSCHAT                    //开启匿名聊天：207
	MT_G_CLOSEANONYMOUTSCHAT                  //关闭匿名聊天：208
	MT_G_NOTICEISCHANGE                       //群公告变动：209
	MT_G_SBBECOMEADMIN                        //某人成为管理：210
	MT_G_SBOUTGOINGADMIN                      //某人被取消管理：211
	MT_G_SBAPPROVALGROUP                      //某人被批准加入了群：212
	MT_G_SBWANTTOJOINGROUP                    //请求入群：213
	MT_G_SBINVITATIONMEJOINGROUP              //被邀请加入群：214
	MT_G_SBINVITATIONSBJOINGROUP              //某人被邀请加入群：215
	MT_G_GROUPISDISSOLVE                      //某群被解散：216
	MT_G_GROUPCARDCHANGE                      //群名片变动: 217
	MT_G_GROUPFILE                            //群文件接收:218
)

// 框架QQ
const (
	MT_QQ_ADD        = 1100 + iota //列表添加了新账号: 1100 （Pro可用）
	MT_QQ_LOGIN                    //QQ登录完成: 1101
	MT_QQ_OFFLINEACT               //QQ被手动离线: 1102 （Pro可用）
	MT_QQ_OFFLINEPAS               //QQ被强制离线: 1103 （Pro可用）
	MT_QQ_DROPLINE                 //QQ掉线: 1104 （Pro可用）
	MT_QQ_2DCC                     //QQ二次数据缓存完成: 1105 （Pro可用）
)

// 秀图 （Pro可用）
const (
	SP_EF_NO         = "0" //秀图无特效
	SP_EF_SHAKE      = "1" //秀图抖动特效
	SP_EF_PHANTOM    = "2" //秀图幻影特效
	SP_EF_BIRTHDAY   = "3" //秀图生日特效
	SP_EF_LOVEU      = "4" //秀图爱你特效
	SP_EF_ADDFRIENDS = "5" //秀图征友特效
)

///队列
const (
	MT_IGNORE    = iota //忽略：0
	MT_CONTINUE         //继续：1
	MT_INTERCEPT        //拦截：2
)

///插件
const (
	MT_P_LOAD    = 12000 + iota //本插件载入：12000
	MT_P_ENABLE                 //用户启用本插件：12001
	MT_P_DISABLE                //用户禁用本插件：12002 无权拒绝
	MT_P_CLICKED                //插件被点击: 12003 点击方式参考子类型.  1=左键单击 2=右键单击
)

// 在线状态
const (
	OT_ONLINE    = 1 + iota //我在线上: 1
	OT_QME                  //Q我把: 2
	OT_LEAVE                //离开: 3
	OT_BUSY                 //忙碌: 4
	OT_NODISTURB            //请勿打扰: 5
	OT_INVISIBLE            //隐身: 6
)
