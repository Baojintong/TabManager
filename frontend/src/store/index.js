import {createStore} from 'vuex'
import {ref} from 'vue'

// Create a new store instance.
const store = createStore({
    state: {
        count: 100,
        labelList: [],
        tabData: {},
        showTabManageId: 0
    },
    mutations: {
        setTabData(state, val) {
            state.tabData = val
        },
        setLabelList(state, val) {
            state.labelList = val
        },
        setShowTabManageId(state, val) {
            state.showTabManageId = val
        }
    },
    getters: {},
    actions: {},
    modules: {}
})
export default store