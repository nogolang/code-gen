<script setup>

import {zhCn} from "element-plus/es/locale/index";
</script>

<template>
  <div>
    <!--中文化处理-->
    <el-config-provider :locale="zhCn">
      <router-view/>
    </el-config-provider>
  </div>
</template>

<style scoped>

</style>
