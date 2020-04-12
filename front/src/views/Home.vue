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
import { MentorDetail } from '@/lib/Protocol'

@Component({
  components: {
    MentorTable
  }
})
export default class MyPage extends Vue {
    private items: Array<MentorDetail> = []
    created (): void {
      this.items = [
        {
          mentorId: '1',
          name: 'A'
        },
        {
          mentorId: '2',
          name: 'B'
        },
        {
          mentorId: '3',
          name: 'C'
        },
        {
          mentorId: '4',
          name: 'D'
        },
        {
          mentorId: '5',
          name: 'E'
        }
      ]
    }
}
</script>

<style scoped>
</style>
