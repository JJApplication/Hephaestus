<template>
  <div class="container mx-auto px-4 py-8">
    <!-- 页面标题 -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-green-400 mb-2">包管理</h1>
      <p class="text-gray-400">管理微服务安装包，查看版本信息和包大小</p>
    </div>

    <!-- 操作按钮区域 -->
    <div class="mb-6 flex flex-wrap gap-4">
      <UButton 
        @click="showUploadModal = true"
        color="primary"
        size="lg"
        icon="i-heroicons-arrow-up-tray"
      >
        上传服务包
      </UButton>
      <UButton 
        @click="startScan"
        :disabled="isScanning"
        :loading="isScanning"
        variant="outline"
        color="primary"
        size="lg"
        icon="i-heroicons-magnifying-glass"
      >
        {{ isScanning ? '扫描中...' : '手动扫描' }}
      </UButton>
    </div>

    <!-- 扫描进度条 -->
    <div v-if="isScanning" class="mb-6">
      <div class="bg-gray-800 rounded-lg p-4">
        <div class="flex items-center justify-between mb-2">
          <span class="text-green-400">扫描进度</span>
          <span class="text-green-300">{{ scanProgress }}%</span>
        </div>
        <div class="w-full bg-gray-700 rounded-full h-2">
          <div 
            class="bg-green-500 h-2 rounded-full transition-all duration-300"
            :style="{ width: scanProgress + '%' }"
          ></div>
        </div>
        <div class="mt-2 text-sm text-gray-400">{{ scanStatus }}</div>
      </div>
    </div>

    <!-- 微服务包列表 -->
    <div class="space-y-4">
      <div 
        v-for="service in services" 
        :key="service.name"
        class="bg-gray-800 rounded-lg border border-gray-700 overflow-hidden"
      >
        <!-- 微服务头部 -->
        <div 
          @click="toggleService(service.name)"
          class="p-4 cursor-pointer hover:bg-gray-750 transition-colors flex items-center justify-between"
        >
          <div class="flex items-center">
            <Icon 
              :name="service.expanded ? 'heroicons:chevron-down' : 'heroicons:chevron-right'" 
              class="w-5 h-5 text-green-400 mr-3"
            />
            <div>
              <h3 class="text-lg font-semibold text-green-400">{{ service.name }}</h3>
              <p class="text-sm text-gray-400">{{ service.packages.length }} 个版本</p>
            </div>
          </div>
          <div class="text-sm text-gray-400">
            最新版本: {{ service.latestVersion || 'N/A' }}
          </div>
        </div>

        <!-- 包版本列表 -->
        <div v-if="service.expanded" class="border-t border-gray-700">
          <div 
            v-for="pkg in service.packages" 
            :key="pkg.id"
            class="p-4 border-b border-gray-700 last:border-b-0 hover:bg-gray-750 transition-colors"
          >
            <div class="flex items-center justify-between">
              <div class="flex-1">
                <div class="flex items-center mb-2">
                  <span class="text-green-300 font-medium mr-4">{{ pkg.version }}</span>
                  <span class="text-xs bg-gray-700 text-gray-300 px-2 py-1 rounded">{{ pkg.size }}</span>
                  <span class="text-xs text-gray-400 ml-4">{{ pkg.uploadTime }}</span>
                </div>
                <div class="text-sm text-gray-400">
                  文件名: {{ pkg.filename }}
                </div>
              </div>
              <div class="flex items-center space-x-2">
                <button 
                  @click="downloadPackage(pkg)"
                  class="text-blue-400 hover:text-blue-300 p-2 rounded transition-colors"
                  title="下载"
                >
                  <Icon name="heroicons:arrow-down-tray" class="w-4 h-4" />
                </button>
                <button 
                  @click="deletePackage(pkg)"
                  class="text-red-400 hover:text-red-300 p-2 rounded transition-colors"
                  title="删除"
                >
                  <Icon name="heroicons:trash" class="w-4 h-4" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-if="services.length === 0" class="text-center py-12">
      <Icon name="heroicons:archive-box" class="w-16 h-16 text-gray-600 mx-auto mb-4" />
      <h3 class="text-lg font-medium text-gray-400 mb-2">暂无服务包</h3>
      <p class="text-gray-500">点击上传按钮添加第一个服务包</p>
    </div>

    <!-- 上传模态框 -->
    <UModal v-model="showUploadModal">
      <UCard :ui="{ ring: '', divide: 'divide-y divide-gray-100 dark:divide-gray-800' }">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-semibold text-green-400">上传服务包</h3>
            <UButton color="gray" variant="ghost" icon="i-heroicons-x-mark-20-solid" class="-my-1" @click="showUploadModal = false" />
          </div>
        </template>
        
        <form @submit.prevent="uploadPackage" class="space-y-4">
          <UFormGroup label="微服务名称" class="mb-4">
            <UInput 
              v-model="uploadForm.serviceName"
              placeholder="输入微服务名称"
              size="lg"
              required
            />
          </UFormGroup>
          
          <UFormGroup label="版本号" class="mb-4">
            <UInput 
              v-model="uploadForm.version"
              placeholder="例如: 1.0.0"
              size="lg"
              required
            />
          </UFormGroup>
          
          <div class="mb-6">
            <UFormGroup label="选择文件" class="mb-4">
              <div class="flex items-center space-x-3">
                <input 
                  ref="fileInput"
                  type="file" 
                  @change="handleFileSelect"
                  accept=".jar,.war,.zip,.tar.gz"
                  class="hidden"
                />
                <UButton 
                  @click="triggerFileSelect"
                  variant="outline"
                  color="primary"
                  icon="i-heroicons-document-arrow-up"
                  size="lg"
                >
                  {{ uploadForm.file ? uploadForm.file.name : '选择文件' }}
                </UButton>
                <span v-if="uploadForm.file" class="text-sm text-gray-400">
                  {{ (uploadForm.file.size / 1024 / 1024).toFixed(1) }} MB
                </span>
              </div>
              <p class="text-xs text-gray-500 mt-2">支持 .jar, .war, .zip, .tar.gz 格式</p>
            </UFormGroup>
          </div>
          
        </form>
        
        <template #footer>
          <div class="flex justify-end space-x-3">
            <UButton 
              variant="ghost"
              color="gray"
              @click="showUploadModal = false"
              size="lg"
            >
              取消
            </UButton>
            <UButton 
              type="submit"
              :loading="uploading"
              :disabled="uploading"
              color="primary"
              size="lg"
              icon="i-heroicons-arrow-up-tray"
              @click="uploadPackage"
            >
              {{ uploading ? '上传中...' : '上传' }}
            </UButton>
          </div>
        </template>
      </UCard>
    </UModal>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

