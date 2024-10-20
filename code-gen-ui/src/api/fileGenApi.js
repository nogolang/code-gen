import request from "@/api/index.js";

export const GenFiles=(checkIds)=> {
	return request.post("/fileGen/genFiles",checkIds)
}

export const FindAll=(query)=> {
	return request.post("/fileGen/findAll",query)
}

export const FindById=(id)=> {
	return request.get("/fileGen/findById/"+id)
}
export const UpdateById=(id,form)=> {
	return request.post("/fileGen/updateById/"+id,form)
}
export const DeleteById=(id,form)=> {
	return request.get("/fileGen/deleteById/"+id,form)
}
export const Add=(form)=> {
	return request.post("/fileGen/add",form)
}
