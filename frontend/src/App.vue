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
    <a-float-button tooltip="个人标签编辑"/>
  </a-float-button-group>
  <a-back-top :style="{ right: '36px' }"/>
</template>

<script setup>
import TabList from "./components/TabList.vue";
import {SettingOutlined} from '@ant-design/icons-vue';


</script>