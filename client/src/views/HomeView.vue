<template>
  <div>
    <div class="navbar bg-base-300">
      <a class="btn btn-ghost normal-case text-xl">Lafted</a>
    </div>
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
