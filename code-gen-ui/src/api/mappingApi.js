

import request from "@/api/index.js";


export const FindAll=(query)=> {
	return request.post("/mapping/findAll",query)
}

export const FindAllFilesByMapId=(id)=> {
	return request.get("/mapping/findAllFilesByMapId/"+id)
}

export const FindAllNoPagination=()=> {
	return request.get("/mapping/findAllNoPagination")
}

export const FindById=(id)=> {
	return request.get("/mapping/findById/"+id)
}
export const UpdateById=(id,form)=> {
	return request.post("/mapping/updateById/"+id,form)
}
export const DeleteById=(id,form)=> {
	return request.get("/mapping/deleteById/"+id,form)
}
export const Add=(form)=> {
	return request.post("/mapping/add",form)
}
