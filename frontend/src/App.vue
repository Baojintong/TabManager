<template>
  <a-row>
    <a-col :span="22" style="margin-top: 10px">
      <TabList/>
    </a-col>
  </a-row>
  <div id="components-back-top-demo-custom">
    <a-back-top>
      <div class="ant-back-top-inner">UP</div>
    </a-back-top>
  </div>
</template>

<script>
import TabList from "./components/TabList.vue";
import TabManage from "./components/TabManage.vue";
import {defineComponent} from "vue";

export default defineComponent({
  components: {TabList}
})


</script>