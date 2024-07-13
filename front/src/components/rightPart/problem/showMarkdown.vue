<template>
  <div>
    <el-button size="mini" @click="change()">{{ bottomText }}</el-button>
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
    props: ["context", "type", "problemId"],
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
        }
    },
    beforeDestroy () {
        EventBus.$off(this.type)
    }
}
</script>

<style>

</style>