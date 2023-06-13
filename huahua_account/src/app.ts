import { createApp } from 'vue'
import './app.scss'
import { createUI } from 'taro-ui-vue3'
const App = createApp({
  onShow (options) {
    console.log(options)
  },
  // 入口组件不需要实现 render 方法，即使实现了也会被 taro 所覆盖
})
import "taro-ui-vue3/dist/style/components/list.scss";
const tuv3 = createUI()
App.use(tuv3)

export default App
