<script setup>
import { ref, defineModel, defineProps } from 'vue';
import InputNumber from 'primevue/inputnumber';
import { useI18n } from 'vue-i18n';
const { t } = useI18n(); 

const model = defineModel();
const props = defineProps(['text']);

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <InputNumber class="w-full" v-model="model" :aria-label :placeholder :min :max :minFranctionDigits />
  </div>
</template>
