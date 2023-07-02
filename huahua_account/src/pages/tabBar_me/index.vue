<template>
  <view class="index">
    <AtList>
      <AtListItem
        :title='info.name'
        arrow='right'
        :thumb='avatarUrl'
        @click="login()"
      />
      <AtListItem title='account books' @click="redirectBook" hasBorder :extraText='info.booksCount' arrow='right'/>
      <AtListItem title='tags' :extraText='info.tagsCount' hasBorder arrow='right' @click="redirectTag"/>
      <!--      <AtListItem title='tags' disabled extraText='4'  hasBorder arrow='right' @click="handleClick"/>-->
    </AtList>
  </view>
</template>

<script>
import {reactive, ref} from 'vue'
import './index.scss'
import Taro from "@tarojs/taro"
import {Books, Login, Tags} from "../../api/api";
import {setToken, setBooks, setTags} from "../../api/common";

export default {
  setup() {
    const info = reactive({
      avatarUrl: "/static/img/tabbar/me.png",
      booksCount: "0",
      tagsCount: "0",
      name: "没有登录"
    })
    const login = () => {
      Taro.login({
          success: async function (res) {
            if (res.code) {
              //发起网络请求
              const {accessToken, avatarUrl: aUrl, name} = await Login({code: res.code, avatarUrl: "", name: "微信用户"})
              setToken(accessToken)
              info.name = name
              info.avatarUrl = aUrl

              const books = await Books()
              setBooks(books)
              info.booksCount = books.length.toString()

              const tags = await Tags()
              setTags(tags)
              info.tagsCount = tags.length.toString()
            } else {
              console.log('登录失败！' + res.errMsg)
            }
          }
        }
      )
    }
    console.log("login")
    login()

    return {
      info,
      login,
      redirectBook: () => {
        Taro.navigateTo({
          url: '/pages/me/book/index'
        })
      },
      redirectTag: () => {
        Taro.navigateTo({
          url: '/pages/me/tag/index'
        })
      },
      handleClick: () => {
        console.log(12333)
      },
      setUserProfile: () => {
        Taro.getUserProfile({
          desc: '获取头像等信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
          success: ({userInfo}) => {
            // 开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
            // todo 保存用户信息
          }
        })
      }

    }
  }
}
</script>
