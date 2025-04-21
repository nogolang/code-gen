

# 格外注意

不要在项目里去生成，因为默认会删除之前生成的文件





# 部署



## 免部署

直接下载release，然后直接打开即可

访问localhost:8088即可



**配置数据库**

你需要设置一个数据库专门用于存储项目本身的表，在config里修改成你自己的数据库连接即可

会自动创建相关的表



## 部署后端

当你下载完源码后，想要部署，需要做下面的步骤



修改数据库连接，在config里修改为你自己的数据库，项目运行的时候会产生几张表，所以你最好提供一个专门的数据库用于存储本项目自动生成的几张表

```go
gorm:
  url: "root:root@tcp(192.168.80.128:3306)/myGenCode?charset=utf8&parseTime=True&loc=Local"

```



**运行项目**

```go
go run .\cmd\
```



## 部署前端

当你下载完源码后，想要部署，需要做下面的步骤

```go
npm install

npm run dev
```



# 管理页面使用教程



## 数据源配置

**首先，你应该先配置一个数据源**

这个配置源是你要生成的库的数据源，在这里指定你要生成的数据库

![image-20241209002018994](./assets/image-20241209002018994.png)



<img src="./assets/image-20241209002039018.png" alt="image-20241209002039018" style="zoom:50%;" />



## **mapping配置**

因为这是一个通用的代码生成器，所以需要你自己配置类型映射

比如数据库varchar类型对应go的类型是string

![image-20241209002132574](./assets/image-20241209002132574.png)



<img src="./assets/image-20241209002150001.png" alt="image-20241209002150001" style="zoom:50%;" />







## 组配置

首先在本地准备你的模板文件，模板的编写看go template语法

<img src="./assets/image-20241209002240597.png" alt="image-20241209002240597" style="zoom:50%;" />



一个组对应多个模板文件

![image-20241209003306290](./assets/image-20241209003306290.png)





新增组的时候，可以通过指定一个根路径，会自动查出该路径下所有的.gohtml结尾的文件

然后填写模板的信息，

- 比如要生成到哪个目录，

- 比如名称的后缀，指定Controller，那么生成user表的时候，最终文件名称就是userController，并且还需要制定文件的后缀，比如.go，那么最终生成就是userController.go，还可以指定表名的驼峰形式，比如       UserController

<img src="./assets/image-20241209003124514.png" alt="image-20241209003124514" style="zoom: 67%;" />







## 生成文件

我们需要指定组对应的表，表可以指定多个

这样一来，一个组下面会有多个模板文件，每一个模板文件都可以应用多个表

![](./assets/image_Elk4TB-hCg.png)



最后便可以生成文件了

![](./assets/image_pFSaHqxKPS.png)



## 以表名生成目录

可以指定{{tableWithSmallCamel}}用来表示 "表的名称"

- {{table}}是原表名
- {{tableWithSmallCamel}}是小驼峰形式
- {{tableWithBigCamel}}是大驼峰形式

<img src="./assets/image-20241209003528963.png" alt="image-20241209003528963" style="zoom: 67%;" />



最终你会在gen目录下生成表名称目录

<img src="./assets/image-20241209003602053.png" alt="image-20241209003602053" style="zoom:67%;" />









# 模板语法

## go tempate

模板采用go tempate

```go
//一些go tempate语法的文档，大家自行去学习
https://cloud.tencent.com/developer/article/1683688


//sprig库，用于在模板页面进行一些字符串处理，比如转换到驼峰
https://masterminds.github.io/sprig/
```



## 可用的字段

你可以直接在页面上使用这些字段，具体可看生成model示例

```go
type Table struct {
	DB                       *gorm.DB
	TableName                string //表名
	IdName                   string //当前表的id的名称和类型，id必须在第1个字段
	IdType                   string //id的类型
	IdNameWithSmallCamel     string  //小驼峰形式的id名称
	IdNameWithBigCamel       string  //大驼峰形式的id名称
	TableNameWithBigCamel    string  //大驼峰表名
	TableNameWithSmallCamel  string  //小驼峰表名
	TableComment             string  //表的注解
	DataBaseName             string  //表所在的数据库名称
	DataBaseNameWithNoPrefix string  //去除前缀的数据库名称
    DataBaseNameWithNoPrefixSmallCamel string  //去除前缀的数据库名称,小驼峰
    DataBaseNameWithNoPrefixBigCamel string  //去除前缀的数据库名称,大驼峰
	Fields                   []field //表的字段
}

// Field代表数据库的字段名称和类型
type field struct {
	FieldName               string //原始字段名，从规则上来说应该设计为蛇形命名
	FieldNameWithBigCamel   string //大驼峰字段名,UserInfo
	FieldNameWithSmallCamel string //小驼峰字段名,userInfo
	FieldType               string //字段类型

	//去掉了后面的括号的原始数据库类型，比如varchar(255)，变成varchar，方便判断
	RawFieldType string
	FieldComment string //字段的注解
}

```







