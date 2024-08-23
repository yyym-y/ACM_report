<!-- 
    用来添加题目的难度类型(点击后展示该表单)
 -->
<template>
  <div>
    <el-dialog
        :visible.sync="vis" width="500px" title="添加难度">
        <el-form :model="forms" label-width="80px">
            <el-form-item label="难度名称">
                <el-input v-model="forms.diff_name"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="Submit">立即添加</el-button>
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
            forms : {
                diff_name : ""
            }
        }
    },
    methods : {
        openBox() {
            this.vis = true
        },
        Submit() {
            this.$api.index.insertDiff(this, {"diff" : this.forms.diff_name});
        }
    }
}
</script>

<style>

</style>