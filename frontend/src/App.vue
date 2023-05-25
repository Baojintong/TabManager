<template>
  <a-row>
    <a-col :span="22">
      <a-row class="createTextRow">
        <CreateText title="新建记事" btn-type="primary"/>
      </a-row>
      <a-divider class="divider"/>
      <div v-for="(items,time) in groupedData">
        <div class="time_div">{{ time }}</div>
        <div v-for="item in items" class="text_div">
<!--          <div style="float:left;"><img :src="item.iconUrl" width="20" height="20" @error="handleImageError(item)" alt="image"></div>-->
          <div v-on:click="openUrl(item.url)" class="text_content_div">
            {{ item.title }}
          </div>
          <div class="text_button_div">
            <a-button type="ghost" shape="circle" size="large">
              <template #icon>
                <EditOutlined/>
              </template>
            </a-button>
            <a-button type="ghost" shape="circle" size="large" style="margin-left: 8px">
              <template #icon>
                <DeleteOutlined/>
              </template>
            </a-button>
          </div>
        </div>
      </div>
    </a-col>
  </a-row>
  <div id="components-back-top-demo-custom">
    <a-back-top>
      <div class="ant-back-top-inner">UP</div>
    </a-back-top>
  </div>
</template>

<script>
import CreateText from "./components/CreateText.vue";
import {GetTabList} from "../wailsjs/go/main/App.js";
import {ref, watch, onMounted, defineComponent} from "vue";
//let props = defineProps(['flush'])

import {EditOutlined, DeleteOutlined} from '@ant-design/icons-vue';
import {notification} from 'ant-design-vue';

/*watch(props, (newFlush)=>{
  getTabList()
})*/
onMounted(() => {
  getTabList()
})

let list = ref()
export default defineComponent({
  components: {
    CreateText,
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
        const groupedData = res.data.reduce((acc, cur) => {
          const time = cur.saveTime;
          if (!acc[time]) {
            acc[time] = [];
          }
          acc[time].push(cur);
          return acc;
        }, {});
        this.groupedData = groupedData
      })
    },
    openUrl(url) {
      window.open(url, '_blank');

    },
    handleImageError(item){
      item.iconUrl="../../image/appicon.png"
    }
  }
});


</script>