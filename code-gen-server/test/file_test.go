package test

import (
	"code-gen/internal/utils/genUtils"
	"testing"
)

func Test_recurisonFile(t *testing.T) {
	files, err := genUtils.RecursionFiles("D:\\myFile\\代码生成器模板\\grpc-kratos\\go")
	if err != nil {
		return
	}
	for _, file := range files {
		t.Log(file)
	}
}
