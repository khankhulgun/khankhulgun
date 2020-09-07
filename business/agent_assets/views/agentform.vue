<template>
    <section class="page">
        <paper-header class="mini">
            <div slot="right">
                <slot name="user-control"></slot>
            </div>

            <div slot="nav">
                <ul>
                    <li v-if="type == 'profile'">
                        <router-link to="">
                            <i class="tu-user"></i>
                            <span>Хувийн мэдээлэл</span>
                        </router-link>
                    </li>
                    <li v-if="type == 'password'">
                        <router-link to="">
                            <i class="tu-user"></i>
                            <span>Нууц үг солих</span>
                        </router-link>
                    </li>
                </ul>
            </div>
            <div slot="tool">
            </div>
        </paper-header>
        <section class="page-agent-form">
            <dataform v-if="type == 'profile'" class="material-form" ref="agentForm" schemaID="user_profile" :editMode="editMode" :onSuccess="onSuccess"/>
            <dataform v-if="type == 'password'" class="material-form" ref="agentForm" schemaID="user_password" :editMode="editMode" :onSuccess="onSuccess"/>
        </section>
    </section>
</template>

<script>
    import pagination from "./pagination"

    export default {
        props:['type'],
        components: {
            'dv-pagination': pagination
        },
        data() {
            return {
                editMode: true,
            }
        },

        mounted() {
            this.editUser(this.$user.id);

        },

        methods: {
            onSuccess(data) {

            },

            editUser(id) {
                this.$refs.agentForm.editModel(id);
            },

            showDefaultAvatar(e) {
                e.target.src = "/images/avatar/no_avatar.svg";
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/agent.scss";
</style>
