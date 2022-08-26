
import {reactive} from 'vue';
import axios from 'axios';

const emails = reactive(new Set());

export const useEmailSelection = function() {
    const toggle = function(email) {
        if(emails.has(email)) {
            emails.delete(email);
        } else {
            emails.add(email);
        }
    }
    const clear = () => {
        emails.clear() // set.clear()
    }

    const forSelected = (fn) => {
        emails.forEach((email) => {
            fn(email)
            axios.put(`http://localhost:1323/emails/${email.id}`, email);
        })
    }

    const addMultiple = (newEmails) => {
        newEmails.forEach((email) => {
            emails.add(email)
        })
    }
    const markRead = () =>  forSelected(email => email.read = true)
    const markUnread = () => forSelected(email => email.read = false)
    const archive = () => { forSelected(email => email.archived = true); clear();}

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