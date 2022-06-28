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
      const filtered = state.properties.filter((property) => property.price > 100);
      return Math.min(...filtered.map((property) => property.price));
    },
    highestPrice: (state) => {
      const filtered = state.properties.filter((property) => property.price > 100);
      return Math.max(...filtered.map((property) => property.price));
    },
    priceDistributionIn500s: (state) => {
      const distribution = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
      state.properties.forEach((property) => {
        const index = Math.floor(property.price / 500);
        if (index <= 9) {
          distribution[index] += 1;
        } else {
          distribution[10] += 1;
        }
      });
      console.log(distribution);
      return distribution;
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
