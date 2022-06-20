import { defineStore } from 'pinia';
import api from '@/api';

const useProperties = defineStore('properties', {
  state: () => ({
    properties: [],
  }),
  getters: {
    averagePrice: (state) => {
      const total = state.properties.reduce((acc, property) => acc + property.price, 0);
      return total / state.properties.length;
    },
  },
  actions: {
    async FetchAllProperties() {
      const { data } = await api.GetAllProperties();
      this.properties = data.properties;
    },
  },
});

export default useProperties;
