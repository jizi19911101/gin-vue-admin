package monkey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	//url = "http://120.25.149.119:8082/api/v1/user/devices?user_id=admin@anonymous.com"
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
	defer resp.Body.Close()

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
	atxAgentAddress := bodyMap["device"].(map[string]interface{})["source"].(map[string]interface{})["atxAgentAddress"]
	fmt.Println(atxAgentAddress, "atxAgentAddressatxAgentAddress")

	// 测试app是否存在
	url = "http://" + atxAgentAddress.(string) + "/shell?user_id=" + userId + "&command=pm%20list%20packages%20-3"
	resp, err = http.Get(url)
	defer resp.Body.Close()

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
	fmt.Println(error, "errorerror")
	if error != nil {
		return errors.New("查询app包列表失败")
	}

	output := bodyMap["output"].(string)
	if !strings.Contains(output, startMonkeyReq.App) {
		return errors.New("此app不存在，请进行确认")
	}
	// shell 命令发起monkey测试
	//url = "http://" + atxAgentAddress.(string) + "/shell/background?user_id=" + userId + "&command="
	//command := "CLASSPATH=/sdcard/monkey.jar:/sdcard/framework.jar exec app_process /system/bin tv.panda.test.monkey.Monkey -p " + startMonkeyReq.App + "  --uiautomatordfs  --running-minutes  " + startMonkeyReq.Duration + "  --throttle 500 -v -v "

	// 生成测试报告

	return nil
}
