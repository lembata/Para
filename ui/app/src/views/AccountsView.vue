<script setup>
import { ref, watch } from 'vue';
import AccountForm from '../components/forms/AccountForm.vue';
import API from '@/api/api.js';


const loadInitial = async () => {
  getPage(1);
}

const getPage = async (page) => {
  const limit = 10;
  API.Accounts.all(limit, (page - 1) * limit)
    .then((result) => {
      if (result.success) {
        //accounts.value = result.data;
        console.log('accounts', result.data);
      } else {
        console.error('Failed to load accounts', result);
        toast.add({ severity: 'error', summary: t('notifications.failedToLoadAccounts'), life: 3000 });
      }
    });
}

watch(() => loadInitial());
</script>

<template>
  <main>
    "This will be the accounts view."
  </main>
</template>
