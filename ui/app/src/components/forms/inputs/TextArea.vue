<script setup>
import { ref, defineModel, defineProps } from 'vue';
import Textarea from 'primevue/textarea';
import { useI18n } from 'vue-i18n';
const { t } = useI18n(); 

const model = defineModel();
const props = defineProps(['text', 'rows', 'cols', 'placeholder', 'ariaLabel', 'disabled']);

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <Textarea class="w-full" v-model="model" :aria-label :placeholder :rows :cols :disabled autoresize/>
    <p class="hidden text-xs italic text-red-500">Please fill out this field.</p>
  </div>
</template>
