import Taro from "@tarojs/taro"
import {getToken} from "./common"
import {BookJournalsReq, CreateBookJournalReq, CreateTagReq, UpdateBookJournalReq} from "./types"

// const domain = "http://maocat.cc"
const domain = "https://ac.maocat.cc"
// const domain = "http://127.0.0.1:8888"
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
  console.log("req", url, method, data, header)
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
const Books = async () => {
  return await req("/book", "get",{} )
}

const Tags = async () => {
  return await req("/tag", "get",{} )
}
const CreateBookJournal = async (c: CreateBookJournalReq) => {
  return await req("/journal", "post",c)
}
const UpdateBookJournal = async (c: UpdateBookJournalReq) => {
  return await req("/journal", "put",c)
}

const DeleteBookJournal = async (id: string) => {
  return await req("/journal", "delete",{id})
}

const BookJournals = async (c:BookJournalsReq) => {
  return await req("/journal", "get", c)
}

const CreateTag = async (c: CreateTagReq) => {
  return await req("/tag", "post",c)
}

export  {Login, Books, Tags, CreateBookJournal, UpdateBookJournal, BookJournals, CreateTag, DeleteBookJournal}

