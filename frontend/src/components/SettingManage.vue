<template>
  <a-float-button tooltip="设置" @click="dialogVisible = true">
    <template #icon>
      <SettingOutlined/>
    </template>
  </a-float-button>
  <a-modal v-model:open="dialogVisible" title="设置">
    <a-form
        name="basic"
        autocomplete="off"
    >
      <a-form-item label="保存目录" name="path">
        <a-input placeholder="file"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup>

import {ref, reactive} from "vue";
import {SettingOutlined} from '@ant-design/icons-vue';

const dialogVisible = ref(false)

const formState = reactive({
  username: '',
  password: '',
  remember: true,
});
const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};

</script>
