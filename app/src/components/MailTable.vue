<template>
<table class="mail-table">
    <tbody>
        <tr v-for="email in unarchivedEmails"
            :key="email.id"
            :class="['clickable', email.read ? 'read' : '']"
            @click="openEmail(email)">
          <td><input type="checkbox"/></td>
          <td>{{email.from}}</td>
          <td>
            <p><strong>{{email.subject}}</strong></p>
          </td>
          <td class="date">{{format(new Date(email.sentAt), 'MMM do yyyy')}}</td>
          <td><button @click="archiveEmail(email)">Archive</button></td>
          </tr>
    </tbody>
  </table>
  <ModalView v-if="openedEmail" @closeModal="openedEmail = null">
    <MailView :email="openedEmail" @changeEmail="changeEmail"/>
  </ModalView>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';
import MailView from './MailView.vue';
import ModalView from './ModalView.vue';

export default {
    async setup() {
        const { data: emails } = await axios.get("http://localhost:1323/emails/");
        return {
            format,
            emails: ref(emails),
            openedEmail: ref(null),
        };
    },
    methods: {
        openEmail(email) {
             this.openedEmail = email;
            if (email) {
                email.read = true;
                this.updateEmail(email);
            }
        },
        archiveEmail(email) {
            email.archived = true;
            this.updateEmail(email);
        },
        updateEmail(email) {
            axios.put(`http://localhost:1323/emails/${email.id}`, email);
        },
        changeEmail({toggleRead, toggleArchive, save, closeModal, changeIndex}) {
            const email = this.openedEmail;
            
            if(toggleRead) { email.read = !email.read }
            if(toggleArchive) { email.archived = !email.archived }
            if(save) { this.updateEmail(email)}
            if(closeModal) { this.openedEmail = null }

            if (changeIndex) {
                const emails = this.unarchivedEmails;
                const currentIndex = emails.indexOf(this.openedEmail)
                const newEmail = emails[currentIndex + changeIndex]
                this.openEmail(newEmail);
            }
        }
    },
    computed: {
        sortedEmails() {
            return this.emails.sort((e1, e2) => {
                return (e1.sentAt < e2.sentAt) ? 1 : -1;
            });
        },
        unarchivedEmails() {
            return this.sortedEmails.filter(e => !e.archived);
        }
    },
    components: { MailView, ModalView }
}
</script>