# 文章网站前端

## 项目简介

一个使用 Vue 3 构建的简约风格文章网站前端。用户可以注册、登录、浏览文章、查看文章详情、点赞文章以及发布新文章。

## 功能特点

- 用户认证 (登录/注册)
- JWT 认证
- 文章列表浏览
- 文章详情查看
- 文章点赞功能
- 发布新文章

## 技术栈

- Vue 3
- Vue Router 4
- Vuex 4
- Axios
- Vite

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

## 项目结构

```
frontend/
├── public/
│   └── index.html
├── src/
│   ├── App.vue                 # 主应用组件
│   ├── main.js                 # 入口文件
│   ├── router/
│   │   └── index.js            # 路由配置
│   ├── store/
│   │   └── auth.js             # 认证状态管理
│   ├── api/
│   │   └── api.js              # API请求封装
│   ├── components/
│   │   ├── Navbar.vue          # 导航栏组件
│   │   ├── ArticlePreview.vue  # 文章预览组件
│   │   └── ArticleForm.vue     # 文章发布表单
│   └── views/
│       ├── LoginView.vue       # 登录页面
│       ├── RegisterView.vue    # 注册页面
│       ├── HomeView.vue        # 主页(文章列表)
│       ├── ArticleDetailView.vue # 文章详情页
│       └── CreateArticleView.vue # 创建文章页面
└── package.json                # 项目依赖
```

## 后端 API

本前端应用与后端 API 一起工作，后端 API 端点包括：

- 用户认证: `/api/auth/login`, `/api/auth/register`
- 文章管理: `/api/articles`, `/api/articles/:id`
- 点赞功能: `/api/articles/:id/likes`
