<template>
  <div class="test-payment">
    <div v-if="state === 'unpaid'">
      <div v-if="invoiceDetail !== undefined">
        <p>支払い金額: {{invoiceDetail.amount}}</p>
        <PaymentCheckoutForm :public-key="paymentPublicKey" @token-created-event="tokenCreatedEvent" @token-failed-event="tokenFailedEvent"></PaymentCheckoutForm>
      </div>
      <div v-if="invoiceDetail === undefined">
        <p>ロード中</p>
      </div>
    </div>
    <div v-if="state === 'complete'">
      <p>支払い金額: {{invoiceDetail.amount}}</p>
      <p>支払い完了</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import PaymentCheckoutForm from '@/components/PaymentCheckoutForm.vue'
import PaymentAPI from '@/lib/PaymentAPI'
import { InvoiceDetail, PaymentState } from '@/lib/Protocol'

@Component({
  components: {
    PaymentCheckoutForm
  }
})
export default class TestPayment extends Vue {
    private api: PaymentAPI = new PaymentAPI()
    private paymentPublicKey: string = process.env.VUE_APP_PAYMENT_PUB_KEY

    private invoiceDetail?: InvoiceDetail = {
      invoiceId: '',
      amount: 100,
      paid: false
    }

    private state: PaymentState = 'unpaid'

    async created () {
      this.invoiceDetail = await this.api.getInvoiceDetail(this.$route.params.invoiceId)
    }

    tokenCreatedEvent (token: string) {
      this.api
        .makePayment(this.$route.params.invoiceId, token)
        // eslint-disable-next-line no-return-assign
        .then(() => this.state = 'complete')
    }

    tokenFailedEvent (message: string) {
      alert(message)
    }
}
</script>
