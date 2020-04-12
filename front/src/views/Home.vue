<template>
  <div class="home">
    <h1>お話しましょう〜</h1>
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
