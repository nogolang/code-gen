import{I as t}from"./index-B3M88loM.js";const d=a=>t.get("/database/findTablesByDatabaseId/"+a),n=()=>t.get("/database/findAllNoPagination"),o=a=>t.post("/database/findAll",a),r=a=>t.post("/database/checkConnect",a),c=a=>t.get("/database/findById/"+a),i=(a,e)=>t.post("/database/updateById/"+a,e),b=(a,e)=>t.get("/database/deleteById/"+a,e),l=a=>t.post("/database/add",a);export{l as A,b as D,c as F,i as U,o as a,n as b,r as c,d};