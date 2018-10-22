const PROTO_PATH = '../../api/api.proto'
const grpc = require('grpc')

const protoLoader = require('@grpc/proto-loader')

const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  }
)

const api = grpc.loadPackageDefinition(packageDefinition)

function HandleGeneric(call, callback) {
  const name = call.request.name
  console.log(`handling request ${name}`)
  const message = `Hello ${name}. You are awesome. Thanks so much for reading!!`
  callback(null, {message})
}

function main() {
  const server = new grpc.Server()
  server.addService(api.Handler.service, {HandleGeneric})
  server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
  server.start();
}

main()

