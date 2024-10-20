

import request from "./index.js";


export const FindAllNoPagination=(query)=> {
   return request.get("/file/findAllNoPagination",query)
}

export const FindAll=(query)=> {
   return request.post("/file/findAll",query)
}
export const FindById=(id)=> {
   return request.get("/file/findById/"+id)
}
export const updateById=(id,form)=> {
   return request.post("/file/updateById/"+id,form)
}
export const deleteById=(id,form)=> {
   return request.get("/file/deleteById/"+id,form)
}
export const add=(form)=> {
   return request.post("/file/add",form)
}


