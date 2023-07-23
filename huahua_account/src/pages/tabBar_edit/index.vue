<template>
  <view>
    <AtForm>
      <picker
        mode='selector'
        rangeKey="name"
        :range="fromData.bookSelector"
        :value="fromData.bookSelectorValue"
        @change="handleBookChange"
        @cancel="handleCancel"
      >
        <view class="at-input at-input__container">
          <text class="at-input__title">账本</text>
          <text class="at-input__input">{{ fromData.bookSelector[fromData.bookSelectorValue].name }}</text>
        </view>

<!--        <AtInput-->
<!--          name='value'-->
<!--          title='账本'-->
<!--          type='text'-->
<!--          placeholder='账本'-->
<!--          :value=""-->
<!--        />-->
      </picker>
      <picker
        mode='selector'
        rangeKey="name"
        :range="fromData.tagSelector"
        :value="fromData.tagSelectorValue"
        @change="handleTagChange"
        @cancel="handleCancel"
      >
        <view class="at-input at-input__container">
          <text class="at-input__title">标签</text>
          <text class="at-input__input">{{ fromData.tagSelector[fromData.tagSelectorValue].name }}</text>
        </view>
<!--        <AtInput-->
<!--          name='value'-->
<!--          title='标签'-->
<!--          type='text'-->
<!--          placeholder='标签'-->
<!--          :value="fromData.tagSelector[fromData.tagSelectorValue].name"-->
<!--        />-->
      </picker>

      <AtInput
        name='value'
        title='消费名称'
        type='text'
        placeholder='消费名称'
        v-model:value="fromData.name"
      />
      <AtInput
        name='value'
        title='金额'
        type='text'
        placeholder='金额(小数保持两位)'
        v-model:value="fromData.amount"
      />
      <AtInput
        name='value'
        title='备注'
        type='text'
        placeholder='备注'
        v-model:value="fromData.record"
      />

      <timePickerColumn
        :start-time="dateSelector.start_time"
        :end-time="dateSelector.now_time"
        :default-time="dateSelector.now_time"
        @result="onResult"
      >
        <view class="at-input at-input__container">
          <text class="at-input__title">时间</text>
          <text class="at-input__input">{{ fromData.date }}</text>
        </view>
      </timePickerColumn>

      <AtButton formType='submit' @click="onSubmit">提交</AtButton>
    </AtForm>

  </view>
</template>
<script>
import {reactive, ref} from 'vue'
import './index.scss'
import Taro from "@tarojs/taro"
import {getTags,getBooks} from "../../api/common";
import {CreateBookJournal, UpdateBookJournal} from "../../api/api";
import moment from "moment";
import timePickerColumn from "../../components/picker/index";
import { AtButton, AtForm, AtInput } from 'taro-ui-vue3'
export default {
  onPullDownRefresh (){
    this.fromData.tagSelector = getTags()
    this.fromData.bookSelector = getBooks()
    console.log('onPullDownRefresh ok')
    Taro.stopPullDownRefresh()
  },
  onShow() {
    const strData = Taro.getCurrentInstance().router.params["data"]
    if (strData){
        const data = JSON.parse(strData)
        this.fromData.name = data.name
        this.fromData.jid = data.id
        this.fromData.amount = data.amount
        this.fromData.record = data.record
        this.fromData.bookSelector.map((item, index)=>{
          if (item.id === data.bookID){
            this.fromData.bookSelectorValue = index
          }
        })
        this.fromData.tagSelector.map((item, index)=>{
          if (item.id === data.tid){
            this.fromData.tagSelectorValue = index
          }
        })
        this.fromData.date = moment(data.date).format('YYYY-MM-DD HH:mm')
    }
  },
  components: {
    timePickerColumn,
    AtForm,
    AtButton,
    AtInput
  },
  setup () {
    const now = new Date()
    const dateSelector = reactive({
      start_time: new Date(now.getFullYear()-1, now.getMonth(), now.getDate(), 0, 0, 0),
      now_time: new Date(now.getFullYear(), now.getMonth(), now.getDate(), 0, 0, 0),
    })

    const tagSelector = getTags()
    const bookSelector = getBooks()
    const fromData = reactive({
      jid : "",
      tagSelector: tagSelector,
      tagSelectorValue: 0,
      bookSelector: bookSelector,
      bookSelectorValue: 0,
      name: "",
      amount: '',
      record: '',
      // selectorKey: 'name',
      date: moment(new Date()).format('YYYY-MM-DD HH:mm'),
    })
    return {
      fromData,
      dateSelector,
      handleBookChange:(e)=>{
        fromData.bookSelectorValue = e.detail.value
      },
      handleTagChange:(e)=>{
        fromData.tagSelectorValue = e.detail.value
      },
      handleCancel:(e)=>{
        console.log('handleCancel', e)
      },
      async onSubmit (event) {
        if (fromData.amount === '' || fromData.name === ''){
          Taro.showToast({
            icon: 'none',
            title: "请填写完整"
          })
          return
        }
        const res  = await UpdateBookJournal({
          id: fromData.jid,
          tid: tagSelector[fromData.tagSelectorValue].id,
          tname: tagSelector[fromData.tagSelectorValue].name,
          amount: fromData.amount,
          name: fromData.name,
          record: fromData.record,
          bookID: bookSelector[fromData.bookSelectorValue].id,
          date: parseInt(moment(fromData.date).format("x"))
        })
        Taro.showToast({
          icon: 'none',
          title: "修改成功"
        })
        Taro.navigateBack()
      },
      onReset (event) {
      },
      onResult (date) {
        date[1] = date[1] - 1
        fromData.date = moment(date).format('YYYY-MM-DD HH:mm')
      },
    }
  }
}
</script>
