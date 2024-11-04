<template>
  <div :id="props.id ? props.id : ''" class="max-w-md mx-auto overflow-y-auto h-screen">
    <div v-for="(video, index) in videos" :key="index" class="bg-white rounded-lg shadow-md overflow-hidden" @click="openModal(video)">
      <v-lazy-image
        :src="video.thumbnail"
        src-placeholder="https://cdn-images-1.medium.com/max/80/1*xjGrvQSXvj72W4zD6IWzfg.jpeg"
        :alt="video.title"
        class="w-full h-48 object-cover"
      />
      <div class="p-4">
        <h2 class="text-lg font-semibold">{{ video.title }}</h2>
        <p class="text-gray-600 text-sm">{{ video.description }}</p>
      </div>
    </div>
  </div>

  <SocialWallVideoView :is-open="isModalOpen" @update:isOpen="isModalOpen = $event" />
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import VLazyImage from 'v-lazy-image'
import SocialWallVideoView from '@components/social-wall/SocialWallVideoView.vue'
import videosAPI from '@apis/videos'

const props = defineProps<{
  id?: string
  // Sets the maximum number of videos to fetch and display at a time.
  maxVideos: number
  // Defines the layout of the videos (e.g., grid, masonry, or carousel).
  gridLayout: 'grid' | 'masonry' | 'carousel'
  // Boolean to show or hide playback controls like play, pause, and volume.
  controls?: boolean
  // Callback for custom handling of video play events.
  onPlay?: (videoId: string) => void
  // Callback for custom handling of video pause events.
  onPause?: (videoId: string) => void
  // Enables or disables autoplay on scroll into view for each video.
  autoplay?: boolean
}>()

const isModalOpen = ref(false)

function openModal(video: any) {
  console.log('video', video)
  isModalOpen.value = true
}

// Sample data for videos
const videos = ref([
  {
    title: 'Vue.js Basics',
    description: 'Learn the basics of Vue.js in this introductory video.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
  {
    title: 'Advanced Vue Techniques',
    description: 'Explore advanced techniques in Vue.js development.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
  {
    title: 'Building a Vue App',
    description: 'Step-by-step guide to building a Vue.js application.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
  {
    title: 'State Management with Vuex',
    description: 'Manage state in your Vue.js apps using Vuex.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
  {
    title: 'Vue Router for Navigation',
    description: 'Implement navigation in your Vue.js apps with Vue Router.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
  {
    title: 'Deploying Vue Apps',
    description: 'Learn how to deploy your Vue.js applications.',
    thumbnail: 'https://cdn-images-1.medium.com/max/1600/1*xjGrvQSXvj72W4zD6IWzfg.jpeg',
  },
])

const posts = ref([])

const fetchPosts = async () => {
  try {
    const response = await videosAPI.listPosts()
    console.log(response)
  } catch (error) {
    console.error('Error fetching posts:', error)
  }
}
onMounted(fetchPosts)
</script>

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
