<template>
  <apexchart
    type="area"
    :options="chartOptions"
    :series="series"
  />
</template>

<script>

export default {
  name: 'PriceAverageChart',
  props: {
    priveAverageTimseries: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      series: [{
        name: 'Average Price',
        data: this.priveAverageTimseries,
      }],
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
          text: 'Rent Average (Timeseries)',
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
};
</script>
