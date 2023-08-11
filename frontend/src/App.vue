<template>
  <a-tabs v-model:activeKey="tabsActiveKey" centered>
    <a-tab-pane key="timeRoller" tab="时间轴">
      <a-collapse :ghost=true>
        <a-collapse-panel key="labelFilter" header="标签筛选">
          <div class="label-list">
            <div class="label" @click="clickTag(0)" :style="{backgroundColor:'#625B5BFF',color:'#ffffff'}">重置</div>
            <div v-for="label in labelList" class="label"
                 :style="{backgroundColor:label.color,color:'#ffffff'}" @click="clickTag(label.id)" :key="label.id">
              {{ label.name }}
            </div>
          </div>
        </a-collapse-panel>
      </a-collapse>
      <a-row>
        <a-col :span="22" style="margin-top: 10px;">
          <TabList/>
        </a-col>
      </a-row>
    </a-tab-pane>
  </a-tabs>


  <a-float-button-group trigger="click" :style="{ right: '36px' }">
    <template #icon>
      <SettingOutlined/>
    </template>
    <LabelManage/>
    <SettingManage/>
  </a-float-button-group>
  <a-back-top :style="{ right: '36px' }"/>
</template>

<script setup>
import TabList from "./components/TabList.vue";
import Collect from "./components/Collect.vue";
import LabelManage from "./components/LabelManage.vue";
import SettingManage from "./components/SettingManage.vue";
import {SettingOutlined} from '@ant-design/icons-vue';
import {onMounted, ref} from "vue";
import {
  useLabelList,
  setLabelList,
  useTabData,
  setTabData
} from "./common.js"

let labelList = useLabelList()
let tabsActiveKey = ref('timeRoller')
let tabData = useTabData()

onMounted(() => {
  setLabelList(labelList)
})

const clickTag = (id) => {
  setTabData(tabData, id)
}


</script>