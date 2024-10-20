<script setup>
import {ref} from "vue";
import * as databaseApi from "@/api/databaseApi.js"

const options = ref([])

const findAll = async () => {
  let res = await databaseApi.FindAllDataBaseNameNoPagination()
  options.value = res.data
}
findAll()

const emit = defineEmits(['afterSelect'])
const handlerSelect = () => {
  emit('afterSelect')
}

defineExpose({
      findAll
    }
)
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
               :key="item"
               :label="item"
               :value="item">

    </el-option>
  </el-select>
</template>

<style scoped lang="scss">

</style>