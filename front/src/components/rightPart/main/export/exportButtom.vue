<template>
  <span>
    <el-button size="mini" icon="el-icon-download" @click="opendia()">导出题目&题解</el-button>
    <exportDialog ref="dialog"></exportDialog>
  </span>
</template>

<script>
import exportDialog from './exportDialog.vue';
export default {
  components : {
    exportDialog
  },
  methods : {
    opendia() {
      this.$refs.dialog.openBox()
    }
  }
}
</script>

<style>

</style>