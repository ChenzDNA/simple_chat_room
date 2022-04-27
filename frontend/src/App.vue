<template>
  <suspense>
    <template #default>
      <n-message-provider>
        <h3 class="title">几把聊天</h3>
        <a class="a-git" href="https://github.com/ChenzDNA/simple_chat_room">Github</a>
        <MainView/>
      </n-message-provider>
    </template>
    <template #fallback>
      <div>loading...</div>
    </template>
  </suspense>
</template>

<script lang="ts">
import { defineComponent } from "vue"
import MainView from '@/views/main/index.vue'

export default defineComponent({
  components: { MainView }
})
</script>

<style scoped>
.title {
  text-align: center;
}

.a-git {
  position: absolute;
  right: 10px;
  top: 5px;
  color: dodgerblue;
}
</style>