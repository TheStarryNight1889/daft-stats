<template>
  <div>
    <div class="grid bg-base-300 grid-cols-2 py-4 px-2 divide-x-2">
      <h1 class=" col-span-2 text-xl font-bold">
        Quick Stats
      </h1>
      <quick-stat
        :title="'Average Rent'"
        :trend-up="isTrendUp('price_average')"
        :difference-percentage="differencePercentage('price_average')"
        :value="`€${averagePrice}`"
      />
      <quick-stat
        :title="'Lowest Rent'"
        :trend-up="isTrendUp('price_low')"
        :difference-percentage="differencePercentage('price_low')"
        :value="`€${lowestPrice}`"
      />
      <quick-stat
        :title="'Highest Rent'"
        :trend-up="isTrendUp('price_high')"
        :difference-percentage="differencePercentage('price_high')"
        :value="`€${highestPrice}`"
      />
      <quick-stat
        :title="'New Rentals'"
        :trend-up="isTrendUp('properties_added')"
        :difference-percentage="differencePercentage('properties_added')"
        :value="`${propertiesAdded}`"
      />
    </div>
    <div class="py-2 px-2 my-2 rounded">
      <PriceDistributionChart :price-distribution="priceDistribution" />
    </div>
    <div class="divider m-0 p-0 bg-base-100 px-4" />
    <div class="py-2 px-2 my-2 rounded text-right">
      <select
        class="select select-bordered select-sm max-w-xs mb-4"
        @change="updateTimeseriesKey($event)"
      >
        <option
          selected
          value="price_average"
        >
          Price Average
        </option>
        <option value="price_high">
          Price High
        </option>
        <option value="price_low">
          Price Low
        </option>
        <option value="properties_added">
          Properties Added
        </option>
      </select>
      <TimeSeriesChart
        :time-series-data="timeseriesData"
        :timeseries-key="timeseriesKey"
      />
    </div>
  </div>
</template>

<script>
import useStats from '../../stores/useStats';
import PriceDistributionChart from './PriceDistributionChart.vue';
import TimeSeriesChart from './TimeSeriesChart.vue';
import QuickStat from './QuickStat.vue';

export default {
  name: 'HomeDashboard',
  components: { PriceDistributionChart, TimeSeriesChart, QuickStat },
  async setup() {
    const statsStore = useStats();
    await statsStore.FetchAllStats();
    return {
      statsStore,
    };
  },
  data() {
    return {
      timeseriesKey: 'price_average',
      timeseriesData: [{
        name: 'Average Price',
        data: [],
      }],
    };
  },
  computed: {
    priceDistribution() {
      return this.statsStore.priceDistributionCurrent;
    },
    averagePrice() {
      return Math.round(this.statsStore.priceAverageCurrent);
    },
    lowestPrice() {
      return this.statsStore.priceLowCurrent;
    },
    highestPrice() {
      return this.statsStore.priceHighCurrent;
    },
    propertiesAdded() {
      return this.statsStore.propertiesAddedCurrent;
    },
  },
  created() {
    this.timeseries();
  },
  methods: {
    isTrendUp(datapoint) {
      return this.statsStore.isTrendUp(datapoint);
    },
    differencePercentage(datapoint) {
      const diff = this.statsStore.differencePercentage(datapoint);
      return (diff * 100).toFixed(2);
    },
    updateTimeseriesKey(datapoint) {
      this.timeseriesKey = datapoint.target.value;
      this.timeseries();
    },
    timeseries() {
      this.timeseriesData[0].data = this.statsStore.timeseries(this.timeseriesKey);
    },
  },
};
</script>
