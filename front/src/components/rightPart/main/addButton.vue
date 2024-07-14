<template>
  <div>
    <el-button icon="el-icon-plus" size="mini"
        style="margin-bottom: 5px;" @click="clickedBottom">添加问题</el-button>
    <detailFromVue :title="'添加问题'" ref="form" :type="'insert'"></detailFromVue>
  </div>
</template>

<script>
import detailFromVue from "./detailFrom.vue"
export default {
    components : {
      detailFromVue
    },
    data() {
      return {
        
      }
    },
    methods : {
      clickedBottom() {
          this.$refs.form.openBox();
      }
    }
}
</script>

<style>

</style>