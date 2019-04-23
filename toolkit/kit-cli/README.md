[build]
go get -u github.com/golang/protobuf/protoc-gen-go
go get -v github.com/LongMarch7/higo/toolkit/kit-cli

[new server]
kit-cli n s admin
kit-cli n s setting -p admin    //create a setting service in admin parent service

[add interface]
modify service.go file


[generate server]
kit-cli g s admin -t grpc
kit-cli g s setting -t grpc -p admin //generate a setting service in admin parent service

[Remarks]
if function has struct  must add "Alias" Keyword