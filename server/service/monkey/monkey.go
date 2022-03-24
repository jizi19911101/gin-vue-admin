package monkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	urls "net/url"
	"strings"
	"text/template"
	"time"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	monkeyReq "github.com/jizi19911101/gin-vue-admin/server/model/monkey/request"
)

type MonkeyService struct {
}

func (monkeyService *MonkeyService) StartMonkey(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	//设备是否占用中
	err := monkeyService.checkDevice(startMonkeyReq)
	if err != nil {
		return errors.New("设备已被占用")
	}

	// 占用设备，
	// todo defer释放命令
	err = monkeyService.useDevice(startMonkeyReq)
	if err != nil {
		return errors.New("设备占用失败")
	}

	// todo 是否请除日志

	// 获取atxAgentAddress等信息
	atxAgentAddress, phoneVersion, err := monkeyService.getAtxAndPhoneInfo(startMonkeyReq)
	if err != nil {
		return errors.New("获取atxAgentAddress等信息失败")
	}

	// 测试app是否存在
	err = monkeyService.checkAppExist(atxAgentAddress, startMonkeyReq)
	if err != nil {
		return errors.New("查询测试app过程出错")
	}

	// shell 命令发起monkey测试
	err = monkeyService.startMonkey(atxAgentAddress, startMonkeyReq)
	if err != nil {
		return errors.New("启动Monkey失败")
	}
	beginTime := time.Now().Format("2006-01-02 15:04:05")

	// 查询测试进程是否结束,10秒查一次

	//LOOP:
	//	subprocess, err := monkeyService.getSubprocess(atxAgentAddress)
	//	if err != nil {
	//		return err
	//	}
	//
	//	for strings.Contains(subprocess, "tv.panda.test.monkey") {
	//		time.Sleep(time.Duration(10) * time.Second)
	//		goto LOOP
	//
	//	}

	// 生成测试报告
	report, err := monkeyService.generateReport(atxAgentAddress, beginTime, phoneVersion, startMonkeyReq)
	if err != nil {
		return err
	}
	fmt.Println(report)

	return nil
}

func (monkeyService *MonkeyService) checkDevice(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://120.25.149.119:8082/api/v1/devices/"
	userId := startMonkeyReq.UserId
	device := startMonkeyReq.Device
	resp, err := http.Get(url + device + "?user_id=" + userId)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	using := bodyMap["device"].(map[string]interface{})["using"]
	if using.(bool) && bodyMap["device"].(map[string]interface{})["userId"] != userId {
		return errors.New("设备占用中")
	}
	return nil
}

func (monkeyService *MonkeyService) useDevice(startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://120.25.149.119:8082/api/v1/user/devices"
	dataStr := "{\"udid\":\"" + startMonkeyReq.Device + "\"}"
	dataJson := []byte(dataStr)
	req, err := http.NewRequest("POST", url+"?user_id="+startMonkeyReq.UserId, bytes.NewBuffer(dataJson))
	if err != nil {
		return errors.New("设备占用失败")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return errors.New("设备占用失败")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	success := bodyMap["success"]
	if !success.(bool) {
		return errors.New("设备占用失败")
	}
	return nil
}

func (monkeyService *MonkeyService) getAtxAndPhoneInfo(startMonkeyReq monkeyReq.StartMonkeyReq) (string, string, error) {
	url := "http://120.25.149.119:8082/api/v1/user/devices/"
	resp, err := http.Get(url + startMonkeyReq.Device + "?user_id=" + startMonkeyReq.UserId)
	defer resp.Body.Close()

	if err != nil {
		return "", "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
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
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	error := bodyMap["error"]
	if error != nil {
		return errors.New("查询app包列表失败")
	}

	output := bodyMap["output"].(string)
	if !strings.Contains(output, startMonkeyReq.App) {
		return errors.New("此app不存在，请进行确认")
	}
	return nil
}

func (monkeyService *MonkeyService) startMonkey(atxAgentAddress string, startMonkeyReq monkeyReq.StartMonkeyReq) error {
	url := "http://" + atxAgentAddress + "/shell/background?user_id=" + startMonkeyReq.UserId + "&command="
	command := "CLASSPATH=/sdcard/monkey.jar:/sdcard/framework.jar exec app_process /system/bin tv.panda.test.monkey.Monkey -p " + startMonkeyReq.App + "  --uiautomatordfs  --running-minutes  " + startMonkeyReq.Duration + "  --throttle 500 -v -v "
	url = url + urls.QueryEscape(command)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap := make(map[string]interface{}, 0)

	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	success := bodyMap["success"]
	if !success.(bool) {
		return errors.New("启动Monkey失败")
	}
	return nil
}
func (monkeyService *MonkeyService) getSubprocess(atxAgentAddress string) (string, error) {
	url := "http://" + atxAgentAddress + "/proc/list"

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	if resp.Status != "200 OK" {
		return "", errors.New("查询测试进程失败")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (monkeyService *MonkeyService) generateReport(atxAgentAddress string, beginTime string, phoneVersion string, startMonkeyReq monkeyReq.StartMonkeyReq) (string, error) {
	htmlPath := global.RedirectConfigFile("tpl.html")
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		return "", err
	}

	url := "http://" + atxAgentAddress + "/packages/" + startMonkeyReq.App + "/info"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return "", err
	}
	appName := bodyMap["data"].(map[string]interface{})["label"].(string)
	appVersion := bodyMap["data"].(map[string]interface{})["versionName"].(string)

	data := struct {
		AppName      string
		AppVersion   string
		Duration     string
		BeginTime    string
		PhoneSystem  string
		PhoneVersion string
	}{
		AppName:      appName,
		AppVersion:   appVersion,
		Duration:     startMonkeyReq.Duration,
		BeginTime:    beginTime,
		PhoneSystem:  "安卓",
		PhoneVersion: phoneVersion,
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
