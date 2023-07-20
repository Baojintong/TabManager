import {GetLabelList, GetTabList,GetConfigList} from "../wailsjs/go/main/App.js";
import {notification} from "ant-design-vue";
import {useStore} from 'vuex'
import {computed} from 'vue'
import {QUERY_ERROR} from "./const.js";



export const Notification = (msg, level = 'error') => {
    return notification[level]({
        message: msg,
        description: msg,
    });
}

export const setLabelList = (labelList) => {
    GetLabelList().then(res => {
        if (res.code !== 200) {
            Notification(QUERY_ERROR)
        }
        labelList.value = res.data
    })
}

export const setTabData = (tabData,labelId=0) => {
    GetTabList(labelId).then(res => {
        if (res.code !== 200) {
            Notification(QUERY_ERROR)
        }
        let list = res.data
        if (Array.isArray(list) && !(list.length === 0)) {
            tabData.value = list.reduce((acc, cur) => {
                const time = cur.saveTime;
                if (!acc[time]) {
                    acc[time] = [];
                }
                acc[time].push(cur);
                return acc;
            }, {})
        } else {
            tabData.value = {}
        }
    })
}

export const resetShowTabManageId = (tabId) => {
    tabId.value = 0
}

export const getConfigList = (configList) => {
    GetConfigList().then(res => {
        if (res.code !== 200) {
            Notification(QUERY_ERROR)
        }
        configList.value = res.data
        console.log("configList.value:",JSON.stringify(configList.value))
    })
}

/**
 * 快捷生成store
 * @returns {(*&{checked: boolean})[]|WritableComputedRef<*>}
 */
export function useLabelList() {
    const store = useStore()

    return computed({
        get() {
            return store.state.labelList.map((label, index) => {
                return {
                    ...label,
                    checked: false
                }
            })
        },
        set(val) {
            store.commit('setLabelList', val)
        }
    })
}

export function useTabData() {
    const store = useStore()

    return computed({
        get() {
            return store.state.tabData
        },
        set(val) {
            store.commit('setTabData', val)
        }
    })
}

export function showTabManageId() {
    const store = useStore()

    return computed({
        get() {
            return store.state.showTabManageId
        },
        set(val) {
            store.commit('setShowTabManageId', val)
        }
    })
}
