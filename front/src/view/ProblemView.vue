<template>
  <div>
    <showMarkdownVue :context="description" :type="'description'"></showMarkdownVue>
    <el-divider></el-divider>
    <showMarkdownVue :context="solution" :type="'solution'"></showMarkdownVue>
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
            this.api.project.changeDescription({
                Problem_id : id,
                text : payload.text
            }).then((res)=> {
                res = res.data;
            })
        });
        EventBus.$on("solution", (payload) => {
            this.api.project.changeSolution({
                Problem_id : id,
                text : payload.text
            }).then((res)=> {
                res = res.data;
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