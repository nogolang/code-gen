<script setup>
import {ref,watch,nextTick} from "vue";
import {Check, Delete, WarningFilled} from "@element-plus/icons-vue";
import * as groupApi from "@/api/groupApi.js"
import {ElMessage} from "element-plus";
import FileStatus from "@/views/fileGen/component/fileStatus.vue";
import {DeleteFileById} from "@/api/groupApi.js";
import FileCamelStatus from "@/views/fileGen/component/fileCamelStatus.vue";
import * as mappingApi from "@/api/mappingApi.js";

const saveDialog= ref(false)
const isUpdate= ref(false)
const Id=ref(0)
const emit=defineEmits(['afterSave'])
const mappingOptions= ref([])

const form = ref({
  //对这次操作的描述
  describe:"",
  genRootDir:"",
  searchRootDir:"",
  fileModels:[],
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

  //打开组的时候查询所有的mapping
  findAllMapping()
}

const findAllMapping=async ()=>{
   let res=await mappingApi.FindAllNoPagination()
  mappingOptions.value=res.data
}

const findById=async (id)=>{
  let res=await groupApi.FindById(id)
  form.value=res.data
}
const FindAllDir=async ()=>{
  let tempForm={
    path:form.value.searchRootDir
  }
  let res=await groupApi.FindAllDir(tempForm)
  form.value.fileModels=res.data
}

const FindAllDirForUpdate=async ()=>{
  let tempForm={
    id:Id.value,
    path:form.value.searchRootDir
  }
  let res=await groupApi.FindAllDirForUpdate(tempForm)
  form.value.fileModels=res.data
}


const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  await groupApi.Add(form.value)
  ElMessage.success("新增成功")

  form.value={
    fileModels:[],
  }
  saveDialog.value=false

  //让父组件更新
  emit('afterSave')
}

const updateForm=async ()=>{
  await groupApi.UpdateById(Id.value,form.value)
  ElMessage.success("更新成功")
  form.value={
    fileModels:[],
  }
  saveDialog.value=false
  emit('afterSave')
}

const cancel=()=>{
  saveDialog.value=false
}

const clear=()=>{
  form.value={
    fileModels:[],
  }
}

//const fileSelectValue=ref("")
//const confirmAddFile=async ()=>{
//  //填充fileAndGroup
//  //但是groupId是在后台添加完group之后才填充的
//  //并非现在填充
//  let fileAndGroup={
//    fileId:0,
//    outDir:"",
//    groupId:Id.value,
//
//    //templateName不是数据库字段，仅仅是展示的
//    templateName:"",
//    fileInfo:{
//      templatePathIsExist:false
//    }
//
//  }
//  if (fileSelectValue.value!==0&& fileSelectValue.value!==""){
//    //根据fileId查询file信息
//    let res=await fileApi.FindById(fileSelectValue.value)
//    //只保留文件名称
//    let endIndex = res.data.templatePath.lastIndexOf('/');
//    fileAndGroup.templateName=res.data.templatePath.substring(endIndex+1)
//    fileAndGroup.fileInfo.templatePathIsExist=res.data.templatePathIsExist
//
//    //赋值id
//    fileAndGroup.fileId=res.data.id
//
//    //给fileAndGroups新增了对象之后
//    //我们下面的collapse才能填充
//    form.value.fileAndGroups.push(fileAndGroup)
//  }
//}


const deleteFileById=async (item)=>{
  //同时在数组里删除选中的元素
  form.value.fileModels.splice(form.value.fileModels.indexOf(item),1)
  await DeleteFileById(item.id)
  ElMessage.success("删除成功")
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
        <el-form-item label="选择要搜索的根目录">
          <el-input v-model="form.searchRootDir"
                    clearable>
          </el-input>
          <el-button v-if="!isUpdate" style="margin-top: 5px"
                     type="danger"
                     @click="FindAllDir">确定搜索</el-button>
          <el-button v-else style="margin-top: 5px"
                     type="danger"
                     @click="FindAllDirForUpdate">重新搜索</el-button>
        </el-form-item>
            <el-collapse>
              <el-collapse-item
                                v-for="item in form.fileModels"
                                :key="item">
                <template #title style="position: relative">
                  <div>{{item.templatePath}}</div>
                  <!--后面加上删除图片-->
                  <el-icon :size="20" style="position: absolute;right: 80px">
                    <Delete @click="deleteFileById(item)"/>
                  </el-icon>

                  <file-status style="margin-left: 20px" :status="item.templatePathIsExist"></file-status>
                </template>

                <!--自定义的图标不显示出来，不然最右侧有一个箭头图标-->
                <template #icon="{ isActive }">
                  <span class="icon-ele">
                    {{ isActive ? '' : '' }}
                  </span>
                </template>

                <el-form-item label="要生成到的目录">
                  <el-input v-model="item.genPath" placeholder="相对于根目录,比如/go/model"></el-input>
                </el-form-item>
                <el-form-item label="名称后缀" >
                  <el-input clearable style="width: 150px"
                            placeholder="比如_controller"
                            v-model="item.nameSuffix"></el-input>
                </el-form-item>
                <el-form-item label="文件后缀">
                  <el-input clearable style="width: 150px"
                            placeholder="比如.go"
                            v-model="item.fileSuffix"></el-input>
                </el-form-item>
                <el-form-item label="mapping">
                  <mapping-select :options="mappingOptions"
                                  v-model="item.mappingId"></mapping-select>
                </el-form-item>
                <el-form-item label="是否驼峰:">
                  <template #label>
                    <div class="myLabel">
                      是否驼峰
                      <el-tooltip
                          effect="dark"
                          content="原表名user_info,小驼峰userInfo,大驼峰UserInfo"
                      >
                        <el-icon>
                          <WarningFilled/>
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                  <file-camel-status v-model.number="item.isCamelCase"></file-camel-status>
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