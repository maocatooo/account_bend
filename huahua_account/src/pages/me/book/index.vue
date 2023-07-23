<template>
  <at-list>
    <at-swipe-action
      v-for="(item, index) in list"
      :key="item.name"
      :options="OPTIONS"
      :isOpened="item.isOpened"
      @click="handleClicked"
      @opened="handleSingle(index)"
    >
      <at-list-item :title="item.name"  hasBorder />
    </at-swipe-action>
  </at-list>

</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs} from "vue"
import Taro from "@tarojs/taro"
import "./index.scss"
import { getBooks } from "../../../api/common";
import { AtListItem, AtList, AtSwipeAction } from 'taro-ui-vue3'

export default defineComponent({
  components: {
    AtListItem,
    AtSwipeAction,
    AtList
  },
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
    console.log(getBooks())
    const books = getBooks().map((item) => {
      item.isOpened = false
      return item
    })

    const currentIndex = ref(0)
    const state = reactive({
      list: books
    })
    function handleClick(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}`)
      console.log(arguments)
    }
    function handleClicked(item, key) {
      console.log("i am me book")
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
