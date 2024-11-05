<script lang="ts" setup>
defineProps<{
  loading?: boolean
  error?: boolean
}>()
</script>

<template>
  <div>
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">Error: {{ error }}</div>
    <div v-else>
      <slot />
    </div>
  </div>
</template>
