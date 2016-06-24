/*
 1. GET request format
 http://sns.jiuzhilan.com/v1/weixin/token?
     nonce=57155157
     &timestamp=1397022061823
     &echostr=dnPdpTZz85
     &signature=90e4c22c90a58f26526c2dd5b6c56c8822edeaa1

 2. POST request format
 http://sns.jiuzhilan.com/v1/weixin/token?
     nonce=57155157
     &timestamp=1397022061823
     &signature=90e4c22c90a58f26526c2dd5b6c56c8822edeaa1

=======================================================================
1. 推送的文本消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1348831860</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[this is a test]]></Content>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         text
Content         文本消息内容
MsgId           消息id，64位整型

=======================================================================
2. 推送图片消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1348831860</CreateTime>
    <MsgType><![CDATA[image]]></MsgType>
    <PicUrl><![CDATA[this is a url]]></PicUrl>
    <MediaId><![CDATA[media_id]]></MediaId>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         image
PicUrl          图片链接
MediaId         图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
MsgId           消息id，64位整型

=======================================================================
3. 推送语音消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1357290913</CreateTime>
    <MsgType><![CDATA[voice]]></MsgType>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Format><![CDATA[Format]]></Format>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         语音为voice
MediaId         语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
Format          语音格式，如amr，speex等
MsgID           消息id，64位整型

=======================================================================
4. 推送视频消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1357290913</CreateTime>
    <MsgType><![CDATA[video]]></MsgType>
    <MediaId><![CDATA[media_id]]></MediaId>
    <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         视频为video
MediaId         视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
ThumbMediaId    视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
MsgId           消息id，64位整型

=======================================================================
5. 推送地理位置消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1351776360</CreateTime>
    <MsgType><![CDATA[location]]></MsgType>
    <Location_X>23.134521</Location_X>
    <Location_Y>113.358803</Location_Y>
    <Scale>20</Scale>
    <Label><![CDATA[位置信息]]></Label>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         location
Location_X      地理位置维度
Location_Y      地理位置经度
Scale           地图缩放大小
Label           地理位置信息
MsgId           消息id，64位整型

=======================================================================
6. 推送链接消息
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1351776360</CreateTime>
    <MsgType><![CDATA[link]]></MsgType>
    <Title><![CDATA[公众平台官网链接]]></Title>
    <Description><![CDATA[公众平台官网链接]]></Description>
    <Url><![CDATA[url]]></Url>
    <MsgId>1234567890123456</MsgId>
</xml>

说明：
ToUserName      接收方微信号
FromUserName    发送方微信号，若为普通用户，则是一个OpenID
CreateTime      消息创建时间
MsgType         消息类型，link
Title           消息标题
Description     消息描述
Url             消息链接
MsgId           消息id，64位整型

=======================================================================
7. 推送关注/取消关注事件
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe]]></Event>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，subscribe(订阅)、unsubscribe(取消订阅)

=======================================================================
8. 推送带参数二维码事件
用户扫描带场景值二维码时，可能推送以下两种事件：
如果用户还未关注公众号，则用户可以关注公众号，关注后微信会将带场景值关注事件推送给开发者。
如果用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者。
(1) 用户未关注时，进行关注后的事件推送
<xml><ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe]]></Event>
    <EventKey><![CDATA[qrscene_123123]]></EventKey>
    <Ticket><![CDATA[TICKET]]></Ticket>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，subscribe
EventKey        事件KEY值，qrscene_为前缀，后面为二维码的参数值
Ticket          二维码的ticket，可用来换取二维码图片

(2) 用户已关注时的事件推送
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[SCAN]]></Event>
    <EventKey><![CDATA[SCENE_VALUE]]></EventKey>
    <Ticket><![CDATA[TICKET]]></Ticket>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，SCAN
EventKey        事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
Ticket          二维码的ticket，可用来换取二维码图片

=======================================================================
9. 推送上报地理位置事件
用户同意上报地理位置后，每次进入公众号会话时，都会在进入时上报地理位置，或在进入会话后每5秒上报一次地理位置，公众号可以在公众平台网站中修改以上设置。上报地理位置时，微信会将上报地理位置事件推送到开发者填写的URL。

<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[LOCATION]]></Event>
    <Latitude>23.137466</Latitude>
    <Longitude>113.352425</Longitude>
    <Precision>119.385040</Precision>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，LOCATION
Latitude        地理位置纬度
Longitude       地理位置经度
Precision       地理位置精度

=======================================================================
10. 推送点击菜单拉取消息时的事件
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[CLICK]]></Event>
    <EventKey><![CDATA[EVENTKEY]]></EventKey>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，CLICK
EventKey        事件KEY值，与自定义菜单接口中KEY值对应

=======================================================================
11. 推送点击菜单跳转链接时的事件
<xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[VIEW]]></Event>
    <EventKey><![CDATA[www.qq.com]]></EventKey>
</xml>

说明：
ToUserName      开发者微信号
FromUserName    发送方帐号（一个OpenID）
CreateTime      消息创建时间 （整型）
MsgType         消息类型，event
Event           事件类型，VIEW
EventKey        事件KEY值，设置的跳转URL

*/

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

