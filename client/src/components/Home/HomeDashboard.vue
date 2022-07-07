<template>
  <div>
    <div class="grid bg-base-300 grid-cols-2 my-4 py-4 px-2 divide-x-2">
      <h1 class=" col-span-2 text-2xl font-bold">
        Quick Stats
      </h1>
      <div class="my-2">
        <h1 class="text-xl">
          Average Rent
          <font-awesome-icon
            v-if="isAverageTrendUp"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isAverageTrendUp"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ averageDifferencePercentage }}%
          {{ isAverageTrendUp ? 'higher' : 'lower' }}
          than last month
        </p>

        <p>
          €{{ averagePrice }}  per dwelling
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          Lowest Rent
          <font-awesome-icon
            v-if="isLowestTrendUp"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isLowestTrendUp"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ lowestDifferencePercentage }}%
          {{ isLowestTrendUp ? 'higher' : 'lower' }}
          than last month
        </p>
        <p>
          €{{ lowestPrice }}
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          Highest Rent
          <font-awesome-icon
            v-if="isHighestTrendUp"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isHighestTrendUp"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ highestDifferencePercentage }}%
          {{ isHighestTrendUp ? 'higher' : 'lower' }}
          than last month
        </p>
        <p>
          €{{ highestPrice }}
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          New Rentals
          <font-awesome-icon
            v-if="isPropertiesAddedTrendUp"
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
          <font-awesome-icon
            v-if="!isPropertiesAddedTrendUp"
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          {{ propertiesAddedDifferencePercentage }}%
          {{ isPropertiesAddedTrendUp ? 'higher' : 'lower' }}
          than last month
        </p>
        <p>
          {{ propertiesAdded }}
        </p>
      </div>
    </div>
    <div class="py-2 px-2 rounded">
      <h1>Price Distribution (Today)</h1>
      <PriceDistributionChart :price-distribution="priceDistribution" />
    </div>
  </div>
</template>

<script>
import useStats from '../../stores/useStats';
import PriceDistributionChart from './PriceDistributionChart.vue';

export default {
  name: 'HomeDashboard',
  components: { PriceDistributionChart },
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
    isAverageTrendUp() {
      return this.statsStore.isAverageTrendUp;
    },
    isLowestTrendUp() {
      return this.statsStore.isLowestTrendUp;
    },
    isHighestTrendUp() {
      return this.statsStore.isHighestTrendUp;
    },
    isPropertiesAddedTrendUp() {
      return this.statsStore.isPropertiesAddedTrendUp;
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
  },
};
</script>
