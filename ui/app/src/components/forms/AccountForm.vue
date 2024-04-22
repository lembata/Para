<script>
import FormInput from './inputs/TextInput.vue'
</script>

<template>
  <form>
    <label class="mb-2">Mandatory fields:</label>
    <hr class="p-3" />
    <div style="width:25px;height: 25px; background: red;"></div>
    <FormInput text="Account Name" name="name" > </FormInput>
    <div style="width:25px;height: 25px; background: red;"></div>

    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Account
      Name </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="text" placeholder="Cash" />
    <p class="hidden text-xs italic text-red-500">Please fill out this field.</p>
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name">
      Currency </label>
    <div class="relative">
      <select
        class="block w-full appearance-none rounded border border-gray-200 bg-gray-200 px-4 py-3 pr-8 leading-tight text-gray-700 focus:border-gray-500 focus:bg-white focus:outline-none"
        id="grid-state">
        <option>Euro €</option>
        <option>Dollars $</option>
        <option>Bulgarian Lev лв.</option>
      </select>
      <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
        <svg class="h-4 w-4 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
          <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
        </svg>
      </div>
    </div>
    <hr class="p-3" />
    <label class="mb-2 mt-10">Optional fields:</label>
    <hr class="p-3" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> IBAN:
    </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="text" placeholder="IBAN" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> BIC:
    </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="text" placeholder="BIC" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Account
      Number: </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="text" placeholder="Account Number" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Opening
      Balance: </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="number" placeholder="100" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Opening
      balance date: </label>
    <input
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="date" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Include
      in net worth: </label>
    <input
      class="mb-3 block rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="checkbox" />
    <label class="mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700" for="grid-first-name"> Notes:
    </label>
    <textarea
      class="mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none"
      id="grid-first-name" type="text"> </textarea>
  </form>
</template>

<style>
.form-label {
  @apply mb-2 block text-xs font-bold uppercase tracking-wide text-gray-700;
}

.form-input {
  @apply mb-3 block w-full appearance-none rounded border bg-gray-200 px-4 py-3 leading-tight text-gray-700 focus:bg-white focus:outline-none;
}
</style>
