<template>
    <section class="page">
        <section class="agent-header">
            <Row :gutter="16" class="tbl-header">
                <Col span="12" class-name="tbl-header-left">
                    <div class="tbl-header-title">
                        <span v-if="!isDeleted">Хэрэглэгчдийн жагсаалт</span>
                        <span v-else>Устгасан хэрэглэгчдийн жагсаалт</span>
                    </div>
                    <div class="tbl-header-count">
                        <a @click="isDeleted = false" :class="isDeleted ? '': 'active'"><i
                            class="mdi mdi-account-group mdi-18px"></i> Нийт ажилчид: <span>{{ users.total }}</span></a>
                        <a @click="isDeleted = true" :class="isDeleted ? 'active': ''"><i
                            class="mdi mdi-account-remove mdi-18px"></i> Устгасан: <span>{{ deletedUsers.total }}</span></a>
                    </div>
                </Col>

                <Col span="12" class-name="tbl-header-right">
                    <div class="tbl-search">
                        <form v-on:submit.prevent="searchUser">
                            <input placeholder="мэдээллээс хайх" v-model="q" class="tbl-search-input"/>
                        </form>
                        <i class="ti ti-search"></i>
                    </div>

                    <Button class="agent-add-btn" type="success" @click="showForm = true">
                        <i class="ti-plus"></i> Хэрэглэгч нэмэх
                    </Button>

                    <slot name="user-control"></slot>
                </Col>
            </Row>
        </section>

        <section class="page-agent">
            <template v-if="isDeleted">
                <template v-if="deletedUsers.total == 0">
                    <div class="no-user-data">
                        <i class="mdi mdi-account-off mdi-48px"></i>
                        <p>Устгасан ажилтан алга байна</p>
                    </div>
                </template>
            </template>

            <template v-else>
                <template v-if="users.total == 0">
                    <div class="no-user-data">
                        <i class="mdi mdi-account-off mdi-48px"></i>
                        <p>Хэрэглэгч алга байна</p>
                    </div>
                </template>
            </template>

            <dv-pagination :pagination="paginateData" :is-deleted="isDeleted" :deleted-model="deletedUsers"
                           :model="users" :query="query" :roles="roles" :is-top="false" :layout="layout"></dv-pagination>

            <Row :gutter="16" class="user-grid-wrapper">
                <Col v-for="user in isDeleted ? deletedUsers.data : users.data" :key="user.id" :xs="24" :sm="12" :md="12"
                     :lg="8">
                    <div class="user-grid">
                        <div class="user-head">
                            <div class="user-avatar">
                                <template
                                    v-if="user.profile != null && user.profile.avatar != null && user.profile.avatar != ''">
                                    <img @error="showDefaultAvatar" :src="user.profile.avatar" alt="" width="80"
                                         class="user-avatar-img">
                                </template>
                                <template v-else>
                                    <img src="/images/avatar/no_avatar.svg" alt="" width="80"
                                         class="user-avatar-img">
                                </template>
                            </div>
                            <div class="user-info">
                                <h3>{{ user.login }}</h3>
                                <span>{{ user.display }}</span>
                            </div>
                            <div class="user-action">
                                <template v-if="isDeleted">
                                    <div class="deleted-user-grid">
                                        <Tooltip content="Сэргээх" placement="top">
                                            <Poptip confirm title="Ажилтныг сэргээхдээ итгэлтэй байна уу?"
                                                    cancelText="Үгүй" okText="Тийм" transfer
                                                    @on-ok="restoreUser(user.id)"
                                                    @on-cancel="">
                                                <a><i class="ti-reload"></i></a>
                                            </Poptip>
                                        </Tooltip>
                                        <Tooltip content="Бүр мөсөн устгах" placement="top">
                                            <Poptip confirm title="Ажилтныг бүр мөсөн устгахдаа итгэлтэй байна уу?"
                                                    cancelText="Үгүй" okText="Тийм" transfer
                                                    @on-ok="deleteUserComplete(user.id)"
                                                    @on-cancel="">
                                                <a><i class="ti-trash"></i></a>
                                            </Poptip>
                                        </Tooltip>
                                    </div>
                                </template>

                                <template v-else>
                                    <Tooltip content="Засах" placement="top">
                                        <a @click="editUser(user.id)"><i class="ti-pencil"></i></a>
                                    </Tooltip>

                                    <Poptip confirm title="Ажилтныг устгахдаа итгэлтэй байна уу?"
                                            cancelText="Үгүй"
                                            okText="Тийм" transfer @on-ok="deleteUser(user.id)"
                                            @on-cancel="">
                                        <a><i class="ti-trash"></i></a>
                                    </Poptip>
                                </template>
                            </div>
                            <div class="user-status">
                                <div v-if="user.status == 0" class="false">
                                    <i class="ti-na"></i>
                                    <span>Бүртгэл баталгаажаагүй</span>
                                </div>

                                <div v-else class="user-status">
                                    <i class="ti-check"></i>
                                    <span>Бүртгэл баталгаажсан</span>
                                </div>
                            </div>
                        </div>

                        <div class="user-content">
                            <ul class="user-content-list">
                                <li>
                                    <div class="user-content-list-content">
                                        <span class="user-content-list-heading">Овог нэр</span>
                                        <span class="user-content-list-data">
                                            <template
                                                v-if="user.first_name != null && user.first_name !=''">{{ user.first_name}}</template>
                                            <template
                                                v-if="user.last_name != null && user.last_name !=''"> {{user.last_name}}</template>
                                            <template v-else><span class="user-no-data">Мэдээлэл алга</span></template>
                                        </span>
                                    </div>
                                </li>
                                <li>
                                    <div class="user-content-list-content">
                                        <span class="user-content-list-heading">Цахим шуудан</span>
                                        <span class="user-content-list-data">
                                            <template
                                                v-if="user.email != null && user.email != ''">{{ user.email}}</template>
                                            <template v-else><span class="user-no-data">Мэдээлэл алга</span></template>
                                        </span>
                                    </div>
                                </li>
                                <li>
                                    <div class="user-content-list-content">
                                        <span class="user-content-list-heading">Гар утас</span>
                                        <span class="user-content-list-data">
                                            <template
                                                v-if="user.phone != null && user.phone != ''">{{ user.phone}}</template>
                                            <template v-else><span class="user-no-data">Мэдээлэл алга</span></template>
                                        </span>
                                    </div>
                                </li>
                                <li>
                                    <div class="user-content-list-content">
                                        <span class="user-content-list-heading">{{ !isDeleted ? 'Үүсгэсэн' : 'Устгсан'}}  огноо</span>
                                        <span class="user-content-list-data">{{ !isDeleted ? setMoment(user.created_at) : user.deleted_at}}</span>
                                    </div>
                                </li>
                            </ul>
                        </div>
                    </div>
                </Col>
            </Row>
        </section>

        <Drawer width="700" class="agent-form" :closable="false" v-model="showForm">
            <dataform ref="agentForm" schemaID="user_form" :editMode="editMode" :onSuccess="onSuccess"/>
        </Drawer>
    </section>
