<template>
  <teleport to="body">
    <div v-if="isOpen" class="modal-overlay">
      <div class="modal-content">
        <slot></slot>
        <button @click="close">Close</button>
      </div>
    </div>
  </teleport>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue'

defineProps({
  isOpen: {
    type: Boolean,
    required: true,
  },
})

const emit = defineEmits(['update:isOpen'])

function close() {
  emit('update:isOpen', false)
}
</script>

<style>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 5px;
}
</style>
