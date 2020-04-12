<template>
  <div class="home">
    <PaymentCheckoutForm public-key="pk_test_6acc32ac21aa1b0f55d5e3b8"></PaymentCheckoutForm>
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
}
</script>
