<script setup>
import {ref} from "vue";
import * as mappingApi from "@/api/mappingApi.js"

//快捷双向绑定
let models = defineModel("modelValue");

const options=ref([])

const findAll=async ()=>{
   let res=await mappingApi.FindAllNoPagination()
   options.value=res.data
}
findAll()

const emit=defineEmits(['afterSelect'])
const handlerSelect=()=>{
  emit('afterSelect')
}

</script>

<template>
 <!--
 filterable可过滤的，我们可以在里面搜索label
 -->
 <el-select
     v-model="models"
     filterable
     @change="handlerSelect"
     size="large"

     placeholder="请选择内容"
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