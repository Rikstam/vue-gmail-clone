<script setup>
import MailTable from './components/MailTable.vue';
import useEmailSelection from './composables/use-email-selection';

const emailSelection = useEmailSelection()

</script>
<template>
  <h1>VMail Inbox</h1>
  <Suspense>
    <template #default>
      <MailTable />
    </template>
    <template #fallback>
      Loading...
    </template>
  </Suspense>
</template>

<style>
</style>
