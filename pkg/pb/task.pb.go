// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateTaskRequest struct {
	X          *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=directive" json:"directive,omitempty"`
}

func (m *CreateTaskRequest) Reset()                    { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()               {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateTaskRequest) GetX() *google_protobuf.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateTaskRequest) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId  *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *CreateTaskResponse) Reset()                    { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()               {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CreateTaskResponse) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type Task struct {
	TaskId     *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId      *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode  *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive  *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor   *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner      *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target     *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId     *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Task) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *google_protobuf.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *google_protobuf.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *google_protobuf.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *google_protobuf.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *google_protobuf.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *google_protobuf.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId []string `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId  []string `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Status []string `protobuf:"bytes,3,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset uint32 `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeTasksRequest) Reset()                    { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()               {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeTasksResponse struct {
	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet    []*Task `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
}

func (m *DescribeTasksResponse) Reset()                    { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()               {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcf, 0x4f, 0x1b, 0x47,
	0x14, 0xc7, 0xbb, 0x06, 0x1b, 0xfc, 0xb6, 0xae, 0xe8, 0x08, 0xda, 0x95, 0x45, 0xc1, 0x5d, 0xf5,
	0x40, 0xa1, 0xec, 0x82, 0x69, 0xa5, 0xca, 0xa8, 0x07, 0x97, 0x1e, 0xea, 0x43, 0x55, 0x64, 0x48,
	0x0e, 0xb9, 0x58, 0xe3, 0xdd, 0xe7, 0x65, 0x8c, 0xd9, 0x59, 0x66, 0xc6, 0xd8, 0xc7, 0x28, 0x87,
	0x48, 0x91, 0x72, 0x22, 0x7f, 0x4b, 0xfe, 0x92, 0xfc, 0x0b, 0x39, 0xe7, 0x6f, 0x88, 0x66, 0x76,
	0x8d, 0x0d, 0x86, 0x64, 0x93, 0x48, 0x39, 0x59, 0xf3, 0xe6, 0xf3, 0xdd, 0xf7, 0x73, 0x9e, 0x01,
	0x14, 0x95, 0xe7, 0x5e, 0x22, 0xb8, 0xe2, 0x04, 0x78, 0x82, 0x71, 0xc2, 0x94, 0x60, 0xe3, 0xea,
	0x46, 0xc4, 0x79, 0x34, 0x40, 0xdf, 0xdc, 0x74, 0x87, 0x3d, 0x7f, 0x24, 0x68, 0x92, 0xa0, 0x90,
	0x29, 0x5b, 0xdd, 0xbc, 0x7b, 0xaf, 0xd8, 0x05, 0x4a, 0x45, 0x2f, 0x92, 0x0c, 0x58, 0xcf, 0x00,
	0x9a, 0x30, 0x9f, 0xc6, 0x31, 0x57, 0x54, 0x31, 0x1e, 0x4f, 0xe4, 0xbf, 0x99, 0x9f, 0x60, 0x37,
	0xc2, 0x78, 0x57, 0x8e, 0x68, 0x14, 0xa1, 0xf0, 0x79, 0x62, 0x88, 0x79, 0xda, 0x7d, 0x67, 0xc1,
	0xf7, 0x47, 0x02, 0xa9, 0xc2, 0x53, 0x2a, 0xcf, 0xdb, 0x78, 0x39, 0x44, 0xa9, 0xc8, 0xaf, 0x60,
	0x75, 0x1c, 0xab, 0x66, 0x6d, 0xd9, 0xf5, 0x75, 0x2f, 0xf5, 0xe6, 0x4d, 0xc2, 0xf1, 0x4e, 0x94,
	0x60, 0x71, 0xf4, 0x98, 0x0e, 0x86, 0xd8, 0xfe, 0x86, 0x1c, 0x40, 0xa9, 0xcf, 0xbb, 0x1d, 0x16,
	0x3a, 0x85, 0x1c, 0x7c, 0xb1, 0xcf, 0xbb, 0xad, 0x90, 0xfc, 0x05, 0xb6, 0x2e, 0x4e, 0x87, 0x06,
	0x3a, 0x16, 0x67, 0x21, 0x87, 0xd2, 0x54, 0xb3, 0x69, 0x78, 0xd2, 0x80, 0x72, 0xc8, 0x04, 0x06,
	0x8a, 0x5d, 0xa1, 0xb3, 0x98, 0x43, 0x3c, 0xc5, 0xdd, 0xa7, 0x16, 0x90, 0xd9, 0x84, 0x65, 0xc2,
	0x63, 0x89, 0xe4, 0x0f, 0x58, 0x32, 0x11, 0xb1, 0x30, 0x57, 0xde, 0x25, 0x0d, 0xb7, 0xc2, 0xcf,
	0xca, 0xde, 0x7d, 0x5d, 0x84, 0x45, 0xed, 0xfc, 0x6b, 0x3a, 0xfd, 0xd2, 0x92, 0xff, 0x0e, 0x25,
	0xa9, 0xa8, 0x1a, 0xca, 0x5c, 0xf5, 0xce, 0x58, 0x72, 0x08, 0x80, 0x42, 0x70, 0xd1, 0x09, 0x78,
	0x88, 0x4e, 0xf1, 0x01, 0xe5, 0xa3, 0x56, 0xac, 0x0e, 0xea, 0x59, 0xa7, 0x0c, 0x7f, 0xc4, 0x43,
	0xbc, 0xdd, 0xe5, 0xd2, 0x27, 0x75, 0x99, 0xfc, 0x09, 0xcb, 0x38, 0xc6, 0x60, 0xa8, 0xb8, 0x70,
	0x96, 0x72, 0x48, 0x6f, 0x68, 0x52, 0x87, 0x22, 0x1f, 0xc5, 0x28, 0x9c, 0xe5, 0x3c, 0xb5, 0x35,
	0xa8, 0x2e, 0x8e, 0xa2, 0x22, 0x42, 0xe5, 0x94, 0xf3, 0xb5, 0x51, 0xb3, 0xba, 0xfb, 0x31, 0x0f,
	0x51, 0xf7, 0x11, 0xf2, 0xc8, 0x34, 0xdc, 0x0a, 0xc9, 0x21, 0xd8, 0x81, 0x99, 0xdf, 0x8e, 0xde,
	0x0b, 0x8e, 0x6d, 0xa4, 0xd5, 0x39, 0xe9, 0xe9, 0x64, 0x69, 0xb4, 0x21, 0xc5, 0xb5, 0x41, 0x8b,
	0xd3, 0xd6, 0xa4, 0xe2, 0x6f, 0x3f, 0x2e, 0x4e, 0x71, 0x6d, 0x70, 0x5f, 0x5a, 0xb0, 0xfa, 0x0f,
	0xca, 0x40, 0xb0, 0xae, 0x79, 0x3c, 0x72, 0xb2, 0x2e, 0x7e, 0x9c, 0x9d, 0xe3, 0x85, 0xad, 0xf2,
	0xcd, 0xa4, 0xae, 0xcd, 0x4c, 0xaa, 0xb6, 0x67, 0xb3, 0xf8, 0xc3, 0xcd, 0x30, 0x2d, 0xa4, 0x78,
	0x36, 0x2e, 0xab, 0x50, 0x1c, 0xb0, 0x0b, 0xa6, 0xcc, 0x8c, 0x55, 0xda, 0xe9, 0x41, 0xd3, 0xbc,
	0xd7, 0x93, 0xa8, 0xcc, 0x00, 0x55, 0xda, 0xd9, 0xc9, 0x45, 0x58, 0xbb, 0x13, 0x4d, 0xf6, 0x96,
	0x37, 0xc1, 0x56, 0x5c, 0xd1, 0x41, 0x27, 0xe0, 0xc3, 0x58, 0x99, 0xa7, 0x55, 0x69, 0x83, 0x31,
	0x1d, 0x69, 0x0b, 0xd9, 0x81, 0x65, 0x13, 0xaf, 0xfe, 0xa6, 0x0e, 0xcc, 0xae, 0xaf, 0x78, 0xd3,
	0x05, 0xed, 0x99, 0xc5, 0x60, 0x32, 0x3a, 0x41, 0x55, 0x7f, 0x51, 0x00, 0x5b, 0x5b, 0xfe, 0xa3,
	0x31, 0x8d, 0x50, 0x90, 0x4b, 0x80, 0xe9, 0xfe, 0x20, 0x3f, 0xcd, 0x0a, 0xe7, 0x16, 0x69, 0x75,
	0xe3, 0xa1, 0xeb, 0x34, 0x54, 0xf7, 0x97, 0xeb, 0x66, 0x85, 0x64, 0xfd, 0xac, 0x69, 0x8f, 0xcf,
	0xde, 0xbc, 0x7d, 0x55, 0xf8, 0xce, 0x2d, 0xfb, 0x57, 0xfb, 0xbe, 0x3e, 0xcb, 0x86, 0xb5, 0x4d,
	0x9e, 0x5b, 0x50, 0xb9, 0x95, 0x2a, 0xa9, 0xcd, 0x7e, 0xf7, 0xbe, 0x9e, 0x54, 0x7f, 0xfe, 0x00,
	0x91, 0x39, 0xdf, 0xbb, 0x6e, 0xae, 0x93, 0x6a, 0x98, 0xdd, 0x19, 0xf7, 0xb2, 0x36, 0x62, 0xea,
	0xac, 0xd6, 0x63, 0x03, 0x85, 0xc2, 0xc4, 0x62, 0x93, 0x69, 0x2c, 0x7f, 0x8f, 0xaf, 0x9b, 0x97,
	0xe4, 0x5f, 0x20, 0xff, 0x27, 0x18, 0x1f, 0x33, 0x81, 0x6c, 0x5c, 0x3b, 0x16, 0xbc, 0x8f, 0x81,
	0x72, 0x77, 0xee, 0xb3, 0x92, 0xb5, 0x33, 0xa5, 0x12, 0xd9, 0xf0, 0xfd, 0x99, 0x60, 0x18, 0xaf,
	0x17, 0xf7, 0xbc, 0x3d, 0x6f, 0x7f, 0xdb, 0xb2, 0xea, 0x2b, 0x34, 0x49, 0x06, 0x2c, 0x30, 0xff,
	0x49, 0x7e, 0x5f, 0xf2, 0xb8, 0x31, 0x67, 0x79, 0x52, 0x48, 0xba, 0xdd, 0x92, 0x99, 0xcd, 0x83,
	0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xf5, 0x6f, 0x8d, 0x55, 0x07, 0x00, 0x00,
}
