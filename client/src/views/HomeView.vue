<template>
  <div>
    <div
      class="text-center py-16 bg-base-300 bg-contain bg-center"
    >
      <h1 class="text-4xl">
        Lafted
      </h1>
      <p class="text-l">
        A simple, dashboard for consuming information about the Irish rental market.
      </p>
    </div>
    <div class="divider m-0 p-0 bg-base-300 px-4" />
    <Suspense>
      <HomeDashboard />
    </Suspense>
  </div>
</template>

<script>
import HomeDashboard from '../components/Home/HomeDashboard.vue';

export default {
  name: 'HomeView',
  components: {
    HomeDashboard,
  },
};
</script>
