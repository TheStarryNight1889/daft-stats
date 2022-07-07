import { defineStore } from 'pinia';
import api from '../api';

const useStats = defineStore('stats', {
  state: () => ({
    stats: [],
  }),
  getters: {
    // TODO remove repeated code
    priceLowCurrent: (state) => state.stats[state.stats.length - 1]
      .price_low,
    priceHighCurrent: (state) => state.stats[state.stats.length - 1]
      .price_high,
    priceAverageCurrent: (state) => state.stats[state.stats.length - 1]
      .price_average,
    priceDistributionCurrent: (state) => state.stats[state.stats.length - 1]
      .price_distribution,
    propertiesAddedCurrent: (state) => state.stats[state.stats.length - 1]
      .properties_added,
    isAverageTrendUp: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return last.price_average > secondLast.price_average;
    },
    isLowestTrendUp: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return last.price_low > secondLast.price_low;
    },
    isHighestTrendUp: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return last.price_high > secondLast.price_high;
    },
    isPropertyAddedTrendUp: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return last.properties_added > secondLast.properties_added;
    },
    averageDifferencePercentage: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return (last.price_average - secondLast.price_average) / last.price_average;
    },
    lowestDifferencePercentage: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return (last.price_low - secondLast.price_low) / last.price_low;
    },
    highestDifferencePercentage: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return (last.price_high - secondLast.price_high) / last.price_high;
    },
    // TODO this will produce infinty if there are no properties
    propertiesAddedDifferencePercentage: (state) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return (last.properties_added - secondLast.properties_added) / last.properties_added;
    },

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
