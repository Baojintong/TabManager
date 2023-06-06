<template>
  <a-button type="ghost" shape="circle" size="large" @click="dialogVisible = true" class="button_edit">
    <template #icon>
      <EditOutlined/>
    </template>
  </a-button>

  <a-modal v-model:open="dialogVisible" title="编辑" @ok="handleOk" @cancel="cancel">
    <a-form
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

      <a-form-item label="选择标签" name="label">
        <a-tag v-for="label in labelList" :color="label.color">
          {{ label.name }}
        </a-tag>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {reactive, ref} from "vue";
import {EditOutlined} from "@ant-design/icons-vue";
import {UpdateTab, GetLabelList} from "../../wailsjs/go/main/App.js";
import { notification } from 'ant-design-vue';

let props = defineProps(['data', 'getTabList', 'labelList'])
let dialogVisible = ref(false)
let item = reactive({})
let originalItem = {}
let labelList = reactive([])

if (props.data !== undefined) {
  item = props.data//reactive({...props.data});
  originalItem = JSON.parse(JSON.stringify(props.data));
}

if (props.labelList !== undefined) {
  labelList = props.labelList
}


const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};
const handleOk = e => {
  dialogVisible.value = false;
  UpdateTab(JSON.stringify(item)).then(res => {
    if (res.code !== 200) {
      notification['error']({
        message: '更新失败',
        description: '错误',
      });
    } else {
      originalItem = JSON.parse(JSON.stringify(item));
      props.getTabList()
    }
  })
};

const cancel = e => {
  Object.assign(item, originalItem)
  props.getTabList()
};
</script>