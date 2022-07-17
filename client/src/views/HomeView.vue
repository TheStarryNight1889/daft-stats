<template>
  <div>
    <div class="text-center py-8 bg-base-300">
      <p>
        <img
          src="logo_transparent.png"
          alt=""
          class="w-2/3 mr-auto ml-auto p-4"
        >
      </p>
      <p class="text-l px-2">
        A simple dashboard for viewing information about the Irish rental market.
      </p>
    </div>
    <div class="divider m-0 p-0 bg-base-300 px-4" />
    <Suspense>
      <HomeDashboard class="bg-white" />
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
