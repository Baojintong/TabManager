<template>
  <a-float-button tooltip="个人标签编辑" @click="dialogVisible = true"/>
  <a-modal v-model:open="dialogVisible" title="标签编辑" @ok="handleOk" @cancel="cancel">
    <a-form
        name="basic"
        autocomplete="off"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
    >
      <a-form-item label="标签名称" name="labelName">
        <a-input v-model:value="label.name.value" placeholder="自定义标签"/>
      </a-form-item>
      <a-form-item label="标签颜色" name="labelColor">
        <ColorPicker :color="label.color.value" @color-change="updateColor"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup>

import {ref, reactive} from "vue";
import {ColorPicker} from 'vue-accessible-color-picker'
import {SaveLabel, UpdateTab} from "../../wailsjs/go/main/App.js";

const dialogVisible = ref(false)

let props = defineProps(['data'])
let label = {}

if (props.data !== undefined) {
  label = props.data
  label = {
    name: ref(props.data.name),
    color: ref(props.data.color)
  }
} else {
  label = {
    name: ref(""),
    color: ref("hsl(177 34.09% 39.77% / 0.49)")
  }
}


const handleOk = e => {
  dialogVisible.value = false;
  let data = {
    name: label.name.value,
    color: label.color.value
  }
  if (data.name === '') {
    data.name = '自定义标签'
  }
  SaveLabel(JSON.stringify(data))
}

const cancel = e => {

}

const onFinish = values => {
  console.log('Success:', values);
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};

const updateColor = e => {
  label.color.value = e.cssColor
}

</script>
