<script setup>
import { ref, defineModel, defineProps } from 'vue';
import InputText from 'primevue/inputtext';
import { useI18n } from 'vue-i18n';
const { t } = useI18n(); 

const model = defineModel();
const props = defineProps(['text']);//, 'name', 'placeholder', "aria-label", "modelValue"]);

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <InputText class="w-full" v-model="model" :aria-label :placeholder />
    <p class="hidden text-xs italic text-red-500">Please fill out this field.</p>
  </div>
</template>