const (
	WEIXIN_CHECK_TOKEN_SQL                   string = "SELECT cid, account_id FROM jzl_weixin_account WHERE token=? AND is_delete=0"
	WEIXIN_AUTO_REPLY_WELCOME_SQL            string = "SELECT msg_type, create_time, article_count, detail, develop_name from jzl_weixin_auto_reply where cid=? AND account_id=? AND is_welcome=1"
	GET_WEIXIN_USER_NAME_SQL                 string = "SELECT name FROM jzl_weixin_account WHERE cid=? AND account_id=? AND is_delete=0"
	GET_WEIXIN_NEWS_URL_BY_KEYWORD_SQL       string = "SELECT DISTINCT weixin_news_addr, weixin_name, create_time, thumb_img_url FROM jzl_weixin_news_search WHERE cid=? AND account_id=? AND (weixin_tags LIKE ? OR weixin_tags LIKE ?) ORDER BY create_time DESC LIMIT 6"
	GET_WEIXIN_NEWS_URL_BY_KEYWORD_COUNT_SQL string = "SELECT COUNT(1) FROM (SELECT DISTINCT weixin_news_addr, weixin_name, create_time, thumb_img_url FROM jzl_weixin_news_search WHERE cid=? AND account_id=? AND (weixin_tags LIKE ? OR weixin_tags LIKE ?) ORDER BY create_time DESC LIMIT 6) AS a"
	GET_WEIXIN_TEXT_REPLY_MESSAGE_SQL        string = "SELECT develop_name, create_time,detail FROM jzl_weixin_auto_reply WHERE cid=? AND account_id=? AND msg_type='text' AND is_welcome=0"
)

type WeixinPrivateMsg struct {
	CustomerId   int
	AccountId    int
	ToUserName   string  `xml:"ToUserName"`
	FromUserName string  `xml:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	Content      string  `xml:"Content"`
	MsgId        string  `xml:"MsgId"`
	PicUrl       string  `xml:"PicUrl"`
	MediaId      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaId string  `xml:"ThumbMediaId"`
	Location_X   float64 `xml:"Location_X"`
	Location_Y   float64 `xml:"Location_Y"`
	Scale        int     `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	Url          string  `xml:"Url"`
	Event        string  `xml:"Event"`
	EventKey     string  `xml:"EventKey"`
	Ticket       string  `xml:"Ticket"`
}

type WeixinProcessor struct {
}

