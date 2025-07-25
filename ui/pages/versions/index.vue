<template>
  <div class="min-h-screen bg-gray-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- 页面标题 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-green-400 text-glow mb-4">版本变更管理</h1>
        <p class="text-green-300">跟踪和管理微服务的版本发布历史</p>
      </div>
      
      <!-- 操作栏 -->
      <div class="mb-8 flex flex-col md:flex-row gap-4 justify-between">
        <div class="flex gap-4">
          <USelect 
            v-model="selectedService"
            :options="serviceOptions"
            placeholder="选择服务"
            class="w-48"
          />
          <USelect 
            v-model="timeFilter"
            :options="timeOptions"
            placeholder="时间范围"
            class="w-32"
          />
        </div>
        <UButton 
          @click="showNewVersionModal = true"
          icon="heroicons:plus"
          color="green"
          size="lg"
        >
          新建版本
        </UButton>
      </div>
      
      <!-- 版本统计 -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">总版本数</p>
              <p class="text-2xl font-bold text-green-400">{{ totalVersions }}</p>
            </div>
            <Icon name="heroicons:code-bracket" class="w-8 h-8 text-green-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">本周发布</p>
              <p class="text-2xl font-bold text-green-400">{{ weeklyReleases }}</p>
            </div>
            <Icon name="heroicons:rocket-launch" class="w-8 h-8 text-green-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">回滚次数</p>
              <p class="text-2xl font-bold text-yellow-400">{{ rollbackCount }}</p>
            </div>
            <Icon name="heroicons:arrow-uturn-left" class="w-8 h-8 text-yellow-400" />
          </div>
        </div>
        
        <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-green-300 text-sm">成功率</p>
              <p class="text-2xl font-bold text-green-400">{{ successRate }}%</p>
            </div>
            <Icon name="heroicons:check-circle" class="w-8 h-8 text-green-400" />
          </div>
        </div>
      </div>
      
      <!-- 版本时间线 -->
      <div class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow">
        <h2 class="text-2xl font-bold text-green-400 mb-6">版本发布时间线</h2>
        
        <div class="relative">
          <!-- 时间线主线 -->
          <div class="absolute left-8 top-0 bottom-0 w-0.5 bg-green-500/30"></div>
          
          <!-- 版本节点 -->
          <div class="space-y-8">
            <div 
              v-for="(version, index) in filteredVersions" 
              :key="version.id"
              class="relative flex items-start"
            >
              <!-- 时间线节点 -->
              <div class="relative z-10">
                <div 
                  class="w-4 h-4 rounded-full border-2 border-green-500 flex items-center justify-center"
                  :class="getVersionStatusClass(version.status)"
                >
                  <div class="w-2 h-2 rounded-full" :class="getVersionDotClass(version.status)"></div>
                </div>
              </div>
              
              <!-- 版本信息卡片 -->
              <div class="ml-6 flex-1">
                <div class="bg-gray-800/50 border border-green-500/20 rounded-lg p-6 hover:border-green-500/40 transition-all duration-300">
                  <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-4">
                    <div class="flex items-center space-x-4">
                      <h3 class="text-xl font-semibold text-green-300">{{ version.service }}</h3>
                      <span class="text-lg font-mono text-green-400">{{ version.version }}</span>
                      <span 
                        class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="getStatusBadgeClass(version.status)"
                      >
                        {{ version.status }}
                      </span>
                    </div>
                    <div class="text-sm text-green-400 mt-2 md:mt-0">
                      {{ version.releaseDate }}
                    </div>
                  </div>
                  
                  <p class="text-green-300 mb-4">{{ version.description }}</p>
                  
                  <!-- 变更详情 -->
                  <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
                    <div class="flex items-center space-x-2">
                      <Icon name="heroicons:plus-circle" class="w-4 h-4 text-green-400" />
                      <span class="text-sm text-green-300">{{ version.additions }} 新增</span>
                    </div>
                    <div class="flex items-center space-x-2">
                      <Icon name="heroicons:pencil-square" class="w-4 h-4 text-blue-400" />
                      <span class="text-sm text-green-300">{{ version.modifications }} 修改</span>
                    </div>
                    <div class="flex items-center space-x-2">
                      <Icon name="heroicons:minus-circle" class="w-4 h-4 text-red-400" />
                      <span class="text-sm text-green-300">{{ version.deletions }} 删除</span>
                    </div>
                  </div>
                  
                  <!-- 发布者信息 -->
                  <div class="flex items-center justify-between">
                    <div class="flex items-center space-x-2">
                      <Icon name="heroicons:user-circle" class="w-5 h-5 text-green-400" />
                      <span class="text-sm text-green-300">{{ version.author }}</span>
                    </div>
                    
                    <!-- 操作按钮 -->
                    <div class="flex space-x-2">
                      <UButton 
                        size="sm" 
                        variant="ghost" 
                        color="green"
                        icon="heroicons:eye"
                        @click="viewVersionDetails(version)"
                      >
                        详情
                      </UButton>
                      <UButton 
                        v-if="version.status === '已发布' && index > 0"
                        size="sm" 
                        variant="ghost" 
                        color="yellow"
                        icon="heroicons:arrow-uturn-left"
                        @click="rollbackVersion(version)"
                      >
                        回滚
                      </UButton>
                      <UButton 
                        v-if="version.status === '待发布'"
                        size="sm" 
                        variant="ghost" 
                        color="blue"
                        icon="heroicons:rocket-launch"
                        @click="deployVersion(version)"
                      >
                        部署
                      </UButton>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 新建版本模态框 -->
    <UModal v-model="showNewVersionModal">
      <div class="p-6">
        <h3 class="text-lg font-semibold text-green-400 mb-4">创建新版本</h3>
        
        <form @submit.prevent="createNewVersion" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">服务名称</label>
            <USelect 
              v-model="newVersion.service"
              :options="serviceOptions"
              placeholder="选择服务"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">版本号</label>
            <UInput 
              v-model="newVersion.version"
              placeholder="例如: v1.2.3"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">变更描述</label>
            <UTextarea 
              v-model="newVersion.description"
              placeholder="描述此版本的主要变更..."
              rows="4"
              required
            />
          </div>
          
          <div class="flex justify-end space-x-3 pt-4">
            <UButton 
              type="button" 
              variant="ghost" 
              @click="showNewVersionModal = false"
            >
              取消
            </UButton>
            <UButton 
              type="submit" 
              color="green"
            >
              创建版本
            </UButton>
          </div>
        </form>
      </div>
    </UModal>
  </div>
