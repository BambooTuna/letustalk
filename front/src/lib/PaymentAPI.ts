import axios, { AxiosError, AxiosResponse } from 'axios'
import { ErrorResponseJson, InvoiceDetail } from '@/lib/Protocol'

export default class PaymentAPI {
  private endpoint!: string
  constructor (endpoint = process.env.VUE_APP_SERVER_ENDPOINT) {
    this.endpoint = endpoint
  }

  getInvoiceDetail (invoiceId: string): Promise<InvoiceDetail> {
    return axios({
      url: this.endpoint + '/invoice/' + invoiceId,
      method: 'get'
    })
      .then((res: AxiosResponse) => res.data)
      // .catch(this.errorHandler)
  }

  pay (invoiceId: string, token: string): Promise<InvoiceDetail> {
    return axios({
      url: this.endpoint + '/pay/' + invoiceId,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: { token: token }
    })
      .then((res: AxiosResponse) => res.data)
      .catch(this.errorHandler)
  }

  private errorHandler<T> (e: AxiosError): Promise<T> {
    const errorResponseJson: ErrorResponseJson = e.response?.data
    return Promise.reject(new Error(errorResponseJson.message))
  }
}
