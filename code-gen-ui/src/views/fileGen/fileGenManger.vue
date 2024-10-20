<script setup>
import {ref, onMounted, handleError} from "vue";

import * as fileGenApi from "@/api/fileGenApi.js";
import {ElMessage, ElMessageBox} from "element-plus";
import SaveFileGenInfo from "@/views/fileGen/component/saveFileGenInfo.vue";

const saveFileGenInfoRef=ref(null)
const tableData=ref([])
const totalCount=ref(0)


const findAll=async ()=>{
  let res=await fileGenApi.FindAll(query.value)
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


//新增组
const add=()=>{
  saveFileGenInfoRef.value.open(0)
}

//修改组
const updateById=(id)=>{
  saveFileGenInfoRef.value.open(id)
}

const deleteById=async (id)=>{
  ElMessageBox.confirm('确定删除该配置吗？','提示',{
    type:'warning'
  }).then(async ()=>{
    //确定删除
    let res=await fileGenApi.DeleteById(id)
    ElMessage.success(res.message)

    await findAll()

  }).catch(()=>{
    console.log("取消删除")
  })

}

//成功保存模板信息后的回调
const afterSave=()=>{
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

const checkIds=ref([])
const handlerSelect=(val)=>{
  checkIds.value=val.map((item)=>{
    return item.id
  })
}

//一键生成所有文件
const genFiles=async ()=>{
  if (checkIds.value.length===0){
    ElMessage.warning("请选择要生成的文件")
    return
  }
  let res=await fileGenApi.GenFiles(checkIds.value)
  ElMessage.success(res.message)
}


</script>

<template>
  <div>
    <p style="color: red;font-size: 20px">不要直接生成在项目目录</p>
    <div style="margin-top: 20px">
      <el-input style="width: 250px" v-model="query.queryStr"
                clearable
                @clear="findAll"
                @keyup.enter="findAll" placeholder="可根据描述,表名模糊搜索"></el-input>
      <el-button style="margin-left: 20px" type="primary"  @click="findAll">查询</el-button>
      <el-button style="margin-left: 20px" type="primary"  @click="add">新增生成配置</el-button>
      <el-button style="margin-left: 20px" type="primary"  @click="genFiles">一键生成</el-button>
    </div>
    <el-table
        style="margin-top: 10px"
        :data="tableData"
        @selection-change="handlerSelect"
    >
      <el-table-column type="selection"></el-table-column>
      <el-table-column label="描述">
        <template #default="scope">
          {{scope.row.describe}}
        </template>
      </el-table-column>
      <el-table-column label="数据库">
        <template #default="scope">
          {{scope.row.databaseName}}
        </template>
      </el-table-column>
      <el-table-column label="要生成的表">
        <template #default="scope">
          {{scope.row.tableNames}}
        </template>
      </el-table-column>
      <el-table-column label="组描述">
        <template #default="scope">
          {{scope.row.groupDescribe}}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="primary"  @click="updateById(scope.row.id)">修改</el-button>
          <el-button type="danger"  @click="deleteById(scope.row.id)">删除</el-button>
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
    <save-file-gen-info ref="saveFileGenInfoRef" @afterSave="afterSave"></save-file-gen-info>
  </div>
</template>

<style scoped lang="scss">
.customTableInfo{
  p{
    margin-bottom: 20px;
    font-size: 16px;
    font-weight: bold;
    font-family: "微软雅黑 Light" ;
  }
}
</style>