func (this *WeixinProcessor) ProcessFunc(version int, resource string, method string, token string, params map[string]string, body []byte, result map[string]interface{}) error {

	if !CheckAuth(token, params) {
		LOG_ERROR("微信认证检查失败")
		//return fmt.Errorf("Check Auth Fail")
	}

	switch method {
	case "GET":
		//微信服务器通过GET请求建立连接
		var echostr string
		echostr, ok := params["echostr"]
		if !ok {
			LOG_ERROR("请求中不包含echostr信息")
			return fmt.Errorf("no echostr info")
		}
		if echostr == "" {
			LOG_ERROR("echostr信息为空")
			return fmt.Errorf("echostr info is null")
		}

		LOG_DEBUG("echostr: %v", echostr)

		LOG_INFO("签名信息验证通过，回复%v跟微信服务器建立连接", echostr)

		//所有验证通过，回复echostr建立连接
		result["echo"] = echostr

	case "POST":
		//根据token信息查找私信对应的cid和account_id
		cid, acc_id, err := this.CheckToken(token)
		if err != nil {
			LOG_ERROR("检查token失败: %v", err)
			return err
		}

		//对接收到微信私信进行解析
		msg := &WeixinPrivateMsg{
			CustomerId: cid,
			AccountId:  acc_id,
		}
		err = xml.Unmarshal(body, msg)
		if err != nil {
			LOG_ERROR("解析消息体失败: %v", err)
			return err
		}

		//处理后的微信消息做JSON编码
		json_msg, err := json.Marshal(msg)
		if err != nil {
			LOG_ERROR("对微信私信做JSON编码失败: %v", err)
			return err
		}

		topic := "WeixinPrivateMsg"
		err = g_nsqProducer.Publish(topic, json_msg)
		if err != nil {
			LOG_ERROR("发布消息[%v]到NSQ主题[%v]失败，失败原因：%v", string(json_msg), topic, err)
			return err
		}

		LOG_INFO("发布消息[%v]到NSQ主题[%v]成功", string(json_msg), topic)

		//假如开发者无法保证在五秒内处理并回复，可以直接回复空串，
		//微信服务器不会对此作任何处理，并且不会发起重试
		//当MsgType为event，并且Event类型为subscribe时，向服务器推送欢迎消息
		if msg.MsgType == "text" {
			//返回关键字回复
			err = SendTextMsg(msg, result)
			if err != nil {
				LOG_ERROR("回复推送类型为text的消息失败。")
				return err
			}
			LOG_INFO("回复推送类型为text的消息成功。")
		}
		if msg.MsgType == "event" && msg.Event == "VIEW" {
			//响应内部账号的菜单点击跳转
			err = SendViewMsg(msg, result)
			if err != nil {
				LOG_DEBUG("回复用户关注跳转信息失败： %v", err)
			}
			LOG_INFO("回复用户关注跳转信息成功。")
		}
		if msg.MsgType == "event" && msg.Event == "CLICK" {
			//响应内部账号的菜单点击跳转
			err = SendClickMsg(msg, result)
			if err != nil {
				LOG_DEBUG("回复用户关注点击信息失败： %v", err)
			}
			LOG_INFO("回复用户关注点击信息成功。")
		}
		if msg.MsgType == "event" && msg.Event == "subscribe" {
			err = SendWelcomeMsg(msg, result)
			if err != nil {
				LOG_ERROR("回复用户关注欢迎信息失败： %v", err)
				return err
			}
			LOG_INFO("回复用户关注欢迎信息成功。")

		} /*else {
			result["echo"] = ""
		}
		*/
	default:
		//其它请求不处理
		LOG_INFO("接收到不符合要求的请求[%v]，忽略掉", method)
		return nil
	}

	return nil
}

func (this *WeixinProcessor) CheckToken(token string) (customer_id, account_id int, err error) {
	row, err := g_mysqladaptor.QueryFormat(WEIXIN_CHECK_TOKEN_SQL, token)
	if err != nil {
		return
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&customer_id, &account_id)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("token %v invalid", token)
		return
	}

	return
}
