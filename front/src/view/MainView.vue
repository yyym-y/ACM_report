<template>
  <div>
    <addProblemButton></addProblemButton>
    <problemTable></problemTable>
  </div>
</template>

<script>
import problemTable from '@/components/rightPart/main/problem/problemTable.vue';
import addProblemButton from '@/components/rightPart/main/problem/addProblemButton.vue';
export default {
    components : {
      problemTable, addProblemButton
    }
}
</script>

<style>

</style>