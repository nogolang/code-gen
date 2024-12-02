<script setup>
import { ref,onMounted } from 'vue'
import {dataBaseStore} from "@/stores/index.js";
import * as fileApi  from "@/api/fileApi.js";
import {ElMessage} from "element-plus";
import * as databaseApi from "@/api/databaseApi.js";

const saveDialog= ref(false)
const isUpdate= ref(false)
const Id=ref(0)
const emit=defineEmits(['afterSave'])


const form=ref({
  describe:"",
  host:'',
  port:'3306',
  dataBaseName:'',
  username:'',
  password:'',
  suffix:'',
  prefix:'',
})

const open = (id)=>{
  Id.value=id
  saveDialog.value=true
  if (id){
    isUpdate.value=true
    //如果有id，那就是更新操作
    //通过id查询出对应的信息
    findById(id)
  }else{
    isUpdate.value=false
    //新增，这里不清空，友好一点,但是id得置空
    //form.value={
    //}
    form.value.id=0
  }
}

const findById=async (id)=>{
  let res=await databaseApi.FindById(id)
  form.value=res.data
}

const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  let res=await databaseApi.Add(form.value)
  ElMessage.success(res.message)

  form.value={}
  saveDialog.value=false

  //让父组件更新
  emit('afterSave')
}

const updateForm=async ()=>{
  let res=await databaseApi.UpdateById(Id.value,form.value)
  ElMessage.success(res.message)
  form.value={}
  saveDialog.value=false
  emit('afterSave')
}

const cancel=()=>{
  saveDialog.value=false
}

const clear=()=>{
  form.value={}
}


const checkConnect=async ()=>{
  let res=await databaseApi.checkConnect(form.value)
  ElMessage.success(res.message)
}

defineExpose({
  open,
})

</script>

<template>
  <div>
    <el-dialog :title="isUpdate?'修改':'新增'" v-model="saveDialog">
      <el-form label-position="left" label-width="auto">
        <el-form-item label="描述:">
          <el-input clearable v-model="form.describe"></el-input>
        </el-form-item>
        <el-form-item label="host:">
          <el-input clearable v-model="form.host"></el-input>
        </el-form-item>
        <el-form-item label="port:">
          <el-input clearable placeholder="3306" v-model="form.port"></el-input>
        </el-form-item>
        <el-form-item label="dataBaseName:">
          <el-input  clearable v-model="form.dataBaseName"></el-input>
        </el-form-item>
        <el-form-item label="username:">
          <el-input clearable v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="password:">
          <el-input clearable type="password"  v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item label="数据库前缀是什么:">
          <el-input clearable type="text" placeholder="比如my_product,前缀就是my_,后面可能要用到'去除指定前缀'" v-model="form.prefix"></el-input>
        </el-form-item>
      </el-form>
      <div style="position: absolute;right: 20px">
        <el-button @click="checkConnect" type="warning">测试连接</el-button>
      </div>
      <template #footer>
        <div style="margin-top: 30px">
        <el-button @click="cancel">取消</el-button>
        <el-button @click="clear">清空</el-button>
        <el-button v-if="isUpdate" type="primary" @click="updateForm">修改</el-button>
        <el-button v-else type="primary" @click="addForm">新增</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>