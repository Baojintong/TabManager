import {createStore} from 'vuex'

// Create a new store instance.
const store = createStore({
    state: {
        count: 100,
        labelList: [],
        tabData: {}
    },
    mutations: {
        setTabData(state, val) {
            state.tabData = val
        },
        setLabelList(state, val) {
            state.labelList = val
        }
    },
    getters: {},
    actions: {},
    modules: {}
})
export default store