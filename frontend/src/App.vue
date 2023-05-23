<template>
  <a-row>
    <a-col :span="22">
      <a-row class="createTextRow">
        <CreateText title="新建记事" btn-type="primary"/>
      </a-row>
      <a-divider class="divider"/>
      <div>
        <div class="time_div">2023-05-19 09:56</div>
        <div v-for="n in 4" class="text_div">
          <div class="text_content_div">
            工作记事{{ n }}
          </div>
          <div class="text_button_div">
            <a-button type="ghost" shape="circle" size="large">
              <template #icon>
                <EditOutlined/>
              </template>
            </a-button>
            <a-button type="ghost" shape="circle" size="large" style="margin-left: 8px">
              <template #icon>
                <DeleteOutlined/>
              </template>
            </a-button>
          </div>
        </div>
      </div>


      <div>
        <div class="time_div">2023-05-18 09:56</div>
        <div v-for="n in 3" class="text_div">
          <div class="text_content_div">
            工作记事{{ n }}
          </div>
          <div class="text_button_div">
            <a-button type="ghost" shape="circle" size="large">
              <template #icon>
                <EditOutlined/>
              </template>
            </a-button>
            <a-button type="ghost" shape="circle" size="large" style="margin-left: 8px">
              <template #icon>
                <DeleteOutlined/>
              </template>
            </a-button>
          </div>
        </div>
      </div>

      <div>
        <div class="time_div">2023-04-18 09:56</div>
        <div v-for="n in 2" class="text_div">
          <div class="text_content_div">
            工作记事{{ n }}
          </div>
          <div class="text_button_div">
            <a-button type="ghost" shape="circle" size="large">
              <template #icon>
                <EditOutlined/>
              </template>
            </a-button>
            <a-button type="ghost" shape="circle" size="large" style="margin-left: 8px">
              <template #icon>
                <DeleteOutlined/>
              </template>
            </a-button>
          </div>
        </div>
      </div>
    </a-col>
  </a-row>
  <div id="components-back-top-demo-custom">
    <a-back-top>
      <div class="ant-back-top-inner">UP</div>
    </a-back-top>
  </div>
</template>

<script>
import CreateText from "./components/CreateText.vue";
import {GetTabList} from "../../wailsjs/go/main/App.js";
import {ref, watch,onMounted} from "vue";

import {EditOutlined, DeleteOutlined} from '@ant-design/icons-vue';
import {notification} from 'ant-design-vue';
import {defineComponent} from 'vue';

let list = ref()
export default defineComponent({
  components: {
    CreateText,
    EditOutlined,
    DeleteOutlined
  },
});

function getTabList() {
  GetTabList().then(res => {
    if (res.code !== 200) {
      notification({
        title: res.msg,
        type: "error",
      })
    }
    list.value = res.data
  })
}


</script>