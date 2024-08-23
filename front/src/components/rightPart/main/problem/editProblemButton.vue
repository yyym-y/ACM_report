<!-- 
    问题列表中编辑问题的组件, 包含编辑问题和查看问题
    1. 点击编辑按钮会打开 problemDetailFrom
    2. 点击查看按钮会打开 Problem 视图查看具体内容

    创建此组件需要以下的 prop :
      data : { Problem_name : "", Problem_url : "", Type_id : (int), Solve_time : ""(yyyy-mm-dd), Diff_id : (int) }
 -->
<template>
  <div>
    <el-button
        size="mini"
        @click="clickedBottom()">编辑</el-button>
    <el-button
        size="mini"
        @click="handleLook(data)">查看</el-button>
    <problemDetailFrom ref="forms" :title="'编辑问题'"
        :type="'updata'" :oldInfo="infos" :problemId="problemId"></problemDetailFrom>
    </div> 
</template>

<script>
import problemDetailFrom from './problemDetailFrom.vue';
export default {
    props : ["data"],
    components : {
      problemDetailFrom
    },
    data() {
      return {
        infos : {
          Problem_name : "",
          Problem_url : "",
          Type_id : null,
          Solve_time : "",
          Diff_id : null
        },
        problemId : null
      }
    },
    methods : {
      clickedBottom() {
          this.$refs.forms.openBox();
      },
      handleLook(data) {
        this.$router.push({ name : "problem", query: data })
      }
    },
    created() {
        let d = this.data
        this.infos.Problem_name = d.Problem_name
        this.infos.Problem_url = d.Problem_url
        this.infos.Type_id = d.Type_id
        this.infos.Solve_time = d.Solve_time
        this.infos.Diff_id = d.Diff_id
        this.problemId = d.Problem_id
    }
}
</script>

<style>

</style>