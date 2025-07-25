# Hephaestus - 微服务部署平台

一个现代化的微服务部署和管理平台，采用暗黑极客绿色主题设计。

## 技术栈

- **前端框架**: Vue 3 + Nuxt 3
- **UI 组件库**: Nuxt UI
- **构建工具**: Vite
- **样式框架**: Tailwind CSS
- **图标**: Heroicons

## 功能特性

### 🏠 首页
- 微服务集群架构可视化
- 3D立方体微服务节点动画
- 数据流动画效果
- 实时系统监控指标
- 网格背景和发光效果

### 🔧 微服务列表
- 服务状态监控
- 资源使用情况展示
- 服务操作管理
- 搜索和筛选功能

### 📋 版本变更管理
- 版本发布时间线
- 变更历史追踪
- 部署状态管理
- 回滚功能支持

### 🚀 部署管理
- 多环境部署支持
- CI/CD 流水线配置
- 部署进度实时监控
- 自动化部署策略

### 📊 日志管理
- 实时日志流监控
- 多级别日志筛选
- 日志搜索和导出
- 可视化统计分析

## 设计特色

### 🎨 视觉设计
- **暗黑主题**: 深色背景配合极客绿色
- **发光效果**: 边框和文字的绿色发光
- **网格背景**: 科技感网格图案
- **动画效果**: 流畅的过渡和悬浮动画

### 🎯 交互体验
- **响应式设计**: 适配各种屏幕尺寸
- **固定导航**: 顶部菜单固定不滚动
- **实时更新**: 数据自动刷新
- **直观操作**: 简洁明了的操作界面

### 🔮 3D 效果
- **立方体微服务**: 3D旋转的服务节点
- **浮动动画**: 上下浮动效果
- **阴影效果**: 立体感阴影
- **数据流**: 箭头流动动画

## 快速开始

### 安装依赖
```bash
npm install
```

### 启动开发服务器
```bash
npm run dev
```

### 构建生产版本
```bash
npm run build
```

### 预览生产版本
```bash
npm run preview
```

## 项目结构

```
ui/
├── assets/
│   └── css/
│       └── main.css          # 主样式文件
├── layouts/
│   └── default.vue           # 默认布局
├── pages/
│   ├── index.vue            # 首页
│   ├── services/
│   │   └── index.vue        # 微服务列表
│   ├── versions/
│   │   └── index.vue        # 版本变更
│   ├── deploy/
│   │   └── index.vue        # 部署管理
│   └── logs/
│       └── index.vue        # 日志管理
├── app.vue                  # 应用根组件
├── nuxt.config.ts          # Nuxt 配置
├── tailwind.config.js      # Tailwind 配置
└── package.json            # 项目依赖
```

## 自定义配置

### 颜色主题
在 `assets/css/main.css` 中修改 CSS 变量：
```css
:root {
  --primary-green: #00ff41;
  --secondary-green: #00cc33;
  --dark-green: #003d0f;
  --bg-dark: #0a0a0a;
  --bg-darker: #050505;
}
```

### 动画效果
在 `tailwind.config.js` 中自定义动画：
```javascript
animation: {
  'float': 'float 3s ease-in-out infinite',
  'pulse': 'pulse 2s ease-in-out infinite',
  'data-flow': 'dataFlow 2s linear infinite',
}
```

## 浏览器支持

- Chrome >= 88
- Firefox >= 85
- Safari >= 14
- Edge >= 88

## 开发说明

### 添加新页面
1. 在 `pages/` 目录下创建 Vue 文件
2. 文件名即为路由路径
3. 自动生成路由配置

### 修改导航菜单
在 `layouts/default.vue` 中修改导航链接

### 添加新组件
在 `components/` 目录下创建可复用组件

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

---

**Hephaestus** - 让微服务部署变得简单而优雅 🚀