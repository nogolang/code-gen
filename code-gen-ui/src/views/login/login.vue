<script setup>
import {ref} from "vue";
import {dataBaseStore} from "@/stores/index.js";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";

import * as userApi from "@/api/userApi.js"
import {User} from "@element-plus/icons-vue/global";
import {Edit} from "@element-plus/icons-vue";

const router=useRouter()

let loginForm = ref({
  username:"",
  password:"",
})
const isLoginPage=ref(true)
const loginFormRef=ref(null)

const login=async ()=>{
  loginFormRef.value.validate().then(async (isValidate)=>{
      //处理token
      // let res=await userApi.login(loginForm.value)
      // userStore().setToken(res.data)
      // ElMessage.success(res.message)

      //模拟一个token
      dataBaseStore().setToken("123456")
      ElMessage.success("登录成功")

      //跳转后后台页面
      router.push("/")
  },err=>{
    ElMessage.warning("请输入正确的用户名或密码")
  })
}



let loginRules={
  username:[
    {required:true,message:"用户名不能为空",trigger:"blur"},
    //前面占了1个位置，所以后面是5-14，代表6-15位
    {pattern:/^[a-zA-Z][a-zA-Z0-9]{5,14}$/,
      message:"用户名必须以字母开头，长度在6-15位之间",
      trigger:"blur"}
  ],
  password:[
    {required:true,message:"密码不能为空",trigger:"blur"},
    //前面[A-Za-z\d]占了1个位置，所以后面是5-14，代表6-15位
    {pattern:/^(?=.*[a-zA-Z])(?=.*\d)[A-Za-z\d]{5,14}$/,
      message:"密码必须包含至少一个字母和一个数字，长度在 6-15 位之间",
      trigger:"blur"}
  ]
}

</script>

<template>
  <div>
    <!--middle垂直居中-->
    <el-row align="middle">
      <el-col :span="12">
          <img class="login-img" src="~@/assets/img/login.png" alt="">
      </el-col>
      <el-col :span="12" class="flex-login" ref="myForm">
        <div class="login-form">
          <h2>{{isLoginPage?"登录":"注册"}}</h2>
          <!--表单需要校验-->
          <el-form :model="loginForm"
                   :rules="loginRules"
                   ref="loginFormRef"
                   require-asterisk-position="right">
            <el-form-item prop="username">
              <el-input clearable :prefix-icon="User"
                        v-model="loginForm.username"
                        type="text"
                        placeholder="请输入用户名">

              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input  :prefix-icon="Edit"
                        v-model="loginForm.password"
                        type="password"
                        :show-password="true"
                        placeholder="请输入密码">
              </el-input>
            </el-form-item>
            <el-form-item>
              <el-button v-if="isLoginPage" @click="login" type="warning" class="login-button">登录</el-button>
              <el-button v-else @click="login" type="primary" class="login-button">注册</el-button>
            </el-form-item>
          </el-form>

          <el-link @click="isLoginPage=!isLoginPage" style="margin-top: 20px" type="primary">
            {{isLoginPage?"没有账号？去注册":"已有账号？去登录"}}
          </el-link>

        </div>
      </el-col>
    </el-row>
  </div>
</template>
<style scoped lang="scss">
//使用vh让其占满整个屏幕
//我们使用了el-col让其只占一半屏幕
.login-img {
  width: 100%;
  height: 100vh;
}

.flex-login {
  display: flex;
  justify-content: center;
}
.login-form {
  width: 300px;
  margin-top: -250px;
  h2{
   margin-bottom: 15px;
  }
}
.login-button{
  margin-top: 20px;
  width: 100%;
}
</style>