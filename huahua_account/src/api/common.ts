import Taro from "@tarojs/taro"
import {ref} from "vue";

export  function getToken(){
  return Taro.getStorageSync('__token')
}

export  function setToken(token){
  Taro.setStorage({
    key:"__token",
    data: token
  })
}


export  function setBooks(books){
  Taro.setStorage({
    key:"__books",
    data: books
  })
}


export  function getBooks(){
  return Taro.getStorageSync('__books') || []
}

export  function setTags(data){
  Taro.setStorage({
    key:"__tags",
    data: data
  })
}
export  function getTags(){
  return Taro.getStorageSync('__tags') || []
}

export const DEL_OPTION_INDEX =  0
export const UPD_OPTION_INDEX =  1
export const OPTIONS = ref([
  {
    text: "删除",
    style: {
      color: '#333',
      backgroundColor: '#F7F7F7'
    }
  },
  {
    text: '编辑',
    style: {
      backgroundColor: '#E93B3D'
    }
  }
])

export const showToast = (titleMsg)=> {
  Taro.showToast({
    icon: 'none',
    title: titleMsg || '点击了按钮',
  })
}
