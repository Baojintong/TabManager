<template>
  <main>
    <a-button :type="btnType" @click="dialogVisible = true" size="large">{{title}}</a-button>
    <a-modal v-model:visible="dialogVisible" title="新建tab" @ok="handleOk">
      <a-form
          :model="formState"
          name="basic"
          autocomplete="off"
          @finish="onFinish"
          @finishFailed="onFinishFailed"
      >
        <a-form-item label="标题" name="title">
          <a-input v-model:value="formState.title" />
        </a-form-item>
        <a-form-item label="链接" name="url">
          <a-input v-model:value="formState.url" />
        </a-form-item>
        <a-form-item label="图标" name="iconUrl">
          <a-input v-model:value="formState.iconUrl" />
        </a-form-item>
        <a-form-item label="描述" name="describe">
          <a-input v-model:value="formState.describe" />
        </a-form-item>
      </a-form>
    </a-modal>
  </main>
</template>

<script setup>
import {ref,defineComponent, reactive} from "vue";
let props = defineProps(['title', 'btnType', 'data'])
const dialogVisible = ref(false)
const formState = reactive({
  title: '',
  url: '',
  iconUrl: '',
  describe: ''
});
const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};

const handleOk = e => {
  console.log(e);
  dialogVisible.value = false;
};
</script>