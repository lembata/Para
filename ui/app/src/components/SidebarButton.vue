<script>
import { RouterLink } from 'vue-router'
//import { useAttrs } from 'vue';

export default {
  components: {
    RouterLink
  },
  props: ['to', 'text'],
  data(props) {
    //console.log('SidebarButton.vue data', props, ctx);//.attrs);
    return {
      text: props.text,
      to: props.to,
    }
  },
  setup() {
  }
}
</script>

<template>
  <li>
    <RouterLink :to="to" class="sidebar-btn">
      <slot name="icon" class="sidebar-icon"></slot>
      <span class="ms-3">{{ text }}</span>
    </ RouterLink>
  </li>
</template>


<style scoped>
.sidebar-btn {
  @apply flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700;
}

.sidebar-icon {
  @apply flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white;
}
</style>
