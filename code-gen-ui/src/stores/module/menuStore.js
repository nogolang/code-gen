import {defineStore} from "pinia";
import {ref} from "vue";


export const menuStore = defineStore('menu',()=>{
    const  expand=ref([])
    const setExpand=(newValue)=>{
        expand.value=newValue
    }

    const getExpand=()=>{
        return expand.value
    }

    //导出给外部使用
    return{
        expand,
        setExpand,
        getExpand
    }

},{
        persist:{
            storage:localStorage,
        },
    }
)