</template>

<script>
    import moment from 'moment';
    import pagination from "./pagination"

    export default {
        components: {
            'dv-pagination': pagination
        },
        data() {
            return {
                showForm: false,
                isDeleted: false,
                loading: false,
                editMode: false,
                q: '',
                users: {
                    total: 0
                },
                deletedUsers: {
                    total: 0
                },
                paginateData: {
                    name: 'хэрэглэгчдээс',
                },
                query: {
                    column: 'id',
                    direction: 'asc',
                    page: 1,
                    per_page: 16,
                    role: 'all'
                },
                roles: []
            }
        },
        created() {
            this.getRoles();
            this.fetchData();
            this.fetchDeletedData();
        },

        methods: {
            setMoment(date) {
                return moment(date).format('YYYY-MM-DD')
            },
            fetchData() {
                axios.get(`/agent/users?role=${this.query.role}&page=${this.query.page}&sort=${this.query.column}&direction=${this.query.direction}`).then(({data}) => {
                    this.users = data;
                })
            },
            fetchDeletedData() {
                axios.get(`/agent/users/deleted?role=${this.query.role}&page=${this.query.page}&sort=${this.query.column}&direction=${this.query.direction}`).then(({data}) => {
                    this.deletedUsers = data;
                })
            },

            onSuccess(data) {
                if (this.editMode) {
                    this.users.data = this.users.data.map(item => {
                        if (item.id == data.id) {
                            item = data;
                        }
                        return item;
                    });
                    this.showForm = false;
                } else {
                    this.users.data.push(data);
                }
            },

            sortData(sort) {
                this.query.direction = sort.order;
                this.query.column = sort.key;
                this.fetchData();
                this.fetchDeletedData();
            },

            editUser(id) {
                this.editMode = true;
                this.showForm = true;

                this.$refs.agentForm.editModel(id);
            },

            deleteUser(id) {
                axios.get('/agent/delete/' + id).then(o => {
                    if (o.status) {
                        this.$Message.success('Хэрэглэгч устгагдлаа')

                        let deletedUser = this.users.data.find(item => item.id == id);
                        this.deletedUsers.data.push(deletedUser);

                        this.users.data = this.users.data.filter(item => item.id != id);

                        this.deletedUsers.total++;
                        this.users.total--;
                    } else {
                        this.$Message.error('Устгах үед алдаа гарлаа');
                    }
                })
            },

            deleteUserComplete(id) {
                axios.get('/agent/delete/complete/' + id).then(o => {
                    if (o.status) {
                        this.$Message.success('Хэрэглэгч устгагдлаа')
                        this.deletedUsers.data = this.deletedUsers.data.filter(item => item.id != id);
                        this.deletedUsers.total--;
                    } else {
                        this.$Message.error('Устгах үед алдаа гарлаа');
                    }
                })
            },

            restoreUser(id) {
                axios.get('/agent/restore/' + id)
                    .then(o => {
                        if (o.status) {
                            this.$Message.success('Хэрэглэгчийн мэдээлэл сэргээгдлээ');
                            let restoredUser = this.deletedUsers.data.find(item => item.id == id);
                            this.users.data.push(restoredUser);
                            this.deletedUsers.data = this.deletedUsers.data.filter(item => item.id != id);

                            this.deletedUsers.total--;
                            this.users.total++;
                        } else {
                            this.$Message.error("Мэдээлэл сэргээхэд алдаа гарлаа!");
                        }
                    })
                    .catch(errors => {
                        this.$Message.error("Мэдээлэл сэргээхэд алдаа гарлаа!");
                    });
            },

            showDefaultAvatar(e) {
                e.target.src = "/images/avatar/no_avatar.svg";
            },


            searchUser() {
                if (this.q == null || this.q == '') {
                    this.fetchData();
                } else {
                    this.handleSearch(this.q);
                }
            },

            handleSearch(q) {
                axios.get('/agent/search/' + q).then(o => {
                    if (o.data.status) {
                        this.users = o.data.data;
                    } else {
                        this.$Message.error("Хайлтанд илэрц олдсонгүй!");
                    }
                });
            },

            getRoles() {
                axios.get('/agent/roles').then(({data}) => {
                    this.roles = data;
                })
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/agent.scss";
</style>
