import Taro from "@tarojs/taro"

export  function getToken(){
  return Taro.getStorageSync('__token')
}

export  function setToken(token){
  Taro.setStorage({
    key:"__token",
    data: token
  })
}
