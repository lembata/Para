<script setup>
import { ref, defineModel, defineProps } from 'vue';
import InputSwitch from 'primevue/inputswitch';
import { useI18n } from 'vue-i18n';
const { t } = useI18n(); 

const model = defineModel();
const props = defineProps(['text', 'ariaLabel', 'disabled']);

</script>

<template>
  <div class="form-input">
    <label class="form-label" for="grid-first-name">
      {{ $t(text) }}: {{ model }}
    </label>
    <InputSwitch v-model="model" :aria-label :disabled/>
  </div>
</template>
