<template>
  <div>
    <showMarkdownVue :context="description" :type="'description'" :problemId="infos.Problem_id"
        :pType="infos.Type" :url="infos.Problem_url"></showMarkdownVue>
    <el-divider></el-divider>
    <showMarkdownVue :context="solution" :type="'solution'" :problemId="infos.Problem_id"></showMarkdownVue>
  </div>
</template>

<script>
import showMarkdownVue from '../components/rightPart/problem/showMarkdown.vue'
import { EventBus } from '@/bus/EventBus'
export default {
    components : {
        showMarkdownVue,
    },
    data() {
        return {
            infos : {},
            detail : {},
            description : "",
            solution : ""
        }
    },
    created() {
        this.infos = this.$route.query
        let id = this.infos.Problem_id
        let title = this.infos.Problem_name
        this.$api.project.queryDetail({Problem_id: id}).then((res) => {
            res = res.data;
            if(res.code == 0) return
            this.detail = res.data
            this.description = this.detail.Description
            this.solution = this.detail.Solution
        })
        EventBus.$on("description", (payload) => {
            console.log(payload)
            console.log(payload.text)
            this.$api.project.changeDescription({
                Problem_id : payload.problemId,
                text : payload.text
            }).then((res)=> {
                res = res.data;
                if(res.code == 0) {
                    this.$message.error("保存失败"); return
                }
                this.$message.success("保存成功");
            })
        });
        EventBus.$on("solution", (payload) => {
            this.$api.project.changeSolution({
                Problem_id : payload.problemId,
                text : payload.text
            }).then((res)=> {
                res = res.data;
                if(res.code == 0) {
                    this.$message.error("保存失败"); return
                }
                this.$message.success("保存成功");
            })
        });
    },
    methods: {
        addTitle(title, context) {
            context = "## " + title + "\n" + context;
            return context;
        }
    }
}
</script>

<style>

</style>