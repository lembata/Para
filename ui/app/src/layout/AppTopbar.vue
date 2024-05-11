<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import { useRouter } from 'vue-router';

const { layoutConfig, onMenuToggle } = useLayout();

console.log(onMenuToggle);

const outsideClickListener = ref(null);
const topbarMenuActive = ref(false);
const router = useRouter();

onMounted(() => {
   bindOutsideClickListener();
});

onBeforeUnmount(() => {
   unbindOutsideClickListener();
});

const logoUrl = computed(() => {
   //${layoutConfig.darkTheme.value ? 'logo-white' : 'logo-dark'}.svg`;
   return `/layout/images/asd.svg`;
});

const onTopBarMenuButton = () => {
   topbarMenuActive.value = !topbarMenuActive.value;
};

const onSettingsClick = () => {
   topbarMenuActive.value = false;
   router.push('/documentation');
};
const topbarMenuClasses = computed(() => {
   return {
      'layout-topbar-menu-mobile-active': topbarMenuActive.value
   };
});

const bindOutsideClickListener = () => {
   if (!outsideClickListener.value) {
      outsideClickListener.value = (event) => {
         if (isOutsideClicked(event)) {
            topbarMenuActive.value = false;
         }
      };
      document.addEventListener('click', outsideClickListener.value);
   }
};

const unbindOutsideClickListener = () => {
   if (outsideClickListener.value) {
      document.removeEventListener('click', outsideClickListener);
      outsideClickListener.value = null;
   }
};

const isOutsideClicked = (event) => {
   if (!topbarMenuActive.value) return;

   const sidebarEl = document.querySelector('.layout-topbar-menu');
   const topbarEl = document.querySelector('.layout-topbar-menu-button');

   return !(sidebarEl.isSameNode(event.target) 
      || sidebarEl.contains(event.target)
      || topbarEl.isSameNode(event.target)
      || topbarEl.contains(event.target));
};

</script>

<template>
   <div class="layout-topbar">
      <router-link to="/" class="layout-topbar-logo">
         <img :src="logoUrl" alt="logo Para" />
         <span>PARA</span>
      </router-link>

      <button class="p-link layout-menu-button layout-topbar-button" @click="onMenuToggle()">
         <i class="pi pi-bars"></i>
      </button>

      <button class="p-link layout-topbar-menu-button layout-topbar-button" @click="onTopBarMenuButton()">
         <i class="pi pi-ellipsis-v"></i>
      </button>

      <div class="layout-topbar-menu" :class="topbarMenuClasses">
         <button @click="onTopBarMenuButton()" class="p-link layout-topbar-button">
            <i class="pi pi-calendar"></i>
         </button>
         <button @click="onTopBarMenuButton()" class="p-link layout-topbar-button">
            <i class="pi pi-user"></i>
         </button>
         <button @click="onSettingsClick()" class="p-link layout-topbar-button">
            <i class="pi pi-cog"></i>
         </button>
      </div>
   </div>
</template>

<style>
.layout-topbar {
   @apply
      fixed
      h-20
      z-50
      left-0
      top-0
      bg-surface-800
      w-full
      px-8
      flex 
      items-center
      shadow-md
      transition-transform;

   .layout-topbar-logo {
      /*TODO: make this color in the theme or something*/
      @apply
            flex
            items-center
            text-xl
            font-medium
            text-gray-100
            rounded-md
         ;

      img {
         @apply h-10 mr-2;
      }
   }

   .layout-topbar-button {
      @apply
         inline-flex
         justify-center
         items-center
         relative
         text-primary-200
         rounded-full
         w-12
         h-12
         cursor-pointer
         transition-colors

         hover:text-primary-500
      ;

         /*
      span {
         @apply text-base hidden;
      }
      */
   }

   .layout-menu-button {
      @apply ml-8;
   }
   

   .layout-topbar-menu {
      @apply
         m-0
         ml-auto
         p-0
         flex
         list-none
      ;

      .layout-topbar-button {
         @apply ml-4;
      }
   }
}


</style>
