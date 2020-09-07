import Vue from 'vue'
import axios from 'axios'
import router from './auth_router'
window.Vue = Vue;
window.axios = axios;

window.axios.defaults.headers.common = {
    'X-Requested-With': 'XMLHttpRequest',
    'X-CSRF-TOKEN': document.querySelector('meta[name="csrf-token"]').getAttribute('content')
};
Vue.config.productionTip = false;


new Vue({
    el: '#app',
    router,
    render: h => h(require(`./views/theme/${window.lambda.theme}/index`))
});
