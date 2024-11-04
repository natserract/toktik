<template>
  <main>
    <div class="video-container">
      <h1>Video Streaming Example</h1>
      <video ref="videoPlayer" controls width="600" :src="videoSrc" @play="onPlay" @pause="onPause" @loadedmetadata="onLoadedMetadata">
        Your browser does not support the video tag.
      </video>
    </div>
  </main>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'

export default {
  name: 'VideoPlayer',
  setup() {
    // Set the source of the video to the Go server endpoint
    const videoSrc = ref('http://localhost:8080/api/v1/feeds/7432034974246669611/stream')
    const videoPlayer = ref(null)

    // Methods for handling video events
    const onPlay = () => console.log('Video playing...')
    const onPause = () => console.log('Video paused.')
    const onLoadedMetadata = () => {
      console.log('Video metadata loaded.')
    }

    // On mount, you could check if the video is accessible or log info
    onMounted(() => {
      console.log('Video component mounted.')
    })

    return {
      videoSrc,
      videoPlayer,
      onPlay,
      onPause,
      onLoadedMetadata,
    }
  },
}
</script>

<style scoped>
.video-container {
  text-align: center;
  margin-top: 20px;
}
</style>
