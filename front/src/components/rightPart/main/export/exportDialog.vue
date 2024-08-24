<template>
  <div>
    <el-dialog title="导出题目 & 题解" :visible.sync="vis" width="35%">
        <el-form :model="form" label-width="80px">
            <el-form-item label="题目范围" v-model="form.interval">
                <el-tooltip effect="dark" placement="top-start">
                <div slot="content">导出单个题目, 题目编号使用逗号键分开(12,14)；导出连续多个题目,题目编号使用短横隔开(2-15)<br/>
                    支持两种模式混用， 例子 ： 13,15,19-20</div>
                    <el-input placeholder="请输入导出范围" v-model="form.interval"></el-input>
                </el-tooltip>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="exportPro()">立即创建</el-button>
            </el-form-item>
        </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
    data() {
        return {
            vis : false,
            form : {
                interval : ""
            }
        }
    },
    methods : {
        openBox() {
            this.vis = true
        },
        exportPro() {
            let ss = this.form.interval
            let data = []
            let base = ss.split(",")
            base.forEach((ele) => {
                if(ele.includes("-")) {
                    let tem = ele.split("-")
                    let ll = parseInt(tem[0])
                    let rr = parseInt(tem[1])
                    for(let i = ll ; i <= rr ; i ++)
                    data.push(parseInt(i))
                } else {
                    data.push(parseInt(ele))
                }
            });
            data = Array.from(new Set(data))
            this.$api.index.exportProblem(this, {
                nums : data
            })
        }
    }
}
</script>

<style>

</style>