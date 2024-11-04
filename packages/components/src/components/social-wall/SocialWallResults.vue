<template>
  <div class="max-w-md mx-auto overflow-y-auto h-screen">
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

  <SocialWallDialog :is-open="isModalOpen" @update:isOpen="isModalOpen = $event" />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import VLazyImage from 'v-lazy-image'
import SocialWallDialog from '@components/social-wall/SocialWallDialog.vue'

const open = ref(false)
const isModalOpen = ref(false)

function openModal(video: any) {
  console.log('video', video)
  isModalOpen.value = true
}

// Sample data for videos
const videos = [
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
]

const items = ref(
  videos.map((v, idx) => ({
    ...v,
    id: idx,
  }))
)
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
