import { mount } from '@vue/test-utils'
import ModalView from '../src/components/ModalView.vue'

test('mount component', async () => {
    expect(ModalView).toBeTruthy()

    const wrapper = mount(ModalView, {})
})  