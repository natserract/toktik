<script setup lang="ts">
import { ref, onMounted } from 'vue'
import VLazyImage from 'v-lazy-image'

import SocialWallVideoView from '@components/social-wall/SocialWallVideoView.vue'
import videosAPI from '@apis/videos'

const props = defineProps<{
  id?: string
  maxVideos: number
  gridLayout: 'grid' | 'masonry' | 'carousel'
  controls?: boolean
  onPlay?: (videoId: string) => void
  onPause?: (videoId: string) => void
  autoplay?: boolean
}>()

const isModalOpen = ref(false)

function openModal(video: any) {
  isModalOpen.value = true
}

const fetchPosts = async () => {
  try {
    const response = await videosAPI.searchVideos('obama', 2)
    console.log(response)
  } catch (error) {
    console.error('Error fetching posts:', error)
  }
}
onMounted(fetchPosts)
</script>

<template>
  <div :id="props.id ? props.id : ''" class="max-w-md mx-auto overflow-y-auto h-screen">
    <!--    <div v-for="(video, index) in videos" :key="index" class="bg-white rounded-lg shadow-md overflow-hidden" @click="openModal(video)">-->
    <!--      <v-lazy-image-->
    <!--        :src="video.thumbnail"-->
    <!--        src-placeholder="https://cdn-images-1.medium.com/max/80/1*xjGrvQSXvj72W4zD6IWzfg.jpeg"-->
    <!--        :alt="video.title"-->
    <!--        class="w-full h-48 object-cover"-->
    <!--      />-->
    <!--      <div class="p-4">-->
    <!--        <h2 class="text-lg font-semibold">{{ video.title }}</h2>-->
    <!--        <p class="text-gray-600 text-sm">{{ video.description }}</p>-->
    <!--      </div>-->
    <!--    </div>-->
  </div>

  <SocialWallVideoView :is-open="isModalOpen" @update:isOpen="isModalOpen = $event" />
</template>

<style>
@import url('@/style.css');

.v-lazy-image {
  filter: blur(10px);
  transition: filter 0.7s;
}
.v-lazy-image-loaded {
  filter: blur(0);
}
</style>
