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
  
<script>
import MailTable from './components/MailTable.vue';

export default {
  name: 'App',
  components: {
    MailTable,
  }
};
</script>

<style>
</style>
