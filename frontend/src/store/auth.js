import { authAPI } from '../api/api';

const state = {
  token: localStorage.getItem('token') || '',
  username: localStorage.getItem('username') || '',
  isAuthenticated: !!localStorage.getItem('token')
};

const getters = {
  isAuthenticated: state => state.isAuthenticated,
  username: state => state.username
};

const actions = {
  async login({ commit }, { username, password }) {
    try {
      const response = await authAPI.login(username, password);
      const token = response.data.token;
      
      // 存储认证信息
      localStorage.setItem('token', token);
      localStorage.setItem('username', username);
      
      commit('SET_AUTH', { token, username });
      return Promise.resolve(response);
    } catch (error) {
      return Promise.reject(error);
    }
  },
  
  async register({ commit }, { username, password }) {
    try {
      const response = await authAPI.register(username, password);
      const token = response.data.token;
      
      // 存储认证信息
      localStorage.setItem('token', token);
      localStorage.setItem('username', username);
      
      commit('SET_AUTH', { token, username });
      return Promise.resolve(response);
    } catch (error) {
      return Promise.reject(error);
    }
  },
  
  logout({ commit }) {
    // 清除认证信息
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    commit('LOGOUT');
  }
};

const mutations = {
  SET_AUTH(state, { token, username }) {
    state.token = token;
    state.username = username;
    state.isAuthenticated = true;
  },
  
  LOGOUT(state) {
    state.token = '';
    state.username = '';
    state.isAuthenticated = false;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
};
