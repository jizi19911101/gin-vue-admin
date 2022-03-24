package monkey

import (
	"bytes"
	"encoding/json"
	"errors"
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

	// 占用设备，defer释放命令
	url = "http://120.25.149.119:8082/api/v1/user/devices"
	dataStr := "{\"udid\":\"" + device + "\"}"
	dataJson := []byte(dataStr)
	req, err := http.NewRequest("POST", url+"?user_id="+userId, bytes.NewBuffer(dataJson))
	if err != nil {
		return errors.New("设备占用失败")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return errors.New("设备占用失败")
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap = make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	success := bodyMap["success"]
	if !success.(bool) {
		return errors.New("设备占用失败")
	}
	// 是否请除日志

	// 获取atxAgentAddress
	url = "http://120.25.149.119:8082/api/v1/user/devices/"
	resp, err = http.Get(url + device + "?user_id=" + userId)

	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap = make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	atxAgentAddress := bodyMap["device"].(map[string]interface{})["source"].(map[string]interface{})["atxAgentAddress"].(string)
	phoneVersion := bodyMap["device"].(map[string]interface{})["properties"].(map[string]interface{})["version"].(string)

	// 测试app是否存在
	url = "http://" + atxAgentAddress + "/shell?user_id=" + userId + "&command=pm%20list%20packages%20-3"
	resp, err = http.Get(url)

	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap = make(map[string]interface{}, 0)
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

	// shell 命令发起monkey测试
	url = "http://" + atxAgentAddress + "/shell/background?user_id=" + userId + "&command="
	command := "CLASSPATH=/sdcard/monkey.jar:/sdcard/framework.jar exec app_process /system/bin tv.panda.test.monkey.Monkey -p " + startMonkeyReq.App + "  --uiautomatordfs  --running-minutes  " + startMonkeyReq.Duration + "  --throttle 500 -v -v "
	url = url + urls.QueryEscape(command)

	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap = make(map[string]interface{}, 0)

	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
	}
	success = bodyMap["success"]
	if !success.(bool) {
		return errors.New("启动Monkey失败")
	}
	pid := bodyMap["pid"]

	beginTime := time.Now().Format("2006-01-02 15:04:05")

	// 查询测试进程是否结束,10秒查一次

LOOP:
	subprocess, err := monkeyService.getSubprocess(atxAgentAddress)
	if err != nil {
		return nil
	}

	for strings.Contains(subprocess, "tv.panda.test.monkey") {
		time.Sleep(time.Duration(10) * time.Second)
		goto LOOP

	}

	// 生成测试报告
	htmlPath := global.RedirectConfigFile("tpl.html")
	t, err := template.ParseFiles(htmlPath)
	if err != nil {
		return nil
	}

	url = "http://" + atxAgentAddress + "/packages/" + startMonkeyReq.App + "/info"
	resp, err = http.Get(url)

	if err != nil {
		return err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyMap = make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(string(body)), &bodyMap)
	if err != nil {
		return err
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
		return nil
	}
	return nil
}

func (monkeyService *MonkeyService) getSubprocess(atxAgentAddress string) (string, error) {
	url := "http://" + atxAgentAddress + "/proc/list"

	resp, err := http.Get(url)
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
