import { onBeforeUnmount } from 'vue';

type KeyBoardKey = 'r' | 'e' | 'k' | 'j' | '[' | ']'

type keyCombo = {
    key: KeyBoardKey;
    fn: () => void;
}

const useKeyDown = (keyCombos: keyCombo[]) => {
    const onKeydown = (event: KeyboardEvent) => {
        console.log(event.key)
        const kc = keyCombos.find(kc => kc.key == event.key)
        if (kc) {
            kc.fn()
        }
    }

    window.addEventListener('keydown', onKeydown)

    onBeforeUnmount(() => {
        window.removeEventListener('keydown', onKeydown)
    })
}

export default useKeyDown