import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import createPersistedstate from 'vuex-persistedstate'

Vue.use(Vuex)
const persistedstate = createPersistedstate({
  paths: ["list"]
})

export default new Vuex.Store({
  plugins: [persistedstate],
  state: {
    list: []
  },
  getters: {
  },
  mutations: {
    setList(state, payload){
      state.list = payload
    }
  },
  actions: {
    fetchNews(store){
      axios.get("https://newsapi.org/v2/everything?q=bitcoin&apiKey=a088ce87f18f40b89adfaa23fd3c7c4e").then((response)=>{
        store.commit("setList", response.data.articles)
      })
    }
  },
  modules: {
  }
})
