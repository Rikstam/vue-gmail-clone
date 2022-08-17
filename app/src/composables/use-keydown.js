import { onBeforeUnmount } from 'vue';

const useKeyDown = (keyCombos) => {
    const onKeydown = (event) => {
        console.log(event.key)
        const kc = keyCombos.find(kc => kc.key == event.key)
        if (kc) {
           kc.fn()
        }
    }
   
   window.addEventListener('keydown', onKeydown)
    
    onBeforeUnmount(()=> {
        window.removeEventListener('keydown', onKeydown)
    })
}

export default useKeyDown