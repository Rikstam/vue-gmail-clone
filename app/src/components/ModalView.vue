<template>
<div class="modal">
    <div class="overlay" @click="emit('closeModal')"></div>
        <div class="modal-card">
        <slot></slot>
    </div>
</div>
</template>

<script setup>

import useKeyDown from '../composables/use-keydown.js'
import { defineEmits } from 'vue'

const emit = defineEmits(['closeModal'])

useKeyDown([
    {key:'Escape',fn: () => {emit('closeModal')}},
    {key:'Enter', fn: () => {console.log('pressed Enter')}}
])
</script>

<style scoped>
</style>