## 生成model示例



**模板文件**

这里涉及到go template语法，希望大家自己去了解

```go
{{$table:=.}}
package model

import (
  "gorm.io/gorm"
  "time"
)

type {{$table.TableNameWithBigCamel}} struct {
{{- range $index, $field := $table.Fields}}
  {{- if eq $index 0}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" gorm:"column:id;primaryKey" comment:"{{$field.FieldComment}}"`
  {{- else}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" gorm:"column:{{$field.FieldName}}" comment:"{{$field.FieldComment}}"`
  {{- end}}
{{- end}}
}

func (receiver *{{$table.TableNameWithBigCamel}}) TableName() string {
  return "{{$table.TableName}}"
}

type {{$table.TableNameWithBigCamel}}Request struct {
{{- range $index, $field := $table.Fields}}
  {{- if ne $index 0}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" binding:"required"`
  {{- end}}
{{- end}}
}

type {{$table.TableNameWithBigCamel}}Response struct {
{{- range $index, $field := $table.Fields}}
  {{- if ne $index 0}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" binding:"required"`
  {{- end}}
{{- end}}
}

type {{$table.TableNameWithBigCamel}}QuerySearch struct {
  Page     int    `json:"page"`
  Size     int    `json:"size"`
  QueryStr string `json:"queryStr"`
}

type {{$table.TableNameWithBigCamel}}ResponseAll struct {
  DataList []*{{$table.TableNameWithBigCamel}}Response `json:"dataList"`
  Total    int64           `json:"total"`
}

func (receiver *{{$table.TableNameWithBigCamel}}Request) Validate() (bool, error) {
  return true, nil
}


```



**最终生成的**

```go
package model

import (
  "gorm.io/gorm"
  "time"
)

type LayoutUserIdCard struct {
    Id int64 `json:"id" gorm:"column:id;primaryKey" comment:""`
    CardName string `json:"cardName" gorm:"column:card_name" comment:""`
    UserId int64 `json:"userId" gorm:"column:user_id" comment:""`
    CreatedAt time.Time `json:"createdAt" gorm:"column:created_at" comment:""`
    UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at" comment:""`

    //生成出来不是万能的，比如这里你就需要改为gorm.DeleteAt
    //或者在customFunc里添加自定义函数处理，自定义函数看下面
    DeletedAt time.Time `json:"deletedAt" gorm:"column:deleted_at" comment:""`
}

func (receiver *LayoutUserIdCard) TableName() string {
  return "layout_user_id_card"
}

type LayoutUserIdCardRequest struct {
    CardName string `json:"cardName" binding:"required"`
    UserId int64 `json:"userId" binding:"required"`
    CreatedAt time.Time `json:"createdAt" binding:"required"`
    UpdatedAt time.Time `json:"updatedAt" binding:"required"`
    DeletedAt time.Time `json:"deletedAt" binding:"required"`
}

type LayoutUserIdCardResponse struct {
    CardName string `json:"cardName" binding:"required"`
    UserId int64 `json:"userId" binding:"required"`
    CreatedAt time.Time `json:"createdAt" binding:"required"`
    UpdatedAt time.Time `json:"updatedAt" binding:"required"`
    DeletedAt time.Time `json:"deletedAt" binding:"required"`
}

type LayoutUserIdCardQuerySearch struct {
  Page     int    `json:"page"`
  Size     int    `json:"size"`
  QueryStr string `json:"queryStr"`
}

type LayoutUserIdCardResponseAll struct {
  DataList []*LayoutUserIdCardResponse `json:"dataList"`
  Total    int64           `json:"total"`
}

func (receiver *LayoutUserIdCardRequest) Validate() (bool, error) {
  return true, nil
}
) Validate() (bool, error) {
  return true, nil
}


```



## **自定义函数**

你可以在CustomFunc 方法里自定义自己的方法

```go
// CustomFunc 自定义函数
var (
  CustomFunc = template.FuncMap{
    //转换到驼峰后，首字母小写
    "lowerFirstCamel": func(str string) string {
      camelStr := strutil.CamelCase(str)
      first := camelStr[:1]
      remain := camelStr[1:]
      first = strings.ToLower(first)
      return first + remain
    },
    "isGormDeleteAt": func(str string) bool {
      //如果是deleteAt字段，那么类型变为gorm.DeletedAt
      if strings.Contains(str, "DeletedAt") ||
        strings.Contains(str, "deleted_at") ||
        strings.Contains(str, "deleteAt") {
        return true
      }
      return false
    },
  }
)
```



**应用isGormDeleteAt**

```go
type {{$table.TableNameWithBigCamel}} struct {
{{- range $index, $field := $table.Fields}}
  {{- if eq $index 0}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" gorm:"column:id;primaryKey" comment:"{{$field.FieldComment}}"`
  {{- else if eq (isGormDeleteAt $field.FieldName) true}}
     //这里如果返回true，那么类型直接就是gorm.DeletedAt
    {{$field.FieldNameWithBigCamel}} gorm.DeletedAt `json:"{{$field.FieldNameWithSmallCamel}}" gorm:"column:{{$field.FieldName}}" comment:"{{$field.FieldComment}}"`
  {{else}}
    {{$field.FieldNameWithBigCamel}} {{$field.FieldType}} `json:"{{$field.FieldNameWithSmallCamel}}" gorm:"column:{{$field.FieldName}}" comment:"{{$field.FieldComment}}"`
  {{- end}}
{{- end}}
}

```



## 常见的mapping



### **go**

```
{
  "bigint":"int64",
  "int":"int32",
  "varchar":"string",
  "longtext":"string",
  "char":"string",
  "datetime":"*time.Time",
  "date":"*time.Time",
  "tinyint":"byte",
  "float":"float32",
  "double":"float64"
}
```



### vue

``` {
  "bigint":"0",
  "varchar":"''",
  "longtext":"''",
  "char":"''",
  "int":"0",
  "datetime":"''",
  "date":"''",
  "tinyint":"0",
  "float":"0",
  "double":"0"
}
```





### proto

```
{
  "bigint":"int64",
  "int":"int32",
  "varchar":"string",
  "longtext":"string",
  "char":"string",
  "datetime":"google.protobuf.timestamp",
  "date":"google.protobuf.timestamp",
  "float":"float",
  "double":"double"
}
```





# 综合案例



## 生成vue form

**mapping的配置**

```json
{
  "bigint":"0",
  "varchar":"''",
  "longtext":"''",
  "char":"''",
  "int":"0",
  "datetime":"''",
  "date":"''",
  "tinyint":"0",
  "float":"0",
  "double":"0"
}
```



**下面生成的form表单**

```js
<el-form ref="formRef" :model="form" :rules="rules" label-position="left" label-width="auto">
{{- range $index, $field := $table.Fields}}
  {{- if ne $index 0}}
    <el-form-item label="{{$field.FieldComment}}" prop="{{$field.FieldNameWithSmallCamel}}">
      {{/*如果是数值类型，那么加上.Number,通过原生类型判断即可*/}}
      {{- if isJsNumberType $field.RawFieldType}}
        <el-input v-model.number="form.{{$field.FieldNameWithSmallCamel}}" clearable/>
      {{- else}}
        <el-input v-model="form.{{$field.FieldNameWithSmallCamel}}" clearable/>
      {{- end}}
    </el-form-item>
  {{- end}}
{{- end}}
</el-form>
```



**上面我们使用的自定义的方法**

```go
"isJsNumberType": func(str string) bool {
    //判断类型是string还是number，因为vue里对于数值类型需要用v-model.number
    if strings.Contains(str, "bigint") ||
        strings.Contains(str, "int") ||
        strings.Contains(str, "float") ||
        strings.Contains(str, "double") ||
        strings.Contains(str, "tinyint") {
        return true
    }
    return false
},
```



**最终生成的**

```js
<el-form ref="formRef" :model="form" :rules="rules" label-position="left" label-width="auto">
    <el-form-item label="标题" prop="title">
        <el-input v-model="form.title" clearable/>
    </el-form-item>
    <el-form-item label="价格" prop="price">
        <el-input v-model.number="form.price" clearable/>
    </el-form-item>
    <el-form-item label="库存" prop="stock">
        <el-input v-model.number="form.stock" clearable/>
    </el-form-item>
    <el-form-item label="分类id" prop="categoryId">
        <el-input v-model.number="form.categoryId" clearable/>
    </el-form-item>
    <el-form-item label="品牌id" prop="brandId">
        <el-input v-model.number="form.brandId" clearable/>
    </el-form-item>
    <el-form-item label="展示图" prop="imgUrl">
        <el-input v-model="form.imgUrl" clearable/>
    </el-form-item>
</el-form>
```





**效果图**

其他的样式之类的，基本是不会变的，便的只会是数据库的字段，所以可以很容易制作一个自己的template

<img src="./assets/image-20241204193332502.png" alt="image-20241204193332502" style="zoom:67%;" />







## 大括号问题

```bash
#假设下面有一个restful请求，我们要用 {{IdName}}替换它
/find/{attrId}

#此时就变成了下面这样，此时就会有问题了，go template编译不通过
/find/{{{IdName}}}

#我们应该在自定义函数里添加以下函数，手动添加括号
"addBrace": func(str string) string {
    return "{" + str + "}"
},
"addBrace2": func(str string) string {
    return "{{" + str + "}}"
},

#然后变成这样，那么最终就会生成/find/{attrId}
/find/{{addBrace $table.IdName}}

#如果是vue的，则可以这样写
批量删除({{addBrace2 "Selection.length"}})
```











