<script setup>
import {ref} from "vue";
import * as databaseApi from "@/api/databaseApi.js"
import {ElMessage} from "element-plus";

//快捷双向绑定
let models = defineModel("modelValue");

const options=ref([])

//findAll方法是外部调用的，选择了数据库之后再调用
const findAll=async (id)=>{
  let res=await databaseApi.FindTablesByDatabaseId(id)
  options.value=res.data
  if (res.data==null){
    ElMessage.warning("无任何可用数据库，请配置")
    return
  }
}

defineExpose({
  findAll,
})


</script>

<template>
  <!--
  filterable可过滤的，我们可以在里面搜索label
  -->
  <el-select
      v-model="models"
      multiple
      filterable
      size="large"

      placeholder="可多选"
      clearable
  >
    <el-option v-for="item in options"
               :key="item"
               :label="item"
               :value="item">
    </el-option>
  </el-select>
</template>

<style scoped lang="scss">

</style>