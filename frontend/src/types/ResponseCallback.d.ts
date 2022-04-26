export interface ResponseCallback {
  // eslint-disable-next-line no-unused-vars
  (response: ResponseType): void
}

export interface XHRCallback {
  // eslint-disable-next-line no-unused-vars
  (response: string): void
}

export interface ResponseType {
  code: number
  msg: string
  data: any
}