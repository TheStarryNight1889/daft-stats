import { createApp } from 'vue';
import { createPinia } from 'pinia';
import './index.css';
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faArrowTrendDown, faArrowTrendUp } from '@fortawesome/free-solid-svg-icons';
import VueApexCharts from 'vue3-apexcharts';

import App from './App.vue';
import router from './router';

library.add(faArrowTrendDown, faArrowTrendUp);

/* add icons to the library */

const pinia = createPinia();
const app = createApp(App).component('font-awesome-icon', FontAwesomeIcon);
app.use(router);
app.use(pinia);
app.use(VueApexCharts);

app.mount('#app');
