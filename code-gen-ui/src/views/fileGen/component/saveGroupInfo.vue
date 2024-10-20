<script setup>
import {ref,watch,nextTick} from "vue";
import {Check, Delete, WarningFilled} from "@element-plus/icons-vue";
import * as groupApi from "@/api/groupApi.js"
import * as fileApi from "@/api/fileApi.js"

import {ElMessage} from "element-plus";
import DatabaseSelect from "@/views/fileGen/component/databaseSelect.vue";
import DatabaseTableSelect from "@/views/fileGen/component/databaseTableSelect.vue";
import MappingSelect from "@/views/fileGen/component/mappingSelect.vue";
import FileSelect from "@/views/fileGen/component/fileSelect.vue";
import FileStatus from "@/views/fileGen/component/fileStatus.vue";

const saveDialog= ref(false)
const isUpdate= ref(false)
const Id=ref(0)
const emit=defineEmits(['afterSave'])


const form = ref({
  //对这次操作的描述
  describe:"",
  rootDir:"",
  fileAndGroups:[],
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
      fileAndGroups:[],
    }
  }
}


const findById=async (id)=>{
  let res=await groupApi.FindById(id)
  form.value=res.data
}

const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  let res=await groupApi.Add(form.value)
  ElMessage.success(res.message)

  form.value={
    fileAndGroups:[],
  }
  saveDialog.value=false

  //让父组件更新
  emit('afterSave')
}

const updateForm=async ()=>{
  let res=await groupApi.UpdateById(Id.value,form.value)
  ElMessage.success(res.message)
  form.value={
    fileAndGroups:[],
  }
  saveDialog.value=false
  emit('afterSave')
}

const cancel=()=>{
  saveDialog.value=false
}

const clear=()=>{
  form.value={
    fileAndGroups:[],
  }
}

const fileSelectValue=ref("")
const confirmAddFile=async ()=>{
  //填充fileAndGroup
  //但是groupId是在后台添加完group之后才填充的
  //并非现在填充
  let fileAndGroup={
    fileId:0,
    outDir:"",

    //templateName不是数据库字段，仅仅是展示的
    templateName:"",
  }

  console.log(fileSelectValue.value)
  if (fileSelectValue.value!==0&& fileSelectValue.value!==""){
    //根据fileId查询file信息
    let res=await fileApi.FindById(fileSelectValue.value)
    //只保留文件名称
    let endIndex = res.data.templatePath.lastIndexOf('/');
    fileAndGroup.templateName=res.data.templatePath.substring(endIndex+1)

    //赋值id
    fileAndGroup.fileId=res.data.id

    //给fileAndGroups新增了对象之后
    //我们下面的collapse才能填充
    form.value.fileAndGroups.push(fileAndGroup)
  }

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
        <el-form-item label="要生成文件的根目录">
          <el-input v-model="form.rootDir"
                    clearable>
          </el-input>
        </el-form-item>
        <el-form-item label="选择模板文件">
          <file-select v-model="fileSelectValue"></file-select>
          <el-button style="margin-top: 5px"
                     type="danger"
                     @click="confirmAddFile">确定添加</el-button>
        </el-form-item>
            <el-collapse>
              <el-collapse-item
                                v-for="item in form.fileAndGroups"
                                :key="item">
                <template #title style="position: relative">
                  <div>{{item.templateName}}</div>
                  <!--后面加上删除图片-->
                  <el-icon :size="20" style="position: absolute;right: 80px">
                    <Delete @click="form.fileAndGroups.splice(form.fileAndGroups.indexOf(item),1)"/>
                  </el-icon>

                  <file-status style="margin-left: 20px" :status="item.fileInfo.templatePathIsExist"></file-status>

                </template>

                <!--自定义的图标不显示出来，不然最右侧有一个箭头图标-->
                <template #icon="{ isActive }">
                  <span class="icon-ele">
                    {{ isActive ? '' : '' }}
                  </span>
                </template>

                <el-form-item label="要生成到的目录">
                  <el-input v-model="item.outDir" placeholder="相对于根目录,比如/go/model"></el-input>
                </el-form-item>
              </el-collapse-item>
            </el-collapse>
      </el-form>
      <template #footer>
        <el-button @click="cancel">取消</el-button>
        <el-button @click="clear">清空</el-button>
        <el-button v-if="isUpdate" type="primary" @click="updateForm">修改</el-button>
        <el-button v-else type="primary" @click="addForm">新增</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.myLabel{
  display: flex;
  align-items: center;
}





</style>