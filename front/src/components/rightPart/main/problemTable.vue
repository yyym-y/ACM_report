<template>
  <div>
    <el-table :data="tableData" style="width: 1000px; max-width: 100" :ref="'table'">
        <el-table-column label="编号" width="100"> 
            <template slot-scope="scope">
                {{ scope.row.Problem_id }}
            </template>
        </el-table-column>
        <el-table-column label="解决日期" width="180">
            <template slot-scope="scope">
                <i class="el-icon-time"></i>
                <span style="margin-left: 10px">{{ scope.row.Solve_time }}</span>
            </template>
        </el-table-column>
        <el-table-column label="题目名称" width="325">
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
        <el-table-column label="类型" width="120"> 
            <template slot-scope="scope">
                <el-tag size="medium">{{ scope.row.Type }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column label="难度" width="100">
            <template slot-scope="scope">
                <el-tag size="medium">{{ scope.row.Diff }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
        <template slot-scope="scope">
            <el-button
            size="mini"
            @click="handleEdit(scope.row)">编辑</el-button>
            <el-button
            size="mini"
            @click="handleLook(scope.row)">查看</el-button>
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
    beforeCreate() {
        this.$api.project.getTable().then(res => {
            res = res.data
            if(res.code == 0) return
            this.tableData = res.data
            console.log(res.data)
        })
    },
    mounted() {
        this.controlSize()
    },
    methods : {
        handleEdit(data) {

        },
        handleLook(data) {
            this.$router.push({ name : "problem", query: data })
        },
        controlSize() {
            const elem = this.$refs["table"]
            window.addEventListener('resize', () => {
                const maxWidth = 1000;
                const width = window.innerWidth > maxWidth? maxWidth : window.innerWidth;
                console.log(elem)
                elem.width = `${width}px`;
            });
        }
    }
}
</script>

<style>

</style>