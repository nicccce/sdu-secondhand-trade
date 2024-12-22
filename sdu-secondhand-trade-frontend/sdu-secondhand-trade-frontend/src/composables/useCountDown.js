import { computed, onUnmounted, ref } from 'vue'
import dayjs from 'dayjs'

export const useCountDown = () => {
    // 1. 响应式的数据
    let timer = null
    const time = ref(0)

    // 格式化时间为 xx分xx秒
    const formatTime = computed(() => {
        const minutes = Math.floor(time.value / 60)
        const seconds = time.value % 60
        return `${String(minutes).padStart(2, '0')}分${String(seconds).padStart(2, '0')}秒`
    })

    // 2. 开启倒计时的函数
    const start = (targetTime) => {
        // 将目标时间转为 Unix 时间戳（秒）
        const targetDate = dayjs(targetTime).add(10, 'minute')
        const currentTime = dayjs()

        // 计算倒计时秒数（目标时间 - 当前时间）
        time.value = targetDate.diff(currentTime, 'second')

        // 判断是否倒计时结束
        if (time.value <= 0) {
            time.value = 0
            return
        }

        // 每隔1秒更新一次
        timer = setInterval(() => {
            time.value--

            // 如果倒计时结束，清除定时器
            if (time.value <= 0) {
                clearInterval(timer)
                time.value = 0
            }
        }, 1000)
    }

    // 组件销毁时清除定时器
    onUnmounted(() => {
        timer && clearInterval(timer)
    })

    return {
        formatTime,
        start
    }
}