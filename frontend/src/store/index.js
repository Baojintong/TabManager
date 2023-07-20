import {createStore} from 'vuex'
import {ref} from 'vue'

// Create a new store instance.
const store = createStore({
    state: {
        labelList: [],
        tabData: {},
        showTabManageId: 0,
        configList:[]
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
        },
        setConfigList(state, val) {
            state.configList = val
        }
    },
    getters: {},
    actions: {},
    modules: {}
})
export default store