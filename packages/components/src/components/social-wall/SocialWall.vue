<script setup lang="ts">
import { provideSocialWallContext } from '@components/social-wall/context'

const props = defineProps<{
  id?: string
  baseUrl: string
}>()

if (!props.baseUrl) {
  throw new Error('Please provide base url. (e.g. http://localhost:8080/api/v1')
}
provideSocialWallContext(props.baseUrl)
</script>

<template>
  <div :id="`social-wall ${props.id ? props.id : ''}`">
    <slot />
  </div>
</template>

<style>
@import url('@/style.css');
</style>
