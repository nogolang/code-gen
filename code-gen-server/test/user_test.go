package test

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"testing"
)

func Test_initDB(t *testing.T) {
	client := resty.New()
	baseUrl := "http://localhost:8001/gen/initDataBase"

	//启用调试
	client.SetDebug(true)

	url := "root:root@tcp(192.168.80.128:3306)/myTest?charset=utf8&parseTime=True&loc=Local"
	//want是自定义的code码，看看是不是能对应我们定义的错误情况
	tests := []struct {
		name string
		want int
		args *model.GormConfig
	}{
		{"测试生成的情况", 200, &model.GormConfig{
			Url: url,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//发送请求
			res, _ := client.R().
				SetBody(tt.args).
				EnableTrace().
				Post(baseUrl)

			//解析出返回值
			var data commonRes.Response
			err := json.Unmarshal(res.Body(), &data)
			if err != nil {
				t.Error("json解析出错", err, string(res.Body()))
				return
			}

			//如果不是我们想要的结果，则代表测试失败，return里层函数
			if data.Code != tt.want {
				t.Error(err, string(res.Body()))
				return
			}

			//如果是我们要的结果，打印出来即可
			t.Log("响应码", res.Status())
			t.Log("响应数据", data)
		})
	}
}

func Test_genFile(t *testing.T) {
	client := resty.New()
	baseUrl := "http://localhost:8001/gen/file"

	//启用调试
	client.SetDebug(true)

	//want是自定义的code码，看看是不是能对应我们定义的错误情况
	tests := []struct {
		name string
		want int
		args *model.FileModel
	}{
		{"测试生成的情况", 200, &model.FileModel{
			DataBaseName: "myTest",
			TableName:    "user",
			TemplatePath: []string{
				"C:/Users/1/Desktop/files/tpl/go/model/model.go.gohtml",
			},
			MappingPath: "C:/Users/1/Desktop/files/mapping/go.json",
			NameSuffix:  "",
			OutDir:      "C:/Users/1/Desktop/files/out",
			IsCamelCase: false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//发送请求
			res, _ := client.R().
				SetBody(tt.args).
				EnableTrace().
				Post(baseUrl)

			//解析出返回值
			var data commonRes.Response
			err := json.Unmarshal(res.Body(), &data)
			if err != nil {
				t.Error("json解析出错", err, string(res.Body()))
				return
			}

			//如果不是我们想要的结果，则代表测试失败，return里层函数
			if data.Code != tt.want {
				t.Error(err, string(res.Body()))
				return
			}

			//如果是我们要的结果，打印出来即可
			t.Log("响应码", res.Status())
			t.Log("响应数据", data)
		})
	}
}
