<template>
  <div class="min-h-screen bg-gray-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- 页面标题 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-green-400 text-glow mb-4">日志管理</h1>
        <p class="text-green-300">实时监控和分析微服务日志信息</p>
      </div>
      
      <!-- 搜索和过滤 -->
      <div class="mb-8">
        <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-4">
          <UInput 
            v-model="searchQuery"
            placeholder="搜索日志内容..."
            icon="heroicons:magnifying-glass"
            class="md:col-span-2"
          />
          <USelect 
            v-model="selectedService"
            :options="serviceOptions"
            placeholder="选择服务"
          />
          <USelect 
            v-model="logLevel"
            :options="logLevelOptions"
            placeholder="日志级别"
          />
          <USelect 
            v-model="timeRange"
            :options="timeRangeOptions"
            placeholder="时间范围"
          />
        </div>
        
        <div class="flex gap-4">
          <UButton 
            @click="refreshLogs"
            icon="heroicons:arrow-path"
            variant="outline"
            color="green"
          >
            刷新
          </UButton>
          <UButton 
            @click="clearLogs"
            icon="heroicons:trash"
            variant="outline"
            color="red"
          >
            清空日志
          </UButton>
          <UButton 
            @click="exportLogs"
            icon="heroicons:arrow-down-tray"
            variant="outline"
            color="blue"
          >
            导出日志
          </UButton>
          <UToggle 
            v-model="autoRefresh"
            color="green"
          />
          <span class="text-green-300 text-sm self-center">自动刷新</span>
        </div>
      </div>
      
      <!-- 日志统计 -->
      <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-8">
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-4 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">总日志数</p>
              <p class="text-xl font-bold text-green-400">{{ totalLogs }}</p>
            </div>
            <Icon name="heroicons:document-text" class="w-6 h-6 text-green-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-red-500/30 rounded-lg p-4 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-red-300 text-sm">错误日志</p>
              <p class="text-xl font-bold text-red-400">{{ errorLogs }}</p>
            </div>
            <Icon name="heroicons:exclamation-triangle" class="w-6 h-6 text-red-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-yellow-500/30 rounded-lg p-4 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-yellow-300 text-sm">警告日志</p>
              <p class="text-xl font-bold text-yellow-400">{{ warningLogs }}</p>
            </div>
            <Icon name="heroicons:exclamation-circle" class="w-6 h-6 text-yellow-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-blue-500/30 rounded-lg p-4 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-blue-300 text-sm">信息日志</p>
              <p class="text-xl font-bold text-blue-400">{{ infoLogs }}</p>
            </div>
            <Icon name="heroicons:information-circle" class="w-6 h-6 text-blue-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-gray-500/30 rounded-lg p-4 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-300 text-sm">调试日志</p>
              <p class="text-xl font-bold text-gray-400">{{ debugLogs }}</p>
            </div>
            <Icon name="heroicons:bug-ant" class="w-6 h-6 text-gray-400" />
          </div>
        </div>
      </div>
      
      <!-- 日志内容 -->
      <div class="bg-black/50 border border-green-500/30 rounded-lg overflow-hidden glow">
        <div class="p-4 border-b border-green-500/20 flex items-center justify-between">
          <h2 class="text-xl font-bold text-green-400">实时日志流</h2>
          <div class="flex items-center space-x-2">
            <div class="w-2 h-2 bg-green-400 rounded-full pulse-animation"></div>
            <span class="text-green-300 text-sm">实时监控中</span>
          </div>
        </div>
        
        <!-- 日志列表 -->
        <div class="h-96 overflow-y-auto bg-gray-900/50" ref="logContainer">
          <div class="p-4 space-y-2">
            <div 
              v-for="log in filteredLogs" 
              :key="log.id"
              class="log-entry font-mono text-sm border-l-2 pl-4 py-2 hover:bg-gray-800/30 transition-colors"
              :class="getLogEntryClass(log.level)"
            >
              <div class="flex items-start space-x-4">
                <span class="text-gray-400 text-xs whitespace-nowrap">{{ log.timestamp }}</span>
                <span 
                  class="px-2 py-0.5 rounded text-xs font-medium whitespace-nowrap"
                  :class="getLogLevelClass(log.level)"
                >
                  {{ log.level }}
                </span>
                <span class="text-green-400 text-xs whitespace-nowrap">{{ log.service }}</span>
                <span class="text-green-300 flex-1 break-all">{{ log.message }}</span>
              </div>
              
              <!-- 展开的详细信息 -->
              <div v-if="log.expanded" class="mt-2 pl-4 border-l border-gray-600">
                <div v-if="log.stack" class="text-red-300 text-xs whitespace-pre-wrap">{{ log.stack }}</div>
                <div v-if="log.context" class="text-gray-400 text-xs mt-1">
                  <strong>上下文:</strong> {{ JSON.stringify(log.context, null, 2) }}
                </div>
              </div>
              
              <!-- 操作按钮 -->
              <div class="mt-1 flex space-x-2">
                <button 
                  @click="toggleLogExpand(log)"
                  class="text-green-400 hover:text-green-300 text-xs"
                >
                  {{ log.expanded ? '收起' : '展开' }}
                </button>
                <button 
                  @click="copyLog(log)"
                  class="text-blue-400 hover:text-blue-300 text-xs"
                >
                  复制
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 日志分析图表 -->
      <div class="mt-8 grid grid-cols-1 md:grid-cols-2 gap-8">
        <!-- 日志级别分布 -->
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <h3 class="text-lg font-semibold text-green-400 mb-4">日志级别分布</h3>
          <div class="space-y-3">
            <div v-for="level in logLevelStats" :key="level.name" class="flex items-center justify-between">
              <div class="flex items-center space-x-2">
                <div 
                  class="w-3 h-3 rounded-full"
                  :class="getLogLevelColor(level.name)"
                ></div>
                <span class="text-green-300">{{ level.name }}</span>
              </div>
              <div class="flex items-center space-x-2">
                <div class="w-20 bg-gray-700 rounded-full h-2">
                  <div 
                    class="h-2 rounded-full transition-all duration-1000"
                    :class="getLogLevelColor(level.name)"
                    :style="{ width: `${level.percentage}%` }"
                  ></div>
                </div>
                <span class="text-green-300 text-sm w-12 text-right">{{ level.count }}</span>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 服务日志统计 -->
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <h3 class="text-lg font-semibold text-green-400 mb-4">服务日志统计</h3>
          <div class="space-y-3">
            <div v-for="service in serviceLogStats" :key="service.name" class="flex items-center justify-between">
              <div class="flex items-center space-x-2">
                <div class="microservice-cube-mini">
                  <div class="cube-face-mini cube-front-mini"></div>
                  <div class="cube-face-mini cube-back-mini"></div>
                  <div class="cube-face-mini cube-right-mini"></div>
                  <div class="cube-face-mini cube-left-mini"></div>
                  <div class="cube-face-mini cube-top-mini"></div>
                  <div class="cube-face-mini cube-bottom-mini"></div>
                </div>
                <span class="text-green-300">{{ service.name }}</span>
              </div>
              <div class="flex items-center space-x-2">
                <div class="w-20 bg-gray-700 rounded-full h-2">
                  <div 
                    class="bg-green-400 h-2 rounded-full transition-all duration-1000"
                    :style="{ width: `${service.percentage}%` }"
                  ></div>
                </div>
                <span class="text-green-300 text-sm w-12 text-right">{{ service.count }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// 页面元数据
useHead({
  title: '日志管理 - Hephaestus',
  meta: [
    { name: 'description', content: '实时监控和分析微服务日志信息' }
  ]
})

// 响应式数据
const searchQuery = ref('')
const selectedService = ref('')
const logLevel = ref('')
const timeRange = ref('1h')
const autoRefresh = ref(true)
const logContainer = ref(null)

// 选项数据
const serviceOptions = [
  { label: '全部服务', value: '' },
  { label: 'User Service', value: 'User Service' },
  { label: 'Order Service', value: 'Order Service' },
  { label: 'Payment Service', value: 'Payment Service' },
  { label: 'Inventory Service', value: 'Inventory Service' },
  { label: 'Notification Service', value: 'Notification Service' }
]

const logLevelOptions = [
  { label: '全部级别', value: '' },
  { label: 'ERROR', value: 'ERROR' },
  { label: 'WARN', value: 'WARN' },
  { label: 'INFO', value: 'INFO' },
  { label: 'DEBUG', value: 'DEBUG' }
]

const timeRangeOptions = [
  { label: '最近1小时', value: '1h' },
  { label: '最近6小时', value: '6h' },
  { label: '最近24小时', value: '24h' },
  { label: '最近7天', value: '7d' }
]

// 模拟日志数据
const logs = ref([
  {
    id: 1,
    timestamp: '2024-01-15 16:30:25.123',
    level: 'ERROR',
    service: 'User Service',
    message: 'Database connection failed: Connection timeout after 30 seconds',
    stack: 'java.sql.SQLException: Connection timeout\n\tat com.example.UserService.connect(UserService.java:45)\n\tat com.example.UserService.getUser(UserService.java:23)',
    context: { userId: 12345, requestId: 'req-abc123' },
    expanded: false
  },
  {
    id: 2,
    timestamp: '2024-01-15 16:30:20.456',
    level: 'WARN',
    service: 'Order Service',
    message: 'High memory usage detected: 85% of heap space used',
    context: { memoryUsage: '85%', heapSize: '2GB' },
    expanded: false
  },
  {
    id: 3,
    timestamp: '2024-01-15 16:30:15.789',
    level: 'INFO',
    service: 'Payment Service',
    message: 'Payment processed successfully for order #ORD-2024-001',
    context: { orderId: 'ORD-2024-001', amount: 299.99, currency: 'CNY' },
    expanded: false
  },
  {
    id: 4,
    timestamp: '2024-01-15 16:30:10.012',
    level: 'DEBUG',
    service: 'Inventory Service',
    message: 'Checking inventory for product SKU: PROD-12345',
    context: { productSku: 'PROD-12345', currentStock: 150 },
    expanded: false
  },
  {
    id: 5,
    timestamp: '2024-01-15 16:30:05.345',
    level: 'INFO',
    service: 'Notification Service',
    message: 'Email notification sent to user@example.com',
    context: { recipient: 'user@example.com', template: 'order_confirmation' },
    expanded: false
  }
])

// 计算属性
const totalLogs = computed(() => logs.value.length)
const errorLogs = computed(() => logs.value.filter(log => log.level === 'ERROR').length)
const warningLogs = computed(() => logs.value.filter(log => log.level === 'WARN').length)
const infoLogs = computed(() => logs.value.filter(log => log.level === 'INFO').length)
const debugLogs = computed(() => logs.value.filter(log => log.level === 'DEBUG').length)

const filteredLogs = computed(() => {
  let filtered = logs.value
  
  if (searchQuery.value) {
    filtered = filtered.filter(log => 
      log.message.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      log.service.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
  }
  
  if (selectedService.value) {
    filtered = filtered.filter(log => log.service === selectedService.value)
  }
  
  if (logLevel.value) {
    filtered = filtered.filter(log => log.level === logLevel.value)
  }
  
  return filtered.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp))
})

