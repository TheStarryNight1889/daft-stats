<template>
  <div>
    <apexchart
      ref="realtimeChart"
      type="area"
      :options="chartOptions"
      :series="timeSeriesData"
      :title="{text: timeseriesKey, align: 'left'}"
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
          text: `${this.timeseriesKey} (Timeseries)`,
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
              // if timeseriesKey contains 'price'
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
};
</script>
