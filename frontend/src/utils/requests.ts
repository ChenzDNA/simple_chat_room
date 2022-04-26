import axios, { AxiosRequestConfig } from "axios"
import { ResponseCallback } from "@/types/ResponseCallback";


export default function (config: AxiosRequestConfig, callback: ResponseCallback = null) {
  const requests = axios.create({
    baseURL: '/api',
  })

  requests.interceptors.request.use(
    (config) => {
      if (config.data && config.method === 'post') {
        const fd = new FormData()
        for (const d in config.data) {
          fd.set(d, config.data[d])
        }
        config.data = fd
      }
      return config
    },
    (error) => {
      console.error(error)
    },
  )

  requests.interceptors.response.use(
    (response) => {
      callback && callback(response.data)
      return response.data
    },
    (error) => {
      console.error(error)
    }
  )
  return requests(config)
}
