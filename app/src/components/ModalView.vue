<template>
<div class="modal">
    <div class="overlay" @click="emit('closeModal')"></div>
        <div class="modal-card">
        <slot></slot>
    </div>
</div>
</template>

<script>

import useKeyDown from '../composables/use-keydown.js'

export default {
    setup(props, { emit }) {
        useKeyDown([
            {key:'Escape',fn: () => {emit('closeModal')}},
            {key:'Enter', fn: () => {console.log('pressed Enter')}}
        ])
        return {
            emit
        }
    }
}
</script>

<style scoped>
</style>