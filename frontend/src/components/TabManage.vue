<template>
  <a-button type="ghost" shape="circle" size="large" @click="dialogVisible = true">
    <template #icon>
      <EditOutlined/>
    </template>
  </a-button>

  <a-modal v-model:visible="dialogVisible" title="编辑" @ok="handleOk">
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
import {ref, defineComponent, reactive} from "vue";
import {EditOutlined} from "@ant-design/icons-vue";
import {UpdateTab} from "../../wailsjs/go/main/App.js";

let props = defineProps(['data'])
const dialogVisible = ref(false)
let item = reactive({
  title: '',
  url: '',
  iconUrl: '',
  describe: ''
});

if (props.data !== undefined) {
  item = props.data
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
};
</script>