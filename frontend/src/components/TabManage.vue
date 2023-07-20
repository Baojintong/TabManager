<template>
  <a-modal v-model:open="dialogVisible" title="编辑" @ok="handleOk" @cancel="cancel">
    <a-form
        name="basic"
        autocomplete="off"
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
        <a-checkable-tag v-for="label in labelList"
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
import {GetTabLabelList, UpdateTab,GetTab} from "../../wailsjs/go/main/App.js";
import {
  Notification,
  setLabelList,
  setTabData,
  useLabelList,
  useTabData,
  resetShowTabManageId,
  showTabManageId
} from "../common.js"
import {UPDATE_ERROR} from "../const.js";

let props = defineProps(['data'])
let dialogVisible = ref(false)
let item = reactive({})
let nextSelectedTags = []
let labelList = useLabelList()
let tabData = useTabData()
let labelIds = []
const tabManageId = showTabManageId()
const state = reactive({
  selectedTags: [],
});


const handleOk = e => {
  dialogVisible.value = false;
  item.labelIds = state.selectedTags
  UpdateTab(JSON.stringify(item)).then(res => {
    if (res.code !== 200) {
      Notification(UPDATE_ERROR)
    } else {
      setTabData(tabData)
    }
  })
  resetShowTabManageId(tabManageId)
};

const cancel = e => {
  //更新tabData数据
  setTabData(tabData)
  resetShowTabManageId(tabManageId)
};

onMounted(() => {
  getTabLabel(tabManageId.value)
  getTab(tabManageId.value)
  dialogVisible.value = true
})
const handleChange = (labelId, checked) => {
  const {
    selectedTags,
  } = state;
  nextSelectedTags = checked ? [...selectedTags, labelId] : selectedTags.filter(t => t !== labelId);
  state.selectedTags = nextSelectedTags;
};

const getTab = (tabManageId) => {
  GetTab(tabManageId).then(res => {
    if (res.code !== 200) {
      Notification('标签获取失败')
    }
    item = reactive(res.data)
  })
}

const getTabLabel = (tabManageId) => {
  GetTabLabelList(tabManageId).then(res => {
    if (res.code !== 200) {
      Notification('标签获取失败')
    }
    let tabLabelList = res.data
    if (Array.isArray(tabLabelList) && tabLabelList.length !== 0) {
      labelIds = tabLabelList.map((it) => {
        return it.labelId;
      });
      state.selectedTags = [...state.selectedTags, ...labelIds]
    }
  })
};
</script>