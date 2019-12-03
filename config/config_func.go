package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

// LoadConfig load configure
func LoadConfig(fname string, cfg interface{}) {
	if fname == "" {
		glog.Fatal("use -c to specify configuration file")
	}

	if !exist(fname) {
		glog.Fatalf("config file %s not exist", fname)
	}

	switch typeOfFile(fname) {
	case "yaml":
		content := readConfig(fname)
		err := yaml.Unmarshal(content, cfg)
		if err != nil {
			glog.Fatalf("%s file parse error", fname)
		}
	case "json":
		content := readConfig(fname)
		err := json.Unmarshal([]byte(content), cfg)
		if err != nil {
			glog.Fatalf("%s file parse error", fname)
		}
	case "":
		glog.Fatal("unknown this config file type")
	}
	glog.V(3).Infoln("config file load successfully")
}

// whether the configuration file exists
func exist(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// only support type yaml or json file
func typeOfFile(f string) string {
	typeList := strings.Split(f, ".")
	if len(typeList) == 0 {
		return ""
	}
	t := typeList[len(typeList)-1]
	if t == "yaml" || t == "yml" {
		return "yaml"
	}
	if t == "json" {
		return "json"
	}
	return ""
}

// reading configuration file
func readConfig(f string) []byte {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		glog.V(1).Infoln("can not read config file")
		return []byte("")
	}
	return content
}
