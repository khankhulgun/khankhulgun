<template>
    <div class="login"
         :style="{
         backgroundSize:'contain',
         backgroundRepeat:'no-repeat',
         backgroundPosition:'bottom' }">
<!--         :style="{ backgroundImage: 'url(' + lambda.bg + ')',-->
<!--         backgroundSize:'contain', -->
<!--         backgroundRepeat:'no-repeat', -->
<!--         backgroundPosition:'bottom' }"-->

        <div class="wrap">
            <div class="content">
                <div class="content-blur"></div>
                <div class="content-color-layer"></div>
                <div id="slideshow">
<!--                    <div class="title">ХӨДӨЛМӨР, НИЙГМИЙН ХАМГААЛЛЫН ЯАМ</div>-->
                    <div class="one">
                        <h2>
                            <img src="/catalog/images/bank.svg" class="bank" alt="">
                            <img src="/catalog/images/havt_logo.svg" alt="">
                        </h2>
<!--                        <p>{{ lang.subtitle }}?</p>-->
                        <p>ЗҮҮН БҮСИЙН БИЗНЕСҮҮДИЙН УЯЛДАА ХОЛБООГ САЙЖРУУЛАХ КАТАЛОГ СИСТЕМ</p>
                    </div>

                </div>
            </div>
            <div class="auth">
                <div class="auth-blur"></div>
                <div class="auth-color-layer"></div>
                <div class="lang-switcher" v-if="lambda.has_language">
                    <a v-for="item in languages" :key="item.index"
                       :class="selectedLang == item.value ? 'active' : ''" href="javascript:void(0)"
                       @click="switchLanguage(item.code)">
                        {{ item.label }}
                    </a>
                </div>
                <router-view :selectedLang="selectedLang">
                    <div class="copyright" style="text-align:center;" slot="copyright">
                        {{copyright}}
                    </div>
                </router-view>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "default",
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
                languages: window.lambda.languages,
                copyright: window.lambda.copyright,
                lambda: window.lambda,
            }
        },

        computed: {
            lang() {
                return window.lambda.static_words[this.selectedLang]
            }
        },

        created() {

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
    @import "../../../scss/theme/default/style";
</style>
