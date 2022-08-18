<template>
<div class="email-display">
    <div>
        <button @click="toggleArchive">{{email.archived ? 'Move to Inbox (e)' : 'Move to Archive (e)'}}</button>
        <button @click="toggleRead">{{email.read ? 'Mark Unread (r)' : 'Mark Read (r)'}}</button>
        <button @click="goNewer">Newer (k)</button>
        <button @click="goOlder">Older (j)</button>
    </div>
    <h2 class="mb-0">Subject: <strong>{{email.subject}}</strong></h2>
    <div><em>From {{email.from}} on  {{format(new Date(email.sentAt), 'MMM do yyyy')}} </em></div>
    <div v-html="marked.parse(email.body)"></div>
</div>
</template>

<script>
import { format } from 'date-fns';
import { marked } from 'marked';
import axios from 'axios';

import useKeyDown from '../composables/use-keydown';

export default {
    setup(props) {
        const email = props.email
        const toggleRead = () => {
            email.read = !email.read
            axios.put(`http://localhost:1323/emails/${email.id}`, email);
        }
        const toggleArchive = () => {
            email.archived = !email.archived
            axios.put(`http://localhost:1323/emails/${email.id}`, email);
        }

        const goNewer = () => {
            console.log('go newer')
        }

        const goOlder = () => {
            console.log('go older')
        }
        useKeyDown([
            {key: 'r', fn: toggleRead},
            {key: 'e', fn: toggleArchive},
        ])
        return {
            format,
            toggleRead,
            toggleArchive,
            goNewer,
            goOlder,
            marked
        }
    },
    props: {
        email: {
            type: Object,
            required: true,
        }
    }
}
</script>