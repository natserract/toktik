<script setup lang="ts">
import { onMounted, computed, ref } from 'vue'
import { RecommendationsAPI } from '@apis/recommendations'
import { FeedsAPI } from '@apis/feeds'
import { useSocialWallContext } from '@components/social-wall/context'
import AsyncDataComponent from '@components/AsyncDataComponent.vue'

const props = defineProps<{
  id?: string
  onTagClick?: (tag: string) => void
}>()

const context = useSocialWallContext()
const tags = computed(() => context.getTags())
const loading = computed(() => context.state.tags.loading)
const error = computed(() => !!context.state.tags.error)

const feedsApi = new FeedsAPI(context.state.baseUrl)
const recommendationsApi = new RecommendationsAPI(context.state.baseUrl)

const onClick = async (tag: string) => {
  props.onTagClick && props.onTagClick(tag)

  const videos = await fetchVideosByTag(tag)
  if (videos) {
    context.setFeeds(videos.data)
  }
}

const fetchVideosByTag = async (tag: string) => {
  try {
    context.state.feeds.loading = true
    return await feedsApi.searchVideos(tag, 5)
  } catch (err: any) {
    context.state.feeds.error = err.message
    console.error('Error fetching feeds:', error)
  } finally {
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

if (context.getTags().length === 0) {
  onMounted(fetchRecommendationTags)
}
</script>

<template>
  <AsyncDataComponent :loading="loading" :error="error">
    <div :id="props.id ? props.id : ''" class="w-full max-w-md mx-auto">
      <div class="flex flex-wrap items-center">
        <button
          v-for="(tag, index) in tags"
          :key="index"
          :class="['m-1 px-3 py-1 rounded-full text-sm font-medium bg-gray-200 text-gray-800 hover:bg-gray-300']"
          @click="onClick(tag)"
        >
          {{ tag }}
        </button>
      </div>
    </div>
  </AsyncDataComponent>
</template>

<style>
@import url('@/style.css');
</style>