</template>

<script setup>
// 页面元数据
useHead({
  title: '版本变更 - Hephaestus',
  meta: [
    { name: 'description', content: '管理微服务版本发布和变更历史' }
  ]
})

// 响应式数据
const selectedService = ref('')
const timeFilter = ref('all')
const showNewVersionModal = ref(false)

const serviceOptions = [
  { label: '全部服务', value: '' },
  { label: 'User Service', value: 'User Service' },
  { label: 'Order Service', value: 'Order Service' },
  { label: 'Payment Service', value: 'Payment Service' },
  { label: 'Inventory Service', value: 'Inventory Service' },
  { label: 'Notification Service', value: 'Notification Service' }
]

const timeOptions = [
  { label: '全部', value: 'all' },
  { label: '本周', value: 'week' },
  { label: '本月', value: 'month' },
  { label: '最近3个月', value: 'quarter' }
]

// 新版本表单数据
const newVersion = ref({
  service: '',
  version: '',
  description: ''
})

// 模拟版本数据
const versions = ref([
  {
    id: 1,
    service: 'User Service',
    version: 'v1.2.3',
    status: '已发布',
    description: '修复用户登录问题，优化性能，新增用户头像上传功能',
    releaseDate: '2024-01-15 14:30',
    author: '张三',
    additions: 15,
    modifications: 8,
    deletions: 3
  },
  {
    id: 2,
    service: 'Order Service',
    version: 'v2.1.0',
    status: '已发布',
    description: '重构订单处理逻辑，支持批量订单处理，优化数据库查询',
    releaseDate: '2024-01-14 16:45',
    author: '李四',
    additions: 25,
    modifications: 12,
    deletions: 5
  },
  {
    id: 3,
    service: 'Payment Service',
    version: 'v1.5.2',
    status: '部署中',
    description: '集成新的支付网关，增强安全性验证',
    releaseDate: '2024-01-15 10:20',
    author: '王五',
    additions: 18,
    modifications: 6,
    deletions: 2
  },
  {
    id: 4,
    service: 'Inventory Service',
    version: 'v1.0.9',
    status: '待发布',
    description: '库存预警功能，自动补货逻辑优化',
    releaseDate: '2024-01-15 09:15',
    author: '赵六',
    additions: 22,
    modifications: 10,
    deletions: 1
  },
  {
    id: 5,
    service: 'Notification Service',
    version: 'v1.3.1',
    status: '发布失败',
    description: '推送服务稳定性改进，支持多渠道推送',
    releaseDate: '2024-01-14 20:30',
    author: '孙七',
    additions: 12,
    modifications: 15,
    deletions: 4
  }
])

