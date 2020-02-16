import { TrafficCountRequest } from "../../utils/traffic-pb/v1/traffic_pb"
import Client from "../../utils/grpc/client"

function Traffic(){
    this.client = Client.getInstance()

    this.request = new TrafficCountRequest()
   
    this.get = (emitter) => {
        const stream = this.client.getTrafficCount(this.request, {})
        stream.on('data', function (response) {
            emitter.emit('newtrafficcount', {
                trafficCount: response.getTrafficcount(),
            })
        })    
    }
    

}

export default Traffic;