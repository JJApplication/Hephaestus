<template>
  <div class="min-h-screen bg-gray-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- 页面标题 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-green-400 text-glow mb-4">微服务列表</h1>
        <p class="text-green-300">管理和监控所有微服务的运行状态</p>
      </div>
      
      <!-- 搜索和过滤 -->
      <div class="mb-8 flex flex-col md:flex-row gap-4">
        <div class="flex-1">
          <UInput 
            v-model="searchQuery"
            placeholder="搜索微服务..."
            icon="heroicons:magnifying-glass"
            class="w-full"
          />
        </div>
        <div class="flex gap-4">
          <USelect 
            v-model="statusFilter"
            :options="statusOptions"
            placeholder="状态筛选"
            class="w-40"
          />
          <UButton 
            @click="refreshServices"
            icon="heroicons:arrow-path"
            variant="outline"
            color="green"
          >
            刷新
          </UButton>
        </div>
      </div>
      
      <!-- 服务统计卡片 -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">总服务数</p>
              <p class="text-2xl font-bold text-green-400">{{ totalServices }}</p>
            </div>
            <Icon name="heroicons:squares-2x2" class="w-8 h-8 text-green-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">运行中</p>
              <p class="text-2xl font-bold text-green-400">{{ runningServices }}</p>
            </div>
            <Icon name="heroicons:play-circle" class="w-8 h-8 text-green-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">已停止</p>
              <p class="text-2xl font-bold text-red-400">{{ stoppedServices }}</p>
            </div>
            <Icon name="heroicons:stop-circle" class="w-8 h-8 text-red-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">错误状态</p>
              <p class="text-2xl font-bold text-yellow-400">{{ errorServices }}</p>
            </div>
            <Icon name="heroicons:exclamation-triangle" class="w-8 h-8 text-yellow-400" />
          </div>
        </div>
      </div>
      
      <!-- 服务列表 -->
      <div class="bg-black/50 border border-green-500/30 rounded-lg overflow-hidden glow">
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead class="bg-green-900/20">
              <tr>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">服务名称</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">状态</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">版本</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">实例数</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">CPU</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">内存</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">最后更新</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">操作</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-green-500/20">
              <tr 
                v-for="service in filteredServices" 
                :key="service.id"
                class="hover:bg-green-900/10 transition-colors"
              >
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="microservice-cube-mini mr-4">
                      <div class="cube-face-mini cube-front-mini"></div>
                      <div class="cube-face-mini cube-back-mini"></div>
                      <div class="cube-face-mini cube-right-mini"></div>
                      <div class="cube-face-mini cube-left-mini"></div>
                      <div class="cube-face-mini cube-top-mini"></div>
                      <div class="cube-face-mini cube-bottom-mini"></div>
                    </div>
                    <div>
                      <div class="text-green-300 font-medium">{{ service.name }}</div>
                      <div class="text-green-400 text-sm">{{ service.description }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4">
                  <span 
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="getStatusClass(service.status)"
                  >
                    <span 
                      class="w-2 h-2 rounded-full mr-1.5"
                      :class="getStatusDotClass(service.status)"
                    ></span>
                    {{ service.status }}
                  </span>
                </td>
                <td class="px-6 py-4 text-green-300">{{ service.version }}</td>
                <td class="px-6 py-4 text-green-300">{{ service.instances }}</td>
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="w-16 bg-gray-700 rounded-full h-2 mr-2">
                      <div 
                        class="bg-green-400 h-2 rounded-full"
                        :style="{ width: `${service.cpu}%` }"
                      ></div>
                    </div>
                    <span class="text-green-300 text-sm">{{ service.cpu }}%</span>
                  </div>
                </td>
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="w-16 bg-gray-700 rounded-full h-2 mr-2">
                      <div 
                        class="bg-green-400 h-2 rounded-full"
                        :style="{ width: `${service.memory}%` }"
                      ></div>
                    </div>
                    <span class="text-green-300 text-sm">{{ service.memory }}%</span>
                  </div>
                </td>
                <td class="px-6 py-4 text-green-300 text-sm">{{ service.lastUpdated }}</td>
                <td class="px-6 py-4">
                  <div class="flex space-x-2">
                    <UButton 
                      size="sm" 
                      variant="ghost" 
                      color="green"
                      icon="heroicons:eye"
                      @click="viewService(service)"
                    />
                    <UButton 
                      size="sm" 
                      variant="ghost" 
                      color="blue"
                      icon="heroicons:cog-6-tooth"
                      @click="configureService(service)"
                    />
                    <UButton 
                      size="sm" 
                      variant="ghost" 
                      :color="service.status === '运行中' ? 'red' : 'green'"
                      :icon="service.status === '运行中' ? 'heroicons:stop' : 'heroicons:play'"
                      @click="toggleService(service)"
                    />
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 页面元数据
useHead({
  title: '微服务列表 - Hephaestus',
  meta: [
    { name: 'description', content: '查看和管理所有微服务的运行状态' }
  ]
})

// 响应式数据
const searchQuery = ref('')
const statusFilter = ref('')

const statusOptions = [
  { label: '全部状态', value: '' },
  { label: '运行中', value: '运行中' },
  { label: '已停止', value: '已停止' },
  { label: '错误', value: '错误' }
]

// 模拟服务数据
const services = ref([
  {
    id: 1,
    name: 'User Service',
    description: '用户管理服务',
    status: '运行中',
    version: 'v1.2.3',
    instances: 3,
    cpu: 45,
    memory: 62,
    lastUpdated: '2024-01-15 14:30'
  },
  {
    id: 2,
    name: 'Order Service',
    description: '订单处理服务',
    status: '运行中',
    version: 'v2.1.0',
    instances: 5,
    cpu: 38,
    memory: 55,
    lastUpdated: '2024-01-15 14:25'
  },
  {
    id: 3,
    name: 'Payment Service',
    description: '支付处理服务',
    status: '运行中',
    version: 'v1.5.2',
    instances: 2,
    cpu: 52,
    memory: 48,
    lastUpdated: '2024-01-15 14:20'
  },
  {
    id: 4,
    name: 'Inventory Service',
    description: '库存管理服务',
    status: '已停止',
    version: 'v1.0.8',
    instances: 0,
    cpu: 0,
    memory: 0,
    lastUpdated: '2024-01-15 13:45'
  },
  {
    id: 5,
    name: 'Notification Service',
    description: '通知推送服务',
    status: '错误',
    version: 'v1.3.1',
    instances: 1,
    cpu: 25,
    memory: 35,
    lastUpdated: '2024-01-15 14:10'
  }
])

// 计算属性
const totalServices = computed(() => services.value.length)
const runningServices = computed(() => services.value.filter(s => s.status === '运行中').length)
const stoppedServices = computed(() => services.value.filter(s => s.status === '已停止').length)
const errorServices = computed(() => services.value.filter(s => s.status === '错误').length)

const filteredServices = computed(() => {
  let filtered = services.value
  
  if (searchQuery.value) {
    filtered = filtered.filter(service => 
      service.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      service.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  if (statusFilter.value) {
    filtered = filtered.filter(service => service.status === statusFilter.value)
  }
  
  return filtered
})

// 方法
const getStatusClass = (status) => {
  switch (status) {
    case '运行中':
      return 'bg-green-900/50 text-green-300 border border-green-500/50'
    case '已停止':
      return 'bg-red-900/50 text-red-300 border border-red-500/50'
    case '错误':
      return 'bg-yellow-900/50 text-yellow-300 border border-yellow-500/50'
    default:
      return 'bg-gray-900/50 text-gray-300 border border-gray-500/50'
  }
}

const getStatusDotClass = (status) => {
  switch (status) {
    case '运行中':
      return 'bg-green-400 pulse-animation'
    case '已停止':
      return 'bg-red-400'
    case '错误':
      return 'bg-yellow-400 pulse-animation'
    default:
      return 'bg-gray-400'
  }
}

const refreshServices = () => {
  // 模拟刷新数据
  console.log('刷新服务列表')
}

const viewService = (service) => {
  console.log('查看服务:', service.name)
}

const configureService = (service) => {
  console.log('配置服务:', service.name)
}

const toggleService = (service) => {
  if (service.status === '运行中') {
    service.status = '已停止'
    service.instances = 0
    service.cpu = 0
    service.memory = 0
  } else {
    service.status = '运行中'
    service.instances = Math.floor(Math.random() * 5) + 1
    service.cpu = Math.floor(Math.random() * 60) + 20
    service.memory = Math.floor(Math.random() * 50) + 30
  }
  service.lastUpdated = new Date().toLocaleString('zh-CN')
}
</script>

<style scoped>
/* 迷你立方体样式 */
.microservice-cube-mini {
  width: 24px;
  height: 24px;
  position: relative;
  transform-style: preserve-3d;
  animation: float 3s ease-in-out infinite;
}

.cube-face-mini {
  position: absolute;
  width: 24px;
  height: 24px;
  background: linear-gradient(135deg, var(--primary-green), var(--secondary-green));
  border: 1px solid var(--primary-green);
  opacity: 0.8;
}

.cube-front-mini {
  transform: rotateY(0deg) translateZ(12px);
}

.cube-back-mini {
  transform: rotateY(180deg) translateZ(12px);
}

.cube-right-mini {
  transform: rotateY(90deg) translateZ(12px);
}

.cube-left-mini {
  transform: rotateY(-90deg) translateZ(12px);
}

.cube-top-mini {
  transform: rotateX(90deg) translateZ(12px);
}

.cube-bottom-mini {
  transform: rotateX(-90deg) translateZ(12px);
}
</style>