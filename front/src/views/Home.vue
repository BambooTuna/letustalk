<template>
  <div class="home">
    <header>
      <div class="title">
        <div class="inner-text">
          〇〇
        </div>
      </div>
      <div class="content">
        <router-link to="/Content" class="inner-text">
        何ができるの
        </router-link>
      </div>
      <div class="for-host">
        <a href="https://forms.gle/87DD7SDCqxDCSyPW6" class="inner-text">
          相談されたい人へ
        </a>
      </div>
      <div class="contact">
        <a href="https://forms.gle/E4LzaHD26aRC3cmDA" class="inner-text">
          問い合わせ
        </a>
      </div>
      <div class="search">
        <div class="inner-box">
          <a href="https://forms.gle/GupmVNu3pRcNXWQS8" class="s-box">
          相談されたい人を探す
          </a>
        </div>
      </div>
    </header>
    <div>
      <h2>〇〇は、相談したい人と相談されたい人を繋ぐ場所です</h2>
      <p>誰しもが悩みを抱えています。でもそれでよくないですか？</p>
      <p></p>
      <div class="find-therapist">
        <div class="inner-box">
          <a href="https://forms.gle/GupmVNu3pRcNXWQS8" class="box-inner-text">
            あなたから相談されたい人を見つける
          </a>
        </div>
      </div>
    </div>
    <!-- <MentorTable :items="items"></MentorTable>
    <a href="/payment/test">決済テスト</a> -->
    <div>
      <p class="typ">私たちにできることはただ一つだけ</p>
      <p>あなたの1番の味方になることです</p>
      <p>まずはあなたの声を聞かせてください</p>
    </div>
    <div class="arrowWrap">
    <div class="arrow">
      <span></span>
    </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import MentorTable from '@/components/MentorTable.vue'
import { AccountDetail } from '@/lib/Protocol'
import RestAPI from '@/lib/RestAPI'

@Component({
  components: {
    MentorTable
  }
})
export default class MyPage extends Vue {
    private api: RestAPI = new RestAPI()
    private items: Array<AccountDetail> = []
    async created () {
      this.items = await this.api.getAllMentor()
    }
}
</script>

<style scoped>
  a {
    text-decoration: none;
  }
  a.inner-text {
    font-size: 12px;
    padding: 0.2em 0.2em;
    color: black;
  }
  a.box-inner-text {
    font-size: 12px;
    padding: 0.2em 0.2em;
    color: white;
  }
  h2 {
    color: #479E9E;
  }
  .title {
    color: #479E9E;
    font-size: 30px;
  }
  header {
    display: flex;
  }
  .inner-text {
    margin: 10px;
  }
  .search {
    margin: auto;
    border-radius: 5px;
    border: solid 1px #479E9E;
    font-size: 12px;
  }
  .content {
    margin: auto;
    padding: 0.2em 0.2em;
    font-size: 12px;
  }
  .for-host {
    margin: auto;
    padding: 0.2em 0.2em;
    font-size: 12px;
  }
  .inner-box {
    margin: 5px;
    padding: 0.5em 0.5em;
  }
  a.s-box {
    margin: 12px;
    padding: 0.5em 0.5em;
    color: #479E9E;
  }
  .contact {
    margin: auto;
    padding: 0.2em 0.2em;
    font-size: 12px;
  }
  .find-therapist {
    border: solid 1px #479E9E;
    margin: auto;
    background-color: #479E9E;
    color: white;
    border-radius: 5px;
    width: 250px;
    height: 30px;
  }
  /* アニメーション設定 */
  .arrowWrap {
    margin-top: 30vh;
  }

  .arrow {
    width: 15px;
    margin: 0 auto;
    -webkit-animation: arrow 0.8s cubic-bezier(0.215, 0.61, 0.355, 1) 0s infinite alternate;
    animation: arrow 0.8s cubic-bezier(0.215, 0.61, 0.355, 1) 0s infinite alternate;
  }

  .arrow span {
    position: relative;
    display: block;
    left: 50%;
  }

  .arrow span:before {
    content: '';
    width: 15px;
    height: 15px;
    border: 0;
    border-top: solid 1px #333;
    border-right: solid 1px #333;
    -webkit-transform: rotate(135deg);
    transform: rotate(135deg);
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    margin: auto;
    box-sizing: border-box;
  }

  @keyframes arrow {
    0% {
      opacity: 1;
      -webkit-transform: translate3d(-50%, 0, 0);
      transform: translate3d(-50%, 0, 0);
    }

    100% {
      opacity: 0.5;
      -webkit-transform: translate3d(-50%, -8px, 0);
      transform: translate3d(-50%, -8px, 0);
    }
  }
  .introduction {
    margin-top: 300px;
  }
</style>
