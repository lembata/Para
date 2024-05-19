<script setup>
import {watch} from 'vue';
//import {useI18n} from 'vue-i18n';
import {useRoute} from 'vue-router'

//const {t} = useI18n();
const route = useRoute();

const loadInitial = async () => {
  console.log('loadInitial on account page');
}

watch(() => 0, loadInitial, {immediate: true})
</script>

<template>
  <main>
    Here will be the data for account: {{ route.params.id }}
  </main>
</template>