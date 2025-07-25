<template>
  <div class="min-h-screen bg-gray-900">
    <!-- 固定顶部导航 -->
    <nav class="fixed top-0 left-0 right-0 z-50 bg-black/90 backdrop-blur-sm border-b border-green-500/30">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <!-- Logo -->
          <div class="flex items-center">
            <NuxtLink to="/" class="text-2xl font-bold text-green-400 text-glow hover:text-green-300 transition-colors">
              Hephaestus
            </NuxtLink>
          </div>
          
          <!-- 导航菜单 -->
          <div class="hidden md:block">
            <div class="ml-10 flex items-baseline space-x-4">
              <NuxtLink 
                to="/" 
                class="nav-link"
                :class="{ 'active': $route.path === '/' }"
              >
                <Icon name="heroicons:home" class="w-4 h-4 mr-2" />
                主页
              </NuxtLink>
              <NuxtLink 
                to="/services" 
                class="nav-link"
                :class="{ 'active': $route.path.startsWith('/services') }"
              >
                <Icon name="heroicons:squares-2x2" class="w-4 h-4 mr-2" />
                微服务列表
              </NuxtLink>
              <NuxtLink 
                to="/versions" 
                class="nav-link"
                :class="{ 'active': $route.path.startsWith('/versions') }"
              >
                <Icon name="heroicons:code-bracket" class="w-4 h-4 mr-2" />
                版本变更
              </NuxtLink>
              <NuxtLink 
                to="/packages" 
                class="nav-link"
                :class="{ 'active': $route.path.startsWith('/packages') }"
              >
                <Icon name="heroicons:archive-box" class="w-4 h-4 mr-2" />
                包管理
              </NuxtLink>
              <NuxtLink 
                to="/deploy" 
                class="nav-link"
                :class="{ 'active': $route.path.startsWith('/deploy') }"
              >
                <Icon name="heroicons:rocket-launch" class="w-4 h-4 mr-2" />
                部署管理
              </NuxtLink>
              <NuxtLink 
                to="/logs" 
                class="nav-link"
                :class="{ 'active': $route.path.startsWith('/logs') }"
              >
                <Icon name="heroicons:document-text" class="w-4 h-4 mr-2" />
                日志管理
              </NuxtLink>
            </div>
          </div>
          
          <!-- 移动端菜单按钮 -->
          <div class="md:hidden">
            <button 
              @click="mobileMenuOpen = !mobileMenuOpen"
              class="text-green-400 hover:text-green-300 focus:outline-none focus:text-green-300"
            >
              <Icon name="heroicons:bars-3" class="w-6 h-6" />
            </button>
          </div>
        </div>
      </div>
      
      <!-- 移动端菜单 -->
      <div v-show="mobileMenuOpen" class="md:hidden bg-black/95 border-t border-green-500/30">
        <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
          <NuxtLink to="/" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:home" class="w-4 h-4 mr-2" />
            主页
          </NuxtLink>
          <NuxtLink to="/services" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:squares-2x2" class="w-4 h-4 mr-2" />
            微服务列表
          </NuxtLink>
          <NuxtLink to="/versions" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:code-bracket" class="w-4 h-4 mr-2" />
            版本变更
          </NuxtLink>
          <NuxtLink to="/packages" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:archive-box" class="w-4 h-4 mr-2" />
            包管理
          </NuxtLink>
          <NuxtLink to="/deploy" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:rocket-launch" class="w-4 h-4 mr-2" />
            部署管理
          </NuxtLink>
          <NuxtLink to="/logs" class="mobile-nav-link" @click="mobileMenuOpen = false">
            <Icon name="heroicons:document-text" class="w-4 h-4 mr-2" />
            日志管理
          </NuxtLink>
        </div>
      </div>
    </nav>
    
    <!-- 主要内容区域 -->
    <main class="pt-16">
      <slot />
    </main>
    
    <!-- 底部版权信息 - 悬浮样式 -->
    <div class="fixed bottom-2 left-1/2 transform -translate-x-1/2 z-40">
      <div class="floating-footer">
        <div class="flex items-center space-x-4">
          <div class="text-sm text-green-300">
            Copyright JJAPPs
          </div>
          <a 
            href="https://renj.io" 
            target="_blank" 
            rel="noopener noreferrer"
            class="footer-link-btn"
          >
            <Icon name="heroicons:sparkles" class="w-4 h-4" />
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const mobileMenuOpen = ref(false)
</script>

<style scoped>
.nav-link {
  @apply flex items-center px-3 py-2 rounded-md text-sm font-medium text-green-300 hover:text-green-400 hover:bg-green-900/20 transition-all duration-200;
}

.nav-link.active {
  @apply text-green-400 bg-green-900/30 glow;
}

.mobile-nav-link {
  @apply flex items-center px-3 py-2 rounded-md text-base font-medium text-green-300 hover:text-green-400 hover:bg-green-900/20 transition-all duration-200;
}

.floating-footer {
  @apply bg-black/90 backdrop-blur-sm border border-green-500/30 rounded-2xl px-4 py-3 shadow-lg;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.3), 0 0 20px rgba(0, 255, 65, 0.1);
}
</style>