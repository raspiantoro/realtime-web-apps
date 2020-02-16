import { StreamClient } from "../stream-pb/v1/stream_grpc_web_pb"

class GrpcClient{

    constructor(){
        if(!GrpcClient.instance){
            GrpcClient.instance = this.createInstance()
        }
    }

    createInstance(){
        const client = new StreamClient('http://localhost:8010')
        return client
    }

    getInstance(){
        return GrpcClient.instance
    }

}

const instance = new GrpcClient();
Object.freeze(instance);

export default instance