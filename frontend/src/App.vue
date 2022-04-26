<template>
  <suspense>
    <template #default>
      <n-message-provider>
        <h3 class="title">几把聊天</h3>
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
</style>