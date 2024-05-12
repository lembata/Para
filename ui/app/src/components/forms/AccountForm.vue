<script setup>
import { ref, defineProps, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import TextInput from './inputs/TextInput.vue'
import TextArea from './inputs/TextArea.vue'
import CurrencyInput from './inputs/CurrencyInput.vue'
import NumberInput from './inputs/NumberInput.vue'
import DateInput from './inputs/DateInput.vue'
import SwitchInput from './inputs/SwitchInput.vue'
import Card from 'primevue/card';
import Button from 'primevue/button';
import ProgressSpinner from 'primevue/progressspinner';
import API from '../../api/api.js';
import { useToast } from 'primevue/usetoast';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();

const props = defineProps(['id']);//{ new: true });

const loading = ref(true);
const working = ref(false);

const accountName = ref('');
const currency = ref('EUR');
const iban = ref('')
const bic = ref('');//model.bic ?? '');
const accountNumber = ref('');
const openingBalance = ref(0);
const openiningBalanceDate = ref(new Date(Date.now()));
const notes = ref('');
const toast = useToast();

const loadInitial = async () => {
  if (!route.params.id) {
    loading.value = false;
  }
  else {
    // load data from API
    API.Accounts.get(route.params.id)
      .then((result) => {
        if (result.success) {
          accountName.value = result.data.accountName;
          currency.value = result.data.currency;
          iban.value = result.data.iban;
          bic.value = result.data.bic;
          accountNumber.value = result.data.accountNumber;
          openingBalance.value = result.data.openingBalance;
          openiningBalanceDate.value = new Date(result.data.openiningBalanceDate);
          notes.value = result.data.notes;
        } else {
          console.error('Failed to load account', result);
          //router.redirect('/accounts');
          toast.add({ severity: 'error', summary: t('notifications.accountDoesntExist'), life: 3000 });
          router.push('/accounts');
          //window.location.replace("/accounts");
        }
      });
    loading.value = false;
  }
}

const submitForm = async () => {
  console.log('submitForm', working);//  accountName.value, currency.value);
  if (loading.value) return;
  if (working.value) return;

  working.value = true;

  data = {
    accountName: accountName.value,
    currency: currency.value,
    iban: iban.value,
    bic: bic.value,
    accountNumber: accountNumber.value,
    openingBalance: openingBalance.value,
    openiningBalanceDate: openiningBalanceDate.value,
    notes: notes.value
  }

  let result;

  if (!route.params.id) {
    result = await API.Accounts.add(data);
  } else {
    result = await API.Accounts.edit(route.params.id, data);
  }



  console.warn('result', result);

  if (result.success) {
    toast.add({ severity: 'success', summary: t('notifications.accountCreated'), life: 3000 });
  }
  else {
    toast.add({ severity: 'error', summary: t('notifications.accountCreationFailed'), life: 3000 });
  }

  working.value = false;
}

watch(() => route.params.id, loadInitial, { immediate: true })
</script>

<template>
  <form class="relative">
    <div v-if="loading" class="overlay w-full h-full block absolute bg-black bg-opacity-20 z-20 "></div>
    <div
      class="grid grid-ctoast.add({ severity: 'success', summary: 'Success Message', detail: 'Message Content', life: 3000 });ols-2 lg:grid-cols-1 gap-4">
      <Card>
        <template #title>
          <label class="mb-2"> {{ $t('forms.manditoryFields') }}:</label>
        </template>
        <template #content>
          <hr class="p-3" />
          <TextInput v-model="accountName" text="forms.accountName" placeholder="form.accountNamePlaceholder"
            :disabled="loading">
          </TextInput>
          <CurrencyInput v-model="currency" text="forms.currency" :disabled="loading"> </CurrencyInput>
        </template>
      </Card>
      <Card>
        <template #title>
          <label class="mb-2"> {{ $t('forms.optionalFields') }}:</label>
        </template>
        <template #content>
          <hr class="p-3" />
          <TextInput v-model="iban" text="forms.IBAN" placeholder="forms.IBAN" :disabled="loading"> </TextInput>
          <TextInput v-model="bic" text="forms.BIC" placeholder="forms.BIC" :disabled="loading"> </TextInput>
          <TextInput v-model="accountNumber" text="forms.accountNumber" placeholder="forms.accountNumberPlaceHolder"
            :disabled="loading" />
          <NumberInput v-model="openingBalance" text="forms.openingBalance" placeholder="1000" :minFractionDigits="2"
            :disabled="loading" />
          <DateInput v-model="openiningBalanceDate" text="forms.openingBalanceDate" name="openingBalanceDate"
            :disabled="loading" />
          <SwitchInput text="forms.IncludeInNetWorth" :disabled="loading" />
          <TextArea text="forms.notes" :disabled="loading" />


        </template>
      </Card>
    </div>
    <Button @click="submitForm" type="button" :disabled="working || loading">
      <ProgressSpinner v-if="working" class="mr-2 w-4 h-4" strokeWidth="5" animationDuration=".5s"
        aria-label="Working" />
      {{ $t('buttons.save') }}
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
