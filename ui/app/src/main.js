import 'primevue/resources/themes/lara-dark-indigo/theme.css'
import './assets/main.css'
import './index.css'
//import 'primevue/resources/themes/aura-light-green/theme.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { i18n } from './locales'
import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';
import DialogService from 'primevue/dialogservice';
import ConfirmationService from 'primevue/confirmationservice';
import Tooltip from 'primevue/tooltip'
import Ripple from 'primevue/ripple'
import Lara from '@/presets/lara'

createApp(App)
  .use(router)
  .use(i18n)
  .use(PrimeVue, {
    unstyled: true,
    ripple: true,
    pt: Lara
  })
  .use(ToastService)
  .use(DialogService)
  .use(ConfirmationService)
  .directive('tooltip', Tooltip)
  .directive('ripple', Ripple)
  .mount('#app')

