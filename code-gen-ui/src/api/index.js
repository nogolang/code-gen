

import request from "axios";
import {dataBaseStore} from "@/stores/index.js";
import {ElMessage} from "element-plus";
import {useRouter} from "vue-router";
import router from '@/router'
import NProgress from 'nprogress';

let myAxios = request.create({
    //api开头，因为要整合到gin里，所以就不写/api了
    baseURL:"",
    timeout:5000
});

myAxios.interceptors.request.use(config=>{
    NProgress.start()

    //携带token到header上,login和register请求除外
    let passUrl=["/user/login","/user/register"]
    if (passUrl.includes(config.url)){
        return config
    }

    //把token赋值给headers，这样就可以随着请求而发送
    //config.headers.Authorization=userStore().getToken();
    return config;
})

//响应拦截器
myAxios.interceptors.response.use(response=>{
    NProgress.done()

    //status 2xx范围内
    //具体要看后台怎么返回，这里后台统一返回了200状态码
    //判断的自定义的错误码

    //如果是模板解析类，则弹出可关闭的提示框
    if (response.data.code===10202 ||
        response.data.code===10201 ||
        response.data.code===10203){
        ElMessage({
            showClose:true,
            message: response.data.message+"\n"+response.data.reason,
            type: 'error',
            duration: 0
        })
        return
    }

    if (response.data.code !== 200) {
        //这里直接提示出去,如果后端没有提示,就是未知错误
        //这里换行后显示reason
        ElMessage.warning(response.data.message+"\n"+response.data.reason)



        //把错误信息抛出去
        return Promise.reject(response.data)
    }



    //单独判定401错误
    //如果code是401，代表登录失效等
    //如果code是4401-5401之间，那就其他权限问题
    //if (response.data.code===401){
    //    userStore().clearToken()
    //    ElMessage.error("token失效,请重新登录")
    //    router.push("/login")
    //    return
    //}else if (response.data.code>=4401 && response.data.code<=5401){
    //    ElMessage.warning(response.data.message+"\n"+response.data.reason)
    //}

    return response.data;
},error=>{
    //status 2xx以外的范围

    //结束NProgress
    NProgress.done()

    //先要判断是不是axios自带的错误，比如timeout，
    //不然要是axios的错误，则response属性是没有的
    if (!error.hasOwnProperty("response")) {
        ElMessage.error(error.message)
        return Promise.reject(error)
    }

    //其他的错误，比如后台发生内部错误，则返回的status是500
    ElMessage.error(error.response.data.message || "未知错误，请联系管理员")
    return Promise.reject(error)
})

export default myAxios