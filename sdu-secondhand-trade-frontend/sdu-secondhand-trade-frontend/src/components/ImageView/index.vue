<script setup>
import { useMouseInElement } from '@vueuse/core';
import { onMounted, ref, watch } from 'vue';

// 图片列表
const props = defineProps({
    imageList: {
        type: Array,
        default: () => []
    }
})

const activeIndex = ref(0)
const enterHandler = (i) => {
    activeIndex.value = i
}

const target = ref(null)
const { elementX, elementY, isOutside } = useMouseInElement(target)

const left = ref(0)
const top = ref(0)

const positionX = ref(0)
const positionY = ref(0)

watch([elementX, elementY, isOutside], () => {
    // 优化性能：如果鼠标没有移入到盒子里面 直接不执行后面的逻辑
    if (isOutside.value) {
        return
    }

    // 有效范围内控制滑块距离
    // 横向
    if (elementX.value > 125 && elementX.value < 375) {
        left.value = elementX.value - 125
    }
    // 纵向
    if (elementY.value > 125 && elementY.value < 375) {
        top.value = elementY.value - 125
    }

    // 处理边界
    if (elementX.value > 375) {
        left.value = 250
    }
    if (elementX.value < 125) {
        left.value = 0
    }

    if (elementY.value > 375) {
        top.value = 250
    }
    if (elementY.value < 125) {
        top.value = 0
    }

    // 控制大图的显示
    positionX.value = -left.value * 2
    positionY.value = -top.value * 2
})

const backgroundSize = ref("1000px 1000px")
const setLargeImageStyle = (imageSrc) => {
    const img = new Image();
    img.src = imageSrc;

    img.onload = () => {
        const imgWidth = img.width;
        const imgHeight = img.height;

        // 计算长边并设置背景大小
        const maxDimension = 1000; // 最大长边

        if (imgWidth > imgHeight) {
            // 如果宽度更长
            backgroundSize.value = `${maxDimension}px auto`; // 长边是宽度
        } else {
            // 如果高度更长
            backgroundSize.value = `auto ${maxDimension}px`; // 长边是高度
        }
    };
};

// 监听图片切换
watch(activeIndex, (newIndex) => {
    setLargeImageStyle(props.imageList[newIndex].url);
});

onMounted(() => {
    setLargeImageStyle(props.imageList[0].url)
})
</script>


<template>
    <div class="goods-image">
        <!-- 左侧大图-->
        <div class="middle" ref="target">
            <img class="middle-image" :src="imageList[activeIndex].url" alt=""/>
            <!-- 蒙层小滑块 -->
            <div class="layer" :style="{ left: `${left}px`, top: `${top}px` }" v-show="!isOutside"></div>
        </div>
        <!-- 小图列表 -->
        <ul class="small">
            <li v-for="(img, i) in imageList" :key="i" @mouseenter="enterHandler(i)"
                :class="{ active: i === activeIndex }">
                <img :src="img.url" alt="" />
            </li>
        </ul>
        <!-- 放大镜大图 -->
        <div class="large" :style="[
            {
                backgroundImage: `url(${imageList[activeIndex].url})`,
                backgroundPositionX: `${positionX}px`,
                backgroundPositionY: `${positionY}px`,
                backgroundSize: `${backgroundSize}`
            },
        ]" v-show="!isOutside"></div>
    </div>
</template>

<style scoped lang="scss">
.goods-image {
    width: 580px;
    height: 500px;
    position: relative;
    display: flex;

    .middle {
        width: 100%;
        height: 100%;
        background: #f8f8f8;
        position: relative;
    }

    .middle-image {
        width: 100%;
        height: 100%;
        object-fit: contain;
        object-position: left top;
        background: #f8f8f8;
    }

    .large {
        position: absolute;
        top: 0;
        left: 512px;
        width: 500px;
        height: 500px;
        z-index: 500;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
        background-repeat: no-repeat;
        background-color: #f8f8f8;
    }

    .layer {
        width: 250px;
        height: 250px;
        background: rgba(0, 0, 0, 0.2);
        // 绝对定位
        left: 0;
        top: 0;
        position: absolute;
    }

    .small {
        width: 80px;

        li {
            width: 68px;
            height: 68px;
            margin-left: 12px;
            margin-bottom: 15px;
            cursor: pointer;
            background: #f8f8f8;

            &:hover,
            &.active {
                border: 2px solid $mainColor;
            }
        }
    }
}
</style>