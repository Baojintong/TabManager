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
        <a-checkable-tag v-for="label in labelList" :color="label.color"
                         :style="{borderColor:label.color}"
                         :checked="state.selectedTags.indexOf(label.id) > -1"
                         @change="checked => handleChange(label.id, checked)"
        >
          {{ label.name }}
        </a-checkable-tag>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import {reactive, ref, onMounted} from "vue";
import {EditOutlined} from "@ant-design/icons-vue";
import {UpdateTab} from "../../wailsjs/go/main/App.js";
import {setLabelList, setTabData, useLabelList, useTabData} from "../common.js"
import {UPDATE_ERROR} from "../const.js";

let props = defineProps(['data'])
let dialogVisible = ref(false)
let item = reactive({})
let originalItem = {}

let labelList = useLabelList()
let tabData = useTabData()

if (props.data !== undefined) {
  item = props.data
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
  UpdateTab(JSON.stringify(item)).then(res => {
    if (res.code !== 200) {
      Notification(UPDATE_ERROR)
    } else {
      originalItem = JSON.parse(JSON.stringify(item));
      setTabData(tabData)
    }
  })
};

const cancel = e => {
  Object.assign(item, originalItem)
  setTabData(tabData)
  state.selectedTags=[]
};

onMounted(() => {
  setLabelList(labelList)
})

const state = reactive({
  selectedTags: [],
});
const handleChange = (labelId, checked) => {
  const {
    selectedTags,
  } = state;
  const nextSelectedTags = checked ? [...selectedTags, labelId] : selectedTags.filter(t => t !== labelId);
  console.log('You are interested in: ', nextSelectedTags);
  state.selectedTags = nextSelectedTags;
};
</script>