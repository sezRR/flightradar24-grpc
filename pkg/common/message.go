package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
)

const (
	GrpcFrameHeaderSize = 5
	GrpcCompressionFlag = 0x00
)

func extractPayload(data []byte) ([]byte, error) {
	trailerStart := bytes.LastIndex(data, []byte("grpc-status:"))
	if trailerStart == -1 {
		return data, nil
	}

	payload := data[:trailerStart]

	if len(payload) < GrpcFrameHeaderSize {
		return nil, fmt.Errorf("invalid gRPC frame header")
	}

	compressionFlag := payload[0]
	if compressionFlag != GrpcCompressionFlag {
		return nil, fmt.Errorf("compressed payload not supported")
	}

	msgLength := binary.BigEndian.Uint32(payload[1:GrpcFrameHeaderSize])
	if len(payload) < GrpcFrameHeaderSize+int(msgLength) {
		return nil, fmt.Errorf("incomplete gRPC message")
	}

	return payload[GrpcFrameHeaderSize : GrpcFrameHeaderSize+int(msgLength)], nil
}

func DecodeGRPCMessage[T proto.Message](data []byte, msg T) (T, error) {
	payload, err := extractPayload(data)
	if err != nil {
		return *new(T), fmt.Errorf("failed to clear gRPC frame: %w", err)
	}

	err = proto.Unmarshal(payload, msg)
	if err != nil {
		return *new(T), fmt.Errorf("failed to unmarshal message: %w", err)
	}

	return msg, nil
}

func EncodeGRPCMessage(msg proto.Message) ([]byte, error) {
	// Serialize the Protobuf message to bytes
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %w", err)
	}

	// Create a buffer to construct the binary output
	var buffer bytes.Buffer

	// Write the compression flag (1 byte, 0x00 for no compression)
	if err := buffer.WriteByte(0x00); err != nil {
		return nil, fmt.Errorf("failed to write compression flag: %w", err)
	}

	// Write the message length (4 bytes, big-endian)
	msgLength := uint32(len(msgBytes))
	if err := binary.Write(&buffer, binary.BigEndian, msgLength); err != nil {
		return nil, fmt.Errorf("failed to write message length: %w", err)
	}

	// Write the serialized Protobuf message
	if _, err := buffer.Write(msgBytes); err != nil {
		return nil, fmt.Errorf("failed to write message bytes: %w", err)
	}

	return buffer.Bytes(), nil
}
