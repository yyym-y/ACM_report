<!-- 
  展示一个问题详细的所有信息

  创建此组件需要以下的 prop :
    title : "" [这个 dislog 的标题]
    type : "" [有两个可选值: 'updata'(点击提交按钮会触发提交事件), 'insert'(点击提交按钮会触发插入事件)]
    oldInfo : { Problem_name : "", Problem_url : "", Type_id : (int), Solve_time : ""(yyyy-mm-dd), Diff_id : (int) }
      [只有在 type 为 updata 的时候才需要传入, 用来初始化列表内容]
    problemId : (int) [题目的 ID 值]
  
  通过 $ref 调用的方法
    openBox() : 显示这个 dialog
    closeBox() : 关闭这个 dialog
 -->
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
          <el-button style="margin-left: 5px;" @click="openAddType()">添加难度</el-button>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="Submit()">提交</el-button>
          <el-button type="danger" @click="deleteSubmit()" v-if="type=='updata'">删除</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <addDiffVue ref="diff"></addDiffVue>
  </div>
</template>

<script>
import addDiffVue from './addProblemDiff.vue'
export default {
    props : ["title", "type", "oldInfo", "problemId"],
    components : {
      addDiffVue
    },
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
      openBox() { this.vis = true },
      closeBox() { this.vis = false },

      loadDiff() {
        this.$api.index.queryAllDiffical(this).then((res) => {this.diffs = res})
      },
      loadType() {
        this.$api.index.queryAllType(this).then((res) => {this.types = res})
      },
      insertSubmit() {
        this.$api.index.insertProblem(this, {data : this.formInfo})
      },
      updataSubmit() {
        this.$api.index.updataProblem(this, {
          data : this.formInfo,
          Problem_id : this.problemId
        })
      },
      deleteSubmit() {
        this.$alert('将删除本题所有相关信息', '删除确认', {
          confirmButtonText: '确定',
          callback: action => {
            if(action != "confirm")
              return;
            this.$api.index.deleteProblem(this, {Problem_id : this.Problem_id});
          }
        });
      },
      Submit() {
        switch(this.type) {
          case "insert": {
            this.insertSubmit(); break;
          }
          case "updata": {
            this.updataSubmit(); break;
          }
        }
      },
      openAddType() {
        this.$refs.diff.openBox()
      }
    },
    created() {
      this.loadDiff()
      this.loadType()
      if(this.type == "updata")
        this.formInfo = this.oldInfo
    }
}
</script>

<style>

</style>