<template>
  <view>
    <AtForm>
      <picker
        mode='selector'
        rangeKey="name"
        :range="fromData.selector"
        :value="fromData.selectorValue"
        @change="handleChange"
        @cancel="handleCancel"
      >
        <at-flex>
          <at-flex-item align="center" :size="3"> 标签</at-flex-item>
          <at-flex-item align="center" :size="9"> {{ fromData.selector[fromData.selectorValue].name }}</at-flex-item>
        </at-flex>
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
        v-model:value="fromData.amount"
      />

      <AtButton formType='submit' @click="onSubmit">提交</AtButton>
    </AtForm>
    <timePickerColumn
      :start-time="dateSelector.start_time"
      :end-time="dateSelector.now_time"
      :default-time="dateSelector.now_time"
      @result="onResult"
    >
      <input
        placeholder="请选择"
        :value="dateSelector.now_time"
      >
    </timePickerColumn>
  </view>
</template>
<script>
import {reactive, ref} from 'vue'
import './index.scss'
import {getTags} from "../../api/common";
import timePickerColumn from "../../components/picker/index";
export default {
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
    const fromData = reactive({
      selector: tagSelector,
      selectorValue: 0,
      amount: '',
      // selectorKey: 'name',
      date: new Date().toDateString(),
    })
    return {
      fromData,
      dateSelector,
      handleChange:(e)=>{
        fromData.selectorValue = e.detail.value
      },
      handleCancel:(e)=>{
        console.log('handleCancel', e)
      },
      onSubmit (event) {
        console.log(fromData)
      },
      onReset (event) {
      },
      onResult (event) {
        console.log('onResult', event.join(':'))
      }
    }
  }
}
</script>
