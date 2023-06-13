<template>
  <at-list>
    <at-swipe-action
      v-for="(item, index) in list"
      :key="item.title"
      :options="OPTIONS"
      :isOpened="item.isOpened"
      @click="handleClicked"
      @opened="handleSingle(index)"
    >
      <at-list-item :title="item.title"  hasBorder />
    </at-swipe-action>
  </at-list>
</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs} from "vue"
import Taro from "@tarojs/taro"
import "./index.scss"

export default defineComponent({
  name: "SwipeActionDemo3",
  setup() {
    const OPTIONS = ref([
      {
        text: '删除',
        style: {
          color: '#333',
          backgroundColor: '#F7F7F7'
        }
      },
      {
        text: '确认',
        style: {
          backgroundColor: '#E93B3D'
        }
      }
    ])
    const currentIndex = ref(0)
    const state = reactive({
      isOpened2: false,
      list: [
        {
          title: 'item1',
          isOpened: false,
        },
        {
          title: 'item2',
          isOpened: false,
        },
        {
          title: 'item3',
          isOpened: false,
        },
        {
          title: 'item4',
          isOpened: false,
        },
        {
          title: 'item5',
          isOpened: false,
        },
        {
          title: 'item6',
          isOpened: false,
        }
      ]
    })
    function handleClick(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}`)
      console.log(arguments)
    }
    function handleClicked(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}, index:${currentIndex.value}`)
      if (item.text == "确认") {
      }else{
        state.list = state.list.filter((_item, key) => key !== currentIndex.value)
      }
    }

    function handleSingle(index) {
      currentIndex.value = index
      state.list = state.list.map((item, key) => {
        item.isOpened = key === index
        return item
      })
    }


    function showToast(titleMsg) {
      Taro.showToast({
        icon: 'none',
        title: titleMsg
      })
    }
    return {
      ...toRefs(state),
      OPTIONS,
      handleClick,
      handleClicked,
      handleSingle
    }
  }
})
</script>
