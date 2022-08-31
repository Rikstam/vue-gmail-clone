
import { reactive } from 'vue';
import axios from 'axios';
import { EmailItem } from '../types';

const emails = reactive(new Set());

export const useEmailSelection = function () {
    const toggle = function (email: EmailItem) {
        if (emails.has(email)) {
            emails.delete(email);
        } else {
            emails.add(email);
        }
    }
    const clear = () => {
        emails.clear() // set.clear()
    }

    const forSelected = ((fn) => {
        emails.forEach((email) => {
            fn(email)
            axios.put(`http://localhost:1323/emails/${email.id}`, email);
        })
    })

    const addMultiple = (newEmails: EmailItem[]) => {
        newEmails.forEach((email) => {
            emails.add(email)
        })
    }
    const markRead = () => forSelected((email: EmailItem) => email.read = true)
    const markUnread = () => forSelected((email: EmailItem) => email.read = false)
    const archive = () => { forSelected((email: EmailItem) => email.archived = true); clear(); }

    return {
        emails,
        toggle,
        clear,
        addMultiple,
        markRead,
        markUnread,
        archive
    }
}

export default useEmailSelection;