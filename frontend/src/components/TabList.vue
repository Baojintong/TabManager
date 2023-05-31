<template>
  <main>
    <div v-for="(items,time) in groupedData">
      <div class="time_div">{{ time }}</div>
      <div v-for="item in items" class="text_div">
        <div v-on:click="openUrl(item.url)" class="text_content_div">
          {{ item.title }}
        </div>
        <div class="text_button_div">
          <TabManage :data="item" :getTabList="getTabList"/>
          <a-button type="ghost" shape="circle" size="large" style="margin-left: 8px">
            <template #icon>
              <DeleteOutlined v-on:click="deleteItem(item)"/>
            </template>
          </a-button>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import {DeleteTab, GetTabList} from "../../wailsjs/go/main/App.js";
import {BrowserOpenURL} from "../../wailsjs/runtime";
import {ref} from "vue";
import {DeleteOutlined, EditOutlined} from '@ant-design/icons-vue';
import {notification} from 'ant-design-vue';
import TabManage from "./TabManage.vue";

let list = ref()
export default {
  components: {
    TabManage,
    EditOutlined,
    DeleteOutlined
  },
  created() {
    setInterval(() => {
      this.getTabList()
    }, 2000)
    this.getTabList()
  },
  data() {
    return {
      groupedData: null
    }
  },
  methods: {
    getTabList() {
      GetTabList().then(res => {
        if (res.code !== 200) {
          notification({
            title: res.msg,
            type: "error",
          })
        }
        //this.list = res.data
        this.groupedData = res.data.reduce((acc, cur) => {
          const time = cur.saveTime;
          if (!acc[time]) {
            acc[time] = [];
          }
          acc[time].push(cur);
          return acc;
        }, {})
      })
    },
    openUrl(url) {
      BrowserOpenURL(url);
    },
    deleteItem(obj) {
      DeleteTab(JSON.stringify(obj))
    }
  }
};
</script>