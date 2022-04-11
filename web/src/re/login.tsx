import axios from "axios";


// type ss{
//     socket :String
// }

/**
 * @JE 连接websocket
 * @return socket套接字句柄
 * */
export function linkSocket() {

    axios.defaults.baseURL = "/api"
    axios.post("/linkSocket").then(()=>{

    })
    return
}
/**
 * @添加socket连接
 * */
function addSocket() {

}