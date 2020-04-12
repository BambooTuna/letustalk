<template>
  <div class="home">
    <PaymentCheckoutForm :public-key="paymentPublicKey" @token-created-event="tokenCreatedEvent" @token-failed-event="tokenFailedEvent"></PaymentCheckoutForm>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import PaymentCheckoutForm from '@/components/PaymentCheckoutForm.vue'

@Component({
  components: {
    PaymentCheckoutForm
  }
})
export default class MyPage extends Vue {
    public paymentPublicKey: string = process.env.VUE_APP_PAYMENT_PUB_KEY

    tokenCreatedEvent (token: string) {
      alert(token)
    }

    tokenFailedEvent (message: string) {
      alert(message)
    }
}
</script>
