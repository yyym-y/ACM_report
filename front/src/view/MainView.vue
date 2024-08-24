<template>
  <div>
    <div class="inline">
      <addProblemButton></addProblemButton>
      <exportButtom class="margin-ll"></exportButtom>
    </div>
    <problemTable></problemTable>
  </div>
</template>

<script>
import problemTable from '@/components/rightPart/main/problem/problemTable.vue';
import addProblemButton from '@/components/rightPart/main/problem/addProblemButton.vue';
import exportButtom from '@/components/rightPart/main/export/exportButtom.vue';
export default {
    components : {
      problemTable, addProblemButton, exportButtom
    }
}
</script>

<style>
.inline {
  display: flex;
  flex-wrap: nowrap;
}
.margin-ll {
  margin-left: 16px;
}
</style>