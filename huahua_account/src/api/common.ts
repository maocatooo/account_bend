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
