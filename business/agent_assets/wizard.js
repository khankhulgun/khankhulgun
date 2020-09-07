import Vue from 'vue'
import iView from 'iview';
import axios from 'axios';
import locale from 'iview/src/locale/lang/mn-MN';
import Wizard from './pages/wizard/wizard.vue';

window.Vue = Vue;
window.axios = axios;
window.axios.defaults.headers.common = {
    'X-Requested-With': 'XMLHttpRequest',
    'X-CSRF-TOKEN' : document.querySelector('meta[name="csrf-token"]').getAttribute('content')
};
Vue.config.productionTip = false;
Vue.use(iView, { locale });

new Vue({
    el: '#wizard',
    extends: Wizard,
});
