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

<script>
import useEmailSelection from '../composables/use-email-selection';
import { computed } from 'vue';
export default {
    setup(props) {
        const emailSelection = useEmailSelection()
        const numberSelected =  computed(() => emailSelection.emails.size)
        const numberEmails = props.emails.length
        const allEmailsSelected = computed(() => numberSelected.value === numberEmails) // need to reference with .value as numberSelected is a reactive property
        const someEmailsSelected = computed(() => numberSelected.value > 0 && numberSelected.value < numberEmails)
        
        const bulkSelect = () => {
            if (allEmailsSelected.value) {
                emailSelection.clear()
            } else {
                emailSelection.addMultiple(props.emails)
            }
        }
        return {
            allEmailsSelected,
            someEmailsSelected,
            bulkSelect,
            emailSelection,
            numberSelected
        }
    },
    props: {
        emails: {
            type: Array,
            required: true
        }
    }
}
</script>