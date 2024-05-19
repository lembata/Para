<script setup>
import {ref, watch} from 'vue';
import {useI18n} from 'vue-i18n';
import {RouterLink} from 'vue-router'
import API from '@/api/api.js';
import Button from 'primevue/button';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Dialog from 'primevue/dialog';
import { formatCurrency } from "@/utils/formatter.js";

const {t} = useI18n();
const accounts = ref([]);
const loading = ref(true);
const showDeleteDialog = ref(false);
const selectedRecord = ref(0);
const conformationText = ref("");

const loadInitial = async () => {
  await getPage(1);
}

const getPage = async (page) => {
  const limit = 10;
  API.Accounts.all(limit, (page - 1) * limit)
      .then((result) => {
        if (result.success) {
          accounts.value = result.data;
          console.log('accounts', result.data);
        } else {
          console.error('Failed to load accounts', result);
          toast.add({severity: 'error', summary: t('notifications.failedToLoadAccounts'), life: 3000});
        }
      });
}

const promptDeleteRecord = (id) => {
  console.log('deleteRecord', id);
  showDeleteDialog.value = true;
  selectedRecord.value = id;
}

const deleteRecord = () => {
  const id = selectedRecord.value;
  selectedRecord.value = 0;
  console.log('deleteRecord', id);

  showDeleteDialog.value = false;
  loading.value = true;

  API.Accounts.delete(id)
      .then((result) => {
        if (result.success) {
          getPage(1);
          toast.add({severity: 'success', summary: t('notifications.recordDeleted'), life: 3000});
        } else {
          console.error('Failed to delete account', result);
          toast.add({severity: 'error', summary: t('notifications.failedToDeleteRecord'), life: 3000});
        }
      });
  loading.value = false;
}

//watch(() => loadInitial(), false);
watch(() => 0, loadInitial, {immediate: true})
</script>

<template>
  <main>
    <DataTable
        :value="accounts"
        tableStyle="min-width: 50rem">
      <!--        :reorderableColumns="true"-->
      <!--               @columnReorder="onColReorder" -->
      <!--               @rowReorder="onRowReorder"-->
      <Column field="id" :header="t('id')" key="field"/>
      <Column field="name" :header="t('forms.accountName')" key="field"/>
      <Column field="balance" :header="t('balance')" key="field">
        <template #body="slotProps">
          {{ formatCurrency(slotProps.data.balance.value, slotProps.data.balance.currency) }}
        </template>
      </Column>
      <Column field="actions" :header="t('actions')" key="actions">
        <template #body="slotProps">
          <RouterLink :to="'/accounts/' + slotProps.data.id">
            <Button :label="$t('details')" icon="pi pi-eye"/>
          </RouterLink>
          <RouterLink :to="'/accounts/edit/' + slotProps.data.id">
            <Button :label="$t('edit')" icon="pi pi-pencil"/>
          </RouterLink>
          <Button class="bg-red-500"
                  :label="$t('delete')"
                  icon="pi pi-trash"
                  @click="promptDeleteRecord(slotProps.data.id)"/>
        </template>
      </Column>
    </DataTable>

    <!-- :position="position" -->
    <Dialog v-model:visible="showDeleteDialog"
            header="Edit Profile"
            class="w-100"
            :modal="true"
            :draggable="true">
      <span class="p-text-secondary block mb-5">Are you sure you want to delete.</span>
      <div class="flex align-items-center gap-3 mb-3">
        <label class="font-semibold w-6rem">Username</label>
        <InputText v-model="conformationText" class="flex-auto" autocomplete="off"/>
      </div>
      <div class="dialog-footer">
        <Button type="button" :label="t('actionCancel')" severity="secondary"
                @click="showDeleteDialog = false"/>
        <Button type="button" :label="t('actionDelete')" severity="danger"
                @click="deleteRecord"/>
      </div>
    </Dialog>
  </main>
</template>

<style>
.dialog-footer {
  @apply
  flex
    /*justify-content-end*/
  gap-2;
}
</style>