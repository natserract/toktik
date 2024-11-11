<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { Skeleton, SkeletonTheme } from 'vue-loading-skeleton'

import { FeedsAPI } from '@apis/feeds'
import { useSocialWallContext } from '@components/social-wall/context'
import SocialWallVideoPlayer from '@components/social-wall/SocialWallVideoPlayer.vue'

const props = defineProps<{
  id?: string
  controls?: boolean
  onPlay?: (videoId: string) => void | undefined
  onPause?: (videoId: string) => void | undefined
}>()

const context = useSocialWallContext()
const videos = computed(() => context.getFeeds())
const loading = computed(() => context.state.feeds.loading)
const error = computed(() => !!context.state.feeds.error)

const feedsApi = new FeedsAPI(context.state.baseUrl)

const fetchVideos = async () => {
  try {
    const response = await feedsApi.searchVideos('Trending', 10)
    context.setFeeds(response.data)
  } catch (err: any) {
    context.state.feeds.error = err.message
    console.error('Error fetching feeds:', err)
  } finally {
    context.state.feeds.loading = false
  }
}

if (context.getFeeds().length === 0) {
  onMounted(fetchVideos)
}
</script>

<template>
  <SkeletonTheme v-if="loading" class="max-w-md" color="#e0e0e0">
    <Skeleton :count="8" height="100px" />
  </SkeletonTheme>
  <div v-else-if="error">Error: {{ error }}</div>
  <div v-else>
    <div v-if="videos && videos.length">
      <div :id="props.id ? props.id : ''" class="max-w-md flex flex-col gap-10">
        <div v-for="(video, index) in videos" :key="index" class="bg-white rounded-lg shadow-md overflow-hidden">
          <SocialWallVideoPlayer
            :url="feedsApi.getStreamVideoUrl(video.id || video.video_id)"
            :cover="video.cover"
            :controls="props?.controls"
            @play="props?.onPlay && props?.onPlay(video.id || video.video_id)"
            @pause="props?.onPause && props?.onPause(video.id || video.video_id)"
          />
        </div>
      </div>
    </div>
    <div v-else>Videos not found!</div>
  </div>
</template>

<style>
@import url('@/style.css');
@import url('vue-loading-skeleton/dist/style.css');
</style>
