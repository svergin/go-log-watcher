<script setup lang="ts">
defineProps<{
  logcontent: string;
}>();
</script>

<template>
  <div class="logwelcome">
    <h1 class="green">Logfile</h1>
  </div>
  <div>
    <input type="text" v-model="logcontent" />
  </div>
</template>