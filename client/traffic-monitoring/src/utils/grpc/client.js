import { TrafficClient } from "../traffic-pb/v1/traffic_grpc_web_pb"

const GrpcClient = (function(){
    var instance;
 
    function createInstance() {
        const client = new TrafficClient('http://localhost:8010');
        return client;
    }
 
    return {
        getInstance: function () {
            if (!instance) {
                instance = createInstance();
            }
            return instance;
        }
    };
})();

export default GrpcClient