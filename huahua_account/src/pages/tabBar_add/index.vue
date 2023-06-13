<template>
  <view>
    <AtForm>
      <picker
        mode='selector'
        :range="selector"
        :value="selectorValue"
        @change="handleChange"
        @cancel="handleCancel"
      >
        <at-flex>
          <at-flex-item align="center" :size="3"> 标签</at-flex-item>
          <at-flex-item align="center" :size="9"> {{ selector[selectorValue] }}</at-flex-item>
        </at-flex>
      </picker>

      <AtInput
        name='value'
        title='金额'
        type='text'
        placeholder='金额(小数保持两位)'
        v-model:value="value"
      />
      <AtInput
        name='value'
        title='备注'
        type='text'
        placeholder='备注'
        v-model:value="value"
      />


      <AtButton formType='submit' @click="onSubmit">提交</AtButton>
    </AtForm>
  </view>
</template>
<script>
import { ref } from 'vue'
import './index.scss'

export default {
  setup () {
    const value = ref('Hello world2')
    const selectorValue = ref(1)
    return {
      selector: ['中国', '美国', '巴西', '日本'],
      selectorValue,
      value,
      handleChange:(e)=>{
        selectorValue.value = e.detail.value
      },
      handleCancel:(e)=>{
        console.log('handleCancel', e)
      },
      onSubmit (event) {
        console.log(value.value)
      },
      onReset (event) {
      }
    }
  }
}
</script>
