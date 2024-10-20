<script setup>
import {ref,onMounted} from "vue";

import SaveDatabase from "@/views/fileGen/component/saveDatabase.vue";
import * as databaseApi from "@/api/databaseApi.js";
import {ElMessage, ElMessageBox} from "element-plus";

const saveDatabaseRef=ref(null)
const tableData=ref([])
const totalCount=ref(0)


const findAll=async ()=>{
  let res=await databaseApi.FindAll(query.value)
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
  saveDatabaseRef.value.open(0)
}

//修改组
const updateById=(id)=>{
  saveDatabaseRef.value.open(id)
}

const deleteById=async (id)=>{
  ElMessageBox.confirm('确定删除该配置吗？','提示',{
    type:'warning'
  }).then(async ()=>{
    //确定删除
    let res=await databaseApi.DeleteById(id)
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

</script>

<template>
  <div>
    <div style="margin-top: 20px">
      <el-input style="width: 250px" v-model="query.queryStr"
                clearable
                @clear="findAll"
                @keyup.enter="findAll" placeholder="可根据数据库名,host等搜索"></el-input>
      <el-button style="margin-left: 20px" type="primary"  @click="findAll">查询</el-button>
      <el-button style="margin-left: 20px" type="primary"  @click="add">新增数据源</el-button>
    </div>
    <el-table
        style="margin-top: 10px"
        :data="tableData"
    >
      <el-table-column label="描述">
        <template #default="scope">
          {{scope.row.describe}}
        </template>
      </el-table-column>
      <el-table-column label="host">
        <template #default="scope">
          {{scope.row.host}}:{{scope.row.port}}
        </template>
      </el-table-column>
      <el-table-column label="数据库名">
        <template #default="scope">
          {{scope.row.dataBaseName}}
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
    <save-database ref="saveDatabaseRef" @afterSave="afterSave"></save-database>
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