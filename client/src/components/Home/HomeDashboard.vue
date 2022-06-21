<template>
  <div class="grid grid-cols-2 stats shadow">
    <div>
              <div class="stat">
          <div class="stat-title">Total Page Views</div>
          <div class="stat-value">89,400</div>
          <div class="stat-desc">21% more than last month</div>
        </div>
    </div>
    <div>
              <div class="stat">
          <div class="stat-title">Total Page Views</div>
          <div class="stat-value">89,400</div>
          <div class="stat-desc">21% more than last month</div>
        </div>
    </div>
    <div>
    </div>
  </div>
</template>

<script>
import useProperties from '@/stores/useProperties';

export default {
  name: 'HomeDashboard',
  setup() {
    const propertiesStore = useProperties();
    propertiesStore.FetchAllProperties();
    return {
      propertiesStore,
    };
  },

};
</script>
