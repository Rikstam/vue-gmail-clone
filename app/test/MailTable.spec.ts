import { describe, it, expect, test } from 'vitest'
import { mount } from '@vue/test-utils'
import MailTable from '../src/components/MailTable.vue'


test('mount component', async () => {
    expect(MailTable).toBeTruthy()

    const wrapper = mount(MailTable, {})
})  
