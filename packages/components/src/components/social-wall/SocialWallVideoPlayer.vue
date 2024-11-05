<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'

const props = defineProps<{
  cover: string
  url: string
  controls?: boolean
  onPlay?: () => void
  onPause?: () => void
}>()

const videoPlayer = ref<HTMLVideoElement | null>(null)
const isPlaying = ref(false)
const currentTime = ref(0)
const duration = ref(0)
const showPoster = ref(true)

const togglePlay = (url: string) => {
  if (videoPlayer.value) {
    if (videoPlayer.value.paused) {
      videoPlayer.value.play()
      isPlaying.value = true
    } else {
      videoPlayer.value.pause()
      isPlaying.value = false
    }
  }
}

const updateProgress = () => {
  if (videoPlayer.value) {
    currentTime.value = videoPlayer.value.currentTime
  }
}

const setDuration = () => {
  if (videoPlayer.value) {
    duration.value = videoPlayer.value.duration
  }
}

const seek = () => {
  if (videoPlayer.value) {
    videoPlayer.value.currentTime = currentTime.value
  }
}

const formatTime = (time: number) => {
  const minutes = Math.floor(time / 60)
  const seconds = Math.floor(time % 60)
  return `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`
}

const onPlay = () => {
  props.onPlay && props.onPlay()
  showPoster.value = false
}

const onPause = () => {
  if (videoPlayer.value && videoPlayer.value.paused) {
    props.onPause && props.onPause()
    showPoster.value = true
  }
}

onMounted(() => {
  if (videoPlayer.value) {
    videoPlayer.value.addEventListener('play', () => (isPlaying.value = true))
    videoPlayer.value.addEventListener('pause', () => (isPlaying.value = false))
  }
})

watch(
  () => props.url,
  (newUrl) => {
    if (videoPlayer.value) {
      isPlaying.value = false
      videoPlayer.value.src = newUrl
      videoPlayer.value.load()
    }
  }
)

watch(
  () => props.cover,
  (newCover) => {
    if (videoPlayer.value) {
      videoPlayer.value.poster = newCover
    }
  }
)
</script>

<template>
  <div class="video-player relative w-full h-[468px] bg-gray-900">
    <video
      ref="videoPlayer"
      class="absolute inset-0 w-full h-full object-cover cursor-pointer"
      :poster="props.cover"
      :controls="props.controls"
      @timeupdate="updateProgress"
      @loadedmetadata="setDuration"
      @play="onPlay"
      @pause="onPause"
    >
      <source :src="props.url" type="video/mp4" />
      Your browser does not support the video tag.
    </video>

    <div v-if="!props.controls" class="controls absolute bottom-0 left-0 right-0 flex items-center justify-between p-2 bg-black bg-opacity-50 text-white">
      <button class="play-pause" @click="togglePlay(props.url)">{{ isPlaying ? 'Pause' : 'Play' }}</button>
      <div class="progress-bar flex-1 mx-2">
        <input v-model="currentTime" type="range" min="0" :max="duration" step="0.1" @input="seek" />
      </div>
      <span>{{ formatTime(currentTime) }} / {{ formatTime(duration) }}</span>
    </div>
  </div>
</template>

<style scoped>
@import url('@/style.css');
</style>
