package test

import (
	"code-gen/internal/utils"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"testing"
)

func Test_IsDev(t *testing.T) {
	t.Log(utils.IsDev())
}

func Test_camel(t *testing.T) {
	fmt.Println(strutil.UpperSnakeCase("user_info"))
	fmt.Println(strutil.CamelCase("user_info"))
	fmt.Println(strutil.UpperFirst(strutil.CamelCase("user_info")))
}
