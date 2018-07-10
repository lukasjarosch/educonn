// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/course/course.proto

/*
Package educonn_course is a generated protocol buffer package.

It is generated from these files:
	proto/course/course.proto

It has these top-level messages:
	CourseEntity
	Request
	CourseResponse
	Error
*/
package educonn_course

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Course service

type CourseClient interface {
	Create(ctx context.Context, in *CourseEntity, opts ...client.CallOption) (*CourseResponse, error)
	Get(ctx context.Context, in *CourseEntity, opts ...client.CallOption) (*CourseResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*CourseResponse, error)
}

type courseClient struct {
	c           client.Client
	serviceName string
}

func NewCourseClient(serviceName string, c client.Client) CourseClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "educonn.course"
	}
	return &courseClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *courseClient) Create(ctx context.Context, in *CourseEntity, opts ...client.CallOption) (*CourseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Course.Create", in)
	out := new(CourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) Get(ctx context.Context, in *CourseEntity, opts ...client.CallOption) (*CourseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Course.Get", in)
	out := new(CourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*CourseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Course.GetAll", in)
	out := new(CourseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Course service

type CourseHandler interface {
	Create(context.Context, *CourseEntity, *CourseResponse) error
	Get(context.Context, *CourseEntity, *CourseResponse) error
	GetAll(context.Context, *Request, *CourseResponse) error
}

func RegisterCourseHandler(s server.Server, hdlr CourseHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Course{hdlr}, opts...))
}

type Course struct {
	CourseHandler
}

func (h *Course) Create(ctx context.Context, in *CourseEntity, out *CourseResponse) error {
	return h.CourseHandler.Create(ctx, in, out)
}

func (h *Course) Get(ctx context.Context, in *CourseEntity, out *CourseResponse) error {
	return h.CourseHandler.Get(ctx, in, out)
}

func (h *Course) GetAll(ctx context.Context, in *Request, out *CourseResponse) error {
	return h.CourseHandler.GetAll(ctx, in, out)
}
