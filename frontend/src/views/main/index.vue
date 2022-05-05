<script lang="ts">
import { defineComponent, reactive, ref } from "vue"
import requests from "@/utils/requests"
import chunked from "@/utils/chunked"
import { useMessage } from "naive-ui"
import ChatMessage from "./components/ChatMessage.vue"

const messageType = {
  error: 'error',
  warning: 'warning',
  success: 'success',
  info: 'info',
  loading: 'loading',
}

export default defineComponent({
  name: "MainView",
  components: { ChatMessage },
  methods: {
    send() {
      this.inputStr && requests({ url: '/chat/send', method: 'post', data: { msg: this.inputStr } })
      this.inputStr = ''
    },
    async quit(): Promise<any> {
      await requests({ url: '/chat/quit', method: 'post' })
      this.createMessage(messageType.success, '退出成功')
    },
    chunkedReq(callback) {
      chunked({ url: '/chat/register', method: 'get' }, callback)
    },
    createMessage(type: string, content: string) {
      this.messageHandler[type](content, { duration: 1000 })
    },
  },
  mounted() {
    let callback = async (res: string) => {
      let data = JSON.parse(res.split('####_split_me_####').pop());
      this.messageList.push({ msg: data.msg, isMe: data.userId === this.userId })
      console.log('length:', res.length)
      if (res.length > 4 * 1024 * 1024) {
        await this.quit()
        await requests({ url: '/user/getid', method: 'post' })
        this.chunkedReq(callback)
      }
    }
    this.chunkedReq(callback)
  },
  async setup() {
    const messageHandler = useMessage()
    const inputStr = ref('')
    const messageList = reactive([])
    let userId: string = (await requests({ url: '/user/getid', method: 'post' })).data

    return {
      messageHandler,
      inputStr,
      messageList,
      userId,
    }
  },
})
</script>

<template>
  <div class="chat-area">
    <n-list bordered class="chat-list">
      <chat-message v-for="(item,index) of messageList" :key="index" :msg="item.msg" :is-me="item.isMe"/>
    </n-list>
    <n-input
      type="textarea"
      size="large"
      class="input-area"
      placeholder="进行一个天的聊"
      :autosize="{
        minRows: 5,
        maxRows: 5,
      }"
      v-model:value="inputStr"
    >
    </n-input>
    <div class="btns">
      <n-button class="send-btn" strong type="error" @click="quit">退出</n-button>
      <n-button class="send-btn" type="info" @click="send">发送</n-button>
    </div>
  </div>
</template>


<style scoped>
.chat-area {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.chat-list {
  width: 100%;
  height: calc(100vh - 230px);
  overflow-y: auto;
  overflow-x: hidden;
}

.input-area {
  margin-top: 10px;
  position: fixed;
  bottom: 55px;
}

.btns {
  position: absolute;
  bottom: 10px;
  right: 10px;
}

.send-btn {
  margin-left: 10px;
}

</style>