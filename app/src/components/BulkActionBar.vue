<template>
    <div>
        <input type="checkbox"
                :checked="allEmailsSelected"
                :class="[someEmailsSelected ? 'partial-check': '']"
                @click="bulkSelect" />
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
            bulkSelect
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