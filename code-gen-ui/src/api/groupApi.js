import request from "@/api/index.js";


export const DeleteFileById=(id)=> {
	return request.delete("/group/deleteFileById/"+id)
}


export const FindAllNoPagination=(query)=> {
	return request.get("/group/findAllNoPagination",query)
}


export const FindAll=(query)=> {
	return request.post("/group/findAll",query)
}

export const FindAllFileByGroupId=(id)=> {
	return request.get("/group/findAllFileByGroupId/"+id)
}


export const FindById=(id)=> {
	return request.get("/group/findById/"+id)
}

export const FindAllDir=(form)=> {
	return request.post("/group/findAllDir",form)
}

export const FindAllDirForUpdate=(form)=> {
	return request.post("/group/findAllDirForUpdate",form)
}
export const UpdateById=(id,form)=> {
	return request.post("/group/updateById/"+id,form)
}
export const DeleteById=(id,form)=> {
	return request.get("/group/deleteById/"+id,form)
}
export const Add=(form)=> {
	return request.post("/group/add",form)
}
