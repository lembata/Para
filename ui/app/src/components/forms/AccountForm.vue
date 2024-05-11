<script setup>
import { ref, defineEmits } from 'vue';
import { useI18n } from 'vue-i18n';
import TextInput from './inputs/TextInput.vue'
import CurrencyInput from './inputs/CurrencyInput.vue'
import NumberInput from './inputs/NumberInput.vue'
import DateInput from './inputs/DateInput.vue'
import Card from 'primevue/card';
import Button from 'primevue/button';
import ProgressSpinner from 'primevue/progressspinner';
import API from '../../api/api.js';

const { t } = useI18n();

const accountName = ref('');
const currency = ref('USD');
const emit = defineEmits(['submit']);
const working = ref(false);

const submitForm = async () => {
  //console.log('submitForm', working, accountName.value, currency.value);
  if(working.value) return;

  working.value = true;

  await API.Accounts.add({ 
    accountName: accountName.value,
    currency: currency.value
  });

  working.value = false;
}
</script>

<template>
  <form>
    <div class="grid grid-cols-2 lg:grid-cols-1 gap-4">
      <Card>
        {{ accountName }}
        <template #content>
          <label class="mb-2"> {{ $t('forms.manditoryFields') }} zzz:</label>
          <hr class="p-3" />
          <TextInput v-model="accountName" text="forms.accountName" placeholder="form.accountNamePlaceholder">
          </TextInput>
          <CurrencyInput text="forms.currency" name="currency"> </CurrencyInput>
        </template>
      </Card>
      <Card>
        <template #content>
          <label class="mb-2 mt-10">Optional fields:</label>
          <hr class="p-3" />
          <TextInput text="forms.IBAN" name="iban" placeholder="form.IBAN"> </TextInput>
          <TextInput text="forms.BIC" name="bic" placeholder="form.BIC"> </TextInput>
          <TextInput text="forms.accountNumber" name="accountNumber" placeholder="form.accountNumberPlaceHolder">
          </TextInput>
          <NumberInput text="forms.openingBalance" name="openingBalance" placeholder="1000" :minFractionDigits="2"/>
          <DateInput text="forms.openingBalanceDate" name="openingBalanceDate" />
          <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name">
            Include
            in net worth: </label>
          <input
            class="mb-3 block rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
            id="grid-first-name" type="checkbox" />
          <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name">
            Notes:
          </label>
          <textarea
            class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
            id="grid-first-name" type="text">
          </textarea>
        </template>
      </Card>
    </div>
    <Button @click="submitForm" type="button" :disabled="working">
      <ProgressSpinner v-if="working" class="mr-2 w-4 h-4" strokeWidth="5" 
        animationDuration=".5s" aria-label="Working" />
      Save
    </Button>
  </form>
</template>

<style>
.form-label {
  @apply mb-2 block text-xs font-bold uppercase tracking-wide text-gray-500;
}

.form-input {
  @apply mb-3 block w-full appearance-none rounded px-4 py-3 leading-tight;

}

.form-select {
  @apply block w-full appearance-none rounded border border-gray-200 bg-gray-200 px-4 py-3 pr-8 leading-tight text-gray-700 focus:border-gray-500 focus:bg-white focus:outline-none;
}
</style>
