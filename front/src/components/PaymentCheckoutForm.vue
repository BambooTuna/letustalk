<template>
  <div class="payment-checkout-form">
    <PayjpCheckout
      :api-key="publicKey"
      text="カードを情報を入力して購入"
      submit-text="購入確定"
      name-placeholder="田中 太郎"
      v-on:created="onTokenCreated"
      v-on:failed="onTokenFailed">
    </PayjpCheckout>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Emit } from 'vue-property-decorator'
import PayjpCheckout from 'vue-payjp-checkout'

@Component({
  components: {
    PayjpCheckout
  }
})
export default class PaymentCheckoutForm extends Vue {
  @Prop()
  private publicKey!: string

  @Emit()
  // eslint-disable-next-line @typescript-eslint/no-empty-function,@typescript-eslint/no-unused-vars
  tokenCreatedEvent (token: string) {
  }

  @Emit()
  // eslint-disable-next-line @typescript-eslint/no-empty-function,@typescript-eslint/no-unused-vars
  tokenFailedEvent (message: string) {
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  onTokenCreated (res: any) {
    this.tokenCreatedEvent(res.id)
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  onTokenFailed (status: string, err: string) {
    this.tokenFailedEvent(status)
  }
}
</script>
