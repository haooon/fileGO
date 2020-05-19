package service

import (
	"io"
	"bufio"
    "os"
    "strings"
    "encoding/json"
    "MANGAGO/logService"
)

var ConfigContent = [...]string{
	"app_name",
	"app_port",
	"app_ip",
	"db_name",
	"db_ip",
	"db_port",
	"db_username",
	"db_password"}

var Config  map[string]string /*创建集合 */

// Config["Version"] = "v1.0"

func init() {
	Config = make(map[string]string)
	file, err := os.Open("go.config")
	if err != nil {
		return 
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for{
		data, _, err := r.ReadLine()
		if err == io.EOF{
			break
		}
		if err != nil {
			logService.OUT("read err", string(err.Error()))
			break
		}
		strdata := string(data)
		logService.OUT(strdata)

		for _,iconfig := range ConfigContent{
			// logService.OUT(string(iconfig))
			if strings.Index(strdata, iconfig) != -1{
				Config[iconfig] = strings.Split(strdata, "=")[1]
			}
		}
	}
	jsonStr, err := json.Marshal(Config)
	logService.OUT(string(jsonStr))
}

func GetConfig(name string) {
	logService.OUT(Config[name])
	return Config[name]
}