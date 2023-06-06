<template>
  <main>
    <div v-for="(items,time) in groupedData.value">
      <div class="time_div">{{ time }}</div>
      <div v-for="item in items" class="text_div" ref="text_div">
        <div v-on:click="openUrl(item.url)" class="text_content_div">
          {{ item.title }}
        </div>
        <div class="text_button_div">
          <TabManage :data="item" :getTabList="getTabList" :labelList="labelList"/>
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
import {DeleteTab, GetLabelList, GetTabList} from "../../wailsjs/go/main/App.js";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {onMounted, reactive, ref, nextTick} from "vue";
import {DeleteOutlined} from '@ant-design/icons-vue';
import TabManage from "./TabManage.vue";
import {notification} from "ant-design-vue";

//赋值方式
let groupedData = reactive({
  value: {}
})

let labelList = reactive([])

const getTabList = () => {
  GetTabList().then(res => {
    if (res.code !== 200) {
      notification['error']({
        message: '数据获取失败',
        description: '错误',
      });
    }
    let list = res.data
    if (Array.isArray(list) && !(list.length === 0)) {
      groupedData.value = list.reduce((acc, cur) => {
        const time = cur.saveTime;
        if (!acc[time]) {
          acc[time] = [];
        }
        acc[time].push(cur);
        return acc;
      }, {})
    } else {
      groupedData.value = {}
    }
  })
}

const getLabel = () => {
  GetLabelList().then(res => {
    if (res.code !== 200) {
      notification['error']({
        message: '标签获取失败',
        description: '错误',
      });
    }
    labelList = res.data
    console.log("--------2"+JSON.stringify(labelList))
  })
  console.log("--------1"+JSON.stringify(labelList))
}

onMounted(async () => {
  //setInterval(() => {
  getTabList()
  //}, 2000);
  getLabel()
})

const openUrl = (url) => {
  BrowserOpenURL(url);
}

const deleteItem = (obj) => {
  DeleteTab(JSON.stringify(obj)).then(res => {
    if (res.code !== 200) {
      notification['error']({
        message: '删除失败',
        description: '错误',
      });
    } else {
      getTabList()
    }
  })
}
</script>