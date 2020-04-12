<template>
  <div class="test-payment">
    <div v-show="state === 'unpaid'">
      <WaitLoading :loadingFlag="loadingFlag">
        <p>支払い金額: {{invoiceDetail.amount}}</p>
        <PaymentCheckoutForm :public-key="paymentPublicKey" @token-created-event="tokenCreatedEvent" @token-failed-event="tokenFailedEvent"></PaymentCheckoutForm>
      </WaitLoading>
    </div>
    <div v-show="state === 'processing'">
      <p>支払い処理中</p>
      <WaitLoading :loadingFlag="loadingFlag"></WaitLoading>
    </div>
    <div v-show="state === 'complete'">
      <p>支払い金額: {{invoiceDetail.amount}}</p>
      <p>支払い完了</p>
    </div>
    <div v-show="state === 'error'">
      <p>エラーが起きました</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import PaymentCheckoutForm from '@/components/PaymentCheckoutForm.vue'
import WaitLoading from '@/components/WaitLoading.vue'
import PaymentAPI from '@/lib/PaymentAPI'
import { InvoiceDetail, PaymentState } from '@/lib/Protocol'

@Component({
  components: {
    PaymentCheckoutForm, WaitLoading
  }
})
export default class TestPayment extends Vue {
    private api: PaymentAPI = new PaymentAPI()
    private paymentPublicKey: string = process.env.VUE_APP_PAYMENT_PUB_KEY
    private invoiceDetail?: InvoiceDetail = {
      invoiceId: '',
      amount: 0,
      paid: false
    }

    private loadingFlag = true
    private state: PaymentState = 'unpaid'

    async created () {
      this.api.getInvoiceDetail(this.$route.params.invoiceId)
      // eslint-disable-next-line no-return-assign
        .then((res) => this.invoiceDetail = res)
      // eslint-disable-next-line no-return-assign
        .catch(() => this.state = 'error')
        // eslint-disable-next-line no-return-assign
        .finally(() => this.loadingFlag = false)
    }

    tokenCreatedEvent (token: string) {
      this.state = 'processing'
      this.loadingFlag = true
      this.api
        .makePayment(this.$route.params.invoiceId, token)
        // eslint-disable-next-line no-return-assign
        .then(() => this.state = 'complete')
        // eslint-disable-next-line no-return-assign
        .catch(() => this.state = 'error')
        // eslint-disable-next-line no-return-assign
        .finally(() => this.loadingFlag = false)
    }

    tokenFailedEvent (message: string) {
      alert(message)
    }
}
</script>
