<template>
  <apexchart
    type="bar"
    title="Rent"
    :options="chartOptions"
    :series="series"
  />
</template>

<script>

export default {
  name: 'PriceDistributionChart',
  props: {
    priceDistribution: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      chartOptions: {
        chart: {
          id: 'vuechart-example',
        },
        xaxis: {
          categories: [
            '€0-€500',
            '€500-€1,000',
            '€1,000-€1,500',
            '€1,500-€2,000',
            '€2,000-€2,500',
            '€2,500-€3,000',
            '€3,000-€3,500',
            '€3,500-€4,000',
            '€4,000-€4,500',
            '€4,500-€5,000',
            '€5,000+',
          ],
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: '100%',
              },
            },
          },
        ],
      },
      series: [
        {
          name: 'series-1',
          data: this.priceDistribution,
        },
      ],
    };
  },
};
</script>
