import axios, { AxiosError, AxiosResponse } from 'axios'
import { AccountDetail, ErrorResponseJson, FreeSchedule, InvoiceDetail } from '@/lib/Protocol'

export default class RestAPI {
  private endpoint!: string
  constructor (endpoint = process.env.VUE_APP_SERVER_ENDPOINT) {
    this.endpoint = endpoint
  }

  getAllMentor (): Promise<Array<AccountDetail>> {
    return axios({
      url: this.endpoint + '/mentor',
      method: 'get'
    })
      .then((res: AxiosResponse) => res.data)
    // .catch(this.errorHandler)
  }

  getFreeSchedule (accountId: string): Promise<Array<FreeSchedule>> {
    return axios({
      url: this.endpoint + '/account/' + accountId + '/schedule',
      method: 'get'
    })
      .then((res: AxiosResponse) => res.data)
    // .catch(this.errorHandler)
  }

  getInvoiceDetail (invoiceId: string): Promise<InvoiceDetail> {
    return axios({
      url: this.endpoint + '/invoice/' + invoiceId,
      method: 'get'
    })
      .then((res: AxiosResponse) => res.data)
      // .catch(this.errorHandler)
  }

  issueAnInvoice (amount: number): Promise<InvoiceDetail> {
    return axios({
      url: this.endpoint + '/invoice',
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: { amount: amount }
    })
      .then((res: AxiosResponse) => res.data)
      .catch(this.errorHandler)
  }

  makePayment (invoiceId: string, token: string): Promise<InvoiceDetail> {
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