const logLevelStats = computed(() => {
  const stats = [
    { name: 'ERROR', count: errorLogs.value },
    { name: 'WARN', count: warningLogs.value },
    { name: 'INFO', count: infoLogs.value },
    { name: 'DEBUG', count: debugLogs.value }
  ]
  
  const total = totalLogs.value
  return stats.map(stat => ({
    ...stat,
    percentage: total > 0 ? Math.round((stat.count / total) * 100) : 0
  }))
})

const serviceLogStats = computed(() => {
  const serviceMap = {}
  logs.value.forEach(log => {
    serviceMap[log.service] = (serviceMap[log.service] || 0) + 1
  })
  
  const stats = Object.entries(serviceMap).map(([name, count]) => ({ name, count }))
  const maxCount = Math.max(...stats.map(s => s.count))
  
  return stats.map(stat => ({
    ...stat,
    percentage: maxCount > 0 ? Math.round((stat.count / maxCount) * 100) : 0
  }))
})

// 方法
const getLogEntryClass = (level) => {
  const classes = {
    'ERROR': 'border-red-500/50 bg-red-900/10',
    'WARN': 'border-yellow-500/50 bg-yellow-900/10',
    'INFO': 'border-blue-500/50 bg-blue-900/10',
    'DEBUG': 'border-gray-500/50 bg-gray-900/10'
  }
  return classes[level] || 'border-gray-500/50'
}

