import { createStore } from 'vuex';

import axios from 'axios'

const store = createStore({
    state() {
        return {
            mails: [],
            mail: {},
            name: '',
            search: '',
            error: false,
            loadingMSG: false,
            total_mails: 0,
            actual_page: 0,
            total_pages: 0,
        }
    },
    mutations: {
        changeMailView(state, payload) {
            state.mail.from = payload.from;
            state.mail.subject = payload.subject;
            state.mail.body = payload.body;
            state.mail.to = payload.to;
            state.mail.category = payload.category;
        },
        changeSearch(state, payload) {
            state.search = payload;
        },
        loadMails(state, payload) {
            state.mails = payload;
        },
        errorMSG(state, payload) {
            state.error = payload;
        },
        loadingMSG(state, payload) {
            state.loadingMSG = payload;
        },
        changeTotalMails(state, payload) {
            state.total_mails = payload;
        },
        changeActualPage(state, payload) {
            state.actual_page = payload;
        },
        changeTotalPages(state, payload) {
            state.total_pages = payload;
        },
        changeName(state, payload) {
            state.name = payload;
        }
    },
    actions: {
        changeMailView(context, payload) {
            context.commit('changeMailView', payload);
        },
        changeName(context, payload) {
            context.commit('changeName', payload);
        },
        loadMails({ commit }, payload) {
            commit('loadingMSG', true);
            axios.get(`http://192.168.1.7:3333/search/${payload}/${this.state.actual_page*20}`)
                .then(response => {
                    if(response.data.hits.total.value >= 100) {
                        commit('changeTotalMails', 100);
                    } else {
                        commit('changeTotalMails', response.data.hits.total.value);
                    } 

                    if(response.data.hits.total.value === 0) {
                        commit('changeTotalPages', 0);
                        commit('errorMSG', true)
                        setTimeout(() => {
                            commit('errorMSG', false);
                        }, 3000);
                    }

                    commit('changeTotalPages', Math.ceil(this.state.total_mails/20));

                const mailsLoaded = response.data.hits.hits.map((hit, index) => ({
                    id: index,
                    from: hit._source["mail.from"],
                    subject: hit._source["mail.subject"],
                    body: hit._source["mail.body"],
                    to: hit._source["mail.to"],
                    category: hit._source["mail.category"],
                }));
                commit('loadMails', mailsLoaded);
                commit('loadingMSG', false);
                })
                .catch(error => {
                console.log(`Error: ${error}`);
                commit('errorMSG', true);
                commit('loadingMSG', false);
                setTimeout(() => {
                    commit('errorMSG', false);
                }, 3000);
            });
        },
        searchMails({ commit }, search) {
            commit('changeActualPage', 0);
            commit('changeSearch', search);
            commit('changeTotalMails', 0);
            commit('changeTotalPages', 0);
            this.dispatch('loadMails', search);
            },
        changeActualPage({ commit }, payload) {
            if(this.state.actual_page + payload >= 0 && this.state.actual_page + payload <= this.state.total_pages) {
                commit('changeActualPage', this.state.actual_page + payload);
                this.dispatch('loadMails', this.state.search);
            }
        }
    },
});

export default store;

