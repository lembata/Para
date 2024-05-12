<script setup>
import { ref, defineModel, defineProps } from 'vue';
import Calendar from 'primevue/calendar';
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const props = defineProps(['text', 'placeholder', 'ariaLabel', 'showOnFocus', 'disabled']);
const model = defineModel();

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <Calendar class="w-full" v-model="model" :aria-label :placeholder :showOnFocus="true" :disabled="disabled">
    </Calendar>
    <p class="hidden text-xs italic text-red-500">Please fill out this field.</p>
  </div>
</template>
