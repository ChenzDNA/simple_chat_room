import { AxiosRequestConfig } from "axios";
import { XHRCallback } from "@/types/ResponseCallback";

const BASE_URL = "http://localhost:8080"

export default function (config: AxiosRequestConfig, callback: XHRCallback = null) {
  const request = new XMLHttpRequest();
  request.withCredentials = true
  request.onreadystatechange = () => {
    if (request.readyState === 3) {
      // console.log('ready state:', request.readyState)
      callback && callback(request.response)
    }
  }
  request.open(config.method, BASE_URL + config.url, true)
  request.send()
}