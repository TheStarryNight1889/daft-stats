<template>
  <div>
    <div class="grid bg-base-300 grid-cols-2 my-4 py-4 px-6 divide-x-2">
      <h1 class=" col-span-2 text-2xl font-bold">
        Quick Stats
      </h1>
      <div class="my-2">
        <h1 class="text-xl">
          Average Rent
          <font-awesome-icon
            class="text-red-400"
            icon="fa-solid fa-arrow-trend-down"
          />
        </h1>
        <p class="text-accent text-xs">
          3.5% lower than last month
        </p>
        <p>
          €{{ averagePrice }}  per dwelling
        </p>
      </div>
      <div class="my-2">
        <h1 class="text-xl">
          Lowest Rent
          <font-awesome-icon
            class="text-green-400"
            icon="fa-solid fa-arrow-trend-up"
          />
        </h1>
        <p class="text-accent text-xs">
          15% higher than last month
        </p>
        <p>
          €{{ lowestPrice }}
        </p>
      </div>
    </div>
    <div class="py-2 px-2 m-2 bg-base-300 rounded">
      <PriceDistributionChart :price-distribution="priceDistribution" />
    </div>
  </div>
</template>

<script>
import useProperties from '../../stores/useProperties';
import PriceDistributionChart from './PriceDistributionChart.vue';

export default {
  name: 'HomeDashboard',
  components: { PriceDistributionChart },
  setup() {
    const propertiesStore = useProperties();
    propertiesStore.FetchAllProperties();
    return {
      propertiesStore,
    };
  },
  computed: {
    priceDistribution() {
      return this.propertiesStore.priceDistributionIn500s;
    },
    averagePrice() {
      return Math.round(this.propertiesStore.averagePrice);
    },
    lowestPrice() {
      return this.propertiesStore.lowestPrice;
    },
  },
};
</script>
