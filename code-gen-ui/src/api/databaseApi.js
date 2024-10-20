import request from "@/api/index.js";


export const FindTablesByDatabaseId=(databaseId)=> {
	return request.get("/database/findTablesByDatabaseId/"+databaseId)
}

export const FindAllNoPagination=()=> {
	return request.get("/database/findAllNoPagination")
}

export const FindAll=(query)=> {
	return request.post("/database/findAll",query)
}

export const checkConnect=(form)=>{
	return request.post("/database/checkConnect",form)
}

export const FindById=(id)=> {
	return request.get("/database/findById/"+id)
}
export const UpdateById=(id,form)=> {
	return request.post("/database/updateById/"+id,form)
}
export const DeleteById=(id,form)=> {
	return request.get("/database/deleteById/"+id,form)
}
export const Add=(form)=> {
	return request.post("/database/add",form)
}
