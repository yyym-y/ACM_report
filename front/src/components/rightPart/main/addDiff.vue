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
            this.$api.project.insertDifficulty({
                "diff" : this.forms.diff_name
            }).then((res) => {
                res = res.data
                if(res.data == 0) {
                    this.$message.err("insert diff error"); return
                }
                this.$message.success("添加难度成功")
            }).catch((err) => {
                this.$message.err("insert diff error");
            });
        }
    }
}
</script>

<style>

</style>