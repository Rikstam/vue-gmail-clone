<template>
    <div class="bulk-action-bar">
       <span class="checkbox">
         <input type="checkbox"
                :checked="allEmailsSelected"
                :class="[someEmailsSelected ? 'partial-check': '']"
                @click="bulkSelect" />
       </span>
       <span class="buttons">
          <button @click="emailSelection.markRead()"
                  :disabled="[...emailSelection.emails].every(e => e.read)">
                  Mark Read</button> <!-- splat the emails set into an array with ... -->
          
          <button @click="emailSelection.markUnread()"
                  :disabled="[...emailSelection.emails].every(e => !e.read)"
          >Mark Unread</button>
          
          <button @click="emailSelection.archive()"
                  :disabled="numberSelected === 0">Archive</button>  
       </span>
    </div>
</template>

<script setup lang="ts">
import useEmailSelection from '../composables/use-email-selection';
import { computed, defineProps } from 'vue';
import { EmailItem } from '../types'


const props = defineProps<{emails: EmailItem[]}>()

const emailSelection = useEmailSelection()
const numberSelected =  computed(() => emailSelection.emails.size)
const numberEmails = computed(() => props.emails.length)
const allEmailsSelected = computed(() => numberSelected.value === numberEmails.value) // need to reference with .value as numberSelected and numberEmails arereactive properties
const someEmailsSelected = computed(() => numberSelected.value > 0 && numberSelected.value < numberEmails.value)

const bulkSelect = () => {
    if (allEmailsSelected.value) {
        emailSelection.clear()
    } else {
        emailSelection.addMultiple(props.emails)
    }
}
</script>