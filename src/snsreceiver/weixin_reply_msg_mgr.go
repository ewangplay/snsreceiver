package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type WeixinTextMsg struct {
	Xml WeixinTextMsgInfo `xml:"xml"`
}

type WeixinTextMsgInfo struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}

type WeixinImageMsg struct {
	Xml WeixinImageMsgInfo `xml:"xml"`
}

type WeixinImageMsgInfo struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Image        string `xml:"Image"`
	//Image        WeixinImage `xml:"Image"`
}

type WeixinImage struct {
	MediaId string `xml:"MediaId"`
}

type WeixinVoiceMsg struct {
	Xml WeixinVoiceMsgInfo `xml:"xml"`
}

type WeixinVoiceMsgInfo struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Voice        string `xml:"Voice"`
	//Voice        WeixinVoice `xml:"Voice"`
}

type WeixinVoice struct {
	MediaId string `xml:"MediaId"`
}

type WeixinVideoMsg struct {
	Xml WeixinVideoMsgInfo `xml:"xml"`
}

type WeixinVideoMsgInfo struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Video        string `xml:"Video"`
	//Video        WeixinVideo `xml:"Video"`
}
type WeixinVideo struct {
	MediaId     string `xml:"MediaId"`
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
}

type WeixinMusicMsg struct {
	Xml WeixinMusicMsgInfo `xml:"xml"`
}
type WeixinMusicMsgInfo struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Music        string `xml:"Music"`
	//Music        WeixinVideo `xml:"Music"`
}

