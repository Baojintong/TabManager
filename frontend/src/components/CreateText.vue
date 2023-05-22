<script setup>
let props = defineProps(['title', 'btnType', 'data'])
</script>

<template>
  <main>
    <a-button :type="btnType" @click="dialogVisible = true" size="large">{{title}}</a-button>
  </main>
</template>

<style scoped>
</style>
