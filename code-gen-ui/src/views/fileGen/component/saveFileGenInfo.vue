<script setup>
import {ref,watch,nextTick} from "vue";
import {Check, Delete, WarningFilled} from "@element-plus/icons-vue";
import * as fileGenApi from "@/api/fileGenApi.js"
import {ElMessage} from "element-plus";
import DatabaseSelect from "@/views/fileGen/component/databaseSelect.vue";
import DatabaseTableSelect from "@/views/fileGen/component/databaseTableSelect.vue";
import GroupSelect from "@/views/fileGen/component/groupSelect.vue";

const saveDialog= ref(false)
const isUpdate= ref(false)
const Id=ref(0)
const emit=defineEmits(['afterSave'])


const form = ref({
  //对这次操作的描述
  describe:"",
  dataBaseId:'',
  tableNames:[],
  groupId:'',
})


const open = async (id)=>{
  Id.value=id
  saveDialog.value=true
  if (id){
    isUpdate.value=true
    //如果有id，那就是更新操作
    //通过id查询出对应的信息
    await findById(id)

    //查询数据库对应的表
    //因为这个select组件我们不是直接findAll的，组件初始化的什么都没做
    await  afterDataBaseSelect()
  }else{
    isUpdate.value=false
    form.value={
      tableNames:[]
    }
  }
}


const findById=async (id)=>{
  let res=await fileGenApi.FindById(id)
  form.value=res.data
}

const addForm=async ()=>{
  //新增，向后台添加这个form
  //但是有些内容得要处理，比如数据库名称
  let res=await fileGenApi.Add(form.value)
  ElMessage.success(res.message)

  form.value={
    tableNames:[],
  }
  saveDialog.value=false

  //让父组件更新
  emit('afterSave')
}

const updateForm=async ()=>{
  let res=await fileGenApi.UpdateById(Id.value,form.value)
  ElMessage.success(res.message)
  form.value={
    tableNames:[],
  }
  saveDialog.value=false
  emit('afterSave')
}

const cancel=()=>{
  saveDialog.value=false
}

const clear=()=>{
  form.value={
    tableNames:[],
  }
}

//选择数据库之后查询表
const databaseTableRef=ref(null)
const afterDataBaseSelect=()=>{
  databaseTableRef.value.findAll(form.value.dataBaseId)
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
        <el-form-item label="数据库连接">
            <database-select @afterSelect="afterDataBaseSelect" v-model="form.dataBaseId"></database-select>
        </el-form-item>
        <el-form-item label="要生成的表">
          <database-table-select :database-id="form.dataBaseId"
                                 ref="databaseTableRef"
                                 v-model="form.tableNamesArr"
                                 ></database-table-select>
        </el-form-item>
        <el-form-item label="选择组">
          <group-select v-model="form.groupId"></group-select>
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
.myLabel{
  display: flex;
  align-items: center;
}

</style>