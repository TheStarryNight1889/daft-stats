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

    // 0 0-500
    // 1 500-1000
    // 2 1000-1500
    // 3 1500-2000
    // 4 2000-2500
    // 5 2500-3000
    // 6 3000-3500
    // 7 3500-4000
    // 8 4000-4500
    // 9 4500-5000
    // 10 5000+
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
