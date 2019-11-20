var PROTO_PATH = __dirname + '/butcherpb/butcher.proto';
var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
// The protoDescriptor object has the full package hierarchy
var butcher = protoDescriptor.butcher;

function GetMeat(call, callback) {
    console.log("incoming!");
    let price = 0;
    switch (call.request.name) {
        case "roastbeef":
            price = 1;
            break;
        case "salami":
            price = 2;
            break;
        default:
            price = 3;
            break
    }

    callback(null, {
        name: call.request.name,
        price: price
    });
}

function setup() {
    var server = new grpc.Server();

    server.addService(butcher.ButcherService.service, {
        GetMeat: GetMeat
    });

    const port = 6050
    console.log("started.. on " + port);
    server.bind('0.0.0.0:' + port, grpc.ServerCredentials.createInsecure());
    server.start();
}

setup();

