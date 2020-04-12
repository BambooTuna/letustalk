<template>
  <div class="home">
    <ul>
      <li v-for="row in items" :key="row.mentorId">
        <p><router-link :to=" '/mentor/' + row.mentorId + '/calendar'">{{row.name}}</router-link></p>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import { MentorDetail } from '@/lib/Protocol'

@Component
export default class MentorTable extends Vue {
    @Prop()
    private items!: Array<MentorDetail>
}

</script>

<style scoped>
  ul, ol {
    padding: 0;
    position: relative;
  }
  ul li, ol li {
    color: #000000;
    border-left: solid 6px #000000;/*左側の線*/
    border-right: solid 6px #000000;/*左側の線*/
    border-bottom: solid 6px #000000;/*左側の線*/
    border-top: solid 6px #000000;/*左側の線*/
    background: rgba(241, 248, 255, 0.26);/*背景色*/
    margin-bottom: 3px;/*下のバーとの余白*/
    line-height: 1.5;
    padding: 0.5em;
    list-style-type: none!important;/*ポチ消す*/
  }
</style>
