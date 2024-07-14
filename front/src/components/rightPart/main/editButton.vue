<template>
  <div>
    <el-button
        size="mini"
        @click="clickedBottom()">编辑</el-button>
    <el-button
        size="mini"
        @click="handleLook(data)">查看</el-button>
    <detailFromVue ref="forms" :title="'编辑问题'"
        :type="'updata'" :oldInfo="infos" :problemId="problemId"></detailFromVue>
    </div> 
</template>

<script>
import detailFromVue from './detailFrom.vue'
export default {
    props : ["data"],
    components : {
      detailFromVue
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
        console.log(this.problemId)
    }
}
</script>

<style>

</style>