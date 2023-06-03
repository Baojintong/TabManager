<template>
  <a-row>
    <a-col :span="22" style="margin-top: 10px;">
      <TabList/>
    </a-col>
  </a-row>


  <a-float-button-group trigger="click" type="primary" :style="{ right: '24px' }">
    <template #icon>
      TOP
    </template>
    <a-float-button />
    <a-float-button>
      <template #icon>
        <div id="components-back-top-demo-custom">
          <a-back-top>
            <div class="ant-back-top-inner">UP</div>
          </a-back-top>
        </div>
      </template>
    </a-float-button>
  </a-float-button-group>
</template>

<script>
import TabList from "./components/TabList.vue";
import {defineComponent} from "vue";

export default defineComponent({
  components: {TabList}
})


</script>