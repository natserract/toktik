<script setup lang="ts">
import { ref } from 'vue'
import { useSocialWallContext } from '@components/social-wall/context'
import { RecommendationsAPI } from '@apis/recommendations'
import { FeedsAPI } from '@apis/feeds'
import debounce from '@utils/debounce'

const props = defineProps<{
  placeholder: string
  debounceTime?: number
  onSearch?: (query: string) => void
}>()

const query = ref('')
const context = useSocialWallContext()

const feedsApi = new FeedsAPI(context.state.baseUrl)
const recommendationsApi = new RecommendationsAPI(context.state.baseUrl)

const fetchVideos = async (query: string, count = 10) => {
  try {
    context.state.feeds.loading = true
    return await feedsApi.searchVideos(query, count)
  } catch (err: any) {
    context.state.feeds.error = err.message
    console.error('Error fetching feeds:', err)
  } finally {
    await fetchRecommendationTags()
    context.state.feeds.loading = false
  }
}

const fetchRecommendationTags = async () => {
  try {
    const response = await recommendationsApi.listTags()
    context.setTags(response.data.tags)
  } catch (err: any) {
    context.state.tags.error = err.message
    console.error('Error fetching tags:', err)
  } finally {
    context.state.tags.loading = false
  }
}

const onInput = async () => {
  props.onSearch && props.onSearch(query.value)

  try {
    if (query.value.length > 2) {
      const videos = await fetchVideos(query.value)
      if (videos) {
        context.setFeeds(videos.data)
      }
    }
  } catch (error) {
    console.error('Error fetching feeds:', error)
  }
}
const debouncedOnInput = debounce(onInput, props.debounceTime || 300)

const onEnter = async (e: any) => {
  try {
    const videos = await fetchVideos(e.target.value)
    if (videos) {
      context.setFeeds(videos.data)
    }
  } catch (error) {
    console.error('Error fetching feeds:', error)
  }
}
</script>

<template>
  <div class="relative w-full max-w-md mx-auto">
    <input
      v-model="query"
      type="text"
      class="w-full p-2 border border-gray-300 rounded"
      :placeholder="props.placeholder ? props.placeholder : 'Search...'"
      @input="debouncedOnInput"
      @keydown.enter="onEnter"
    />
  </div>
</template>

<style>
@import url('@/style.css');
</style>
