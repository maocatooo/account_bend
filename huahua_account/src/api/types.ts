interface CreateBookJournalReq {
  amount:string
  record:string
  tid: string
  tname :string
  bookID :string
  name :string
  date:Number
}
interface UpdateBookJournalReq extends CreateBookJournalReq{
  id:string
}

interface BookJournalsReq {
  bookID :string
  date:Number
}

interface CreateTagReq {
  name:string
}


export {CreateBookJournalReq,UpdateBookJournalReq, BookJournalsReq, CreateTagReq}
