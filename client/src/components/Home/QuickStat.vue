<template>
  <div class="my-2">
    <h1 class="text-xl">
      {{ title }}
      <font-awesome-icon
        v-if="differencePercentage == 0.00"
        class="text-blue-400"
        icon="fa-solid fa-arrow-right-long"
      />
      <font-awesome-icon
        v-else-if="trendUp"
        class="text-green-400"
        icon="fa-solid fa-arrow-trend-up"
      />
      <font-awesome-icon
        v-else-if="!trendUp"
        class="text-red-400"
        icon="fa-solid fa-arrow-trend-down"
      />
    </h1>
    <p
      v-if="differencePercentage != 0.00"
      class="text-accent text-xs"
    >
      {{ differencePercentage }}%
      {{ trendUp ? 'higher' : 'lower' }}
      than yesterday
    </p>
    <p
      v-else
      class="text-accent text-xs"
    >
      No change since yesterday
    </p>
    <p>
      {{ value }}
    </p>
  </div>
</template>

<script>
export default {
  name: 'QuickStat',
  props: {
    title: {
      type: String,
      required: true,
    },
    trendUp: {
      type: Boolean,
      required: true,
    },
    differencePercentage: {
      type: Number,
      required: true,
    },
    value: {
      type: Number,
      required: true,
    },
  },
};
</script>
