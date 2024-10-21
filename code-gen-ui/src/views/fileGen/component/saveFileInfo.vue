<script setup>
import {ref,watch,nextTick} from "vue";
import {Check, Delete, WarningFilled} from "@element-plus/icons-vue";
import * as fileApi from "@/api/fileApi.js"
import {ElMessage} from "element-plus";
import MappingSelect from "@/views/fileGen/component/mappingSelect.vue";
import FileCamelStatus from "@/views/fileGen/component/fileCamelStatus.vue";

const saveDialog= ref(false)
const isUpdate= ref(false)
const fileId=ref(0)

const form = ref({
  //对这次操作的描述
  describe:"",
  tableName:"",
  templatePath:"",
  templatePathIsExist:false,
  nameSuffix:"",
  fileSuffix:"",
  //生成出来的文件名称是蛇形还是驼峰
  isCamelCase:"",

  //属于哪个mapping
  mappingId:"",
})


const open = (id)=>{
  fileId.value=id
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
  let res=await fileApi.FindById(id)
  form.value=res.data
}

const emit=defineEmits(["afterSaveFile"])

const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  let res=await fileApi.add(form.value)
  ElMessage.success(res.message)

  form.value={

  }
  saveDialog.value=false

  //让父组件更新
  emit("afterSaveFile")
}

const updateForm=async ()=>{
  let res=await fileApi.updateById(fileId.value,form.value)
  ElMessage.success(res.message)
  form.value={

  }
  saveDialog.value=false
  emit("afterSaveFile")
}

const cancel=()=>{
   saveDialog.value=false
}
const clear=()=>{
  form.value={

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
      <el-form-item label="模板文件路径:">
        <el-input v-model="form.templatePath"
                  clearable
                  placeholder="填写模板文件路径">
        </el-input>
      </el-form-item>
      <el-form-item label="名称后缀" >
        <el-input clearable style="width: 150px"
                  placeholder="比如_controller"
                  v-model="form.nameSuffix"></el-input>
      </el-form-item>
      <el-form-item label="文件后缀">
        <el-input clearable style="width: 150px"
                  placeholder="比如.go"
                  v-model="form.fileSuffix"></el-input>
      </el-form-item>
      <el-form-item label="mapping">
        <mapping-select v-model="form.mappingId"></mapping-select>
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
        <file-camel-status v-model.number="form.isCamelCase"></file-camel-status>
      </el-form-item>
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
  //.el-form-item__label{
  //  align-items: center;
  //}

  .myLabel{
    display: flex;
    align-items: center;
  }

</style>