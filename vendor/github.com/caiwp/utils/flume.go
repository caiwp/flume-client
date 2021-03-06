// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Status int64

const (
	Status_OK      Status = 0
	Status_FAILED  Status = 1
	Status_ERROR   Status = 2
	Status_UNKNOWN Status = 3
)

func (p Status) String() string {
	switch p {
	case Status_OK:
		return "OK"
	case Status_FAILED:
		return "FAILED"
	case Status_ERROR:
		return "ERROR"
	case Status_UNKNOWN:
		return "UNKNOWN"
	}
	return "<UNSET>"
}

func StatusFromString(s string) (Status, error) {
	switch s {
	case "OK":
		return Status_OK, nil
	case "FAILED":
		return Status_FAILED, nil
	case "ERROR":
		return Status_ERROR, nil
	case "UNKNOWN":
		return Status_UNKNOWN, nil
	}
	return Status(0), fmt.Errorf("not a valid Status string")
}

func StatusPtr(v Status) *Status {
	return &v
}

func (p Status) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *Status) UnmarshalText(text []byte) error {
	q, err := StatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *Status) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = Status(v)
	return nil
}

func (p *Status) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Headers
//  - Body
type ThriftFlumeEvent struct {
	Headers map[string]string `thrift:"headers,1,required" db:"headers" json:"headers"`
	Body    []byte            `thrift:"body,2,required" db:"body" json:"body"`
}

func NewThriftFlumeEvent() *ThriftFlumeEvent {
	return &ThriftFlumeEvent{}
}

func (p *ThriftFlumeEvent) GetHeaders() map[string]string {
	return p.Headers
}

func (p *ThriftFlumeEvent) GetBody() []byte {
	return p.Body
}
func (p *ThriftFlumeEvent) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetHeaders bool = false
	var issetBody bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
			issetHeaders = true
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
			issetBody = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetHeaders {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Headers is not set"))
	}
	if !issetBody {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Body is not set"))
	}
	return nil
}

func (p *ThriftFlumeEvent) ReadField1(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.Headers = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val1 = v
		}
		p.Headers[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *ThriftFlumeEvent) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Body = v
	}
	return nil
}

func (p *ThriftFlumeEvent) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ThriftFlumeEvent"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ThriftFlumeEvent) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("headers", thrift.MAP, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:headers: ", p), err)
	}
	if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRING, len(p.Headers)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Headers {
		if err := oprot.WriteString(string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:headers: ", p), err)
	}
	return err
}

func (p *ThriftFlumeEvent) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("body", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:body: ", p), err)
	}
	if err := oprot.WriteBinary(p.Body); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.body (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:body: ", p), err)
	}
	return err
}

func (p *ThriftFlumeEvent) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftFlumeEvent(%+v)", *p)
}

type ThriftSourceProtocol interface {
	// Parameters:
	//  - Event
	Append(event *ThriftFlumeEvent) (r Status, err error)
	// Parameters:
	//  - Events
	AppendBatch(events []*ThriftFlumeEvent) (r Status, err error)
}

type ThriftSourceProtocolClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewThriftSourceProtocolClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ThriftSourceProtocolClient {
	return &ThriftSourceProtocolClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewThriftSourceProtocolClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ThriftSourceProtocolClient {
	return &ThriftSourceProtocolClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Event
func (p *ThriftSourceProtocolClient) Append(event *ThriftFlumeEvent) (r Status, err error) {
	if err = p.sendAppend(event); err != nil {
		return
	}
	return p.recvAppend()
}

func (p *ThriftSourceProtocolClient) sendAppend(event *ThriftFlumeEvent) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("append", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ThriftSourceProtocolAppendArgs{
		Event: event,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ThriftSourceProtocolClient) recvAppend() (value Status, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "append" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "append failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "append failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "append failed: invalid message type")
		return
	}
	result := ThriftSourceProtocolAppendResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Events
func (p *ThriftSourceProtocolClient) AppendBatch(events []*ThriftFlumeEvent) (r Status, err error) {
	if err = p.sendAppendBatch(events); err != nil {
		return
	}
	return p.recvAppendBatch()
}

func (p *ThriftSourceProtocolClient) sendAppendBatch(events []*ThriftFlumeEvent) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("appendBatch", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ThriftSourceProtocolAppendBatchArgs{
		Events: events,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ThriftSourceProtocolClient) recvAppendBatch() (value Status, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "appendBatch" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "appendBatch failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "appendBatch failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "appendBatch failed: invalid message type")
		return
	}
	result := ThriftSourceProtocolAppendBatchResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ThriftSourceProtocolProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ThriftSourceProtocol
}

func (p *ThriftSourceProtocolProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ThriftSourceProtocolProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ThriftSourceProtocolProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewThriftSourceProtocolProcessor(handler ThriftSourceProtocol) *ThriftSourceProtocolProcessor {

	self6 := &ThriftSourceProtocolProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["append"] = &thriftSourceProtocolProcessorAppend{handler: handler}
	self6.processorMap["appendBatch"] = &thriftSourceProtocolProcessorAppendBatch{handler: handler}
	return self6
}

func (p *ThriftSourceProtocolProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type thriftSourceProtocolProcessorAppend struct {
	handler ThriftSourceProtocol
}

func (p *thriftSourceProtocolProcessorAppend) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ThriftSourceProtocolAppendArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ThriftSourceProtocolAppendResult{}
	var retval Status
	var err2 error
	if retval, err2 = p.handler.Append(args.Event); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing append: "+err2.Error())
		oprot.WriteMessageBegin("append", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("append", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type thriftSourceProtocolProcessorAppendBatch struct {
	handler ThriftSourceProtocol
}

func (p *thriftSourceProtocolProcessorAppendBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ThriftSourceProtocolAppendBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("appendBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ThriftSourceProtocolAppendBatchResult{}
	var retval Status
	var err2 error
	if retval, err2 = p.handler.AppendBatch(args.Events); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing appendBatch: "+err2.Error())
		oprot.WriteMessageBegin("appendBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("appendBatch", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Event
type ThriftSourceProtocolAppendArgs struct {
	Event *ThriftFlumeEvent `thrift:"event,1" db:"event" json:"event"`
}

func NewThriftSourceProtocolAppendArgs() *ThriftSourceProtocolAppendArgs {
	return &ThriftSourceProtocolAppendArgs{}
}

var ThriftSourceProtocolAppendArgs_Event_DEFAULT *ThriftFlumeEvent

func (p *ThriftSourceProtocolAppendArgs) GetEvent() *ThriftFlumeEvent {
	if !p.IsSetEvent() {
		return ThriftSourceProtocolAppendArgs_Event_DEFAULT
	}
	return p.Event
}
func (p *ThriftSourceProtocolAppendArgs) IsSetEvent() bool {
	return p.Event != nil
}

func (p *ThriftSourceProtocolAppendArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Event = &ThriftFlumeEvent{}
	if err := p.Event.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Event), err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("event", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:event: ", p), err)
	}
	if err := p.Event.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Event), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:event: ", p), err)
	}
	return err
}

func (p *ThriftSourceProtocolAppendArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftSourceProtocolAppendArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ThriftSourceProtocolAppendResult struct {
	Success *Status `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewThriftSourceProtocolAppendResult() *ThriftSourceProtocolAppendResult {
	return &ThriftSourceProtocolAppendResult{}
}

var ThriftSourceProtocolAppendResult_Success_DEFAULT Status

func (p *ThriftSourceProtocolAppendResult) GetSuccess() Status {
	if !p.IsSetSuccess() {
		return ThriftSourceProtocolAppendResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ThriftSourceProtocolAppendResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ThriftSourceProtocolAppendResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		temp := Status(v)
		p.Success = &temp
	}
	return nil
}

func (p *ThriftSourceProtocolAppendResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("append_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ThriftSourceProtocolAppendResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftSourceProtocolAppendResult(%+v)", *p)
}

// Attributes:
//  - Events
type ThriftSourceProtocolAppendBatchArgs struct {
	Events []*ThriftFlumeEvent `thrift:"events,1" db:"events" json:"events"`
}

func NewThriftSourceProtocolAppendBatchArgs() *ThriftSourceProtocolAppendBatchArgs {
	return &ThriftSourceProtocolAppendBatchArgs{}
}

func (p *ThriftSourceProtocolAppendBatchArgs) GetEvents() []*ThriftFlumeEvent {
	return p.Events
}
func (p *ThriftSourceProtocolAppendBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchArgs) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*ThriftFlumeEvent, 0, size)
	p.Events = tSlice
	for i := 0; i < size; i++ {
		_elem8 := &ThriftFlumeEvent{}
		if err := _elem8.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem8), err)
		}
		p.Events = append(p.Events, _elem8)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("appendBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("events", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:events: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Events)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Events {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:events: ", p), err)
	}
	return err
}

func (p *ThriftSourceProtocolAppendBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftSourceProtocolAppendBatchArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ThriftSourceProtocolAppendBatchResult struct {
	Success *Status `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewThriftSourceProtocolAppendBatchResult() *ThriftSourceProtocolAppendBatchResult {
	return &ThriftSourceProtocolAppendBatchResult{}
}

var ThriftSourceProtocolAppendBatchResult_Success_DEFAULT Status

func (p *ThriftSourceProtocolAppendBatchResult) GetSuccess() Status {
	if !p.IsSetSuccess() {
		return ThriftSourceProtocolAppendBatchResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ThriftSourceProtocolAppendBatchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ThriftSourceProtocolAppendBatchResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		temp := Status(v)
		p.Success = &temp
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("appendBatch_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ThriftSourceProtocolAppendBatchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ThriftSourceProtocolAppendBatchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ThriftSourceProtocolAppendBatchResult(%+v)", *p)
}

// FlumeClient : client to use
type FlumeClient struct {
	host             string
	port             int
	tSocket          *thrift.TSocket
	transport        thrift.TTransport
	transportFactory thrift.TTransportFactory
	protocolFactory  *thrift.TCompactProtocolFactory

	thriftClient *ThriftSourceProtocolClient
}

func NewFlumeClient(host string, port int) *FlumeClient {
	return &FlumeClient{host: host, port: port}
}

func (fc *FlumeClient) connect() error {
	var tSocket *thrift.TSocket
	var err error
	tSocket, err = thrift.NewTSocketTimeout(net.JoinHostPort(fc.host, strconv.Itoa(fc.port)), 10*time.Second)
	if err != nil {
		return err
	}

	fc.tSocket = tSocket
	fc.transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	fc.protocolFactory = thrift.NewTCompactProtocolFactory()

	fc.clientConn()

	return nil
}

func (fc *FlumeClient) clientConn() error {
	fc.transport = fc.transportFactory.GetTransport(fc.tSocket)
	fc.thriftClient = NewThriftSourceProtocolClientFactory(fc.transport, fc.protocolFactory)
	err := fc.transport.Open()
	if err != nil {
		return err
	}
	return nil
}

func (fc *FlumeClient) AppendBatch(events []*ThriftFlumeEvent) error {
	fc.connect()
	defer fc.destroy()

	return fc.innerSend(func() (Status, error) {
		return fc.thriftClient.AppendBatch(events)
	})
}

func (fc *FlumeClient) Append(event *ThriftFlumeEvent) error {
	fc.connect()
	defer fc.destroy()

	return fc.innerSend(func() (Status, error) {
		return fc.thriftClient.Append(event)
	})
}

func (fc *FlumeClient) innerSend(sendfunc func() (Status, error)) error {
	if !fc.transport.IsOpen() {
		err := fc.clientConn()
		if nil != err {
			return err
		}
	}

	status, err := sendfunc()
	if err != nil {
		return err
	}

	if status != Status_OK {
		return errors.New("deliver fail ! |" + status.String())
	}
	return nil
}

func (fc *FlumeClient) destroy() {
	err := fc.transport.Close()
	if err != nil {
		log.Println(err.Error())
	}
}
