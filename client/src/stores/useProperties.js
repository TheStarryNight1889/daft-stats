import { defineStore } from 'pinia';
import api from '../api';

const useProperties = defineStore('properties', {
  state: () => ({
    properties: [],
  }),
  getters: {
  },
  actions: {
    async FetchAllProperties() {
      const { data } = await api.GetAllProperties();
      this.properties = data.properties;
    },
  },
});

export default useProperties;
