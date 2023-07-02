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
      <picker
        mode='selector'
        rangeKey="name"
        :range="bookSelector"
        :value="bookSelectorValue"
        @change="handleBookChange"
        @cancel="handleCancel"
      >
        <at-list-item :title="bookSelector[bookSelectorValue].name" hasBorder/>
      </picker>
      <at-swipe-action
        v-for="(item, index) in bookJournals"
        :key="item.id"
        :options="OPTIONS"
        :isOpened="item.isOpened"
        @click="(item, key) => handleClicked(item, key ,index)"
        @opened="handleSingle(index)"
      >
        <at-list-item :title="item.tname" hasBorder/>
      </at-swipe-action>
    </at-list>
  </view>


</template>

<script lang="ts">
import {defineComponent, reactive, ref, onMounted} from "vue"
import Taro from "@tarojs/taro"
import "./index.scss"
import { getBooks } from "../../api/common";
import {BookJournals} from "../../api/api";

export default defineComponent({
  name: "SwipeActionDemo22",
  async  mounted() {
    console.log('onMounted')
    const data:[] = await  BookJournals(this.bookSelector[0].id)
    data.map((item) => {
      item.isOpened = false
      return item
    })
    console.log(data)
    this.bookJournals.push(...data)
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
        text: '详情',
        style: {
          backgroundColor: '#E93B3D'
        }
      }
    ])
    const dateSel = ref('2022-06')
    const bookSelector = getBooks()
    let bookJournals =  reactive([])

    function handleClick(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}`)
      console.log(arguments)
    }

    function handleClicked(item, key, index) {
      showToast(`点击了${item.text}按钮，key: ${key}, index:${index}`)
      console.log(key === 1)
      if (key === 1) {
      } else {
        bookJournals = bookJournals.filter((_item, key) => key !== index)
      }
    }

    function handleSingle(index) {
      bookJournals = bookJournals.map((item, key) => {
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
    const bookSelectorValue = ref(0)
    return {
      bookJournals,
      OPTIONS,
      dateSel,
      bookSelector,
      bookSelectorValue,
      handleClick,
      handleDateChange,
      handleClicked,
      handleBookChange:(e)=>{
        bookSelectorValue.value = e.detail.value
      },
      handleCancel:(e)=>{
        console.log('handleCancel', e)
      },
      handleSingle
    }
  }
})
</script>
