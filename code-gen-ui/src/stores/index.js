

//导出默认的pinia实例
import {createPinia} from "pinia";
let pinia = createPinia();

//初始化持久化插件
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
pinia.use(piniaPluginPersistedstate)
export default pinia

//导出所有的模块
export * from "./module/dataBaseStore.js";

