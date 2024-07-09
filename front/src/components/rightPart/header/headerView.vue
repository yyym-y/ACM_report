<template>
  <div class="top-head">
    <search_barVue></search_barVue>
  </div>
</template>

<script>
import search_barVue from './search_bar.vue'
export default {
    components : {
        search_barVue
    }
}
</script>

<style scoped>
.top-head {
    height: 97%;
    border-bottom-style: solid;
    border-bottom-color: #C0C4CC;
    border-bottom-width: 1px;
}
</style>