<template>
  <div>
    <a-menu
      v-model:openKeys="state.openKeys"
      v-model:selectedKeys="state.selectedKeys"
      mode="horizontal"
      theme="dark"
      :inline-collapsed="state.collapsed"
      :items="items"
    ></a-menu>
  </div>
</template>
<script lang="ts" setup>
import { reactive, watch, } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const state = reactive({
  collapsed: false,
  selectedKeys: [""],
  openKeys: [""],
  preOpenKeys: [""],
});
const items = reactive([
  {
    key: "Index",
    label: "首页",
    title: "首页",
    onClick: () => router.push("/Index")
  },
  {
    key: "Tool1",
    label: "检测资产存活",
    title: "检测资产存活",
    onClick: () => router.push("/Tool1")
  },
  {
    key: "Tool2",
    label: "工具2",
    title: "工具2",
    onClick: () => router.push("/Tool2")
  },
]);
watch(
  () => state.openKeys,
  (_val, oldVal) => {
    state.preOpenKeys = oldVal;
  }
);
</script>