const getLogLevelClass = (level) => {
  const classes = {
    'ERROR': 'bg-red-900/50 text-red-300 border border-red-500/50',
    'WARN': 'bg-yellow-900/50 text-yellow-300 border border-yellow-500/50',
    'INFO': 'bg-blue-900/50 text-blue-300 border border-blue-500/50',
    'DEBUG': 'bg-gray-900/50 text-gray-300 border border-gray-500/50'
  }
  return classes[level] || 'bg-gray-900/50 text-gray-300'
}

const getLogLevelColor = (level) => {
  const colors = {
    'ERROR': 'bg-red-400',
    'WARN': 'bg-yellow-400',
    'INFO': 'bg-blue-400',
    'DEBUG': 'bg-gray-400'
  }
  return colors[level] || 'bg-gray-400'
}

const toggleLogExpand = (log) => {
  log.expanded = !log.expanded
}

const copyLog = (log) => {
  const logText = `[${log.timestamp}] ${log.level} ${log.service}: ${log.message}`
  navigator.clipboard.writeText(logText)
  console.log('日志已复制到剪贴板')
}

const refreshLogs = () => {
  // 模拟新日志
  const newLog = {
    id: logs.value.length + 1,
    timestamp: new Date().toISOString().replace('T', ' ').substring(0, 23),
    level: ['ERROR', 'WARN', 'INFO', 'DEBUG'][Math.floor(Math.random() * 4)],
    service: serviceOptions[Math.floor(Math.random() * (serviceOptions.length - 1)) + 1].value,
    message: '模拟日志消息 - ' + Math.random().toString(36).substring(7),
    context: { requestId: 'req-' + Math.random().toString(36).substring(7) },
    expanded: false
  }
  
  logs.value.unshift(newLog)
  
  // 自动滚动到顶部
  nextTick(() => {
    if (logContainer.value) {
      logContainer.value.scrollTop = 0
    }
  })
}

