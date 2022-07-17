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
    isTrendUp: (state) => (datapoint) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return last[datapoint] > secondLast[datapoint];
    },
    differencePercentage: (state) => (datapoint) => {
      const last = state.stats[state.stats.length - 1];
      const secondLast = state.stats[state.stats.length - 2];
      return (last[datapoint] - secondLast[datapoint]) / last[datapoint];
    },

    priceAverageTimeseries: (state) => {
      // return an object of the form [{x:, y:}, {x:, y:}];
      const timeseries = [];
      state.stats.forEach((stat) => {
        timeseries.push({
          x: new Date(stat.timestamp.split('.')[0].replace(/-/g, '/')),
          y: stat.price_average,
        });
      });
      return timeseries;
    },

  },

  actions: {
    async FetchAllStats() {
      const { data } = await api.GetAllStats();
      this.stats = data.stats;
    },
  },
});

export default useStats;
