<script setup>
  let props = defineProps({
    status:Boolean
  });
</script>

<template>
  <div>
    <el-tooltip
        effect="dark"
        :content="props.status?'文件有效':'文件不存在'"
    >
      <div class="status" :style="{backgroundColor:props.status?'green':'gray'}">
      </div>
    </el-tooltip>
  </div>
</template>

<style scoped lang="scss">
.status{
  width: 15px;
  height: 15px;
  border-radius: 50%;
}

</style>