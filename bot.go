package main

import (
	"os"
	"strconv"

	"github.com/FloatTech/ZeroBot-Plugin/kanban" // 在最前打印 banner

	// ---------以下插件均可通过前面加 // 注释，注释后停用并不加载插件--------- //
	// ----------------------插件优先级按顺序从高到低---------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------高优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv高优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	//               vvvvvvvvvvvvvv高优先级区vvvvvvvvvvvvvv               //
	//                      vvvvvvv高优先级区vvvvvvv                      //
	//                          vvvvvvvvvvvvvv                          //
	//                               vvvv                               //

	// webctrl "github.com/FloatTech/zbputils/control/web"           // web 后端控制

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chat" // 基础词库

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/sleep_manage" // 统计睡眠时间

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/atri" // ATRI词库

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/manager" // 群管

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/thesaurus" // 词典匹配回复

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^高优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^高优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^高优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------高优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------中优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	//               vvvvvvvvvvvvvv中优先级区vvvvvvvvvvvvvv               //
	//                      vvvvvvv中优先级区vvvvvvv                      //
	//                          vvvvvvvvvvvvvv                          //
	//                               vvvv                               //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/acgimage"       // 随机图片与AI点评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_false"       // 服务器监控
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/aiwife"         // 随机老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/b14"            // base16384加解密
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/baidu"          // 百度一下
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili"       // 查询b站用户信息
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili_parse" // b站视频链接解析
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/book_review"    // 哀伤雪刃吧推书记录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/cangtoushi"     // 藏头诗
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/choose"         // 选择困难症帮手
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/chouxianghua"   // 说抽象话
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/coser"          // 三次元小姐姐
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/cpstory"        // cp短打
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/danbooru"       // DeepDanbooru二次元图标签识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/diana"          // 嘉心糖发病
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/drift_bottle"   // 漂流瓶
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/emojimix"       // 合成emoji
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/fortune"        // 运势
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/funny"          // 笑话
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/gif"            // 制图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/github"         // 搜索GitHub仓库
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/hs"             // 炉石
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/image_finder"   // 关键字搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/jandan"         // 煎蛋网无聊图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/job"            // 定时指令触发器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/juejuezi"       // 绝绝子生成器
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/lolicon"        // lolicon 随机图片
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu"           // 摸鱼
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/moyu_calendar"  // 摸鱼人日历
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"          // 点歌
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativesetu"     // 本地涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nativewife"     // 本地老婆
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nbnhhsh"        // 拼音首字母缩写释义工具
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/novel"          // 铅笔小说网搜索
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/nsfw"           // nsfw图片识别
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/omikuji"        // 浅草寺求签
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/reborn"         // 投胎
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/runcode"        // 在线运行代码
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/saucenao"       // 以图搜图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/scale"          // 叔叔的AI二次元图片放大
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/score"          // 分数
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/setutime"       // 来份涩图
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shadiao"        // 沙雕app
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/shindan"        // 测定
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/tracemoe"       // 搜番
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/translation"    // 翻译
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/vtb_quotation"  // vtb语录
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wangyiyun"      // 网易云音乐热评
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/wordle"         // 猜单词
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ymgal"          // 月幕galgame

	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/wtf"            // 鬼东西
	// _ "github.com/FloatTech/ZeroBot-Plugin/plugin/bilibili_push"  // b站推送

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^中优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^中优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------中优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// ----------------------------低优先级区---------------------------- //
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvvvvvvvvvvvvvvvv //
	//               vvvvvvvvvvvvvv低优先级区vvvvvvvvvvvvvv               //
	//                      vvvvvvv低优先级区vvvvvvv                      //
	//                          vvvvvvvvvvvvvv                          //
	//                               vvvv                               //

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/curse" // 骂人

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/ai_reply" // 人工智能回复

	//                               ^^^^                               //
	//                          ^^^^^^^^^^^^^^                          //
	//                      ^^^^^^^低优先级区^^^^^^^                      //
	//               ^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^               //
	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^低优先级区^^^^^^^^^^^^^^^^^^^^^^^^^^^^ //
	// ----------------------------低优先级区---------------------------- //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	//                                                                  //
	// -----------------------以下为内置依赖，勿动------------------------ //
	"github.com/FloatTech/zbputils/process"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"github.com/wdvxdr1123/ZeroBot/message"

	"github.com/Mrs4s/go-cqhttp/coolq"
	"github.com/Mrs4s/go-cqhttp/modules/servers"
	// -----------------------以上为内置依赖，勿动------------------------ //
)

func init() {
	arg := os.Args
	var qqs []string
	if len(arg) > 1 {
		for _, a := range arg {
			i, err := strconv.ParseUint(a, 10, 64)
			if err == nil {
				qqs = append(qqs, strconv.FormatUint(i, 10))
			}
		}
	}
	driver.RegisterServer(func(s string, f func(driver.CQBot)) {
		servers.RegisterCustom(s, func(c *coolq.CQBot) { f((*CQBot)(c)) })
	})
	driver.NewFuncallClient("zbp", newcaller, func(f *driver.FCClient) {
		// 帮助
		zero.OnFullMatchGroup([]string{"/help", ".help", "菜单"}, zero.OnlyToMe).SetBlock(true).FirstPriority().
			Handle(func(ctx *zero.Ctx) {
				ctx.SendChain(message.Text(kanban.Banner))
			})
		zero.OnFullMatch("查看zbp公告", zero.OnlyToMe, zero.AdminPermission).SetBlock(true).FirstPriority().
			Handle(func(ctx *zero.Ctx) {
				ctx.SendChain(message.Text(kanban.Kanban()))
			})
		zero.Run(
			zero.Config{
				NickName:      []string{"椛椛", "ATRI", "atri", "亚托莉", "アトリ"},
				CommandPrefix: "/",
				// SuperUsers 某些功能需要主人权限，可通过以下两种方式修改
				// SuperUsers: []string{"12345678", "87654321"}, // 通过代码写死的方式添加主人账号
				SuperUsers: qqs, // 通过命令行参数的方式添加主人账号
				Driver:     []zero.Driver{f},
			},
		)
		process.GlobalInitMutex.Unlock()
	})
}
