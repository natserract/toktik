<script setup lang="ts">
import { SocialWall, SocialWallTags, SocialWallVideos, SocialWallSearch } from '@toktik/components'
import { defineComponent } from 'vue'

const BASE_URL = import.meta.env.VITE_BASE_API_URL
defineComponent({
  components: {
    'social-wall': SocialWall,
    'social-wall-search': SocialWallSearch,
    'social-wall-tags': SocialWallTags,
    'social-wall-videos': SocialWallVideos,
  },
})
</script>

<template>
  <main class="relative pb-20">
    <social-wall :base-url="BASE_URL">
      <header class="sticky top-0 z-40 w-full z-10 py-2 bg-[#003] border-[#00003370] overflow-hidden">
        <social-wall-search placeholder="Search video..."></social-wall-search>
      </header>

      <div class="max-w-7xl relative mx-auto px-4 sm:px-6 md:px-8">
        <div class="relative w-full max-w-md mx-auto">
          <div class="sticky bottom-0 top-[3.5rem] z-40 py-5 bg-white">
            <div class="flex flex-col">
              <h3 class="font-bold pl-2 mb-2">Tags:</h3>
              <social-wall-tags></social-wall-tags>
            </div>
          </div>
          <div class="max-w-3xl mx-auto pt-10 px-3">
            <social-wall-videos></social-wall-videos>
          </div>
        </div>
      </div>
    </social-wall>
  </main>
</template>

<style scoped></style>
