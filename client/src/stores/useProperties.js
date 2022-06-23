import { defineStore } from 'pinia';
import api from '../api';

const useProperties = defineStore('properties', {
  state: () => ({
    properties: [],
  }),
  getters: {
    averagePrice: (state) => {
      const total = state.properties.reduce((acc, property) => acc + property.price, 0);
      return total / state.properties.length;
    },
    lowestPrice: (state) => {
      // exclude less than 100
      const filtered = state.properties.filter((property) => property.price > 100);
      return Math.min(...filtered.map((property) => property.price));
    },
    highestPrice: (state) => {
      // exclude less than 100
      const filtered = state.properties.filter((property) => property.price > 100);
      return Math.max(...filtered.map((property) => property.price));
    },
  },
  actions: {
    async FetchAllProperties() {
      const { data } = await api.GetAllProperties();
      console.log(data);
      this.properties = data.properties;
    },
  },
});

export default useProperties;
