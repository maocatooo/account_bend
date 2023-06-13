<template>
  <view class="index">
    <AtList>
        <AtListItem
          :title='title'
          arrow='right'
          :thumb='avatarUrl'
          @click="login"
        />
      <AtListItem title='account books' @click="redirectBook" hasBorder extraText='10'  arrow='right'/>
      <AtListItem title='tags'  extraText='4' hasBorder  arrow='right' @click="redirectTag"/>
      <AtListItem title='tags' disabled extraText='4'  hasBorder arrow='right' @click="handleClick"/>
    </AtList>
  </view>
</template>

<script>
import { ref } from 'vue'
import './index.scss'
import Taro from "@tarojs/taro"
import {AcBooks, Login} from "../../api/api";
import {setToken} from "../../api/common";
export default {
  setup () {
    const avatarUrl = ref('/static/img/tabbar/me.png')
    const title = ref('没有登录')
    return {
      avatarUrl,
      title,
      redirectBook:()=>{
        Taro.navigateTo({
          url:'/pages/me/book/index'
        })
      },
      redirectTag:()=>{
        Taro.navigateTo({
          url:'/pages/me/tag/index'
        })
      },
       handleClick: ()=>{
        console.log(12333)
      },

      login:()=>{
        Taro.getUserProfile({
          desc: '获取头像等信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
          success: ({userInfo}) => {
            // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
            console.log(userInfo)
            Taro.login({
              success: async function (res) {
                console.log("-------", res)
                if (res.code) {
                  //发起网络请求
                  const {accessToken, avatarUrl:aUrl, name} = await Login({code:res.code, avatarUrl:userInfo.avatarUrl, name:userInfo.nickName})

                  setToken(accessToken)
                  title.value = name
                  avatarUrl.value = aUrl
                  const  ac_books = await AcBooks()
                  console.log("======")
                  console.log(ac_books)

                } else {
                  console.log('登录失败！' + res.errMsg)
                }
              }
            })
          }
        })


      }
    }
  }
}
</script>
