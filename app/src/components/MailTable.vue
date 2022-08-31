<template>

<button @click="selectScreen('inbox')"
        :disabled="selectedScreen == 'inbox'"
>Inbox</button>

<button @click="selectScreen('archive')"
        :disabled="selectedScreen == 'archive'"
>Archived</button>

<BulkActionBar :emails="filteredEmails" />
<table class="mail-table">
    <tbody>
        <tr v-for="email in filteredEmails"
            :key="email.id"
            :class="['clickable', email.read ? 'read' : '']"
            >
          <td><input type="checkbox"
                    @click="emailSelection.toggle(email)"
                    :checked="emailSelection.emails.has(email)"/>
          </td>
          <td @click="openEmail(email)">{{email.from}}</td>
          <td @click="openEmail(email)">
            <p><strong>{{email.subject}}</strong></p>
          </td>
          <td class="date" @click="openEmail(email)">{{format(new Date(email.sentAt), 'MMM do yyyy')}}</td>
          <td><button @click="archiveEmail(email)">Archive</button></td>
          </tr>
    </tbody>
  </table>
  <ModalView v-if="openedEmail" @closeModal="openedEmail = null">
    <MailView :email="openedEmail" @changeEmail="changeEmail"/>
  </ModalView>
</template>

<script setup lang="ts">
import { ref, computed, reactive } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';
import MailView from '../components/MailView.vue';
import ModalView from '../components/ModalView.vue';
import useEmailSelection from '../composables/use-email-selection'
import BulkActionBar from './BulkActionBar.vue';

import { EmailItem, changeEmailEventParams } from '../types';

const result = await axios.get("http://localhost:1323/emails/");
const emails = reactive<EmailItem[]>(result.data)
const openedEmail = ref<EmailItem | null>(null)
const emailSelection = useEmailSelection()
const selectedScreen = ref('inbox')

const selectScreen = (newScreen: string) => {
        selectedScreen.value = newScreen
        emailSelection.clear()
}

const openEmail = (email: EmailItem) => {
    openedEmail.value = email;
    if (email) {
        email.read = true;
        updateEmail(email);
    }
}

const archiveEmail = (email: EmailItem) => {
    email.archived = true;
    updateEmail(email);
}

const updateEmail = (email: EmailItem) => {
    axios.put(`http://localhost:1323/emails/${email.id}`, email);
}

const changeEmail = ({toggleRead, toggleArchive, save, closeModal, changeIndex}: changeEmailEventParams) => {
    if (openedEmail) {
        const email = openedEmail;
        if(toggleRead) { email.value.read = !email.value.read }

        if(toggleArchive) { email.value.archived = !email.archived }

        if(save) { updateEmail(email.value)}
    }
    if(closeModal) { openedEmail.value = null }

    if (changeIndex) {
        const currentIndex = filteredEmails.value.indexOf(openedEmail.value)
        console.log(currentIndex)
        console.log(changeIndex)

        const newEmail = filteredEmails.value[currentIndex + changeIndex]
       
        openEmail(newEmail);
    }
}
    
const sortedEmails = computed(() => {
    return emails.sort((e1, e2) => {
        return (e1.sentAt < e2.sentAt) ? 1 : -1;
    })
})

const filteredEmails = computed(() => {
    if(selectedScreen.value == 'inbox') {
        return sortedEmails.value.filter(e => !e.archived);
    } else {
        return sortedEmails.value.filter(e => e.archived);
    }
})
    

</script>