<template>
  <div :id="`social-wall ${props.id ? props.id : ''}`">
    <slot />
  </div>
</template>

<script setup lang="ts">
import { provideSocialWallContext } from '@components/social-wall/context'

const props = defineProps<{
  id?: string
}>()

provideSocialWallContext()
</script>

<style>
@import url('@/style.css');
</style>
