<!-- 
    添加题目的按钮
 -->
<template>
  <div>
    <el-button icon="el-icon-plus" size="mini"
        style="margin-bottom: 5px;" @click="clickedBottom">添加问题</el-button>
    <problemDetailFrom :title="'添加问题'" ref="form" :type="'insert'"></problemDetailFrom>
  </div>
</template>

<script>
import problemDetailFrom from './problemDetailFrom.vue';
export default {
    components : {
      problemDetailFrom
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