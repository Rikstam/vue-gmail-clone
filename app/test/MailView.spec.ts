import { describe, it, expect } from "vitest"
import { mount } from '@vue/test-utils'
import { format } from 'date-fns';

import MailView from '../src/components/MailView.vue'

const mockEmail = {
    "id": 1,
    "from": "team@vuemastery.com",
    "subject": "What's up with Vue 3.0? Here's how to find out from Evan You",
    "body": "The opening keynote of VueConf US this year was Evan You (the creator of Vue), giving his State of the Vuenion address. He walked us through the journey of getting Vue 3 from a prototype to a reality the past year. He also dove into Vue's overall growth in the community.",
    "sentAt": "2020-03-27T18:25:43.511Z",
    "archived": false,
    "read": false
}

describe('MailView', () => {
    it('renders properly', () => {
        const wrapper = mount(MailView, {
            props: {
                email: mockEmail
            }
        })
        const subject = wrapper.get('[data-test="email-subject"]')
        const from = wrapper.get('[data-test="email-from"]')
        const body = wrapper.get('[data-test="email-body"]')
        expect(wrapper.classes()).toContain('email-display')
        expect(from.text()).toBe(`From ${mockEmail.from} on ${format(new Date(mockEmail.sentAt), 'MMM do yyyy')}`)
        expect(subject.text()).toBe(`Subject: ${mockEmail.subject}`)
        expect(body.text()).toBe(`${mockEmail.body}`)
    })
})