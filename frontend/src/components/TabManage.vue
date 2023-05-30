<template>
  <a-button type="ghost" shape="circle" size="large" @click="dialogVisible = true">
    <template #icon>
      <EditOutlined/>
    </template>
  </a-button>

  <a-modal v-model:visible="dialogVisible" title="编辑" @ok="handleOk" @cancel="cancel">
    <a-form
        :model="formState"
        name="basic"
        autocomplete="off"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
    >
      <a-form-item label="标题" name="title">
        <a-input v-model:value="item.title"/>
      </a-form-item>
      <a-form-item label="链接" name="url">
        <a-input v-model:value="item.url" :readonly="true" :disabled="true"/>
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-input v-model:value="item.describe"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {ref, reactive} from "vue";
import {EditOutlined} from "@ant-design/icons-vue";
import {UpdateTab} from "../../wailsjs/go/main/App.js";
import {deepEqual} from "../utils.js";

let props = defineProps(['data', 'getTabList'])
const dialogVisible = ref(false)
let item = {}
let originalItem = {}

if (props.data !== undefined) {
  item = reactive({...props.data});
  originalItem = JSON.parse(JSON.stringify(props.data));
}


const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};

const handleOk = e => {
  dialogVisible.value = false;
  UpdateTab(JSON.stringify(item))
  props.getTabList()
};

const cancel = e => {
  Object.assign(item, originalItem)
  props.getTabList()
};
</script>