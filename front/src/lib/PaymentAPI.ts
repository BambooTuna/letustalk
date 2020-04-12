export default class PaymentAPI {
  private endpoint!: string
  constructor(endpoint = process.env.VUE_APP_SERVER_ENDPOINT) {
    this.endpoint = endpoint
  }
}
