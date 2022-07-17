<template>
  <div>
    <div class="grid bg-base-300 grid-cols-2 py-4 px-2 divide-x-2">
      <h1 class=" col-span-2 text-xl font-bold">
        Quick Stats
      </h1>
      <div class="my-2">
        <h1 class="text-xl">
          Average Rent
          <font-awesome-icon
            v-if="isTrendUp('price_average')"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isTrendUp('price_average')"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ averageDifferencePercentage }}%
          {{ isTrendUp('price_average') ? 'higher' : 'lower' }}
          than yesterday
        </p>

        <p>
          €{{ averagePrice }}  per dwelling
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          Lowest Rent
          <font-awesome-icon
            v-if="isTrendUp('price_low')"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isTrendUp('price_low')"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ lowestDifferencePercentage }}%
          {{ isTrendUp('price_low') ? 'higher' : 'lower' }}
          than yesterday
        </p>
        <p>
          €{{ lowestPrice }}
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          Highest Rent
          <font-awesome-icon
            v-if="isTrendUp('price_high')"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isTrendUp('price_high')"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ highestDifferencePercentage }}%
          {{ isTrendUp('price_high') ? 'higher' : 'lower' }}
          than yesterday
        </p>
        <p>
          €{{ highestPrice }}
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          New Rentals
          <font-awesome-icon
            v-if="isTrendUp('properties_added')"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isTrendUp('properties_added')"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ propertiesAddedDifferencePercentage }}%
          {{ isTrendUp('properties_added') ? 'higher' : 'lower' }}
          than yesterday
        </p>
        <p>
          {{ propertiesAdded }}
        </p>
      </div>
    </div>
    <div class="py-2 px-2 my-2 rounded">
      <PriceDistributionChart :price-distribution="priceDistribution" />
    </div>
    <div class="divider m-0 p-0 bg-base-100 px-4" />
    <div class="py-2 px-2 my-2 rounded">
      <PriceAverageChart :prive-average-timseries="priceAverageTimeseries" />
    </div>
  </div>
</template>

<script>
import useStats from '../../stores/useStats';
import PriceDistributionChart from './PriceDistributionChart.vue';
import PriceAverageChart from './PriceAverageChart.vue';

export default {
  name: 'HomeDashboard',
  components: { PriceDistributionChart, PriceAverageChart },
  async setup() {
    const statsStore = useStats();
    await statsStore.FetchAllStats();
    return {
      statsStore,
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
    averageDifferencePercentage() {
      const diff = this.statsStore.averageDifferencePercentage;
      return (diff * 100).toFixed(2);
    },
    lowestDifferencePercentage() {
      const diff = this.statsStore.lowestDifferencePercentage;
      return (diff * 100).toFixed(2);
    },
    highestDifferencePercentage() {
      const diff = this.statsStore.highestDifferencePercentage;
      return (diff * 100).toFixed(2);
    },
    propertiesAddedDifferencePercentage() {
      const diff = this.statsStore.propertiesAddedDifferencePercentage;
      return (diff * 100).toFixed(2);
    },
    priceAverageTimeseries() {
      return this.statsStore.priceAverageTimeseries;
    },
  },
  methods: {
    isTrendUp(datapoint) {
      return this.statsStore.isTrendUp(datapoint);
    },
  },
};
</script>
