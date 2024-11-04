<template>
  <div id="social-wall">
    <slot />
  </div>
</template>

<script setup lang="ts">
defineProps<{
  // Defines the theme of the social wall (e.g., 'light' or 'dark') for customization.
  theme: 'light' | 'dark'
  // Specifies the layout style for displaying videos (e.g., 'grid' or 'list').
  layout: 'grid' | 'list'
}>()
</script>

<style>
@import url('@/style.css');
</style>
