import { defineStore } from 'pinia';
import api from '../api';

const useStats = defineStore('stats', {
  state: () => ({
    stats: [],
  }),
  getters: {
    priceLowCurrent: (state) => state.stats[state.stats.length - 1]
      .price_low,
    priceHighCurrent: (state) => state.stats[state.stats.length - 1]
      .price_high,
    priceAverageCurrent: (state) => state.stats[state.stats.length - 1]
      .price_average,
    priceDistributionCurrent: (state) => state.stats[state.stats.length - 1]
      .price_distribution,

  },

  actions: {
    async FetchAllStats() {
      const { data } = await api.GetAllStats();
      this.stats = data.stats;
      console.log(this.stats);
    },
  },
});

export default useStats;
