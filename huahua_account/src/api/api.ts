import Taro from "@tarojs/taro"
import {getToken} from "./common"

// const domain = "http://maocat.cc"
const domain = "http://127.0.0.1:8888"
const  req = async function (
   url:string,
   method:any,
   data?:any,

){

  const header = {

  }
  const token = getToken()

  if (token){
    header["Authorization"] =  token
  }

  const query = {
    url: `${domain}${url}`,
    data,
    header,
    method
  }

  const {data:res} = await  Taro.request(query)

  return res
}


const Login = async ({code, avatarUrl, name}) => {
  return await req("/login", "post",{code,avatarUrl, name} )
}
const AcBooks = async () => {
  return await req("/ac_book", "get",{} )
}
export  {Login, AcBooks}
