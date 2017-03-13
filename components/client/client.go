package client

import (
    "errors"
    "flume-client/components/flume"
    "git.apache.org/thrift.git/lib/go/thrift"
    "log"
    "net"
    "strconv"
    "time"
)

type FlumeClient struct {
    host             string
    port             int
    tSocket          *thrift.TSocket
    transport        thrift.TTransport
    transportFactory thrift.TTransportFactory
    protocolFactory  *thrift.TCompactProtocolFactory

    thriftClient     *flume.ThriftSourceProtocolClient
}

func NewFlumeClient(host string, port int) *FlumeClient {
    return &FlumeClient{host: host, port: port}
}

func (fc *FlumeClient) connect() error {
    var tSocket *thrift.TSocket
    var err error
    tSocket, err = thrift.NewTSocketTimeout(net.JoinHostPort(fc.host, strconv.Itoa(fc.port)), 10 * time.Second)
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
    fc.thriftClient = flume.NewThriftSourceProtocolClientFactory(fc.transport, fc.protocolFactory)
    err := fc.transport.Open()
    if err != nil {
        return err
    }
    return nil
}

func (fc *FlumeClient) AppendBatch(events []*flume.ThriftFlumeEvent) error {
    fc.connect()
    defer fc.destroy()

    return fc.innerSend(func() (flume.Status, error) {
        return fc.thriftClient.AppendBatch(events)
    })
}

func (fc *FlumeClient) Append(event *flume.ThriftFlumeEvent) error {
    fc.connect()
    defer fc.destroy()

    return fc.innerSend(func() (flume.Status, error) {
        return fc.thriftClient.Append(event)
    })
}

func (fc *FlumeClient) innerSend(sendfunc func() (flume.Status, error)) error {
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

    if status != flume.Status_OK {
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
