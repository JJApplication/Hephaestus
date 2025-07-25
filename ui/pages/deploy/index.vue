<template>
  <div class="min-h-screen bg-gray-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- 页面标题 -->
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-green-400 text-glow mb-4">部署管理</h1>
        <p class="text-green-300">管理微服务的部署流程和环境配置</p>
      </div>
      
      <!-- 快速操作 -->
      <div class="mb-8 flex flex-col md:flex-row gap-4">
        <UButton 
          @click="showDeployModal = true"
          icon="heroicons:rocket-launch"
          color="green"
          size="lg"
        >
          新建部署
        </UButton>
        <UButton 
          @click="refreshDeployments"
          icon="heroicons:arrow-path"
          variant="outline"
          color="green"
        >
          刷新状态
        </UButton>
        <UButton 
          @click="showPipelineModal = true"
          icon="heroicons:cog-6-tooth"
          variant="outline"
          color="blue"
        >
          流水线配置
        </UButton>
      </div>
      
      <!-- 环境概览 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div 
          v-for="env in environments" 
          :key="env.name"
          class="bg-black/50 border border-green-500/30 rounded-lg p-6 glow hover:border-green-500/50 transition-all duration-300"
        >
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-xl font-semibold text-green-300">{{ env.name }}</h3>
            <span 
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
              :class="getEnvStatusClass(env.status)"
            >
              <span 
                class="w-2 h-2 rounded-full mr-1.5"
                :class="getEnvDotClass(env.status)"
              ></span>
              {{ env.status }}
            </span>
          </div>
          
          <div class="space-y-3">
            <div class="flex justify-between">
              <span class="text-green-400 text-sm">运行服务</span>
              <span class="text-green-300 font-medium">{{ env.runningServices }}/{{ env.totalServices }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-400 text-sm">CPU 使用率</span>
              <span class="text-green-300 font-medium">{{ env.cpuUsage }}%</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-400 text-sm">内存使用率</span>
              <span class="text-green-300 font-medium">{{ env.memoryUsage }}%</span>
            </div>
            <div class="flex justify-between">
              <span class="text-green-400 text-sm">最后部署</span>
              <span class="text-green-300 font-medium text-xs">{{ env.lastDeploy }}</span>
            </div>
          </div>
          
          <div class="mt-4 pt-4 border-t border-green-500/20">
            <UButton 
              @click="viewEnvironment(env)"
              variant="ghost"
              color="green"
              size="sm"
              block
            >
              查看详情
            </UButton>
          </div>
        </div>
      </div>
      
      <!-- 部署历史 -->
      <div class="bg-black/50 border border-green-500/30 rounded-lg overflow-hidden glow">
        <div class="p-6 border-b border-green-500/20">
          <div class="flex items-center justify-between">
            <h2 class="text-2xl font-bold text-green-400">部署历史</h2>
            <div class="flex gap-4">
              <USelect 
                v-model="envFilter"
                :options="envFilterOptions"
                placeholder="环境筛选"
                class="w-32"
              />
              <USelect 
                v-model="statusFilter"
                :options="statusFilterOptions"
                placeholder="状态筛选"
                class="w-32"
              />
            </div>
          </div>
        </div>
        
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead class="bg-green-900/20">
              <tr>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">部署ID</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">服务</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">版本</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">环境</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">状态</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">进度</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">开始时间</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">耗时</th>
                <th class="px-6 py-4 text-left text-sm font-medium text-green-300">操作</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-green-500/20">
              <tr 
                v-for="deployment in filteredDeployments" 
                :key="deployment.id"
                class="hover:bg-green-900/10 transition-colors"
              >
                <td class="px-6 py-4">
                  <span class="font-mono text-green-400">#{{ deployment.id }}</span>
                </td>
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="microservice-cube-mini mr-3">
                      <div class="cube-face-mini cube-front-mini"></div>
                      <div class="cube-face-mini cube-back-mini"></div>
                      <div class="cube-face-mini cube-right-mini"></div>
                      <div class="cube-face-mini cube-left-mini"></div>
                      <div class="cube-face-mini cube-top-mini"></div>
                      <div class="cube-face-mini cube-bottom-mini"></div>
                    </div>
                    <span class="text-green-300 font-medium">{{ deployment.service }}</span>
                  </div>
                </td>
                <td class="px-6 py-4">
                  <span class="font-mono text-green-400">{{ deployment.version }}</span>
                </td>
                <td class="px-6 py-4">
                  <span 
                    class="inline-flex items-center px-2 py-1 rounded text-xs font-medium"
                    :class="getEnvBadgeClass(deployment.environment)"
                  >
                    {{ deployment.environment }}
                  </span>
                </td>
                <td class="px-6 py-4">
                  <span 
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="getDeployStatusClass(deployment.status)"
                  >
                    <span 
                      class="w-2 h-2 rounded-full mr-1.5"
                      :class="getDeployDotClass(deployment.status)"
                    ></span>
                    {{ deployment.status }}
                  </span>
                </td>
                <td class="px-6 py-4">
                  <div class="flex items-center">
                    <div class="w-20 bg-gray-700 rounded-full h-2 mr-2">
                      <div 
                        class="h-2 rounded-full transition-all duration-1000"
                        :class="getProgressBarClass(deployment.status)"
                        :style="{ width: `${deployment.progress}%` }"
                      ></div>
                    </div>
                    <span class="text-green-300 text-sm">{{ deployment.progress }}%</span>
                  </div>
                </td>
                <td class="px-6 py-4 text-green-300 text-sm">{{ deployment.startTime }}</td>
                <td class="px-6 py-4 text-green-300 text-sm">{{ deployment.duration }}</td>
                <td class="px-6 py-4">
                  <div class="flex space-x-2">
                    <UButton 
                      size="sm" 
                      variant="ghost" 
                      color="green"
                      icon="heroicons:eye"
                      @click="viewDeploymentLogs(deployment)"
                    />
                    <UButton 
                      v-if="deployment.status === '部署中'"
                      size="sm" 
                      variant="ghost" 
                      color="red"
                      icon="heroicons:stop"
                      @click="stopDeployment(deployment)"
                    />
                    <UButton 
                      v-if="deployment.status === '失败'"
                      size="sm" 
                      variant="ghost" 
                      color="blue"
                      icon="heroicons:arrow-path"
                      @click="retryDeployment(deployment)"
                    />
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
    
    <!-- 新建部署模态框 -->
    <UModal v-model="showDeployModal">
      <div class="p-6">
        <h3 class="text-lg font-semibold text-green-400 mb-4">新建部署</h3>
        
        <form @submit.prevent="createDeployment" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">服务</label>
            <USelect 
              v-model="newDeployment.service"
              :options="serviceOptions"
              placeholder="选择服务"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">版本</label>
            <USelect 
              v-model="newDeployment.version"
              :options="versionOptions"
              placeholder="选择版本"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">目标环境</label>
            <USelect 
              v-model="newDeployment.environment"
              :options="envOptions"
              placeholder="选择环境"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">部署策略</label>
            <USelect 
              v-model="newDeployment.strategy"
              :options="strategyOptions"
              placeholder="选择部署策略"
              required
            />
          </div>
          
          <div>
            <label class="block text-sm font-medium text-green-300 mb-2">备注</label>
            <UTextarea 
              v-model="newDeployment.notes"
              placeholder="部署备注信息..."
              rows="3"
            />
          </div>
          
          <div class="flex justify-end space-x-3 pt-4">
            <UButton 
              type="button" 
              variant="ghost" 
              @click="showDeployModal = false"
            >
              取消
            </UButton>
            <UButton 
              type="submit" 
              color="green"
            >
              开始部署
            </UButton>
          </div>
        </form>
      </div>
    </UModal>
    
    <!-- 流水线配置模态框 -->
    <UModal v-model="showPipelineModal">
      <div class="p-6">
        <h3 class="text-lg font-semibold text-green-400 mb-4">CI/CD 流水线配置</h3>
        
        <div class="space-y-6">
          <!-- 流水线阶段 -->
          <div>
            <h4 class="text-md font-medium text-green-300 mb-3">流水线阶段</h4>
            <div class="space-y-3">
              <div 
                v-for="(stage, index) in pipelineStages" 
                :key="stage.name"
                class="flex items-center justify-between p-3 bg-gray-800/50 rounded border border-green-500/20"
              >
                <div class="flex items-center space-x-3">
                  <span class="text-green-400 font-mono">{{ index + 1 }}</span>
                  <Icon :name="stage.icon" class="w-5 h-5 text-green-400" />
                  <span class="text-green-300">{{ stage.name }}</span>
                </div>
                <UToggle v-model="stage.enabled" color="green" />
              </div>
            </div>
          </div>
          
          <!-- 自动化设置 -->
          <div>
            <h4 class="text-md font-medium text-green-300 mb-3">自动化设置</h4>
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <span class="text-green-300">自动部署到测试环境</span>
                <UToggle v-model="autoDeployTest" color="green" />
              </div>
              <div class="flex items-center justify-between">
                <span class="text-green-300">自动回滚失败部署</span>
                <UToggle v-model="autoRollback" color="green" />
              </div>
              <div class="flex items-center justify-between">
                <span class="text-green-300">部署前健康检查</span>
                <UToggle v-model="healthCheck" color="green" />
              </div>
            </div>
          </div>
        </div>
        
        <div class="flex justify-end space-x-3 pt-6">
          <UButton 
            variant="ghost" 
            @click="showPipelineModal = false"
          >
            取消
          </UButton>
          <UButton 
            color="green"
            @click="savePipelineConfig"
          >
            保存配置
          </UButton>
        </div>
      </div>
    </UModal>
  </div>
</template>

<script setup>
// 页面元数据
useHead({
  title: '部署管理 - Hephaestus',
  meta: [
    { name: 'description', content: '管理微服务部署流程和环境配置' }
  ]
})

// 响应式数据
const showDeployModal = ref(false)
const showPipelineModal = ref(false)
const envFilter = ref('')
const statusFilter = ref('')

// 环境数据
const environments = ref([
  {
    name: '开发环境',
    status: '正常',
    runningServices: 8,
    totalServices: 8,
    cpuUsage: 35,
    memoryUsage: 45,
    lastDeploy: '2小时前'
  },
  {
    name: '测试环境',
    status: '正常',
    runningServices: 7,
    totalServices: 8,
    cpuUsage: 28,
    memoryUsage: 38,
    lastDeploy: '1天前'
  },
  {
    name: '生产环境',
    status: '正常',
    runningServices: 8,
    totalServices: 8,
    cpuUsage: 52,
    memoryUsage: 68,
    lastDeploy: '3天前'
  }
])

// 部署历史数据
const deployments = ref([
  {
    id: '2024011501',
    service: 'User Service',
    version: 'v1.2.3',
    environment: '生产环境',
    status: '成功',
    progress: 100,
    startTime: '2024-01-15 14:30',
    duration: '3分25秒'
  },
  {
    id: '2024011502',
    service: 'Order Service',
    version: 'v2.1.0',
    environment: '测试环境',
    status: '部署中',
    progress: 65,
    startTime: '2024-01-15 15:45',
    duration: '1分30秒'
  },
  {
    id: '2024011503',
    service: 'Payment Service',
    version: 'v1.5.2',
    environment: '开发环境',
    status: '失败',
    progress: 45,
    startTime: '2024-01-15 16:20',
    duration: '2分10秒'
  }
])

// 表单选项
const serviceOptions = [
  { label: 'User Service', value: 'User Service' },
  { label: 'Order Service', value: 'Order Service' },
  { label: 'Payment Service', value: 'Payment Service' },
  { label: 'Inventory Service', value: 'Inventory Service' },
  { label: 'Notification Service', value: 'Notification Service' }
]

const versionOptions = [
  { label: 'v1.2.3', value: 'v1.2.3' },
  { label: 'v2.1.0', value: 'v2.1.0' },
  { label: 'v1.5.2', value: 'v1.5.2' }
]

const envOptions = [
  { label: '开发环境', value: '开发环境' },
  { label: '测试环境', value: '测试环境' },
  { label: '生产环境', value: '生产环境' }
]

const strategyOptions = [
  { label: '滚动更新', value: 'rolling' },
  { label: '蓝绿部署', value: 'blue-green' },
  { label: '金丝雀发布', value: 'canary' }
]

const envFilterOptions = [
  { label: '全部环境', value: '' },
  { label: '开发环境', value: '开发环境' },
  { label: '测试环境', value: '测试环境' },
  { label: '生产环境', value: '生产环境' }
]

const statusFilterOptions = [
  { label: '全部状态', value: '' },
  { label: '成功', value: '成功' },
  { label: '部署中', value: '部署中' },
  { label: '失败', value: '失败' }
]

// 新部署表单
const newDeployment = ref({
  service: '',
  version: '',
  environment: '',
  strategy: '',
  notes: ''
})

// 流水线配置
const pipelineStages = ref([
  { name: '代码检出', icon: 'heroicons:code-bracket', enabled: true },
  { name: '单元测试', icon: 'heroicons:beaker', enabled: true },
  { name: '代码扫描', icon: 'heroicons:shield-check', enabled: true },
  { name: '构建镜像', icon: 'heroicons:cube', enabled: true },
  { name: '安全扫描', icon: 'heroicons:lock-closed', enabled: true },
  { name: '部署测试', icon: 'heroicons:rocket-launch', enabled: true }
])

const autoDeployTest = ref(true)
const autoRollback = ref(true)
const healthCheck = ref(true)

// 计算属性
const filteredDeployments = computed(() => {
  let filtered = deployments.value
  
  if (envFilter.value) {
    filtered = filtered.filter(d => d.environment === envFilter.value)
  }
  
  if (statusFilter.value) {
    filtered = filtered.filter(d => d.status === statusFilter.value)
  }
  
  return filtered
})

// 方法
const getEnvStatusClass = (status) => {
  return status === '正常' 
    ? 'bg-green-900/50 text-green-300 border border-green-500/50'
    : 'bg-red-900/50 text-red-300 border border-red-500/50'
}

const getEnvDotClass = (status) => {
  return status === '正常' 
    ? 'bg-green-400 pulse-animation'
    : 'bg-red-400 pulse-animation'
}

const getEnvBadgeClass = (env) => {
  const classes = {
    '开发环境': 'bg-blue-900/50 text-blue-300',
    '测试环境': 'bg-yellow-900/50 text-yellow-300',
    '生产环境': 'bg-red-900/50 text-red-300'
  }
  return classes[env] || 'bg-gray-900/50 text-gray-300'
}

const getDeployStatusClass = (status) => {
  const classes = {
    '成功': 'bg-green-900/50 text-green-300 border border-green-500/50',
    '部署中': 'bg-blue-900/50 text-blue-300 border border-blue-500/50',
    '失败': 'bg-red-900/50 text-red-300 border border-red-500/50'
  }
  return classes[status] || 'bg-gray-900/50 text-gray-300 border border-gray-500/50'
}

const getDeployDotClass = (status) => {
  const classes = {
    '成功': 'bg-green-400',
    '部署中': 'bg-blue-400 pulse-animation',
    '失败': 'bg-red-400'
  }
  return classes[status] || 'bg-gray-400'
}

const getProgressBarClass = (status) => {
  const classes = {
    '成功': 'bg-green-400',
    '部署中': 'bg-blue-400',
    '失败': 'bg-red-400'
  }
  return classes[status] || 'bg-gray-400'
}

const refreshDeployments = () => {
  console.log('刷新部署状态')
}

const viewEnvironment = (env) => {
  console.log('查看环境:', env.name)
}

const viewDeploymentLogs = (deployment) => {
  console.log('查看部署日志:', deployment.id)
}

const stopDeployment = (deployment) => {
  deployment.status = '已停止'
  console.log('停止部署:', deployment.id)
}

const retryDeployment = (deployment) => {
  deployment.status = '部署中'
  deployment.progress = 0
  console.log('重试部署:', deployment.id)
}

const createDeployment = () => {
  const deployment = {
    id: Date.now().toString(),
    service: newDeployment.value.service,
    version: newDeployment.value.version,
    environment: newDeployment.value.environment,
    status: '部署中',
    progress: 0,
    startTime: new Date().toLocaleString('zh-CN'),
    duration: '0秒'
  }
  
  deployments.value.unshift(deployment)
  showDeployModal.value = false
  
  // 模拟部署进度
  const progressInterval = setInterval(() => {
    if (deployment.progress < 100) {
      deployment.progress += Math.floor(Math.random() * 15) + 5
      if (deployment.progress >= 100) {
        deployment.progress = 100
        deployment.status = '成功'
        clearInterval(progressInterval)
      }
    }
  }, 1000)
  
  // 重置表单
  newDeployment.value = {
    service: '',
    version: '',
    environment: '',
    strategy: '',
    notes: ''
  }
}

const savePipelineConfig = () => {
  console.log('保存流水线配置')
  showPipelineModal.value = false
}
</script>

<style scoped>
/* 迷你立方体样式 */
.microservice-cube-mini {
  width: 20px;
  height: 20px;
  position: relative;
  transform-style: preserve-3d;
  animation: float 3s ease-in-out infinite;
}

.cube-face-mini {
  position: absolute;
  width: 20px;
  height: 20px;
  background: linear-gradient(135deg, var(--primary-green), var(--secondary-green));
  border: 1px solid var(--primary-green);
  opacity: 0.8;
}

.cube-front-mini {
  transform: rotateY(0deg) translateZ(10px);
}

.cube-back-mini {
  transform: rotateY(180deg) translateZ(10px);
}

.cube-right-mini {
  transform: rotateY(90deg) translateZ(10px);
}

.cube-left-mini {
  transform: rotateY(-90deg) translateZ(10px);
}

.cube-top-mini {
  transform: rotateX(90deg) translateZ(10px);
}

.cube-bottom-mini {
  transform: rotateX(-90deg) translateZ(10px);
}
</style>