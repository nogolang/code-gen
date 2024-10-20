<script setup>
import {ref,onMounted} from "vue";

import SaveGroupInfo from "@/views/fileGen/component/saveGroupInfo.vue";
import * as groupApi from "@/api/groupApi.js"
import {ElMessage, ElMessageBox} from "element-plus";
import {Delete, Edit} from "@element-plus/icons-vue";
import SaveFileInfo from "@/views/fileGen/component/saveFileInfo.vue";
import FileStatus from "@/views/fileGen/component/fileStatus.vue";

const saveGroupRef=ref(null)
const tableData=ref([])
const totalCount=ref(0)


const findAll=async ()=>{
  let res=await groupApi.FindAll(query.value)
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
  saveGroupRef.value.open(0)
}

//修改组
const updateById=(id)=>{
  saveGroupRef.value.open(id)
}

const deleteById=async (id)=>{
  ElMessageBox.confirm('确定删除该配置吗？','提示',{
    type:'warning'
  }).then(async ()=>{
    //确定删除
    let res=await groupApi.DeleteById(id)
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

const saveFileInfoRef=ref(null)
//打开修改文件的窗口
const openFileUpdateDialog=(id)=>{
  saveFileInfoRef.value.open(id)
}

//保存文件信息后的回调
const afterSaveFile=()=>{
  findAll()
}

</script>

<template>
  <div>
    <div style="margin-top: 20px">
      <el-input style="width: 250px" v-model="query.queryStr"
                clearable
                @clear="findAll"
                @keyup.enter="findAll"
                placeholder="可根据描述,表名模糊搜索"></el-input>
      <el-button style="margin-left: 20px" type="primary"  @click="findAll">查询</el-button>
      <el-button style="margin-left: 20px" type="primary"  @click="add">新增组</el-button>
    </div>

    <el-table
        style="margin-top: 10px"
        :data="tableData"
    >
      <!--可展开的表格-->
      <el-table-column type="expand">
        <template #default="props">
          <el-collapse>
              <el-collapse-item  v-for="item in props.row.fileAndGroups" :key="item">
                <template #title style="position: relative">
                  <div>{{item.templateName}}</div>
                  <el-icon :size="20" style="position: absolute;right: 80px">
                    <Edit  @click="openFileUpdateDialog(item.fileInfo.id)"/>
                  </el-icon>
                </template>

                <!--自定义的图标不显示出来，不然最右侧有一个箭头图标-->
                <template #icon="{ isActive }">
                  <span class="icon-ele">
                    {{ isActive ? '' : '' }}
                  </span>
                </template>

                <div class="customTableInfo">
                  <p>文件描述: {{item.fileInfo.describe}}</p>
                  <p style="display: flex;align-items: center">
                    模板路径: {{item.fileInfo.templatePath}}
                    <file-status style="margin-left: 20px" :status="item.fileInfo.templatePathIsExist"></file-status>
                  </p>
                  <p>名称后缀: {{item.fileInfo.nameSuffix}}</p>
                  <p>文件后缀: {{item.fileInfo.fileSuffix}}</p>
                </div>
              </el-collapse-item>
          </el-collapse>
        </template>
      </el-table-column>
      <el-table-column label="描述">
        <template #default="scope">
          {{scope.row.describe}}
        </template>
      </el-table-column>
      <el-table-column label="生成到的根目录">
        <template #default="scope">
          {{scope.row.rootDir}}
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
    <save-group-info ref="saveGroupRef" @afterSave="afterSave"></save-group-info>

    <save-file-info
        ref="saveFileInfoRef"
        @afterSaveFile="afterSaveFile"></save-file-info>
  </div>
</template>

<style scoped lang="scss">
.customTableInfo{
  p{
    margin-bottom: 20px;
    font-weight: bold;
    font-size: 14px;
    font-family: "微软雅黑 Light" ;
  }
}

</style>