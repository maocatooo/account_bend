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
        <AtInput
          name='value'
          title='账本'
          type='text'
          placeholder='账本'
          :value="fromData.bookSelector[fromData.bookSelectorValue].name"
        />
      </picker>
      <picker
        mode='selector'
        rangeKey="name"
        :range="fromData.tagSelector"
        :value="fromData.tagSelectorValue"
        @change="handleTagChange"
        @cancel="handleCancel"
      >
        <AtInput
          name='value'
          title='标签'
          type='text'
          placeholder='标签'
          :value="fromData.tagSelector[fromData.tagSelectorValue].name"
        />
      </picker>

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
        <AtInput
          name='value'
          title='时间'
          type='text'
          placeholder='时间'
          v-model:value="fromData.date"
        />
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
import {CreateBookJournal} from "../../api/api";
import moment from "moment";
import timePickerColumn from "../../components/picker/index";

export default {
  onPullDownRefresh (){
    this.fromData.tagSelector = getTags()
    this.fromData.bookSelector = getBooks()
    console.log('onPullDownRefresh ok')
    Taro.stopPullDownRefresh()
  },
  components: {
    timePickerColumn
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
      tagSelector: tagSelector,
      tagSelectorValue: 0,
      bookSelector: bookSelector,
      bookSelectorValue: 0,
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
        const res  = await CreateBookJournal({
          tid: tagSelector[fromData.tagSelectorValue].id,
          tname: tagSelector[fromData.tagSelectorValue].name,
          amount: fromData.amount,
          record: fromData.record,
          bookID: bookSelector[fromData.bookSelectorValue].id,
          date: parseInt(moment(fromData.date).format("x"))
        })
        console.log(res)
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
