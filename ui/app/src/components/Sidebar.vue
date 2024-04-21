<script setup>
import { RouterLink } from 'vue-router'
import SidebarButton from './SidebarButton.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
</script>

<template>
  <button data-drawer-target="default-sidebar" data-drawer-toggle="default-sidebar" aria-controls="default-sidebar"
    type="button"
    class="inline-flex items-center p-2 mt-2 ms-3 text-sm text-gray-500 rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
    <span class="sr-only">Open sidebar</span>
    <svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
      <path clip-rule="evenodd" fill-rule="evenodd"
        d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z">
      </path>
    </svg>
  </button>

  <aside id="default-sidebar" class="sidebar-container transition-transform -translate-x-full" aria-label="Sidebar">
    <div class="sidebar">
      <ul class="space-y-2 font-medium">
        <SidebarButton text="Dashboard" href="/">
          <template #icon>
            <DocumentationIcon />
          </template>
        </SidebarButton>
        <SidebarButton text="Accounts" href="/accounts">
          <template #icon>
            <CommunityIcon />
          </template>
        </SidebarButton>
        <SidebarButton text="Settings" href="/settings">
          <template #icon>
            <SupportIcon />
          </template>
        </SidebarButton>
      </ul>
    </div>
  </aside>
</template>

<style scoped>
.sidebar-container {
  @apply fixed top-0 left-0 z-40 w-64 h-screen sm:translate-x-0;
}

.sidebar {
  @apply h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800;
}

.sidebar-btn {
  @apply flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700;
}

.sidebar-icon {
  @apply flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white;
}

.sidebar-btn-text {
  @apply flex-1 ms-3 whitespace-nowrap;
}
</style>
