<script setup>
import {ElMessage} from "element-plus";
import {dataBaseStore} from "@/stores/index.js";
import {menuStore} from "@/stores/module/menuStore.js";
import {ref} from "vue";
import {useRoute, useRouter} from "vue-router";


//菜单项
let menuList = [
  {
    title: '文件管理', icon: "Lock", path: "/fileGen", children: [
      {title: '数据库配置', icon: "Lock", path: '/fileGen/database'},
      {title: 'mapping配置', icon: "Lock", path: '/fileGen/mapping'},
      //{title: '文件模板配置', icon: "Lock", path: '/fileGen/fileManager'},
      {title: '组配置', icon: "Lock", path: '/fileGen/fileGroup'},
      {title: '文件生成', icon: "Lock", path: '/fileGen/genFile'},
    ]
  },
]



const router=useRouter()
const route=useRoute()
const handleAvatar=async (commandArg)=>{
  switch (commandArg){
    case "logout":
      ElMessage.success("退出")
      dataBaseStore().clearToken()
      router.push("/login")
      break;
  }
}

//处理头像图片
let url = new URL('/src/assets/img/avatar.png',import.meta.url);
let avatarUrl=url.href


//进来就用token解析出用户信息
// const userInfo=ref(null)
// const findById = async () => {
//   let res = await userApi.getUserInfo()
//   userInfo.value = res.data
// }
// findById()




//默认展开的index，这个应该放到store里，把状态保存起来
const defaultOpeneds=ref([
])

//页面创建的时候从store里获取状态
defaultOpeneds.value=menuStore().getExpand()


//sub-menu 展开的回调
//参数是展开的index
const openSub=(MenuOpenEvent)=>{
  defaultOpeneds.value.push(MenuOpenEvent)

  //保存到store里
  menuStore().setExpand(defaultOpeneds.value)
}

//sub-menu 收起的回调
const closeSub=(MenuCloseEvent)=>{
  //去除数组里的内容
  let start = defaultOpeneds.value.indexOf(MenuCloseEvent);
  defaultOpeneds.value.splice(start,1)

  //保存到store里
  menuStore().setExpand(defaultOpeneds.value)
}


</script>

<template>
  <el-container>
    <el-aside width="20%">
      <a href="/">
        <img class="logoClass" src="@/assets/img/logo.png" alt=""/>
      </a>
    <!--左侧菜单-->
      <!--
          default-openeds 默认的展开index,是一个数组，指定el-sub-menu的index
              @open sub-menu展开的回调函数
          default-active 默认激活的index
          active-text-color,点击标题后的激活颜色，需要设置index,需要设置2个以上的菜单才能看到效果
          router,开启router模式，这样会以index作为router的path进行跳转
          -->
      <el-menu
          :default-active="route.path"
          :default-openeds="defaultOpeneds"
          @open="openSub"
          @close="closeSub"
          router
      >
        <div v-for="(item) in menuList">
          <!--
          如果没有子菜单，直接显示即可，无需渲染children了
          主要还是因为el-menu-item和el-sub-menu是2个元素，不然其实无需判断有没有children
          -->
          <el-menu-item v-if="!item.children" :index="item.path" :key="item.path">
            <template #title>
              <!-- 这里要用名称的方式显示图片，所以要用component-->
              <el-icon>
                <component :is="item.icon" style="margin-right: 5px"></component>
              </el-icon>
              <span>{{ item.title }}</span>
            </template>
          </el-menu-item>
          <!--如果有子菜单-->
          <el-sub-menu v-else :index="item.path">

            <!--利用插槽展示内容-->
            <template #title>
              <el-icon>
                <component :is="item.icon" style="margin-right: 5px"></component>
              </el-icon>
              <span>{{ item.title }}</span>
            </template>
            <!--循环子菜单-->
            <div v-for="(child) in item.children">
              <el-menu-item :index="child.path" :key="child.path">
                <template #title>
                  <el-icon>
                    <component :is="child.icon" style="margin-right: 5px"></component>
                  </el-icon>
                  <!--<i :class="child.icon" style="margin-right: 5px"></i>-->
                  <span>{{ child.title }}</span>
                </template>
              </el-menu-item>
            </div>
          </el-sub-menu>
        </div>
      </el-menu>

    </el-aside>

    <el-container>
      <el-header>
         <div class="avatar-class">
           <!--<span>当前登录人: {{ userInfo.username }}</span>-->
           <span>当前登录人: admin</span>
           <el-dropdown
               @command="handleAvatar"
           >
             <!--嵌套头像-->
             <el-avatar class="el-dropdown-link" :size="50" shape="circle" :src="avatarUrl"></el-avatar>

             <!--下拉菜单的子项-->
             <template #dropdown>
               <el-dropdown-menu>
                 <el-dropdown-item command="logout">退出登录</el-dropdown-item>
               </el-dropdown-menu>
             </template>
           </el-dropdown>
         </div>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>

  </el-container>
</template>

<style scoped lang="scss">

//设置侧边栏，高度占满屏幕
//背景黑色，文字白色
.el-aside{
  height: 100vh;
  background-color: black;
  color: white;

  .el-menu{
    //菜单的文字颜色
     --el-menu-text-color: white;

    //菜单的背景颜色
     --el-menu-bg-color:black;

    //激活菜单项的颜色
     --el-menu-active-color:red;

    //鼠标移动到上面显示的颜色
    --el-menu-hover-bg-color:gray;
  }
}

.logoClass {
  width: 100%;
  height: 100px;
}
.el-header{
  height: 100px;
  //让头像左右分开，然后垂直居中
  //这里height100%是为了让div100%高度
  .avatar-class{
    height: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
}

.el-dropdown{
  .el-dropdown-link:focus{
    outline: none;
  }
}


</style>