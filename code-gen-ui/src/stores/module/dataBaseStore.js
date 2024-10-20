import {defineStore} from "pinia";
import {ref} from "vue";


export const dataBaseStore = defineStore('dataBase',()=>{
    const dataBaseConfig=ref({})

    const setDataBaseConfig=(newValue)=>{
        dataBaseConfig.value=newValue
    }

    const getDataBaseConfig=()=>{
        return dataBaseConfig.value
    }

    const clearDataBaseConfig=()=>{
        dataBaseConfig.value=""
    }

    //导出给外部使用
    return{
        dataBaseConfig,
        setDataBaseConfig,
        getDataBaseConfig,
        clearDataBaseConfig
    }

},{
        persist:{
            storage:localStorage,
        },
    }
)