// 响应式数据
const services = ref([])
const showUploadModal = ref(false)
const isScanning = ref(false)
const scanProgress = ref(0)
const scanStatus = ref('')
const uploading = ref(false)

const uploadForm = ref({
  serviceName: '',
  version: '',
  file: null
})

// 模拟数据
const mockServices = [
  {
    name: 'user-service',
    expanded: false,
    latestVersion: '2.1.0',
    packages: [
      {
        id: 1,
        version: '2.1.0',
        filename: 'user-service-2.1.0.jar',
        size: '15.2 MB',
        uploadTime: '2024-01-15 14:30:00'
      },
      {
        id: 2,
        version: '2.0.5',
        filename: 'user-service-2.0.5.jar',
        size: '14.8 MB',
        uploadTime: '2024-01-10 09:15:00'
      }
    ]
  },
  {
    name: 'order-service',
    expanded: false,
    latestVersion: '1.5.2',
    packages: [
      {
        id: 3,
        version: '1.5.2',
        filename: 'order-service-1.5.2.jar',
        size: '22.1 MB',
        uploadTime: '2024-01-12 16:45:00'
      }
    ]
  }
]

// 方法
const toggleService = (serviceName) => {
  const service = services.value.find(s => s.name === serviceName)
  if (service) {
    service.expanded = !service.expanded
  }
}

const triggerFileSelect = () => {
  const fileInput = document.querySelector('input[type="file"]')
  if (fileInput) {
    fileInput.click()
  }
}

const startScan = async () => {
  isScanning.value = true
  scanProgress.value = 0
  scanStatus.value = '开始扫描...'
  
  // 模拟扫描进度
  const steps = [
    { progress: 20, status: '扫描本地目录...' },
    { progress: 40, status: '分析包信息...' },
    { progress: 60, status: '检查版本号...' },
    { progress: 80, status: '更新数据库...' },
    { progress: 100, status: '扫描完成' }
  ]
  
  for (const step of steps) {
    await new Promise(resolve => setTimeout(resolve, 1000))
    scanProgress.value = step.progress
    scanStatus.value = step.status
  }
  
  setTimeout(() => {
    isScanning.value = false
    // 这里可以重新加载数据
  }, 500)
}

const handleFileSelect = (event) => {
  uploadForm.value.file = event.target.files[0]
}

const uploadPackage = async () => {
  uploading.value = true
  
  try {
    // 模拟上传
    await new Promise(resolve => setTimeout(resolve, 2000))
    
    // 添加到列表
    const serviceName = uploadForm.value.serviceName
    let service = services.value.find(s => s.name === serviceName)
    
    if (!service) {
      service = {
        name: serviceName,
        expanded: true,
        latestVersion: uploadForm.value.version,
        packages: []
      }
      services.value.push(service)
    }
    
    const newPackage = {
      id: Date.now(),
      version: uploadForm.value.version,
      filename: uploadForm.value.file.name,
      size: (uploadForm.value.file.size / 1024 / 1024).toFixed(1) + ' MB',
      uploadTime: new Date().toLocaleString('zh-CN')
    }
    
    service.packages.unshift(newPackage)
    service.latestVersion = uploadForm.value.version
    
    // 重置表单
    uploadForm.value = {
      serviceName: '',
      version: '',
      file: null
    }
    showUploadModal.value = false
    
  } catch (error) {
    console.error('上传失败:', error)
  } finally {
    uploading.value = false
  }
}

const downloadPackage = (pkg) => {
  // 模拟下载
  console.log('下载包:', pkg.filename)
}

const deletePackage = (pkg) => {
  if (confirm(`确定要删除 ${pkg.filename} 吗？`)) {
    services.value.forEach(service => {
      const index = service.packages.findIndex(p => p.id === pkg.id)
      if (index > -1) {
        service.packages.splice(index, 1)
        // 更新最新版本
        if (service.packages.length > 0) {
          service.latestVersion = service.packages[0].version
        } else {
          service.latestVersion = null
        }
      }
    })
    
    // 移除空的服务
    services.value = services.value.filter(service => service.packages.length > 0)
  }
}

// 生命周期
onMounted(() => {
  // 加载模拟数据
  services.value = mockServices
})
</script>

<style scoped>
.container {
  max-width: 1200px;
}
</style>