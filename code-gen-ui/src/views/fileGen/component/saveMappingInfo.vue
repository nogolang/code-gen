<script setup>
import { ref,onMounted } from 'vue'
import {dataBaseStore} from "@/stores/index.js";
import * as fileApi  from "@/api/fileApi.js";
import {ElMessage} from "element-plus";
import * as mappingApi from "@/api/mappingApi.js";
import { Mode } from 'vanilla-jsoneditor'

const saveDialog= ref(false)
const isUpdate= ref(false)
const Id=ref(0)
const emit=defineEmits(['afterSave'])
import JsonEditorVue from 'json-editor-vue'

const form=ref({
  describe:"",
  content:"",
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
    form.value={
    }
  }
}

const findById=async (id)=>{
  let res=await mappingApi.FindById(id)
  form.value=res.data
}

const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  let res=await mappingApi.Add(form.value)
  ElMessage.success(res.message)

  form.value={}
  saveDialog.value=false

  //让父组件更新
  emit('afterSave')
}

const updateForm=async ()=>{
  let res=await mappingApi.UpdateById(Id.value,form.value)
  ElMessage.success(res.message)
  form.value={}
  saveDialog.value=false
  emit('afterSave')
}

const cancel=()=>{
  saveDialog.value=false
}



defineExpose({
  open,
})

</script>

<template>
  <div>
    <el-dialog :title="isUpdate?'修改':'新增'" v-model="saveDialog">
      <el-form label-position="left" label-width="auto">
        <el-form-item label="描述:" style="width: 100%">
          <el-input clearable v-model="form.describe"></el-input>
        </el-form-item>

      <!--这里加入json edit组件-->
        <el-form-item>
          <JsonEditorVue
           style="width: 100%;height: 500px"
           v-model="form.content"
           :mode="Mode.text"
          >
          </JsonEditorVue>
        </el-form-item>
      </el-form>
      <template #footer>
        <div style="margin-top: 30px">
          <el-button @click="cancel">取消</el-button>
          <el-button v-if="isUpdate" type="primary" @click="updateForm">修改</el-button>
          <el-button v-else type="primary" @click="addForm">新增</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>