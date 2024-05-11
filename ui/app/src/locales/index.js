import { createI18n } from 'vue-i18n'
import messages from '@intlify/unplugin-vue-i18n/messages'

export const i18n = createI18n({
  locale: 'bg',
  leagacy: false,
  fallbackLocale: 'en',
  globalInjection: true,
  messages
});
