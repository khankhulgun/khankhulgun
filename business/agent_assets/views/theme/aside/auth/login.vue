<template>
    <div class="form-content">

        <div class="logo with-text" v-if="lambda.logoText != ''">
            <img :src="lambda.logo" alt="">

            <span>{{ lambda.logoText }}</span>
        </div>
        <div class="logo" v-else>
            <img :src="lambda.logo" alt="">
        </div>
        <h2>{{ lang.loginTitle }}</h2>
        <form v-on:submit.prevent="onSubmit" id="authForm" method="post" class="login-form">
            <div class="form-element input">
                <input type="text" v-model="credentials.login" :disabled="loading"
                       :placeholder="lang.username">
                <span class="icon user"></span>
            </div>
            <div class="form-element input">
                <input type="password" v-model="credentials.password" :disabled="loading"
                       :placeholder="lang.password">
                <span class="icon pass"></span>
            </div>
            <div class="form-element">
                <input type="checkbox" class="checkbox" id="remember_me">
                <label for="remember_me">{{ lang.remember }}</label>
            </div>
            <div class="form-element action">
                <button id="submit" class="button" :disabled="loading">
                    <span id="submitTxt">{{ lang.login }}</span>
                    <span class="loader">
                                        <div class="sk-fading-circle" v-if="loading">
                                            <div class="sk-circle1 sk-circle"></div>
                                            <div class="sk-circle2 sk-circle"></div>
                                            <div class="sk-circle3 sk-circle"></div>
                                            <div class="sk-circle4 sk-circle"></div>
                                            <div class="sk-circle5 sk-circle"></div>
                                            <div class="sk-circle6 sk-circle"></div>
                                            <div class="sk-circle7 sk-circle"></div>
                                            <div class="sk-circle8 sk-circle"></div>
                                            <div class="sk-circle9 sk-circle"></div>
                                            <div class="sk-circle10 sk-circle"></div>
                                            <div class="sk-circle11 sk-circle"></div>
                                            <div class="sk-circle12 sk-circle"></div>
                                        </div>
                                    </span>
                </button>
                <p>
                    <router-link class="forgot" to="/forgot">{{ lang.forgot }}</router-link>
                </p>
            </div>
        </form>

        <div id="msg">
            <span v-if="isSuccess" class="success">{{ lang.loginSuccess }}</span>
            <span v-if="isError" class="error">{{ lang.loginError}}</span>
        </div>
    </div>
</template>

<script>


    export default {
        props: ['selectedLang'],
        name: "aside-login",
        data() {
            return {
                loading: false,
                isSuccess: false,
                isError: false,
                credentials: {
                    login: null,
                    password: null
                },
                lambda:window.lambda
            }
        },
        computed: {
            lang() {
                return window.lambda.static_words[this.selectedLang]
            }
        },
        methods: {
            onSubmit() {
                this.isSuccess = false;
                this.isError = false;
                if (!this.loading) {
                    this.loading = true;
                    axios.post('/auth/login', this.credentials).then(({data}) => {
                        setTimeout(() => {
                            this.loading = false;
                            if (data.status) {
                                this.isSuccess = true;
                                setTimeout(() => {
                                    window.location = data.path;
                                }, 600)
                            } else {
                                this.isError = true;
                            }
                        }, 1000);
                    }).catch(e => {
                        setTimeout(() => {
                            this.loading = false;
                            this.isError = true;
                        }, 1000);
                    })
                }
            },

            switchLanguage(val) {
                this.selectedLang = val;
                localStorage.setItem("lang", val);

            }
        }
    }
</script>
