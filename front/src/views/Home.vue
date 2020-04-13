<template>
  <div class="home">
    <h1>お話しましょう〜</h1>
    <h2>メンター一覧</h2>
    <MentorTable :items="items"></MentorTable>
    <a href="/payment/test">決済テスト</a>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MentorTable from '@/components/MentorTable.vue'
import { AccountDetail } from '@/lib/Protocol'
import PaymentAPI from '@/lib/PaymentAPI'

@Component({
  components: {
    MentorTable
  }
})
export default class MyPage extends Vue {
    private api: PaymentAPI = new PaymentAPI()
    private items: Array<AccountDetail> = []
    async created () {
      this.items = await this.api.getAllMentor()
    }
}
</script>

<style scoped>
</style>
