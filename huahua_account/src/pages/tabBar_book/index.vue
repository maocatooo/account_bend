<template>
  <view>
    <at-list>
      <picker
        mode='date'
        fields="month"
        :value="pageData.dateSel"
        @change="handleDateChange"
      >
        <view class="at-list__item">
          <text class="">{{ pageData.dateSel }}</text>
          <text class="">{{ pageData.allAmount }}</text>
        </view>
      </picker>
      <picker
        mode='selector'
        rangeKey="name"
        :range="pageData.bookSelector"
        :value="pageData.bookSelectorValue"
        @change="this.handleBookChange"
        @cancel="this.handleCancel"
      >
        <at-list-item
          :title="pageData.bookSelector&&pageData.bookSelector.length>0? pageData.bookSelector[pageData.bookSelectorValue].name:''"
          hasBorder/>
      </picker>
      <view v-for="(item1, index) in pageData.data" :key="index">
        <view class="at-list__item day_info">
          <text>{{ item1.date }}</text>
          <text>{{ item1.allAmount.toString() }}</text>
        </view>
        <at-swipe-action
          v-for="(item, index) in item1.data"
          :key="item.id"
          :options="pageData.OPTIONS"
          :isOpened="item.isOpened"
          @click="(opt_item, key) => handleClicked(opt_item, item, key ,index)"
          @opened="(index) => handleSingle(item)"
        >
          <view class="at-list__item">
            <text class="j_name">{{ item.name }}</text>
            <text class="j_tname"> ({{ item.tname }})</text>
            <text class="j_date"> {{ formatTime(item.date) }}</text>
            <text class="j_amount"> {{ item.amount }}</text>
          </view>
        </at-swipe-action>
      </view>

    </at-list>
  </view>


</template>

<script lang="ts">
import {defineComponent, reactive} from "vue"
import Decimal from "decimal.js"
import Taro from "@tarojs/taro"
import "./index.scss"
import moment from "moment";
import {getBooks, OPTIONS, DEL_OPTION_INDEX, UPD_OPTION_INDEX} from "../../api/common";
import {AtListItem, AtList, AtSwipeAction} from 'taro-ui-vue3'
import {BookJournals, Books, DeleteBookJournal} from "../../api/api";

export default defineComponent({
  components: {
    AtListItem,
    AtSwipeAction,
    AtList
  },
  methods: {
    handleClicked(opt_item, item, key, index) {
      console.log(`点击了${opt_item.text}按钮，key: ${key}, index:${index}`)
      if (key === DEL_OPTION_INDEX) {
        console.log(item)
        DeleteBookJournal(item.id).then(() => {
          this.getBookJournals()
        })
      } else if (key === UPD_OPTION_INDEX) {

        console.log(`/pages/tabBar_edit/index?data=${JSON.stringify(item)}`)
        Taro.navigateTo({
          url: `/pages/tabBar_edit/index?data=${JSON.stringify(item)}`
        })
      }
    },
    async getBookJournals() {
      if (!this.pageData.bookSelector[0]) {
        return
      }
      const data: [] = await BookJournals({
        bookID: this.pageData.bookSelector[0].id,
        date: parseInt(moment(this.pageData.dateSel).format("x"))
      })
      this.setPageData_data(data)
    },
    handleBookChange(e) {
      this.pageData.bookSelectorValue = e.detail.value
    },
    async handleDateChange(e) {
      this.pageData.dateSel = e.detail.value
      await this.getBookJournals()
    },
    formatTime(t) {
      return moment(t).format('HH:mm')
    },
    setPageData_data(data) {
      this.pageData.data = []
      this.pageData.allAmount = new Decimal(0)
      if (!data || data.length === 0) {
        return
      }
      let c = {}
      data.map((obj) => {
        const date = moment(obj.date)
        const day = date.format("D")
        if (!c[day]) {
          c[day] = {
            data: [],
            allAmount: new Decimal(0),
            date: date.format("MM-DD"),
          };
          this.pageData.data.push(c[day])
        }
        c[day].data.push(obj);

        c[day].allAmount = c[day].allAmount.add(new Decimal(obj.amount));
        this.pageData.allAmount = this.pageData.allAmount.add(new Decimal(obj.amount));
        return obj
      })
    }
  },
  async onPullDownRefresh() {
    this.pageData.bookSelector = getBooks()
    await this.getBookJournals()
    Taro.stopPullDownRefresh()
  },
  async onShow() {
    console.log('onShow')
    this.pageData.bookSelector = await Books()
    await this.getBookJournals()
  },
  setup() {
    let pageData = reactive({
      data: [],
      OPTIONS,
      allAmount: new Decimal(0),
      bookSelectorValue: 0,
      bookSelector: [],
      dateSel: moment().format('YYYY-MM')
    })

    function handleSingle(item) {
      pageData.data.map((item1) => {
        item1.data = item1.data.map((item2) => {
          if (item2.id === item.id) {
            item2.isOpened = true
            return item2
          }
          item2.isOpened = false
          return item2
        })
        return item1
      })
    }

    return {
      pageData,
      handleCancel: (e) => {
        console.log('handleCancel', e)
      },
      handleSingle
    }
  }
})
</script>
