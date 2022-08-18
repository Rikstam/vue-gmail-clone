<template>
  <h1>VMail Inbox</h1>
  <h1>{{emailSelection.emails.size}} emails selected</h1>

  <Suspense>
    <template #default>
      <MailTable />
    </template>
    <template #fallback>
      Loading...
    </template>
  </Suspense>
</template>
  
<script>
import MailTable from './components/MailTable.vue';
import useEmailSelection from './composables/use-email-selection';
export default {
  name: 'App',
  components: {
    MailTable,
  },
  setup() {
    return {
        emailSelection: useEmailSelection()
      }
  }
};
</script>

<style>
</style>
