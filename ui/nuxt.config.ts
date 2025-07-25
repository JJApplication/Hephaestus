export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: [
    '@nuxt/ui'
  ],
  ui: {
    global: true,
    icons: ['heroicons']
  },
  css: ['~/assets/css/main.css'],
  app: {
    head: {
      title: 'Hephaestus - 微服务部署平台',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: '现代化微服务部署和管理平台' }
      ]
    }
  }
})