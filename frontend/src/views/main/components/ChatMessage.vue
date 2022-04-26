<template>
  <div class="chat-message__out">
    <n-avatar
      class="chat-avatar"
      :style="{float:props.isMe?'right':'left'}"
    />
    <div
      class="chat-message"
      :style="msgStyle"
    >
      {{ props.msg }}
    </div>
  </div>
</template>

<script setup>
import { computed, defineProps } from "vue";

let props = defineProps({
  msg: String,
  isMe: Boolean,
})

const msgStyle = computed(() => {
  if (props.isMe) {
    return {
      float: 'right',
      border: '1px solid rgb(32,128,240)',
      color: 'white',
      background: 'rgb(32,128,240)',
    }
  }
  return {
    float: 'left',
    border: '1px solid lightgray',
    color: 'black',
    background: 'white',
  }
})
</script>

<style scoped>
.chat-message {
  white-space: pre-wrap;
  max-width: 250px;
  word-break: normal;
  word-wrap: break-word;
  padding: 5px;
  border-radius: 5px;
}

.chat-message__out {
  width: 100%;
  overflow: hidden;
  margin: 10px 0;
}

.chat-avatar {
  margin: 0 10px 10px;
}
</style>