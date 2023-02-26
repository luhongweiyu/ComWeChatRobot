package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// // login check
// WECHAT_IS_LOGIN := 0                         // 登录检查				//: {},

// // self info
// WECHAT_GET_SELF_INFO := 1                    // 获取个人信息			//: {},

// // send message
// WECHAT_MSG_SEND_TEXT := 2                    // 发送文本				//: {"wxid": "","msg": ""},
// //# wxids需要以`,`分隔，例如`wxid1,wxid2,wxid3```
// WECHAT_MSG_SEND_AT := 3                      // 发送群艾特				//: {"chatroom_id":"","wxids": "","msg": "", "auto_nickname": 1},
// WECHAT_MSG_SEND_CARD := 4                    // 分享好友名片			//: {"receiver":"","shared_wxid":"","nickname":""},
// WECHAT_MSG_SEND_IMAGE := 5                   // 发送图片				//WECHAT_MSG_SEND_IMAGE: {"receiver":"","img_path":""},
// WECHAT_MSG_SEND_FILE := 6                    // 发送文件				//WECHAT_MSG_SEND_FILE: {"receiver":"","file_path":""},
// WECHAT_MSG_SEND_ARTICLE := 7                 // 发送xml文章				//: {"wxid":"","title":"","abstract":"","url":"","img_path":""},
// WECHAT_MSG_SEND_APP := 8                     // 发送小程序				//: {"wxid":"","appid":""},

// // receive message
// WECHAT_MSG_START_HOOK := 9                   // 开启接收消息HOOK，只支持socket监听		//: {"port": 10808},
// WECHAT_MSG_STOP_HOOK := 10                   // 关闭接收消息HOOK		//: {},
// WECHAT_MSG_START_IMAGE_HOOK := 11            // 开启图片消息HOOK		//: {"save_path":""},
// WECHAT_MSG_STOP_IMAGE_HOOK := 12             // 关闭图片消息HOOK		//: {},
// WECHAT_MSG_START_VOICE_HOOK := 13            // 开启语音消息HOOK		//: {"save_path":""},
// WECHAT_MSG_STOP_VOICE_HOOK := 14             // 关闭语音消息HOOK		//: {},

// // contact
// WECHAT_CONTACT_GET_LIST := 15                // 获取联系人列表			//: {},
// WECHAT_CONTACT_CHECK_STATUS := 16            // 检查是否被好友删除		//: {"wxid":""},
// WECHAT_CONTACT_DEL := 17                     // 删除好友				//: {"wxid":""},
// WECHAT_CONTACT_SEARCH_BY_CACHE := 18         // 从内存中获取好友信息	//: {"wxid":""},
// WECHAT_CONTACT_SEARCH_BY_NET := 19           // 网络搜索用户信息		//: {"keyword":""},
// WECHAT_CONTACT_ADD_BY_WXID := 20             // wxid加好友				//: {"wxid":"","msg":""},
// WECHAT_CONTACT_ADD_BY_V3 := 21               // v3数据加好友			//: {"v3":"","msg":"","add_type": 0x6},
// WECHAT_CONTACT_ADD_BY_PUBLIC_ID := 22        // 关注公众号				//: {"public_id":""},
// WECHAT_CONTACT_VERIFY_APPLY := 23            // 通过好友请求			//: {"v3":"","v4":""},
// WECHAT_CONTACT_EDIT_REMARK := 24             // 修改备注				//: {"wxid":"","remark":""},

// // chatroom
// WECHAT_CHATROOM_GET_MEMBER_LIST := 25        // 获取群成员列表			//: {"chatroom_id":""},
// WECHAT_CHATROOM_GET_MEMBER_NICKNAME := 26    // 获取指定群成员昵称		//: {"chatroom_id":"","wxid":""},
// 												//# wxids需要以`,`分隔，例如`wxid1,wxid2,wxid3`
// WECHAT_CHATROOM_DEL_MEMBER := 27             // 删除群成员				//: {"chatroom_id":"","wxids":""},
// //# wxids需要以`,`分隔，例如`wxid1,wxid2,wxid3`
// WECHAT_CHATROOM_ADD_MEMBER := 28             // 添加群成员				//: {"chatroom_id":"","wxids":""},
// WECHAT_CHATROOM_SET_ANNOUNCEMENT := 29       // 设置群公告				//: {"chatroom_id":"","announcement":""},
// WECHAT_CHATROOM_SET_CHATROOM_NAME := 30      // 设置群聊名称			//: {"chatroom_id":"","chatroom_name":""},
// WECHAT_CHATROOM_SET_SELF_NICKNAME := 31      // 设置群内个人昵称		//: {"chatroom_id":"","nickname":""},

