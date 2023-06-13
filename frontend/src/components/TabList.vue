<template>
  <main>
    <div v-for="(items,time) in tabData">
      <div class="time_div">{{ time }}</div>
      <div v-for="item in items" class="text_div" ref="text_div">
        <div v-on:click="openUrl(item.url)" class="text_content_div">
          {{ item.title }}
        </div>
        <div class="text_button_div">
          <TabManage :data="item"/>
          <a-button type="ghost" shape="circle" size="large" class="button_del">
            <template #icon>
              <DeleteOutlined v-on:click="deleteItem(item)"/>
            </template>
          </a-button>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
import {DeleteTab} from "../../wailsjs/go/main/App.js";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {onMounted} from "vue";
import {DeleteOutlined} from '@ant-design/icons-vue';
import TabManage from "./TabManage.vue";
import {useLabelList, setLabelList, Notification, useTabData, setTabData} from "../common.js"
import {DELETE_ERROR} from "../const.js";


const labelList = useLabelList()
const tabData = useTabData()

onMounted(() => {
  resetTabData()
  setLabelList(labelList)
  window.runtime.EventsOn('flushTabs',resetTabData)
})

const openUrl = (url) => {
  BrowserOpenURL(url);
}

const resetTabData = () =>{
  setTabData(tabData)
}

const deleteItem = (obj) => {
  DeleteTab(JSON.stringify(obj)).then(res => {
    if (res.code !== 200) {
      Notification(DELETE_ERROR)
    } else {
      resetTabData()
    }
  })
}
</script>