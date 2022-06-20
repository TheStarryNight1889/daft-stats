<template>
    <h1>HomeComponent</h1>
    {{propertiesStore.averagePrice}}
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
