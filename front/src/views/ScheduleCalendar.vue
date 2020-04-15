<template>
  <div class="schedule-calendar">
    <MonthCalendarTable :items=freeSchedule @reservation-event="reservationEvent"></MonthCalendarTable>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { FreeSchedule } from '@/lib/Protocol'
import MonthCalendarTable from '@/components/MonthCalendarTable.vue'
import RestAPI from '@/lib/RestAPI'

@Component({
  components: {
    MonthCalendarTable
  }
})
export default class ScheduleCalendar extends Vue {
    private api: RestAPI = new RestAPI()
    private freeSchedule: Array<FreeSchedule> = []

    async created () {
      this.freeSchedule = await this.api.getFreeSchedule(this.$route.params.accountId)
    }

    reservationEvent (scheduleId: string) {
      alert('scheduleId: ' + scheduleId + 'を予約しようとしています')
    }
}

</script>

<style scoped>
</style>
