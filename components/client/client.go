package client

import (
	"errors"
    "flume-client/components/flume"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"net"
	"strconv"
	"time"
)

type Status int32

const (
	STATUS_INIT  Status = 0
	STATUS_READY Status = 1
	STATUS_DEAD  Status = 2
)

type FlumeClient struct {
	host             string
	port             int
	tsocket          *thrift.TSocket
	transport        thrift.TTransport
	transportFactory thrift.TTransportFactory
	protocolFactory  *thrift.TCompactProtocolFactory

	thriftclient *flume.ThriftSourceProtocolClient
	status       Status //连接状态
}

func NewFlumeClient(host string, port int) *FlumeClient {
	return &FlumeClient{host: host, port: port, status: STATUS_INIT}
}

func (fc *FlumeClient) IsAlive() bool {
	return fc.status == STATUS_READY
}

func (fc *FlumeClient) Connect() error {

	var tsocket *thrift.TSocket
	var err error
	//创建一个物理连接
	tsocket, err = thrift.NewTSocketTimeout(net.JoinHostPort(fc.host, strconv.Itoa(fc.port)), 10*time.Second)
	if nil != err {
		log.Printf("FLUME_CLIENT|CREATE TSOCKET|FAIL|%s|%s\n", fc.HostPort(), err)
		return err
	}

	fc.tsocket = tsocket
	fc.transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//TLV 方式传输
	fc.protocolFactory = thrift.NewTCompactProtocolFactory()

	fc.clientConn()
	fc.status = STATUS_READY
	go fc.checkAlive()

	return nil
}

func (fc *FlumeClient) clientConn() error {
	//使用非阻塞io来传输
	fc.transport = fc.transportFactory.GetTransport(fc.tsocket)
	fc.thriftclient = flume.NewThriftSourceProtocolClientFactory(fc.transport, fc.protocolFactory)
	if err := fc.transport.Open(); nil != err {
		log.Printf("FLUME_CLIENT|CREATE THRIFT CLIENT|FAIL|%s|%s", fc.HostPort(), err)
		return err
	}
	return nil
}

func (fc *FlumeClient) checkAlive() {
	for fc.status != STATUS_DEAD {
		//休息1s
		time.Sleep(1 * time.Second)
		isOpen := fc.tsocket.IsOpen()
		if !isOpen {
			fc.status = STATUS_DEAD
			log.Printf("flume : %s:%d is Dead", fc.host, fc.port)
			break
		}

	}
}

func (fc *FlumeClient) AppendBatch(events []*flume.ThriftFlumeEvent) error {
	return fc.innerSend(func() (flume.Status, error) {
		return fc.thriftclient.AppendBatch(events)
	})
}

func (fc *FlumeClient) Append(event *flume.ThriftFlumeEvent) error {
	return fc.innerSend(func() (flume.Status, error) {
		return fc.thriftclient.Append(event)
	})
}

func (fc *FlumeClient) innerSend(sendfunc func() (flume.Status, error)) error {
	if fc.status == STATUS_DEAD {
		return errors.New("FLUME_CLIENT|DEAD|" + fc.HostPort())
	}

	//如果transport关闭了那么久重新打开
	if !fc.transport.IsOpen() {
		//重新建立thriftclient
		err := fc.clientConn()
		if nil != err {
			log.Printf("FLUME_CLIENT|SEND EVENT|CLIENT CONN|CREATE FAIL|%s|%s\n", fc.HostPort(), err.Error())
			return err
		}
	}

	status, err := sendfunc()
	if nil != err {
		log.Printf("FLUME_CLIENT|SEND EVENT|FAIL|%s|%s|STATUS:%s\n", fc.HostPort(), err.Error(), status)
		status = flume.Status_ERROR
		fc.status = STATUS_DEAD
	}

	//如果没有成功则向上抛出
	if status != flume.Status_OK {
		return errors.New("deliver fail ! " + fc.HostPort() + "|" + status.String())
	}
	return nil
}

func (fc *FlumeClient) Destroy() {

	fc.status = STATUS_DEAD
	err := fc.transport.Close()
	if nil != err {
		log.Println(err.Error())
	}

}

func (fc *FlumeClient) HostPort() string {
	return fmt.Sprintf("[%s:%d-%d]", fc.host, fc.port, fc.status)
}
