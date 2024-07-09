<!-- 
  左侧的布局组件 
-->
<template>
    <div>
      <iconVue></iconVue>
      <nav_menuVue></nav_menuVue>
    </div>
</template>
  
<script>
    import iconVue from './util/icon.vue'
    import nav_menuVue from './util/nav_menu.vue'
    export default {
        name : "Left_View",
        components : {
            iconVue, nav_menuVue
        }
    }
</script>
  
<style>
  
</style>