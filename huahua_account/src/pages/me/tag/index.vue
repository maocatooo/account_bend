<template>
  <view>
    <at-list>
      <at-swipe-action
        v-for="(item, index) in state.list"
        :key="item.name"
        :options="state.OPTIONS"
        :isOpened="item.isOpened"
        @click="handleClicked"
        @opened="handleSingle(index)"
      >
        <at-list-item :title="item.name" hasBorder/>
      </at-swipe-action>
    </at-list>
    <AtButton formType='submit' @click="popLayer">新增</AtButton>
    <view class="addIsOpened">
      <AtFloatLayout
        ref="atFloatLayout"
        :style="{top: -state.keyboardHeight + 'px'}"
        :is-opened="state.addIsOpened"
        @close="closeLayer"
      >
        <AtInput
          name='value'
          title='标签名称'
          type='text'
          placeholder='标签名称'
          v-model:value="state.fromData.name"
        />
        <AtButton formType='submit' @click="tagAdd">新增</AtButton>
      </AtFloatLayout>
    </view>
  </view>
</template>

<script lang="ts">
import {AtListItem, AtInput, AtList, AtSwipeAction, AtButton, AtFloatLayout} from 'taro-ui-vue3'
import Taro from "@tarojs/taro"
import {defineComponent, reactive} from "vue"
import "./index.scss"
import {setTags, OPTIONS, showToast} from "../../../api/common";
import {CreateTag, Tags} from "../../../api/api";

export default defineComponent({
  components: {
    AtListItem,
    AtButton,
    AtSwipeAction,
    AtFloatLayout,
    AtInput,
    AtList
  },
  mounted() {
    this.getTags()
  },
  methods: {
    popLayer() {
      this.state.fromData.name = ""
      this.state.addIsOpened = true
      Taro.onKeyboardHeightChange((res) => {
        console.log("键盘弹起",res)
        this.state.keyboardHeight = res.height
      })

    },
    closeLayer() {
      this.state.addIsOpened = false
    },
    async tagAdd() {
      if (!this.state.fromData.name) {
        showToast("请输入标签名称")
        return
      }
      if (this.state.list.filter((item) => item.name === this.state.fromData.name).length > 0){
        showToast("标签名称已存在"  )
        return
      }
      await CreateTag({name: this.state.fromData.name})
      this.closeLayer()
      await this.getTags()
      showToast("新增成功")
    },
    async getTags() {
      const tags = await Tags()
      setTags(tags)
      this.state.list = tags.map((item) => {
        item.isOpened = false
        return item
      })
    },
  },
  setup() {
    const state = reactive({
      list: [],
      currentIndex: 0,
      addIsOpened: false,
      keyboardHeight: 0,
      OPTIONS,
      fromData: {
        name: ""
      }
    })


    function handleClick(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}`)
      console.log(arguments)
    }

    function handleClicked(item, key) {
      showToast(`点击了${item.text}按钮，key: ${key}, index:${state.currentIndex}`)
      if (item.text == "确认") {
      } else {
        state.list = state.list.filter((_item, key) => key !== state.currentIndex)
      }
    }

    function handleSingle(index) {
      state.currentIndex = index
      state.list = state.list.map((item, key) => {
        item.isOpened = key === index
        return item
      })
    }

    return {
      state,
      handleClick,
      handleClicked,
      handleSingle,
      showToast,
    }
}
})
</script>
