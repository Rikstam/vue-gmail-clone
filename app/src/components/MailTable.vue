<template>
<table class="mail-table">
    <tbody>
        <tr v-for="email in unarchivedEmails"
            :key="email.id"
            :class="['clickable', email.read ? 'read' : '']"
            @click="readEmail(email)">
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
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { format } from 'date-fns';

export default {
    async setup() {

    const {data: emails} = await axios.get('http://localhost:1323/emails/');
   
    return {
      format,
      emails: ref(emails),
    }
  },
  methods: {
    readEmail(email) {
      email.read = true;
      this.updateEmail(email);
       
    },
    archiveEmail(email) {
      email.archived = true;
      this.updateEmail(email);
    },
    updateEmail(email) {
      axios.put(`http://localhost:1323/emails/${email.id}`, email)
    }
  },
  computed: {
    sortedEmails() {
      return this.emails.sort((e1, e2) => {
        return e1.sentAt < e2.sentAt ? 1 : -1
      })
    },
    unarchivedEmails() {
      return this.sortedEmails.filter(e => !e.archived)
    }
  }
}
</script>