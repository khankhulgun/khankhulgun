<template>
    <div class="dv-pagination">
        <div class="dv-pagination-info">
            <div class="dv-pagination-info-sort">
                <Select class="dv-pagination-sort-select roles-select" clearable @on-change="filterByRole"
                        placeholder="Зэрэглэлээр шүүх"
                        label-in-value
                        style="width:200px">
                    <Option v-for="item in roles" :key="item.index" :value="item.id">
                        <span>{{ item.display_name }}</span>
                    </Option>
                </Select>

                <Select class="dv-pagination-sort-select" @on-change="sortSelect" placeholder="Эрэмбэлэх" label-in-value
                        style="width:200px">
                    <Option value="idASC">
                        <span>ID-гаар:</span>
                        <span class="dv-sort-direction">&uarr;өсөхөөр</span>
                    </Option>
                    <Option value="idDESC">
                        <span>ID-гаар:</span>
                        <span class="dv-sort-direction">&darr;буурах</span>
                    </Option>
                    <Option value="loginASC">
                        <span>Нэвтрэх нэрээр:</span>
                        <span class="dv-sort-direction">&uarr;өсөхөөр</span>
                    </Option>
                    <Option value="loginDESC">
                        <span>Нэвтрэх нэрээр:</span>
                        <span class="dv-sort-direction">&darr;буурах</span>
                    </Option>
                    <Option value="createdASC">
                        <span>Нэмсэн огноо:</span>
                        <span class="dv-sort-direction">&uarr;өсөхөөр</span>
                    </Option>
                    <Option value="createdDESC">
                        <span>Нэмсэн огноо:</span>
                        <span class="dv-sort-direction">&darr;буурах</span>
                    </Option>
                </Select>
            </div>
            <span class="page-info">Нийт <strong>{{isDeleted ? deletedModel.total : model.total}}</strong> {{ pagination.name }}<strong> {{isDeleted ? deletedModel.from : model.from}} - {{isDeleted ? deletedModel.to : model.to}}</strong> харуулж байна</span>
        </div>
        <div class="dv-pagination-control">
            <div class="dv-per-page ivu-page-options-elevator">
                <span>Хуудсанд</span>
                <input v-model.number.lazy="query.per_page" @keyup.enter="perPage">
                <span>-г харуулана</span>
            </div>
            <div>
                <Page :total="isDeleted ? deletedModel.total : model.total" size="small" :current="query.page"
                      :page-size="query.per_page" show-elevator @on-change="changePage" class-name="dv-control"></Page>
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        name: "Pagination",
        props: {
            pagination: {},
            isDeleted: Boolean,
            deletedModel: {},
            model: {},
            isTop: Boolean,
            layout: '',
            query: {
                per_page: Number
            },
            roles: []
        },
//        props: ['pagination', 'isDeleted', 'deletedModel', 'model', 'isTop', 'layout', 'query'],
        methods: {
            perPage() {
//                this.query.per_page = val;
                this.$parent.fetchData();
                this.$parent.fetchDeletedData();
            },
            changePage(page) {
                this.query.page = page;
                this.$parent.fetchData();
                this.$parent.fetchDeletedData();
            },
            sortSelect(sort) {
                if (sort.value === 'idASC') {
                    this.query.direction = 'asc';
                    this.query.column = 'id';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
                else if (sort.value === 'idDESC') {
                    this.query.direction = 'desc';
                    this.query.column = 'id';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
                else if (sort.value === 'loginASC') {
                    this.query.direction = 'asc';
                    this.query.column = 'login';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
                else if (sort.value === 'loginDESC') {
                    this.query.direction = 'desc';
                    this.query.column = 'login';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
                else if (sort.value === 'createdASC') {
                    this.query.direction = 'asc';
                    this.query.column = 'created_at';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
                else if (sort.value === 'createdDESC') {
                    this.query.direction = 'desc';
                    this.query.column = 'created_at';
                    this.$parent.fetchData();
                    this.$parent.fetchDeletedData();
                }
            },

            filterByRole(item) {
                if (typeof item == 'undefined' || item == undefined) {
                    this.query.role = 'all';
                } else {
                    this.query.role = item.value;
                }
                this.$parent.fetchData();
                this.$parent.fetchDeletedData();
            }
        }
    }
</script>
