import Router from 'vue-router'
import Vue from 'vue'
Vue.use(Router)


function load(component) {
    return require(`./views/theme/${window.lambda.theme}/${component}`)
}
let routes = [
    {path: '/', redirect:'/login'},
    {path: '/login', component: load("auth/login.vue")},
    {path: '/forgot', component: load("auth/password/forgot.vue")},
    {path: '/password-reset', component: load("auth/password/password_reset.vue")},
];


export default new Router({
    mode: 'history', //hash,
    base: '/auth/',
    routes,
})
