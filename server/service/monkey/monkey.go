package monkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	urls "net/url"
	"strings"
	"time"

	"github.com/jizi19911101/gin-vue-admin/server/model/monkey"

	"go.uber.org/zap"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	monkeyReq "github.com/jizi19911101/gin-vue-admin/server/model/monkey/request"
)

type MonkeyService struct {
}

func (monkeyService *MonkeyService) StartMonkey(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	//设备是否占用中
	err := monkeyService.checkDevice(startMonkeyReq)
	if err != nil {
		global.GVA_LOG.Error("查询设备是否占用时出错", zap.Error(err))
		return err
	}

	// 占用设备，
	// todo defer释放命令
	err = monkeyService.useDevice(startMonkeyReq)
	if err != nil {
		global.GVA_LOG.Error("占用设备出错", zap.Error(err))
		return err
	}

	// todo 是否请除日志

	// 获取atxAgentAddress等信息
	atxAgentAddress, phoneVersion, err := monkeyService.getAtxAndPhoneInfo(startMonkeyReq)
	if err != nil {
		global.GVA_LOG.Error("获取atxAgentAddress等信息出错", zap.Error(err))
		return err
	}

	// 测试app是否存在
	err = monkeyService.checkAppExist(atxAgentAddress, startMonkeyReq)
	if err != nil {
		global.GVA_LOG.Error("查询测试app是否存在时出错", zap.Error(err))
		return err
	}

	// shell 命令发起monkey测试
	err = monkeyService.startMonkey(atxAgentAddress, startMonkeyReq)
	if err != nil {
		global.GVA_LOG.Error("发起monkey测试时出错", zap.Error(err))
		return err
	}
	beginTime := time.Now().Format("2006-01-02 15:04:05")

	// 生成测试报告
	go monkeyService.generateReport(atxAgentAddress, beginTime, phoneVersion, startMonkeyReq)

	return nil
}

