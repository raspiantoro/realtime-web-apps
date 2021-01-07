import { StreamClient } from "../stream-pb/v1/stream_grpc_web_pb"

class GrpcClient{

    constructor(){
        if(!GrpcClient.instance){
            GrpcClient.instance = this.createInstance()
        }
    }

    createInstance(){
        const host = process.env.REACT_APP_GRPC_HOST
        const port = process.env.REACT_APP_GRPC_PORT
        const url = 'http://' + host + ':' + port
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