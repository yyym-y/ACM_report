<template>
  <div>
    <showMarkdown :context="description" :saveBusKey="'description'" :problemId="infos.Problem_id" :ref="'desc'">
        <el-button size="mini" @click="clawer()" >爬取题目</el-button>
    </showMarkdown>
    <el-divider></el-divider>
    <showMarkdown :context="solution" :saveBusKey="'solution'" :problemId="infos.Problem_id"></showMarkdown>
  </div>
</template>

<script>
import showMarkdown from '@/components/rightPart/mdUtil/showMarkdown.vue';
import { EventBus } from '@/bus/EventBus'
export default {
    name : "ProblemView",
    components : {
        showMarkdown,
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
        this.$api.problem.queryMdDetail(this, {Problem_id: id}).then((res) => {
            this.detail = res
            this.description = this.detail.Description
            this.solution = this.detail.Solution
        })

        EventBus.$on("description", (payload) => {
            this.$api.problem.changeDescription(this, {
                Problem_id : payload.problemId,
                text : payload.text
            });
        });
        EventBus.$on("solution", (payload) => {
            this.$api.problem.changeSolution(this, {
                Problem_id : payload.problemId,
                text : payload.text
            })
        });
    },
    destroyed() {
        EventBus.$off("description");
        EventBus.$off("solution");
    },
    methods : {
        clawer() {
            this.$api.problem.problemCrawler(this, {
                Type : this.infos.Type,
                Problem_url : this.infos.Problem_url
            }).then((res) => {
                if(res != null) {
                    console.log(res)
                    this.$refs.desc.changeContext(res)
                }
            });
        }
    }
}
</script>

<style>

</style>