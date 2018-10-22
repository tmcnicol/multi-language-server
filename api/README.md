# API
This folder contains the protofile and the Go Protocol buffers.

To compile the protofile for go

1. install protoc and grpc, see the blog post or use the google
2.  From the this directory run
  `grpc --go_out=plugins=grpc:. *.protoc`

