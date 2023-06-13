<template>
  <view>
    <at-list>
      <picker
        mode='date'
        fields="month"
        :value="dateSel"
        @change="handleDateChange"
      >
        <at-list-item :title="dateSel" hasBorder/>
      </picker>

      <at-swipe-action
        v-for="(item, index) in list"
        :key="item.title"
        :options="OPTIONS"
        :isOpened="item.isOpened"
        @click="(item, key) => handleClicked(item, key ,index)"
        @opened="handleSingle(index)"
      >
        <at-list-item :title="item.title" hasBorder/>
      </at-swipe-action>
    </at-list>
  </view>


</template>

<script lang="ts">
import {defineComponent, reactive, ref, toRefs} from "vue"
import Taro from "@tarojs/taro"
import "./index.scss"

export default defineComponent({
  name: "SwipeActionDemo22",
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
        text: '详情',
        style: {
          backgroundColor: '#E93B3D'
        }
      }
    ])
    const dateSel = ref('2022-06')
    const state = reactive({
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

    function handleClicked(item, key, index) {
      showToast(`点击了${item.text}按钮，key: ${key}, index:${index}`)
      console.log(key === 1)
      if (key === 1) {
      } else {
        state.list = state.list.filter((_item, key) => key !== index)
      }
    }

    function handleSingle(index) {
      state.list = state.list.map((item, key) => {
        item.isOpened = key === index
        return item
      })
    }
    function handleDateChange(e) {
      dateSel.value = e.detail.value
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
      dateSel,
      handleClick,
      handleDateChange,
      handleClicked,
      handleSingle
    }
  }
})
</script>
