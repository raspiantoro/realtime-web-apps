import { StreamClient } from "../stream-pb/v1/stream_grpc_web_pb"

class GrpcClient{

    constructor(){
        if(!GrpcClient.instance){
            GrpcClient.instance = this.createInstance()
        }
    }

    createInstance(){
        const host = "stream.realtime" //process.env.REACT_APP_GRPC_HOST
        const port = "80" //process.env.REACT_APP_GRPC_PORT
        const url = 'http://' + host + ':' + port
        console.log("ADDRESS: " + url)
        const client = new StreamClient(url)
        return client
    }

    getInstance(){
        return GrpcClient.instance
    }

}

const instance = new GrpcClient();
Object.freeze(instance);

export default instance