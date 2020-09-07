<template>
    <div class="form-wrap">
        <div class="form-content">
            <h2>{{ lang.passwordReset }}</h2>
            <form v-on:submit.prevent="onSubmit"  method="post" class="login-form">
                <div class="form-element input">
                    <input type="text" id="code" v-model="credentials.code" name="code" autocomplete="off"
                           :placeholder="lang.passwordResetCode">
                    <span class="icon pass"></span>
                </div>
                <div class="form-element input">
                    <input type="text" id="email" v-model="credentials.email" name="email" autocomplete="off"
                           :placeholder="lang.email">
                    <span class="icon user"></span>
                </div>
                <div class="form-element input">
                    <input type="password" v-model="credentials.password" :disabled="loading"
                           :placeholder="lang.password">
                    <span class="icon pass"></span>
                </div>
                <div class="form-element input">
                    <input type="password" v-model="credentials.password_confirm" :disabled="loading"
                           :placeholder="lang.passwordConfirm">
                    <span class="icon pass"></span>
                </div>
                <div class="form-element ">
                    <button id="submit" class="button" :disabled="loading" style="width: 100%">
                        <span id="submitTxt">{{ lang.passwordReset }}</span>
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

                </div>
            </form>

            <div id="msg">
                <span v-if="isSuccess" class="success">{{ successMsg }}</span>
                <span v-if="isError" class="error">{{ errorMsg }}</span>
            </div>
        </div>

        <slot name="copyright"></slot>
    </div>
</template>

<script>


    export default {
        props:["selectedLang"],
        name: "password-reset",
        data() {
            return {
                loading: false,
                isSuccess: false,
                isError: false,
                errorMsg:'',
                successMsg:'',
                credentials: {
                    code: null,
                    email: null,
                    password: null,
                    password_confirm: null,
                },

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
            onSubmit() {
                this.isSuccess = false;
                this.isError = false;
                if (!this.loading) {
                    this.loading = true;
                    axios.post('/auth/password-reset', {...this.credentials, lang:this.selectedLang}).then(({data}) => {
                        setTimeout(() => {
                            this.loading = false;
                            if (data.status) {
                                this.isSuccess = true;
                                this.successMsg = data.msg
                                setTimeout(()=>{
                                    this.$router.push('login');
                                }, 2000)
                            } else {
                                this.isError = true;
                                this.errorMsg = 'И-мэйл илгээх үед алдаа гарлаа'
                            }
                        }, 1000);
                    }).catch(e => {

                        this.errorMsg = e.response.data.error;
                        setTimeout(() => {
                            this.loading = false;
                            this.isError = true;
                        }, 1000);
                    })
                }
            },

        }
    }
</script>