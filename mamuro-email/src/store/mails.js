import { createStore } from 'vuex';

import axios from 'axios'

const store = createStore({
    state() {
        return {
            mails: [],
            mail: {},
            index: 'arnold-j',
            search: '',
            error: false,
            loadingMSG: false,
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
        loadMails(state, payload) {
            state.mails = payload;
        },
        errorMSG(state, payload) {
            state.error = payload;
        },
        loadingMSG(state, payload) {
            state.loadingMSG = payload;
        }
    },
    actions: {
        changeMailView(context, payload) {
            context.commit('changeMailView', payload);
        },

        searchMails({ commit }, search) {
            commit('loadingMSG', true);
            axios.get(`http://192.168.1.7:3333/search/arnold-j/${search}`)
                .then(response => {
                const mailsLoaded = response.data.hits.hits.map((hit, index) => ({
                    id: index,
                    from: hit._source.from,
                    subject: hit._source.subject,
                    body: hit._source.body,
                    to: hit._source.to,
                    category: hit._source.category,
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
    },
});

export default store;

