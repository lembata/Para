<script setup>
import { ref, defineModel, defineProps } from 'vue';
import Dropdown from 'primevue/dropdown';
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const props = defineProps(['text', 'invalid', 'aria-lavel']);
const model = defineModel('');

const currencies = [
  { value: 'EUR', name: 'Euro', symbol: '€' },
  { value: 'USD', name: 'Dollars', symbol: '$' },
  { value: 'BGN', name: 'Bulgarian Lev', symbol: 'лв.' }
]

</script>

<template>
  <div class="form-input">
    <label class="form-label"> {{ $t(text) }} {{ model }}</label>
    <div class="relative">
      <Dropdown v-model="model" :options="currencies" optionLabel="name" :invalid :aria-label
        placeholder="Select a currency" class="w-full md:w-14rem" />
    </div>
  </div>
</template>
