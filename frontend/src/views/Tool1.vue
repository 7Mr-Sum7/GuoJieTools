<template>
  <div class="container">
    <a-textarea
      class="source"
      v-model:value="data.addresslist"
      placeholder="待检测资产(输入内容，用换行、逗号或空格分隔)"
      :auto-size="{ minRows: 10, maxRows: 10 }"
    />
    <a-button
      class="check"
      type="primary"
      @click="testalife"
      :disabled="data.disabled"
      >检测资产</a-button
    >
    <a-textarea
      class="response"
      v-model:value="data.response"
      placeholder="资产存活情况"
      :auto-size="{ minRows: 10, maxRows: 10 }"
    />
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from "vue";
import { notification } from "ant-design-vue";
import { Lifestyle } from "../../wailsjs/go/main/App";
const data = reactive({
  addresslist: "",
  response: "",
  disabled: false,
});

const testalife = async () => {
  data.disabled = true; // 禁用按钮
  try {
    const resp = await Lifestyle(data.addresslist);
    if (resp.code == 200) {
      let result = resp.data;
      data.response = result.resultcontent;
      openNotificationWithIcon("success", result.resultlen);
    } else {
      notification["error"]({
        message: `error`,
        description: `请求错误`,
      });
    }
  } catch (error) {
    notification["error"]({
      message: `error`,
      description: `请求错误`,
    });
  }
  data.disabled = false;
};

const openNotificationWithIcon = (
  type: "success" | "info" | "error" | "warning",
  number: number
) => {
  notification[type]({
    message: `${type}`,
    description: `共计检测到${number}个存活资产`,
  });
};
</script>

<style lang="scss" scoped>
.container {
  width: 90%;
  margin: 0 auto;
  .source {
    width: 100%;
  }
  .check {
    margin: 0 auto;
    display: block;
  }
  .response {
    width: 100%;
  }
}
</style>
