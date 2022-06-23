<template>
  <div>
    <div class="navbar bg-base-100">
      <a class="btn btn-ghost normal-case text-xl">DaftStats</a>
    </div>
    <HomeDashboard />
  </div>
</template>

<script>
import HomeDashboard from "../components/Home/HomeDashboard.vue";

export default {
  name: "HomeView",
  components: {
    HomeDashboard,
  },
};
</script>
