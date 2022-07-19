<template>
  <div>
    <apexchart
      ref="timeseriesChart"
      type="area"
      :options="chartOptions"
      :series="timeSeriesData"
    />
  </div>
</template>

<script>

export default {
  name: 'TimeSeriesChart',
  props: {
    timeSeriesData: {
      type: Object,
      required: true,
    },
    timeseriesKey: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      keyMap: {
        price_average: 'Average Price',
        price_low: 'Lowest Price',
        price_high: 'Highest Price',
        properties_added: 'Properties Added',
      },
      chartOptions: {
        chart: {
          type: 'line',
          stacked: false,
          height: 350,
          zoom: {
            type: 'x',
            enabled: true,
            autoScaleYaxis: true,
          },
          toolbar: {
            autoSelected: 'zoom',
          },
        },
        stroke: {
          curve: 'straight',
        },
        dataLabels: {
          enabled: false,
        },
        markers: {
          size: 0,
        },
        title: {
          text: 'Average Price (Timeseries)',
          align: 'left',
        },
        fill: {
          type: 'gradient',
          gradient: {
            shadeIntensity: 1,
            inverseColors: false,
            opacityFrom: 0.5,
            opacityTo: 0,
            stops: [0, 90, 100],
          },
        },
        yaxis: {
          labels: {
            formatter(val) {
              return `€${(val).toFixed(0)}`;
            },
          },
          title: {
            text: 'Price',
          },
        },
        xaxis: {
          type: 'datetime',
        },
        tooltip: {
          shared: false,
          y: {
          },
        },
      },
    };
  },
  computed: {
    getTitle() {
      return this.keyMap[this.timeseriesKey];
    },
  },
  watch: {
    timeseriesKey() {
      this.updateOptions();
    },
  },
  methods: {
    updateOptions() {
      this.$refs.timeseriesChart.updateOptions({
        title: {
          text: `${this.getTitle} (Timeseries)`,
        },
      });
    },
  },
};
</script>
