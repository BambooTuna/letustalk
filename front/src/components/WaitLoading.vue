<template>
    <div class="wait_loading">
        <div v-show="loadingFlag" class="loader"></div>
        <div v-show="!loadingFlag" class="itemContainer">
            <slot></slot>
        </div>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator'

@Component
export default class WaitLoading extends Vue {
    @Prop() public loadingFlag!: boolean
}
</script>

<style scoped>
    .loader {
        margin: 100px auto;
        font-size: 25px;
        width: 1em;
        height: 1em;
        border-radius: 50%;
        position: relative;
        text-indent: -9999em;
        -webkit-animation: load5 1.1s infinite ease;
        animation: load5 1.1s infinite ease;
        -webkit-transform: translateZ(0);
        -ms-transform: translateZ(0);
        transform: translateZ(0);
    }
    @-webkit-keyframes load5 {
        0%,
        100% {
            box-shadow: 0em -2.6em 0em 0em #0433ff, 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.5), -1.8em -1.8em 0 0em rgba(4,51,255, 0.7);
        }
        12.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.7), 1.8em -1.8em 0 0em #0433ff, 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.5);
        }
        25% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.5), 1.8em -1.8em 0 0em rgba(4,51,255, 0.7), 2.5em 0em 0 0em #0433ff, 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        37.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.5), 2.5em 0em 0 0em rgba(4,51,255, 0.7), 1.75em 1.75em 0 0em #0433ff, 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        50% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.5), 1.75em 1.75em 0 0em rgba(4,51,255, 0.7), 0em 2.5em 0 0em #0433ff, -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        62.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.5), 0em 2.5em 0 0em rgba(4,51,255, 0.7), -1.8em 1.8em 0 0em #0433ff, -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        75% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.5), -1.8em 1.8em 0 0em rgba(4,51,255, 0.7), -2.6em 0em 0 0em #0433ff, -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        87.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.5), -2.6em 0em 0 0em rgba(4,51,255, 0.7), -1.8em -1.8em 0 0em #0433ff;
        }
    }
    @keyframes load5 {
        0%,
        100% {
            box-shadow: 0em -2.6em 0em 0em #0433ff, 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.5), -1.8em -1.8em 0 0em rgba(4,51,255, 0.7);
        }
        12.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.7), 1.8em -1.8em 0 0em #0433ff, 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.5);
        }
        25% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.5), 1.8em -1.8em 0 0em rgba(4,51,255, 0.7), 2.5em 0em 0 0em #0433ff, 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        37.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.5), 2.5em 0em 0 0em rgba(4,51,255, 0.7), 1.75em 1.75em 0 0em #0433ff, 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        50% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.5), 1.75em 1.75em 0 0em rgba(4,51,255, 0.7), 0em 2.5em 0 0em #0433ff, -1.8em 1.8em 0 0em rgba(4,51,255, 0.2), -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        62.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.5), 0em 2.5em 0 0em rgba(4,51,255, 0.7), -1.8em 1.8em 0 0em #0433ff, -2.6em 0em 0 0em rgba(4,51,255, 0.2), -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        75% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.5), -1.8em 1.8em 0 0em rgba(4,51,255, 0.7), -2.6em 0em 0 0em #0433ff, -1.8em -1.8em 0 0em rgba(4,51,255, 0.2);
        }
        87.5% {
            box-shadow: 0em -2.6em 0em 0em rgba(4,51,255, 0.2), 1.8em -1.8em 0 0em rgba(4,51,255, 0.2), 2.5em 0em 0 0em rgba(4,51,255, 0.2), 1.75em 1.75em 0 0em rgba(4,51,255, 0.2), 0em 2.5em 0 0em rgba(4,51,255, 0.2), -1.8em 1.8em 0 0em rgba(4,51,255, 0.5), -2.6em 0em 0 0em rgba(4,51,255, 0.7), -1.8em -1.8em 0 0em #0433ff;
        }
    }
</style>
