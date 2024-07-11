<template>
  <div>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column label="解决日期" width="180">
            <template slot-scope="scope">
                <i class="el-icon-time"></i>
                <span style="margin-left: 10px">{{ scope.row.Solve_time }}</span>
            </template>
        </el-table-column>
        <el-table-column label="题目名称" width="250">
        <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
            <p>{{ scope.row.Problem_name }}</p>
            <a :href="scope.row.Problem_url">访问原题链接</a>
            <div slot="reference" class="name-wrapper">
                {{ scope.row.Problem_name }}
            </div>
            </el-popover>
        </template>
        </el-table-column>
        <el-table-column label="操作">
        <template slot-scope="scope">
            <el-button
            size="mini"
            @click="handleEdit(scope.row)">编辑</el-button>
            <el-button
            size="mini"
            @click="handleDelete(scope.row)">查看</el-button>
        </template>
        </el-table-column>
  </el-table>

  </div>
</template>

<script>
export default {
    data() {
        return {
            tableData : {}
        }
    },
    created() {
        this.$api.project.getTable().then(res => {
            res = res.data
            if(res.code == 0) return
            this.tableData = res.data
            console.log(res)
        })
    },
    methods : {
        handleEdit(data) {

        },
        handleDelete(data) {
            
        }
    }
}
</script>

<style>

</style>