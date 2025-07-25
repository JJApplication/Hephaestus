<template>
  <div class="min-h-screen bg-gray-900 grid-bg">
    <!-- 英雄区域 -->
    <section class="relative py-20 px-4">
      <div class="max-w-7xl mx-auto text-center">
        <h1 class="text-5xl md:text-7xl font-bold text-green-400 text-glow mb-6">
          Hephaestus
        </h1>
        <p class="text-xl md:text-2xl text-green-300 mb-12">
          现代化微服务部署与管理平台
        </p>
        
        <!-- 系统状态指示器 -->
        <div class="flex justify-center items-center space-x-8 mb-16">
          <div class="flex items-center space-x-2">
            <div class="w-3 h-3 bg-green-400 rounded-full pulse-animation"></div>
            <span class="text-green-300">系统运行中</span>
          </div>
          <div class="flex items-center space-x-2">
            <div class="w-3 h-3 bg-green-400 rounded-full pulse-animation" style="animation-delay: 0.5s"></div>
            <span class="text-green-300">{{ activeServices }} 个服务在线</span>
          </div>
          <div class="flex items-center space-x-2">
            <div class="w-3 h-3 bg-green-400 rounded-full pulse-animation" style="animation-delay: 1s"></div>
            <span class="text-green-300">{{ totalRequests }} 请求/分钟</span>
          </div>
        </div>
      </div>
    </section>
    
    <!-- 微服务集群视图 -->
    <section class="py-20 px-4">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-3xl font-bold text-green-400 text-center mb-16 text-glow">
          微服务集群架构
        </h2>
        
        <!-- 集群视图容器 -->
        <div class="relative bg-black/50 rounded-lg p-8 border border-green-500/30 glow">
          <!-- 网关层 -->
          <div class="mb-16">
            <h3 class="text-xl font-semibold text-green-300 mb-8 text-center">API 网关</h3>
            <div class="flex justify-center">
              <div class="relative">
                <div class="gateway-node">
                  <Icon name="heroicons:shield-check" class="w-8 h-8 text-green-400" />
                  <span class="text-sm text-green-300 mt-2">Gateway</span>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 数据流箭头 -->
          <div class="flex justify-center mb-8">
            <div class="arrow-container">
              <div class="arrow data-flow" style="width: 200px;"></div>
            </div>
          </div>
          
          <!-- 微服务层 -->
          <div class="mb-16">
            <h3 class="text-xl font-semibold text-green-300 mb-8 text-center">微服务集群</h3>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-8 justify-items-center">
              <div 
                v-for="(service, index) in microservices" 
                :key="service.name"
                class="microservice-container"
                :style="{ animationDelay: `${index * 0.5}s` }"
              >
                <div class="microservice-cube" :class="getServiceStatusClass(service.status)">
                  <div class="cube-face cube-front"></div>
                  <div class="cube-face cube-back"></div>
                  <div class="cube-face cube-right"></div>
                  <div class="cube-face cube-left"></div>
                  <div class="cube-face cube-top"></div>
                  <div class="cube-face cube-bottom"></div>
                </div>
                <div class="cube-shadow"></div>
                <div class="mt-4 text-center">
                  <div class="text-sm font-medium" :class="getServiceTextClass(service.status)">{{ service.name }}</div>
                  <div class="text-xs mt-1" :class="getServiceTextClass(service.status)">
                    <span class="inline-block w-2 h-2 rounded-full mr-1 pulse-animation" :class="getServiceIndicatorClass(service.status)"></span>
                    {{ service.status }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 数据流箭头向下 -->
          <div class="flex justify-center mb-8">
            <div class="arrow-container">
              <div class="arrow data-flow" style="width: 200px; animation-delay: 1s;"></div>
            </div>
          </div>
          
          <!-- 基础设施层 -->
          <div>
            <h3 class="text-xl font-semibold text-green-300 mb-8 text-center">基础设施</h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
              <div 
                v-for="(infra, index) in infrastructure" 
                :key="infra.name"
                class="infra-node"
                :style="{ animationDelay: `${index * 0.3}s` }"
              >
                <div class="bg-green-900/30 border border-green-500/50 rounded-lg p-6 text-center hover:bg-green-900/50 transition-all duration-300 glow">
                  <Icon :name="infra.icon" class="w-12 h-12 text-green-400 mx-auto mb-4" />
                  <h4 class="text-lg font-medium text-green-300 mb-2">{{ infra.name }}</h4>
                  <p class="text-sm text-green-400">{{ infra.description }}</p>
                  <div class="mt-4 flex justify-center items-center space-x-2">
                    <div class="w-2 h-2 bg-green-400 rounded-full pulse-animation"></div>
                    <span class="text-xs text-green-300">运行中</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    
    <!-- 实时监控数据 -->
    <section class="py-20 px-4">
      <div class="max-w-7xl mx-auto">
        <h2 class="text-3xl font-bold text-green-400 text-center mb-16 text-glow">
          实时监控
        </h2>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="metric-card">
            <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 text-center glow">
              <Icon name="heroicons:cpu-chip" class="w-12 h-12 text-green-400 mx-auto mb-4" />
              <h3 class="text-lg font-medium text-green-300 mb-2">CPU 使用率</h3>
              <div class="text-3xl font-bold text-green-400 mb-2">{{ cpuUsage }}%</div>
              <div class="w-full bg-gray-700 rounded-full h-2">
                <div 
                  class="bg-green-400 h-2 rounded-full transition-all duration-1000"
                  :style="{ width: `${cpuUsage}%` }"
                ></div>
              </div>
            </div>
          </div>
          
          <div class="metric-card">
            <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 text-center glow">
              <Icon name="heroicons:circle-stack" class="w-12 h-12 text-green-400 mx-auto mb-4" />
              <h3 class="text-lg font-medium text-green-300 mb-2">内存使用率</h3>
              <div class="text-3xl font-bold text-green-400 mb-2">{{ memoryUsage }}%</div>
              <div class="w-full bg-gray-700 rounded-full h-2">
                <div 
                  class="bg-green-400 h-2 rounded-full transition-all duration-1000"
                  :style="{ width: `${memoryUsage}%` }"
                ></div>
              </div>
            </div>
          </div>
          
          <div class="metric-card">
            <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 text-center glow">
              <Icon name="heroicons:signal" class="w-12 h-12 text-green-400 mx-auto mb-4" />
              <h3 class="text-lg font-medium text-green-300 mb-2">网络流量</h3>
              <div class="text-3xl font-bold text-green-400 mb-2">{{ networkTraffic }} MB/s</div>
              <div class="w-full bg-gray-700 rounded-full h-2">
                <div 
                  class="bg-green-400 h-2 rounded-full transition-all duration-1000"
                  :style="{ width: `${Math.min(networkTraffic * 10, 100)}%` }"
                ></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
// 页面元数据
useHead({
  title: 'Hephaestus - 微服务部署平台',
  meta: [
    { name: 'description', content: '现代化微服务部署和管理平台首页' }
  ]
})

// 响应式数据
const activeServices = ref(8)
const totalRequests = ref(1247)
const cpuUsage = ref(45)
const memoryUsage = ref(62)
const networkTraffic = ref(8.3)

// 微服务数据
const microservices = ref([
  { name: 'User Service', status: '已停止' },
  { name: 'Order Service', status: '运行中' },
  { name: 'Payment Service', status: '运行中' },
  { name: 'Inventory Service', status: '未知' },
  { name: 'Notification Service', status: '运行中' },
  { name: 'Analytics Service', status: '未知' },
  { name: 'Auth Service', status: '运行中' },
  { name: 'File Service', status: '已停止' }
])

// 获取服务状态对应的立方体样式类
const getServiceStatusClass = (status) => {
  switch (status) {
    case '运行中':
      return 'cube-running'
    case '未知':
      return 'cube-unknown'
    case '已停止':
      return 'cube-stopped'
    default:
      return 'cube-running'
  }
}

// 获取服务状态对应的文本样式类
const getServiceTextClass = (status) => {
  switch (status) {
    case '运行中':
      return 'text-green-300'
    case '未知':
      return 'text-yellow-300'
    case '已停止':
      return 'text-red-300'
    default:
      return 'text-green-300'
  }
}

// 获取服务状态对应的指示器样式类
const getServiceIndicatorClass = (status) => {
  switch (status) {
    case '运行中':
      return 'bg-green-400'
    case '未知':
      return 'bg-yellow-400'
    case '已停止':
      return 'bg-red-400'
    default:
      return 'bg-green-400'
  }
}

// 基础设施数据
const infrastructure = ref([
  {
    name: 'Redis Cache',
    description: '高性能缓存服务',
    icon: 'heroicons:bolt'
  },
  {
    name: 'PostgreSQL',
    description: '主数据库服务',
    icon: 'heroicons:circle-stack'
  },
  {
    name: 'Message Queue',
    description: '异步消息处理',
    icon: 'heroicons:queue-list'
  }
])

// 模拟实时数据更新
onMounted(() => {
  const updateMetrics = () => {
    cpuUsage.value = Math.floor(Math.random() * 30) + 40
    memoryUsage.value = Math.floor(Math.random() * 20) + 50
    networkTraffic.value = (Math.random() * 5 + 5).toFixed(1)
    totalRequests.value = Math.floor(Math.random() * 200) + 1200
  }
  
  // 每5秒更新一次数据
  const interval = setInterval(updateMetrics, 5000)
  
  // 组件卸载时清除定时器
  onUnmounted(() => {
    clearInterval(interval)
  })
})
</script>

<style scoped>
.gateway-node {
  @apply flex flex-col items-center justify-center w-24 h-24 bg-green-900/30 border-2 border-green-500 rounded-lg glow float-animation;
}

.microservice-container {
  @apply relative flex flex-col items-center;
  perspective: 1000px;
}

.infra-node {
  @apply float-animation;
}

.metric-card {
  @apply float-animation;
}

.arrow-container {
  @apply relative overflow-hidden;
  width: 200px;
  height: 10px;
}

/* 3D立方体旋转动画 */
@keyframes rotate3d {
  0% {
    transform: rotateX(0deg) rotateY(0deg);
  }
  100% {
    transform: rotateX(360deg) rotateY(360deg);
  }
}

.microservice-cube {
  animation: float 3s ease-in-out infinite, rotate3d 10s linear infinite;
  position: relative;
  width: 60px;
  height: 60px;
  transform-style: preserve-3d;
}

.cube-face {
  position: absolute;
  width: 60px;
  height: 60px;
  border: 2px solid;
  opacity: 0.8;
}

/* 运行中状态 - 绿色 */
.cube-running .cube-face {
  background: linear-gradient(45deg, #10b981, #059669);
  border-color: #10b981;
  box-shadow: 0 0 20px rgba(16, 185, 129, 0.5);
}

/* 未知状态 - 淡黄色 */
.cube-unknown .cube-face {
  background: linear-gradient(45deg, #fbbf24, #f59e0b);
  border-color: #fbbf24;
  box-shadow: 0 0 20px rgba(251, 191, 36, 0.5);
}

/* 已停止状态 - 淡红色 */
.cube-stopped .cube-face {
  background: linear-gradient(45deg, #f87171, #ef4444);
  border-color: #f87171;
  box-shadow: 0 0 20px rgba(248, 113, 113, 0.5);
}

.cube-front {
  transform: rotateY(0deg) translateZ(30px);
}

.cube-back {
  transform: rotateY(180deg) translateZ(30px);
}

.cube-right {
  transform: rotateY(90deg) translateZ(30px);
}

.cube-left {
  transform: rotateY(-90deg) translateZ(30px);
}

.cube-top {
  transform: rotateX(90deg) translateZ(30px);
}

.cube-bottom {
  transform: rotateX(-90deg) translateZ(30px);
}

.cube-shadow {
  position: absolute;
  top: 80px;
  left: 50%;
  transform: translateX(-50%);
  width: 40px;
  height: 40px;
  background: radial-gradient(ellipse, rgba(0, 0, 0, 0.3), transparent);
  border-radius: 50%;
  animation: shadowFloat 3s ease-in-out infinite;
}

@keyframes shadowFloat {
  0%, 100% {
    transform: translateX(-50%) scale(1);
    opacity: 0.3;
  }
  50% {
    transform: translateX(-50%) scale(1.1);
    opacity: 0.2;
  }
}

/* 数据流动画增强 */
@keyframes dataFlowEnhanced {
  0% {
    transform: translateX(-100%) scaleX(0);
    opacity: 0;
  }
  20% {
    transform: translateX(-50%) scaleX(1);
    opacity: 1;
  }
  80% {
    transform: translateX(50%) scaleX(1);
    opacity: 1;
  }
  100% {
    transform: translateX(100%) scaleX(0);
    opacity: 0;
  }
}

.data-flow {
  animation: dataFlowEnhanced 3s ease-in-out infinite;
}
</style>