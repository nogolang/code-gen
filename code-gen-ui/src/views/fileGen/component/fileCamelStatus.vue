<script setup>
import {ref} from "vue";

let model=defineModel("modelValue")

const options=ref([
    {id:1,name:"原表名"},
    {id:2,name:"小驼峰表名"},
    {id:3,name:"大驼峰表名"},
])

</script>

<template>
  <!--
  filterable可过滤的，我们可以在里面搜索label
  -->
  <el-select
      v-model="model"
      placeholder="选择形式"
  >
    <el-option v-for="item in options"
               :key="item.id"
               :label="item.name"
               :value="item.id">
    </el-option>
  </el-select>
</template>

<style scoped lang="scss">

</style>