<template>
    <div class="login" :style="{ backgroundImage: 'url(' + lambda.bg + ')' }">
        <div class="content">
            <div class="content-layer"></div>
            <div class="title">
                <h2 style="max-width: 600px;">{{ lang.title }}</h2>
                <p>{{ lang.subtitle }}</p>
            </div>
        </div>
        <div class="auth">
            <div class="lang-switcher" v-if="lambda.has_language">
                <a v-for="item in languages" :key="item.index"
                   :class="selectedLang == item.code ? 'active' : ''" href="javascript:void(0)"
                   @click="switchLanguage(item.code)">
                    {{ item.label }}
                </a>
            </div>
            <div class="form-wrap">
                <router-view :selectedLang="selectedLang"></router-view>
                <div class="copyright" style="width:70%; text-align:center;">{{ copyright }}</div>
            </div>
        </div>
    </div>
</template>

<script>

    export default {
        name: "aside",
        data() {
            return {
                loading: false,
                isSuccess: false,
                isError: false,
                credentials: {
                    login: null,
                    password: null
                },
                selectedLang: localStorage.getItem("lang") == null ? 'mn' : localStorage.getItem("lang"),
                languages:window.lambda.languages,
                copyright:window.lambda.copyright,
                lambda:window.lambda,
                styleObj: {
                    backgroundImage: window.lambda.bg + ' !important'
                }
            }
        },

        computed: {
            lang() {
                return window.lambda.static_words[this.selectedLang]
            }
        },

        methods: {
            switchLanguage(val) {
                this.selectedLang = val;
                localStorage.setItem("lang", val);

            }
        }
    }
</script>

<style lang="scss">
    @import "../../../scss/theme/aside/style";
</style>

