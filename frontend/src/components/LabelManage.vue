<template>
  <a-float-button tooltip="个人标签编辑" @click="dialogVisible = true"/>
  <a-modal v-model:open="dialogVisible" title="标签编辑" @ok="handleOk" @cancel="cancel">
    <a-form
        name="basic"
        autocomplete="off"
    >
      <a-form-item label="标签名称" name="labelName">
        <a-input v-model:value="label.name.value" placeholder="自定义标签" show-count :maxlength="20"/>
      </a-form-item>
      <a-form-item label="标签颜色" name="labelColor">
        <ColorPicker :color="label.color.value" @color-change="updateColor"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup>

import {ref} from "vue";
import {ColorPicker} from 'vue-accessible-color-picker'
import {SaveLabel,} from "../../wailsjs/go/main/App.js";
import {setLabelList, useLabelList} from "../common.js"
import {SAVE_ERROR} from "../const.js";

const dialogVisible = ref(false)
const labelList = useLabelList()
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
  SaveLabel(JSON.stringify(data)).then(res => {
    if (res.code !== 200) {
      Notification(SAVE_ERROR)
    } else {
      setLabelList(labelList)
    }
  })
}

const updateColor = e => {
  label.color.value = e.cssColor
}

</script>