const clearLogs = () => {
  logs.value = []
}

const exportLogs = () => {
  const logData = filteredLogs.value.map(log => ({
    timestamp: log.timestamp,
    level: log.level,
    service: log.service,
    message: log.message
  }))
  
  const blob = new Blob([JSON.stringify(logData, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `logs-${new Date().toISOString().split('T')[0]}.json`
  a.click()
  URL.revokeObjectURL(url)
}

// 自动刷新
let refreshInterval = null

watch(autoRefresh, (newValue) => {
  if (newValue) {
    refreshInterval = setInterval(refreshLogs, 5000)
  } else {
    if (refreshInterval) {
      clearInterval(refreshInterval)
      refreshInterval = null
    }
  }
})

// 组件挂载时启动自动刷新
onMounted(() => {
  if (autoRefresh.value) {
    refreshInterval = setInterval(refreshLogs, 5000)
  }
})

// 组件卸载时清除定时器
onUnmounted(() => {
  if (refreshInterval) {
    clearInterval(refreshInterval)
  }
})
</script>

<style scoped>
/* 迷你立方体样式 */
.microservice-cube-mini {
  width: 16px;
  height: 16px;
  position: relative;
  transform-style: preserve-3d;
  animation: float 3s ease-in-out infinite;
}

.cube-face-mini {
  position: absolute;
  width: 16px;
  height: 16px;
  background: linear-gradient(135deg, var(--primary-green), var(--secondary-green));
  border: 1px solid var(--primary-green);
  opacity: 0.8;
}

.cube-front-mini {
  transform: rotateY(0deg) translateZ(8px);
}

.cube-back-mini {
  transform: rotateY(180deg) translateZ(8px);
}

.cube-right-mini {
  transform: rotateY(90deg) translateZ(8px);
}

.cube-left-mini {
  transform: rotateY(-90deg) translateZ(8px);
}

.cube-top-mini {
  transform: rotateX(90deg) translateZ(8px);
}

.cube-bottom-mini {
  transform: rotateX(-90deg) translateZ(8px);
}

/* 日志条目动画 */
.log-entry {
  animation: slideInLeft 0.3s ease-out;
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* 自定义滚动条 */
.overflow-y-auto::-webkit-scrollbar {
  width: 6px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: var(--bg-darker);
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background: var(--primary-green);
  border-radius: 3px;
}

.overflow-y-auto::-webkit-scrollbar-thumb:hover {
  background: var(--secondary-green);
}
</style>