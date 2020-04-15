<template>
  <div class="month-calendar-table">
    <div class="calender-component">
      <div class="calender-header">
        <div class="arrow" v-on:click="changeMonth(0)"> ＜ </div>
        <div class="current-date">{{ dateLabel }}</div>
        <div class="arrow" v-on:click="changeMonth(1)"> ＞ </div>
      </div>
      <div class="calender-body">
        <table class="calender-panel-list">
          <tr>
            <th class="calender-panel_space">日</th>
            <th class="calender-panel_space">月</th>
            <th class="calender-panel_space">火</th>
            <th class="calender-panel_space">水</th>
            <th class="calender-panel_space">木</th>
            <th class="calender-panel_space">金</th>
            <th class="calender-panel_space">土</th>
          </tr>
          <tr v-for="k in (Math.floor(calendarDates.length / 7) + 1)" :key="k">
            <td v-for="data in dropArray(k * 7 - 7)" :key="data.date">
              <div class="calender-panel_space" v-if="data === 0"></div>
              <div class="calender-panel"  v-on:click="selectDate(getDateString(data.date))" v-bind:class="toDay === data.date ? 'selected' : ''" v-if="data !== 0">
                <div class="calender-date">{{ data.dateNumber }}</div>
                <div class="calender-schedule">{{ (data.scheduleNumber > 0) ? '○' : '×' }}</div>
              </div>
            </td>
          </tr>
        </table>
      </div>
      <div class="calender-footer">
        <div class="calender-footer_todo" v-for="schedule in items" v-show="getDateString(schedule.from) === selectedDate" :key="schedule.scheduleId">
          <div class="calender-footer_todo-period">時間: {{ getDateDetailString(schedule.from) }} 〜 {{ getDateDetailString(schedule.to) }}</div>
          <div class="calender-footer_todo-unitPrice">費用: {{ schedule.unitPrice }} 円</div>
          <button @click="onClick(schedule.scheduleId)">予約する</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Emit, Watch } from 'vue-property-decorator'
import { FreeSchedule } from '@/lib/Protocol'
import moment, { Moment } from 'moment'

@Component
export default class MonthCalendarTable extends Vue {
    @Prop()
    public items!: Array<FreeSchedule>

    public calendarDates: Array<CalendarDates | number> = []
    public dateLabel = ''
    public selectedMonth: Moment = moment()
    public selectedDate = moment().format('YYYY-MM-DD')
    public toDay: string = this.selectedDate

    selectDate (id: string) {
      this.selectedDate = id
    }

    changeMonth (num: number) {
      if (num === 0) {
        this.selectedMonth = moment(this.selectedMonth).subtract(1, 'months')
      } else {
        this.selectedMonth = moment(this.selectedMonth).add(1, 'months')
      }
    }

    created () {
      this.selectedMonth = moment()
      this.toDay = moment().format('YYYY-MM-DD')
      this.selectedDate = this.toDay
    }

    dropArray (n: number): Array<CalendarDates | number> {
      const arr = this.calendarDates
      const start = n + 1
      const end = (start + 7 > arr.length) ? arr.length : start + 7
      return arr.slice(start, end)
    }

    getDateString (v: string): string {
      return moment(v).format('YYYY-MM-DD')
    }

    getDateDetailString (v: string): string {
      return moment(v).format('YYYY-MM-DD HH:mm:ss')
    }

    @Watch('items')
    itemsUpdated () {
      this.watchSelectedMonth()
    }

    @Watch('selectedMonth')
    watchSelectedMonth () {
      this.dateLabel = moment(this.selectedMonth).format('YYYY年MM月')
      this.calendarDates = []
      for (let i = 0; i < moment(this.selectedMonth).startOf('month').day(); i++) {
        this.calendarDates[i] = 0
      }
      for (let i = 0; i < moment(this.selectedMonth).daysInMonth(); i++) {
        const date = moment(this.selectedMonth).startOf('month').add(i, 'day').format('YYYY-MM-DD')
        let scheduleNumber = 0
        for (const k in this.items) {
          if (this.getDateString(this.items[k].from) === date) {
            scheduleNumber++
          }
        }
        this.calendarDates.push({
          date: date,
          dateNumber: i + 1,
          scheduleNumber: scheduleNumber
        })
      }
    }

    @Emit()
    // eslint-disable-next-line @typescript-eslint/no-empty-function,@typescript-eslint/no-unused-vars
    reservationEvent (scheduleId: string) {
    }

    onClick (scheduleId: string) {
      this.reservationEvent(scheduleId)
    }
}

type CalendarDates = {
    date: string;
    dateNumber: number;
    scheduleNumber: number;
}

</script>

<style scoped>
  .calender {
    width: 336px;
  }

  .calender-component {
    min-height: 320px;
    border: solid 1px gray;
    padding: 24px;
    box-sizing: border-box;
    width: 100%;
  }

  .calender-header {
    display: flex;
    justify-content: space-between;
  }

  .arrow {
    font-size: 16px;
    color: gray;
    user-select: none;
    cursor: pointer;
    text-align: center;
    width: 24px;
    height: 24px;
  }

  .arrow:hover {
    background-color: silver;
    border-radius: 4px;
  }

  .current-date {
    color: gray;
    user-select: none;
  }

  .calender-body {
    margin-top: 24px;
  }

  .calender-panel-list {
    /*display: flex;*/
    /*flex-wrap: wrap;*/
    padding: 24px;
    justify-content: center;
    width: 100%;
  }

  .calender-panel {
    padding: 8px 0px;
    text-align: center;
    color: gray;
    font-size: 14px;
    user-select: none;
    list-style: none;
  }

  .calender-panel:hover {
    background-color: silver;
    border-radius: 4px;
  }

  .calender-date {
    cursor: pointer;
  }

  .selected {
    background-color: #f0f0f0;
    border-radius: 4px;
  }

  .calender-schedule {
    cursor: pointer;
    text-align:center;
    margin: 0px 8px;
    line-height: 25px;
  }

  .calender-panel_space {
    width: 40px;
    list-style: none;
  }

  .calender-footer_todo {
    border-top: 1px gray solid;
    padding-top: 12px;
    margin-top: 8px;
  }

  .calender-footer_todo-period {
    font-size: 18px;
    color: gray;
    word-break: break-word;
    text-decoration: underline;
  }

  .calender-footer_todo-unitPrice {
    font-size: 20px;
    color: #ff5336;
    word-break: break-word;
  }

  button {
    display: inline-block;
    max-width: 180px;
    text-align: left;
    background-color: #ff3308;
    font-size: 16px;
    color: #FFF;
    text-decoration: none;
    font-weight: bold;
    padding: 10px 24px;
    border-radius: 4px;
    border-bottom: 4px solid #d30e01;
  }
  button:active {
    transform: translateY(4px);
    border-bottom: none;
  }
</style>
