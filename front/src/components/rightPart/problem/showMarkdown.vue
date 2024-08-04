<template>
  <div>
    <el-button size="mini" @click="change()">{{ bottomText }}</el-button>
    <el-button size="mini" @click="clawer()" v-if="type == 'description'">爬取题目</el-button>
    <mavon-editor
        class="markdown"
        :value="context"
        :subfield="ifEdit"
        defaultOpen="preview"
        :toolbarsFlag="ifTool"
        :editable="true"
        :scrollStyle="true"
        style="min-height:20px; margin-top: 10px;"
        :toolbars="toolbars"
        @save="HandleSave"
        ref="md"
    ></mavon-editor>
  </div>
</template>

<script>
import { EventBus } from '@/bus/EventBus'
export default {
    props: ["context", "type", "problemId", "url", "pType"],
    data() {
        return {
            toolbars: {
                    trash: true, // 清空
                    save: true, // 保存
                    navigation: true // 导航目录
            },
            ifEdit : false,
            ifTool : false,
            bottomText : "编辑文本",
            contextT : ""
        }
    },
    methods : {
        change() {
            if(this.ifTool) {
                this.ifEdit = false; this.bottomText = "编辑文本"; this.ifTool = false
            } else {
                this.ifEdit = true; this.bottomText = "关闭编辑"; this.ifTool = true
            }
        },
        HandleSave() {
            EventBus.$emit(this.type, {
                problemId : this.problemId,
                text : this.$refs.md.d_value
            })
        },
        clawer() {
            this.$api.project.problemCrawler({
                Type : this.pType,
                Problem_url : this.url
            }).then((res) => {
                res = res.data
                console.log(res)
                if(res.code == 0) {
                    this.$message.error("爬取失败，该类型爬虫不存在"); return
                }
                this.$message.success("爬取成功")
                res = res.data
                this.$refs.md.d_value = res
                console.log(res)
            }).catch((err) => {});
        }
    },
    beforeDestroy () {
        EventBus.$off(this.type)
    },
    created() {
        this.contextT = this.context
    }
}
</script>

<style>

</style>