// 计算属性
const totalVersions = computed(() => versions.value.length)
const weeklyReleases = computed(() => {
  const oneWeekAgo = new Date()
  oneWeekAgo.setDate(oneWeekAgo.getDate() - 7)
  return versions.value.filter(v => new Date(v.releaseDate) > oneWeekAgo).length
})
const rollbackCount = computed(() => 3) // 模拟数据
const successRate = computed(() => {
  const successful = versions.value.filter(v => v.status === '已发布').length
  return Math.round((successful / versions.value.length) * 100)
})

const filteredVersions = computed(() => {
  let filtered = versions.value
  
  if (selectedService.value) {
    filtered = filtered.filter(v => v.service === selectedService.value)
  }
  
  if (timeFilter.value !== 'all') {
    const now = new Date()
    let cutoffDate
    
    switch (timeFilter.value) {
      case 'week':
        cutoffDate = new Date(now.getTime() - 7 * 24 * 60 * 60 * 1000)
        break
      case 'month':
        cutoffDate = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000)
        break
      case 'quarter':
        cutoffDate = new Date(now.getTime() - 90 * 24 * 60 * 60 * 1000)
        break
    }
    
    filtered = filtered.filter(v => new Date(v.releaseDate) > cutoffDate)
  }
  
  return filtered.sort((a, b) => new Date(b.releaseDate) - new Date(a.releaseDate))
})

// 方法
const getVersionStatusClass = (status) => {
  switch (status) {
    case '已发布':
      return 'bg-green-900/50'
    case '部署中':
      return 'bg-blue-900/50'
    case '待发布':
      return 'bg-yellow-900/50'
    case '发布失败':
      return 'bg-red-900/50'
    default:
      return 'bg-gray-900/50'
  }
}

const getVersionDotClass = (status) => {
  switch (status) {
    case '已发布':
      return 'bg-green-400'
    case '部署中':
      return 'bg-blue-400 pulse-animation'
    case '待发布':
      return 'bg-yellow-400'
    case '发布失败':
      return 'bg-red-400'
    default:
      return 'bg-gray-400'
  }
}

const getStatusBadgeClass = (status) => {
  switch (status) {
    case '已发布':
      return 'bg-green-900/50 text-green-300 border border-green-500/50'
    case '部署中':
      return 'bg-blue-900/50 text-blue-300 border border-blue-500/50'
    case '待发布':
      return 'bg-yellow-900/50 text-yellow-300 border border-yellow-500/50'
    case '发布失败':
      return 'bg-red-900/50 text-red-300 border border-red-500/50'
    default:
      return 'bg-gray-900/50 text-gray-300 border border-gray-500/50'
  }
}

const viewVersionDetails = (version) => {
  console.log('查看版本详情:', version)
}

const rollbackVersion = (version) => {
  console.log('回滚版本:', version)
}

const deployVersion = (version) => {
  version.status = '部署中'
  console.log('部署版本:', version)
}

const createNewVersion = () => {
  const version = {
    id: versions.value.length + 1,
    service: newVersion.value.service,
    version: newVersion.value.version,
    status: '待发布',
    description: newVersion.value.description,
    releaseDate: new Date().toLocaleString('zh-CN'),
    author: '当前用户',
    additions: Math.floor(Math.random() * 20) + 5,
    modifications: Math.floor(Math.random() * 15) + 3,
    deletions: Math.floor(Math.random() * 5) + 1
  }
  
  versions.value.unshift(version)
  showNewVersionModal.value = false
  
  // 重置表单
  newVersion.value = {
    service: '',
    version: '',
    description: ''
  }
}
</script>

<style scoped>
/* 时间线样式增强 */
.timeline-node {
  @apply relative z-10;
}

.timeline-node::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: var(--primary-green);
  opacity: 0.3;
  animation: pulse 2s ease-in-out infinite;
}
</style>