type WeixinMusic struct {
	//MediaId string `xml:"MediaId"`
	Title        string `xml:"Title"`
	Description  string `xml:"Description"`
	MusicUrl     string `xmL:"MusicUrl"`
	HQMusicUrl   string `xml:"HQMusicUrl"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}

type WeixinNewsMsg struct {
	Xml WeixinNewsMsgInfo `xml:"xml"`
}
type WeixinNewsMsgInfo struct {
	ToUserName   string     `xml:"ToUserName"`
	FromUserName string     `xml:"FromUserName"`
	CreateTime   int64      `xml:"CreateTime"`
	MsgType      string     `xml:"MsgType"`
	ArticleCount int64      `xml:"ArticleCount"`
	Articles     WeixinNews `xml:"Articles"`
}

type WeixinNews struct {
	Item []WeixinNewsInfo `xml:"item"`
}

type WeixinNewsInfo struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

func SendTextMsg(msg *WeixinPrivateMsg, result map[string]interface{}) error {
	textMsg := &WeixinTextMsg{}
	textMsg.Xml.ToUserName = msg.FromUserName
	rows, err := g_mysqladaptor.QueryFormat(GET_WEIXIN_TEXT_REPLY_MESSAGE_SQL, msg.CustomerId, msg.AccountId)
	if err != nil {
		LOG_ERROR("get weixin text reply message error: %v", err)
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&textMsg.Xml.FromUserName, &textMsg.Xml.CreateTime, &textMsg.Xml.Content)
		if err != nil {
			LOG_ERROR("scan weixin text reply message error: %v", err)
			return err
		}
	}
	textMsg.Xml.MsgType = "text"
	textMsg.Xml.CreateTime = time.Now().Unix()
	x, err := xml.Marshal(textMsg.Xml)
	if err != nil {
		return err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinTextMsgInfo", "xml", -1)
	LOG_DEBUG("echo string: %v", ret)
	result["echo"] = ret

	/*
		ret, err := buildNewsReplyMsg(msg.FromUserName, msg.ToUserName, 1, 1, "")
		if err != nil {
			LOG_ERROR("build news reply msg error: %v", err)
		}
		ret = strings.Replace(ret, "WeixinTextMsgInfo", "xml", -1)
		LOG_DEBUG("echo string: %v", ret)
		result["echo"] = ret
	*/
	return nil
}

func buildNewsReplyMsg(to_user_name, develop_name string, create_time, article_count int64, detail string) (string, error) {
	res := &WeixinNewsMsg{}
	var weixinnews WeixinNews
	var news1 WeixinNewsInfo
	news1.Title = "测试标题"
	news1.Description = "测试描述"
	news1.PicUrl = "http://app.jiuzhilan.com/assets/images/login-logo.png"
	news1.Url = "app.jiuzhilan.com" + "?" + "openid=" + to_user_name
	weixinnews.Item = append(weixinnews.Item, news1)
	res.Xml.Articles = weixinnews
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "news"
	res.Xml.ArticleCount = article_count
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinNewsMsgInfo", "xml", -1)
	return ret, nil
}
func buildImageReplyMsg(to_user_name, develop_name string, create_time int64, detail string) (string, error) {
	res := &WeixinImageMsg{}
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "image"
	res.Xml.Image = detail
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinImageMsgInfo", "xml", -1)
	return ret, nil
}
func buildVoiceReplyMsg(to_user_name, develop_name string, create_time int64, detail string) (string, error) {
	res := &WeixinVoiceMsg{}
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "Voice"
	res.Xml.Voice = detail
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinVoiceMsgInfo", "xml", -1)
	return ret, nil
}
func buildVideoReplyMsg(to_user_name, develop_name string, create_time int64, detail string) (string, error) {
	res := &WeixinVideoMsg{}
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "Video"
	res.Xml.Video = detail
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinVideoMsgInfo", "xml", -1)
	return ret, nil
}

func buildMusicReplyMsg(to_user_name, develop_name string, create_time int64, detail string) (string, error) {
	res := &WeixinMusicMsg{}
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "Music"
	res.Xml.Music = detail
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinMusicMsgInfo", "xml", -1)
	return ret, nil
}

func buildTextReplyMsg(to_user_name, develop_name string, create_time int64, detail string) (string, error) {
	res := &WeixinTextMsg{}
	res.Xml.ToUserName = to_user_name
	res.Xml.FromUserName = develop_name
	res.Xml.CreateTime = create_time
	res.Xml.MsgType = "text"
	res.Xml.Content = detail
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return "", err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinTextMsgInfo", "xml", -1)
	return ret, nil
}

func SendViewMsg(msg *WeixinPrivateMsg, result map[string]interface{}) error {
	return nil
}

func SendClickMsg(msg *WeixinPrivateMsg, result map[string]interface{}) error {
	//从EventKey中读取cid,account_id,关键字的key,并将key字段值对应到关键字，在表中搜索关键字，并得到对应的url
	res := &WeixinNewsMsg{}
	var keyword1, keyword2 string
	var key string
	var article_count int64
	cid := int64(msg.CustomerId)
	account_id := int64(msg.AccountId)
	//获取artile_count
	//获取url
	key = msg.EventKey
	LOG_INFO("cid, account_id, key: %v, %v, %v", cid, account_id, key)
	switch key {
	case "k1":
		keyword1 = "微信"
		keyword2 = "微博"
	case "k2":
		keyword1 = "SEO"
		keyword2 = "SEM"
	case "k3":
		keyword1 = "自动化"
		keyword2 = "分群"
	case "k4":
		keyword1 = "内容"
		keyword2 = "邮件"
	case "k5":
		keyword1 = "流量"
		keyword2 = "流量"
	case "k6":
		keyword1 = "hot"
		keyword2 = "hot"
	default:
		keyword1 = ""
		keyword2 = ""
	}
	keyword1 = "%%" + keyword1 + "%%"
	keyword2 = "%%" + keyword2 + "%%"
	article_count, err := getArticleCount(cid, account_id, keyword1, keyword2)
	if err != nil {
		return err
	}
	LOG_INFO("article count: %v", article_count)
	if article_count == 0 {
		err = SendWelcomeMsg(msg, result)
		if err != nil {
			LOG_ERROR("send welcome message error: %v", err)
			return err
		}
		return nil
	}
	rows, err := g_mysqladaptor.QueryFormat(GET_WEIXIN_NEWS_URL_BY_KEYWORD_SQL, cid, account_id, keyword1, keyword2)
	if err != nil {
		return err
	}
	defer rows.Close()
	var news WeixinNews
	create_time := "1900-01-01 00:00:00"
	for rows.Next() {
		//新建图文消息结构体
		var newsinfo WeixinNewsInfo
		var weixin_news_addr, weixin_name, img_url string
		err = rows.Scan(&weixin_news_addr, &weixin_name, &create_time, &img_url)
		if err != nil {
			return err
		}
		newsinfo.Title = weixin_name
		newsinfo.Url = weixin_news_addr
		newsinfo.PicUrl = img_url
		LOG_INFO("news info: title, url, img_url:%v,%v, %v", newsinfo.Title, newsinfo.Url, newsinfo.PicUrl)
		news.Item = append(news.Item, newsinfo)
	}
	LOG_DEBUG("Articles: %v", news)
	res.Xml.Articles = news
	res.Xml.ToUserName = msg.FromUserName
	res.Xml.FromUserName = "jiuzhilan"
	//res.Xml.FromUserName = "gh_3ec16dc93702"
	tm, _ := time.Parse("2006-01-02 15:04:05", create_time)
	tmInt := tm.Unix()
	res.Xml.CreateTime = tmInt
	res.Xml.MsgType = "news"
	res.Xml.ArticleCount = article_count
	LOG_DEBUG("res: %v", res)
	x, err := xml.Marshal(res.Xml)
	if err != nil {
		return err
	}
	ret := string(x)
	ret = strings.Replace(ret, "WeixinNewsMsgInfo", "xml", -1)
	result["echo"] = ret
	LOG_DEBUG("echo string: %v", ret)
	return nil
}

func getArticleCount(cid, account_id int64, keyword1, keyword2 string) (int64, error) {
	var count int64
	row, err := g_mysqladaptor.QueryFormat(GET_WEIXIN_NEWS_URL_BY_KEYWORD_COUNT_SQL, cid, account_id, keyword1, keyword2)
	if err != nil {
		return count, err
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return count, err
		}
	}
	return count, nil
}

func SendWelcomeMsg(msg *WeixinPrivateMsg, result map[string]interface{}) error {
	//从表中取出欢迎信息，存到result["echo"]字段中
	row, err := g_mysqladaptor.QueryFormat(WEIXIN_AUTO_REPLY_WELCOME_SQL, msg.CustomerId, msg.AccountId)
	if err != nil {
		return err
	}
	defer row.Close()
	echoStr, err := buildReplyMsg(row, msg)
	if err != nil {
		LOG_ERROR("build reply msg err : %v", err)
		return err
	}
	fmt.Println("echo string: ", echoStr)
	LOG_DEBUG("echo string: %v", echoStr)
	result["echo"] = echoStr

	return nil
}

func buildReplyMsg(row *sql.Rows, msg *WeixinPrivateMsg) (string, error) {
	var ret string
	if row.Next() {
		var msg_type, detail, develop_name string
		var create_time, article_count int64
		err := row.Scan(&msg_type, &create_time, &article_count, &detail, &develop_name)
		if err != nil {
			return "", err
		}
		defer row.Close()

		switch msg_type {
		case "text":
			ret, err = buildNewsReplyMsg(msg.FromUserName, develop_name, create_time, article_count, detail)
		case "news":
			ret, err = buildNewsReplyMsg(msg.FromUserName, develop_name, create_time, article_count, detail)
		case "image":
			ret, err = buildImageReplyMsg(msg.FromUserName, develop_name, create_time, detail)
		case "voice":
			ret, err = buildVoiceReplyMsg(msg.FromUserName, develop_name, create_time, detail)
		case "video":
			ret, err = buildVideoReplyMsg(msg.FromUserName, develop_name, create_time, detail)
		case "music":
			ret, err = buildMusicReplyMsg(msg.FromUserName, develop_name, create_time, detail)
		default:
			return "", fmt.Errorf("Unsupported message type.")
		}
		if err != nil {
			LOG_ERROR("build %v reply msg error: %v", msg_type, err)
			return ret, err
		}

	}
	return ret, nil
}
