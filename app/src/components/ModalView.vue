<template>
<div class="modal">
    <div class="overlay" @click="emit('closeModal')"></div>
        <div class="modal-card">
        <slot></slot>
    </div>
</div>
</template>

<script>
import { onBeforeMount } from 'vue';
export default {
    setup(props, { emit }) {
        const onKeydown = (event) => {

            if (event.key == 'Escape') {
                emit('closeModal');
            }
        }
        window.addEventListener('keydown', onKeydown)
        onBeforeMount(()=> {
            window.removeEventListener('keydown', onKeydown)
        })
        return {
            emit
        }
    }
}
</script>

<style scoped>
</style>