// // database
// WECHAT_DATABASE_GET_HANDLES := 32            // 获取数据库句柄			//: {},
// WECHAT_DATABASE_BACKUP := 33                 // 备份数据库				//: {"db_handle":0,"save_path":""},
// WECHAT_DATABASE_QUERY := 34                  // 数据库查询				//: {"db_handle":0,"sql":""},

// // version
// WECHAT_SET_VERSION := 35                     // 修改微信版本号			//: {"version": "3.7.0.30"},

// // log
// WECHAT_LOG_START_HOOK := 36                  // 开启日志信息HOOK		//: {},
// WECHAT_LOG_STOP_HOOK := 37                   // 关闭日志信息HOOK		//: {},

// // browser
// WECHAT_BROWSER_OPEN_WITH_URL := 38           // 打开微信内置浏览器		//: {"url": "https://www.baidu.com/"},
// WECHAT_GET_PUBLIC_MSG := 39                  // 获取公众号历史消息		//: {"public_id": "","offset": ""},

// WECHAT_MSG_FORWARD_MESSAGE := 40             // 转发消息				//: {"wxid": "filehelper","msgid": 2 ** 64 - 1},
// WECHAT_GET_QRCODE_IMAGE := 41                // 获取二维码				//: {},
// WECHAT_GET_A8KEY := 42                       // 获取A8Key				//: {"url":""},
// WECHAT_MSG_SEND_XML := 43                    // 发送xml消息				//: {"wxid":"filehelper","xml":"","img_path":""},
// WECHAT_LOGOUT := 44                          // 退出登录				//: {},
// WECHAT_GET_TRANSFER := 45                    // 收款					//: {"wxid":"","transcationid":"","transferid":""},
// WECHAT_MSG_SEND_EMOTION := 46                // 发送表情				//: {"wxid":"","img_path":""},
// WECHAT_GET_CDN := 47                         // 下载文件、视频、图片	//: {"msgid":2 ** 64 - 1},
var URL消息监听端口 string = "8687"
var URL发送指令端口 string = "8686"
var URL外部访问地址 string = "127.0.0.1:8080"

func 发送消息(id string, msg string) {
	var website string = fmt.Sprintf("http://127.0.0.1:%v/api/?type=%v", URL发送指令端口, 2)
	b, err := json.Marshal(map[string]string{"wxid": id, "msg": msg})
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}
	resp, err := http.Post(website, "application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Cannot connect the server:", err)
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println("HTML content:", string(body))
	} else {
		fmt.Println("Cannot read from connected http server:", err)
	}
}

func 消息处理(消息内容 string) {
	type m2 struct {
		Message string `json:"message"`
		Wxid    string `json:"wxid"`
		Sender  string `json:"sender"`
	}
	m := m2{"1", "2", "3"}
	err := json.Unmarshal([]byte(消息内容), &m)
	if err != nil {
		fmt.Println("消息错误:", 消息内容, err)
		return
	}
	fmt.Println(m)
	if m.Message == "查询wxid" && m.Sender != "" {
		fmt.Println(m.Sender)
		发送消息(m.Sender, m.Sender)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024 * 10]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
		消息处理(recvStr)
	}
}

func 启动消息监听() {
	listen, err := net.Listen("tcp", "127.0.0.1:"+URL消息监听端口)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
func 启动http() {
	r := gin.Default()
	r.POST("/msg", func(c *gin.Context) {
		target := c.DefaultPostForm("target", "")
		msg := c.DefaultPostForm("msg", "")
		fmt.Printf("目标:%-25v,消息:%v\n", target, msg)
		if target == "" || msg == "" {
			fmt.Println("目标或者内容错误")
			return
		}
		发送消息(target, msg)
	})
	r.Run(URL外部访问地址) // 监听并在 0.0.0.0:8080 上启动服务

}
func 启动定时获取数据库指令() {
	var website string = "http://www.future.org.cn"
	if resp, err := http.Get(website); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println("HTML content:", string(body))
		} else {
			fmt.Println("Cannot read from connected http server:", err)
		}
	} else {
		fmt.Println("Cannot connect the server:", err)
	}

}
func main() {
	go 启动消息监听()
	go 启动http()
	// go 启动定时获取数据库指令()
	// 监听消息
	// 获取网页消息指令
	// 获取http消息指令
	for {
		time.Sleep(time.Second)

	}
}
