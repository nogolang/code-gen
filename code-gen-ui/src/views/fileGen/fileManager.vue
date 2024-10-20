<script setup>
import {ref,onMounted} from "vue";

import SaveFileInfo from "@/views/fileGen/component/saveFileInfo.vue";
import * as fileApi from "@/api/fileApi.js"
import {ElMessage, ElMessageBox} from "element-plus";
import FileStatus from "@/views/fileGen/component/fileStatus.vue";

const saveFileRef=ref(null)
const tableData=ref([])
const totalCount=ref(0)


const findAll=async ()=>{
  let res=await fileApi.FindAll(query.value)
  tableData.value=res.data.data
  totalCount.value=res.data.total
}


onMounted(()=>{
  //进入页面就查询
  findAll()
})

const query=ref({
  page:1,
  size:5,
  queryStr:''
})


//新增生成配置
const addTemplateConfig=()=>{
  saveFileRef.value.open(0)
}

//修改配置
const updateTemplateConfig=(id)=>{
  saveFileRef.value.open(id)
}


const deleteTemplateConfig=async (id)=>{
  ElMessageBox.confirm('确定删除该配置吗？','提示',{
    type:'warning'
    }).then(async ()=>{
      //确定删除
      let res=await fileApi.deleteById(id)
      ElMessage.success(res.message)

      await findAll()

    }).catch(()=>{
    console.log("取消删除")
  })

}

//成功保存模板信息后的回调
const afterSaveFile=()=>{
  findAll()
}

//分页方法
const handleSizeChange=(currentSize)=>{
  query.value.size=currentSize
  findAll()
}

const handlePageChange=(currentPage)=>{
  query.value.page=currentPage
  findAll()
}

</script>

<template>
 <div>
   <div style="margin-top: 20px">
     <el-input style="width: 250px" 
               v-model="query.queryStr" @keyup.enter="findAll"
               clearable
               @clear="findAll"
               placeholder="可以根据描述,后缀,模糊查询"></el-input>
     <el-button style="margin-left: 20px" type="primary"  @click="findAll">查询</el-button>
     <el-button style="margin-left: 20px" type="primary"  @click="addTemplateConfig">新增配置</el-button>
   </div>

   <el-table
       style="margin-top: 10px"
       :data="tableData"
   >
     <el-table-column label="描述"  width="150px">
       <template #default="scope">
         {{scope.row.describe}}
       </template>
     </el-table-column>
     <el-table-column label="名称后缀"  width="100px">
       <template #default="scope">
         {{scope.row.nameSuffix}}
       </template>
     </el-table-column>
     <el-table-column label="文件后缀" width="100px">
       <template #default="scope">
         {{scope.row.fileSuffix}}
       </template>
     </el-table-column>
     <el-table-column label="模板文件路径">
       <template #default="scope">
         {{scope.row.templatePath}}
       </template>
     </el-table-column>
     <el-table-column align="center"
                      width="110px"
                      label="文件是否存在">
       <template #default="scope">
         <file-status
             class="tableCellCenter"
             :status="scope.row.templatePathIsExist"
         ></file-status>
       </template>
     </el-table-column>
     <el-table-column label="操作">
       <template #default="scope">
         <el-button type="primary"  @click="updateTemplateConfig(scope.row.id)">修改</el-button>
         <el-button type="danger"  @click="deleteTemplateConfig(scope.row.id)">删除</el-button>
       </template>
     </el-table-column>
   </el-table>
   <el-pagination
       v-model:current-page="query.page"
       v-model:page-size="query.size"
       :page-sizes="[1, 5, 10]"
       layout="total, sizes, prev, pager, next, jumper"
       :total="totalCount"
       @size-change="handleSizeChange"
       @current-change="handlePageChange"
   />
   <save-file-info ref="saveFileRef" @afterSaveFile="afterSaveFile"></save-file-info>
 </div>
</template>

<style  scoped lang="scss">
.customTableInfo{
  p{
    margin-bottom: 20px;
    font-size: 16px;
    font-weight: bold;
    font-family: "微软雅黑 Light" ;
  }
}

.tableCellCenter{
  position: absolute;
  left: 50%;
  transform: translateX(-50%)
}

</style>