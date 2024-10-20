<script setup>
import {ref} from "vue";
import * as fileApi from "@/api/fileApi.js"

let model=defineModel("modelValue")

const options=ref([])

const findAll=async ()=>{
  let res=await fileApi.FindAllNoPagination()
  options.value=res.data
}
findAll()
</script>

<template>
  <!--
  filterable可过滤的，我们可以在里面搜索label
  -->
  <el-select
      v-model="model"
      filterable
      placeholder="选择模板文件"
      clearable
  >
    <el-option v-for="item in options"
               :key="item.id"
               :label="item.describe"
               :value="item.id">
    </el-option>
  </el-select>
</template>

<style scoped lang="scss">

</style>