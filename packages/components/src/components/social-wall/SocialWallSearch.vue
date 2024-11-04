<template>
  <div class="relative w-full max-w-md mx-auto">
    <input
      v-model="query"
      type="text"
      class="w-full p-2 border border-gray-300 rounded"
      placeholder="Search..."
      @input="onInput"
      @focus="showSuggestions = true"
      @blur="hideSuggestions"
    />
    <ul v-if="showSuggestions && filteredSuggestions.length" class="absolute z-10 w-full bg-white border border-gray-300 rounded shadow-lg">
      <li v-for="(suggestion, index) in filteredSuggestions" :key="index" class="p-2 hover:bg-gray-100 cursor-pointer" @mousedown.prevent="selectSuggestion(suggestion)">
        {{ suggestion }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { RouterView } from 'vue-router'

const props = defineProps<{
  // Customizable placeholder text for the search input.
  placeholder: string
  // Delay time in milliseconds before search results are fetched.
  debounceTime: number
  // Callback function to handle the search query input by the user.
  onSearch: (query: string) => void
}>()

const query = ref('')
const showSuggestions = ref(false)

const suggestions = ref<string[]>(['Apple', 'Banana', 'Cherry', 'Date', 'Elderberry', 'Fig', 'Grape', 'Honeydew'])

const filteredSuggestions = computed(() => {
  if (!query.value) {
    return []
  }
  const lowerCaseQuery = query.value.toLowerCase()
  return suggestions.value.filter((suggestion) => suggestion.toLowerCase().includes(lowerCaseQuery))
})

const onInput = (e) => {
  props.onSearch(e)
  showSuggestions.value = true
}

const hideSuggestions = () => {
  setTimeout(() => {
    showSuggestions.value = false
  }, 100)
}

const selectSuggestion = (suggestion: string) => {
  query.value = suggestion
  showSuggestions.value = false
}
</script>

<style>
@import url('@/style.css');
</style>
