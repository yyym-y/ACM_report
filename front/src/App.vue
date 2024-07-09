<template>
  <div>
    <el-container>
        <el-aside width="200px">
            <leftViewVue></leftViewVue>
        </el-aside>
        <el-container>
            <el-header>
              <headerViewVue></headerViewVue>
            </el-header>
            <el-main>
              <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
  </div>
</template>

<script>
import leftViewVue from './components/leftPart/leftView.vue'
import headerViewVue from './components/rightPart/header/headerView.vue'
export default {
  name: 'App',
  components: {
    leftViewVue, headerViewVue
  }
}
</script>

<style>
  .el-header {
    color: #333;
    line-height: 60px;
    background-color: #f7f8f9;
    min-width: 700px;
  }
  .el-aside {
    color: #333;
    text-align: center;
    line-height: 200px;
    height: 100vh;
    border-right-style: solid;
    border-right-color: #C0C4CC;
    border-right-width: 1px;

  }
  .el-main {
    color: #333;
    background-color: #f7f8f9;
  }
</style>
