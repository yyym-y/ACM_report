<!-- 
    添加题目的来源类型(点击后展示该表单)
 -->
<template>
    <div>
      <el-dialog
          :visible.sync="vis" width="500px" title="添加类型">
          <el-form :model="form" label-width="80px">
              <el-form-item label="类型名称">
                  <el-input v-model="forms.type_name"></el-input>
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
                  type_name : ""
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