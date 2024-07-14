<template>
  <div>
    <el-dialog :title="title" width="500px" :visible.sync="vis">
      <el-form :model="formInfo" label-width="80px">
        <el-form-item label="题目名称">
          <el-input v-model="formInfo.Problem_name"></el-input>
        </el-form-item>
        <el-form-item label="题目链接">
          <el-input v-model="formInfo.Problem_url"></el-input>
        </el-form-item>
        <el-form-item label="解决时间">
          <el-date-picker v-model="formInfo.Solve_time" type="date" placeholder="选择日期"
          value-format="yyyy-MM-dd"></el-date-picker>
        </el-form-item>
        <el-form-item label="题目类型">
          <el-select v-model="formInfo.Type_id" placeholder="请选择">
            <el-option v-for="type in types" :key="'type' + type.Type_id"
              :label="type.Type_name" :value="type.Type_id">
            </el-option>
          </el-select>
          <el-button style="margin-left: 5px;">添加类型</el-button>
        </el-form-item>
        <el-form-item label="难度选择">
          <el-select v-model="formInfo.Diff_id" placeholder="请选择">
            <el-option v-for="diff in diffs" :key="'diff' + diff.Diff_id"
              :label="diff.Diff_name" :value="diff.Diff_id">
            </el-option>
          </el-select>
          <el-button style="margin-left: 5px;">添加难度</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="Submit()">提交</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
    props : ["title", "type"],
    data() {
      return {
        vis : false,
        formInfo : {
          Problem_name : "",
          Problem_url : "",
          Type_id : null,
          Solve_time : "",
          Diff_id : null
        },
        types : [],
        diffs : []
      }
    },
    methods : {
      openBox() {
        this.vis = true
      },
      closeBox() {
        this.vis = false
      },
      readType() {
        this.$api.project.queryAllType().then((res) => {
          res = res.data
          if(res.code == 0) {
            this.$message.err("type read error"); return
          }
          this.types = res.data
        }).catch((err) => {
          this.$message.err("type read error");
        });
      },
      readDiff() {
        this.$api.project.queryAllDiffical().then((res) => {
          res = res.data
          if(res.code == 0) {
            this.$message.err("type diff error"); return
          }
          this.diffs = res.data
        }).catch((err) => {
          this.$message.err("type diff error");
        });
      },
      insertsubmit() {
        console.log(this.formInfo)
        this.$api.project.insertProblem(
          {data : this.formInfo}
        ).then((res)=> {
          res = res.data;
          if(res.code == 0) {
            this.$message.error("insert error"); return;
          }
          this.$message.success("问题添加成功")
        })
      },
      Submit() {
        switch(this.type) {
          case "insert": {
            this.insertsubmit(); break;
          }
        }
      }
    },
    created() {
      this.readDiff(); this.readType();
    }
}
</script>

<style>

</style>