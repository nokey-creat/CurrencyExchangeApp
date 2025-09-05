import axios from 'axios';

// 创建 axios 实例
const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  }
});

// 请求拦截器 - 添加认证令牌
api.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = token;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 认证相关 API
export const authAPI = {
  login: (username, password) => {
    return api.post('/auth/login', { username, password });
  },
  register: (username, password) => {
    return api.post('/auth/register', { username, password });
  }
};

// 文章相关 API
export const articleAPI = {
  getArticles: () => {
    return api.get('/articles');
  },
  getArticleById: (id) => {
    return api.get(`/articles/${id}`);
  },
  createArticle: (articleData) => {
    return api.post('/articles', articleData);
  },
  likeArticle: (id) => {
    return api.post(`/articles/${id}/likes`);
  },
  getArticleLikes: (id) => {
    return api.get(`/articles/${id}/likes`);
  }
};

export default api;
