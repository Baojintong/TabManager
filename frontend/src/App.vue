<template>
  <a-tabs v-model:activeKey="activeKey" centered>
    <a-tab-pane key="timeRoller" tab="时间轴">
      <a-row>
        <a-col :span="22" style="margin-top: 10px;">
          <TabList/>
        </a-col>
      </a-row>
    </a-tab-pane>
    <a-tab-pane key="collect" tab="我的收藏" force-render>
      <Collect/>
    </a-tab-pane>
  </a-tabs>


  <a-float-button-group trigger="click" :style="{ right: '36px' }">
    <template #icon>
      <SettingOutlined/>
    </template>
    <LabelManage/>
  </a-float-button-group>
  <a-back-top :style="{ right: '36px' }"/>
</template>

<script setup>
import TabList from "./components/TabList.vue";
import LabelManage from "./components/LabelManage.vue";
import Collect from "./components/Collect.vue";
import {SettingOutlined} from '@ant-design/icons-vue';
import {ref} from "vue";

const activeKey = ref('timeRoller')


</script>