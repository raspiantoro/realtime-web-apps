import { StreamRequest } from "../../utils/stream-pb/v1/stream_pb"
import GrpcClient from "../../utils/grpc/client"

class Stream{

    constructor(emitter){
        this.emitter = emitter
        this.client = GrpcClient.getInstance()
        this.request = new StreamRequest()
    }
    
    emit(response, emitter){
        emitter.emit('newstream', {
            count: response.getCount(),
        })
    }
    
    get(){
        const stream = this.client.getDataStream(this.request, {})
        const emit = this.emit
        const emitter = this.emitter
        stream.on('data', function (response) {
            emit(response, emitter)
        })    
    }

}

export default Stream;