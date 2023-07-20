<template>
  <a-float-button tooltip="设置" @click="dialogVisible = true">
    <template #icon>
      <SettingOutlined/>
    </template>
  </a-float-button>
  <a-modal v-model:open="dialogVisible" title="设置" @ok="handleOk">
    <a-form
        name="basic"
        autocomplete="off"
    >
      <a-form-item v-for="config in configList" v-model:label="config.describe" v-model:name="config.key">
        <a-input v-model:value="config.value"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup>

import {ref, reactive, onMounted} from "vue";
import {SettingOutlined} from '@ant-design/icons-vue';
import {Notification, setTabData} from "../common.js"
import {GetConfigList, SaveConfig} from "../../wailsjs/go/main/App.js";
import {QUERY_ERROR, UPDATE_ERROR} from "../const.js";

const dialogVisible = ref(false)
let configList = reactive([])

const handleOk = e => {
  console.log("Success:", JSON.stringify(configList))
  SaveConfig(JSON.stringify(configList)).then(res => {
    if (res.code !== 200) {
      Notification(UPDATE_ERROR)
    } else {
      //getConfigList()
    }
  })
}

onMounted(() => {
  getConfigList()
})

const getConfigList = () => {
  GetConfigList().then(res => {
    if (res.code !== 200) {
      Notification(QUERY_ERROR)
    }
    configList = reactive(res.data)
  })
}

</script>
