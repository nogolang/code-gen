package genUtils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
)

// ReadFile 读取文件所有内容
func ReadFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return "", errors.WithMessage(err, "打开文件出错")
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		return "", errors.WithMessage(err, "读取文件失败")
	}
	return string(all), err
}

func JsonToMap(jsonStr string) (map[string]string, error) {
	var mp map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mp)
	if err != nil {
		return nil, errors.Wrap(err, "json解析出错")
	}
	return mp, nil
}

func WindowsPathToLinux(path string) string {
	//去掉俩边的双引号，因为windows直接复制，是有双引号的
	path = strings.Trim(path, "\"")
	replace := strings.Replace(path, "\\", "/", -1)
	return replace
}
