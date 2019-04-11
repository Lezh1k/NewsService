PROTO_MSG_INC=commons/pb
PROTO_MSG_TARGET=$(PROTO_MSG_INC)/message.proto
PROTO_MSG_TARGET_DIR=$(PROTO_MSG_INC)

instal_proto_gen:
	go get -u github.com/golang/protobuf/protoc-gen-go

proto_message:
	protoc -I=$(PROTO_MSG_INC) --go_out=$(PROTO_MSG_TARGET_DIR) $(PROTO_MSG_TARGET)

proto_clean:
	rm -f $(PROTO_MSG_TARGET_DIR)/message.pb.go
