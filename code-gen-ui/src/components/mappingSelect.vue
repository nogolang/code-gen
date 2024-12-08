<script setup>
import {ref} from "vue";
import * as mappingApi from "@/api/mappingApi.js"

//快捷双向绑定
let models = defineModel("modelValue");
let props=defineProps({
  options:{
    type:Array,
  }
})

//因为后续每个文件都需要，会导致查询多次
//我们只需要在打开组的时候去查询，然后把数据放进来即可
//const options=ref([])
//const findAll=async ()=>{
//   let res=await mappingApi.FindAllNoPagination()
//   options.value=res.data
//}

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
   <el-option v-for="item in props.options"
              :key="item.id"
              :label="item.describe"
              :value="item.id">
   </el-option>
 </el-select>
</template>

<style scoped lang="scss">

</style>