func (monkeyService *MonkeyService) checkDevice(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://120.25.149.119:8082/api/v1/devices/"
	userId := startMonkeyReq.UserId
	device := startMonkeyReq.Device
	resp, err := http.Get(url + device + "?user_id=" + userId)

	if err != nil {
		global.GVA_LOG.Error("checkDevice请求url出错", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("checkDevice http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return errors.New("checkDevice wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("checkDevice读取body出错", zap.Error(err))
		return err

	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("checkDevice反序列化body出错", zap.Error(err))
		return err
	}
	using := bodyMap["device"].(map[string]interface{})["using"]
	if using.(bool) && bodyMap["device"].(map[string]interface{})["userId"] != userId {
		err = errors.New("设备占用中")
		global.GVA_LOG.Error("checkDevice查到设备占用中", zap.Error(err))
		return err
	}
	return nil
}

func (monkeyService *MonkeyService) useDevice(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://120.25.149.119:8082/api/v1/user/devices"
	dataStr := "{\"udid\":\"" + startMonkeyReq.Device + "\"}"
	dataJson := []byte(dataStr)
	req, err := http.NewRequest("POST", url+"?user_id="+startMonkeyReq.UserId, bytes.NewBuffer(dataJson))
	if err != nil {
		global.GVA_LOG.Error("useDevice NewRequest出错", zap.Error(err))
		return errors.New("设备占用失败")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("useDevice发起请求出错", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("useDevice http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return errors.New("useDevice wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("useDevice读取body出错", zap.Error(err))
		return err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("useDevice反序列化body出错", zap.Error(err))
		return err
	}
	success := bodyMap["success"]
	if !success.(bool) {
		err = errors.New("设备占用失败")
		global.GVA_LOG.Error("useDevice设备占用失败", zap.Error(err))

		return err
	}
	return nil
}

func (monkeyService *MonkeyService) getAtxAndPhoneInfo(startMonkeyReq monkeyReq.StartMonkeyReq) (string, string, error) {
	url := "http://120.25.149.119:8082/api/v1/user/devices/"
	resp, err := http.Get(url + startMonkeyReq.Device + "?user_id=" + startMonkeyReq.UserId)
	if err != nil {
		global.GVA_LOG.Error("getAtxAndPhoneInfo发起请求失败", zap.Error(err))
		return "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("getAtxAndPhoneInfo http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return "", "", errors.New("getAtxAndPhoneInfo wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("getAtxAndPhoneInfo读取body失败", zap.Error(err))
		return "", "", err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("getAtxAndPhoneInfo反序列化body失败", zap.Error(err))
		return "", "", err
	}
	atxAgentAddress := bodyMap["device"].(map[string]interface{})["source"].(map[string]interface{})["atxAgentAddress"].(string)
	phoneVersion := bodyMap["device"].(map[string]interface{})["properties"].(map[string]interface{})["version"].(string)
	return atxAgentAddress, phoneVersion, nil
}

func (monkeyService *MonkeyService) checkAppExist(atxAgentAddress string, startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://" + atxAgentAddress + "/shell?user_id=" + startMonkeyReq.UserId + "&command=pm%20list%20packages%20-3"
	resp, err := http.Get(url)

	if err != nil {
		global.GVA_LOG.Error("checkAppExist发起请求失败", zap.Error(err))
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("checkAppExist http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return errors.New("checkAppExist wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("checkAppExist读取body失败", zap.Error(err))
		return err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("checkAppExist反序列化body失败", zap.Error(err))
		return err
	}
	error := bodyMap["error"]
	if error != nil {
		err = errors.New("查询app包列表失败")
		global.GVA_LOG.Error("checkAppExist查询app包列表失败", zap.Error(err))
		return err
	}

	output := bodyMap["output"].(string)
	if !strings.Contains(output, startMonkeyReq.App) {
		err = errors.New("此app不存在，请进行确认")
		global.GVA_LOG.Error("checkAppExist查询到app不存在，请进行确认", zap.Error(err))
		return err
	}
	return nil
}

func (monkeyService *MonkeyService) startMonkey(atxAgentAddress string, startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://" + atxAgentAddress + "/shell/background?user_id=" + startMonkeyReq.UserId + "&command="
	command := "CLASSPATH=/sdcard/monkey.jar:/sdcard/framework.jar exec app_process /system/bin tv.panda.test.monkey.Monkey -p " + startMonkeyReq.App + "  --uiautomatordfs  --running-minutes  " + startMonkeyReq.Duration + "  --throttle 500 -v -v "
	url = url + urls.QueryEscape(command)

	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("startMonkey发起请求失败", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("startMonkey http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return errors.New("startMonkey wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("startMonkey读取body失败", zap.Error(err))
		return err
	}
	bodyMap := make(map[string]interface{}, 0)

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("startMonkey反序列化body失败", zap.Error(err))
		return err
	}
	success := bodyMap["success"]
	if !success.(bool) {
		err = errors.New("启动Monkey失败")
		global.GVA_LOG.Error("startMonkey启动Monkey失败", zap.Error(err))
		return err
	}
	return nil
}
func (monkeyService *MonkeyService) getSubprocess(atxAgentAddress string) (string, error) {
	url := "http://" + atxAgentAddress + "/proc/list"

	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("getSubprocess请求url失败", zap.Error(err))
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("getSubprocess http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return "", errors.New("getSubprocess wrong resp statusCode")
	}

	if resp.StatusCode != 200 {
		err = errors.New("查询测试进程失败")
		global.GVA_LOG.Error("getSubprocess查询测试进程失败", zap.Error(err))
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("getSubprocess读取body失败", zap.Error(err))
		return "", err
	}
	return string(body), nil
}

func (monkeyService *MonkeyService) generateReport(atxAgentAddress string, beginTime string, phoneVersion string, startMonkeyReq monkeyReq.StartMonkeyReq) {

	// 查询测试进程是否结束,10秒查一次

LOOP:
	subprocess, err := monkeyService.getSubprocess(atxAgentAddress)
	if err != nil {
		global.GVA_LOG.Error("generateReport getSubprocess失败", zap.Error(err))
	}

	for strings.Contains(subprocess, "tv.panda.test.monkey") {
		time.Sleep(time.Duration(10) * time.Second)
		goto LOOP

	}
	// 拉取手机的崩溃日志
	logStr, err := monkeyService.pullCrashLog(atxAgentAddress)
	if err != nil {
		global.GVA_LOG.Error("pullCrashLog失败", zap.Error(err))
	}
	// 生成测试报告html
	url := "http://" + atxAgentAddress + "/packages/" + startMonkeyReq.App + "/info"
	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("generateReport请求url失败", zap.Error(err))
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("generateReport http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("generateReport读取body失败", zap.Error(err))
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		global.GVA_LOG.Error("generateReport反序列化body失败", zap.Error(err))
	}
	appName := bodyMap["data"].(map[string]interface{})["label"].(string)
	appVersion := bodyMap["data"].(map[string]interface{})["versionName"].(string)

	// 保存报告到数据库
	report := monkey.MonkeyReport{
		Name:         startMonkeyReq.Report,
		AppName:      appName,
		AppVersion:   appVersion,
		Duration:     startMonkeyReq.Duration,
		BeginTime:    beginTime,
		PhoneSystem:  "安卓",
		PhoneVersion: phoneVersion,
		Log:          logStr,
	}
	err = global.GVA_DB.Create(&report).Error

	if err != nil {
		global.GVA_LOG.Error("generateReport保存报告失败", zap.Error(err))
	}

}

func (monkeyService *MonkeyService) pullCrashLog(atxAgentAddress string) (string, error) {
	url := "http://" + atxAgentAddress + "/raw/sdcard/crash-dump.log"
	resp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("pullCrashLog请求url失败", zap.Error(err))
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		global.GVA_LOG.Error("pullCrashLog http resp statusCode is "+string(resp.StatusCode), zap.Error(err))
		return "", errors.New("pullCrashLog wrong resp statusCode")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("pullCrashLog读取body失败", zap.Error(err))
		return "", err
	}
	return string(body), nil
}

func (monkeyService *MonkeyService) ReportList(info monkeyReq.ReportSearch) (error, []monkey.MonkeyReport, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&monkey.MonkeyReport{})
	var total int64
	reportList := make([]monkey.MonkeyReport, 0)

	if info.Name != "" {
		db.Where("name = ?", info.Name)
	}

	if err := db.Count(&total).Error; err != nil {
		return err, nil, 0
	}
	err := db.Limit(limit).Offset(offset).Order("ID desc").Find(&reportList).Error
	return err, reportList, total

}

func (monkeyService *MonkeyService) ReportContent(htmlReq monkeyReq.HtmlReq) (error, monkey.MonkeyReport) {
	id := htmlReq.ID
	db := global.GVA_DB.Model(&monkey.MonkeyReport{})
	reportContent := monkey.MonkeyReport{}
	err := db.Where("id = ? ", id).Find(&reportContent).Error
	if err != nil {
		return err, reportContent

	}
	return nil, reportContent

}
