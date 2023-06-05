<template>
  <a-row>
    <a-col :span="22" style="margin-top: 10px;">
      <TabList/>
    </a-col>
  </a-row>


  <a-float-button-group trigger="click" :style="{ right: '36px' }">
    <template #icon>
      <SettingOutlined />
    </template>
    <LabelManage/>
  </a-float-button-group>
  <a-back-top :style="{ right: '36px' }"/>
</template>

<script setup>
import TabList from "./components/TabList.vue";
import LabelManage from "./components/LabelManage.vue";
import {SettingOutlined} from '@ant-design/icons-vue';


</script>