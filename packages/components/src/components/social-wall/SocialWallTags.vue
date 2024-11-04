<template>
  <div class="w-full max-w-md mx-auto">
    <div class="flex flex-wrap items-center">
      <button
        v-for="(tag, index) in tags"
        :key="index"
        :class="['m-1 px-3 py-1 rounded-full text-sm font-medium', selectedTags.includes(tag) ? 'bg-blue-500 text-white' : 'bg-gray-200 text-gray-800 hover:bg-gray-300']"
        @click="toggleTag(tag)"
      >
        {{ tag }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { createApp, h, ref } from 'vue'

const props = defineProps<{
  // Callback function triggered when a tag is clicked, allowing custom tag-based filtering.
  onTagClick: (tag: string) => void
}>()

// Sample data for tags
const tags = ref<string[]>(['JavaScript', 'Vue.js', 'React', 'Angular', 'TypeScript', 'Node.js', 'CSS', 'HTML'])

// Reactive state to track selected tags
const selectedTags = ref<string[]>([])

// Method to toggle tag selection
const toggleTag = (tag: string) => {
  props.onTagClick && props.onTagClick(tag)

  const index = selectedTags.value.indexOf(tag)
  if (index === -1) {
    selectedTags.value.push(tag)
  } else {
    selectedTags.value.splice(index, 1)
  }
}
</script>

<style>
@import url('@/style.css');
</style>
