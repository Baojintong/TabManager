import {GetLabelList, GetTabList} from "../wailsjs/go/main/App.js";
import {notification} from "ant-design-vue";
import {useStore} from 'vuex'
import {computed} from 'vue'
import {QUERY_ERROR} from "./const.js";

export const setLabelList = (labelList) => {
    GetLabelList().then(res => {
        if (res.code !== 200) {
            Notification('标签获取失败')
        }
        labelList.value = res.data
    })
}

export const Notification = (msg, level = 'error') => {
    return notification[level]({
        message: msg,
        description: '错误',
    });
}

export const setTabData = (tabData) => {
    GetTabList().then(res => {
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
