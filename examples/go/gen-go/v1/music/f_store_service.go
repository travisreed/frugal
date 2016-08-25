// Autogenerated by Frugal Compiler (1.14.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package music

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/Workiva/frugal/lib/go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type FStore interface {
	BuyAlbum(ctx *frugal.FContext, ASIN string, acct string) (r *Album, err error)
	EnterAlbumGiveaway(ctx *frugal.FContext, email string, name string) (r bool, err error)
}

type FStoreClient struct {
	transport       frugal.FTransport
	protocolFactory *frugal.FProtocolFactory
	oprot           *frugal.FProtocol
	mu              sync.Mutex
	methods         map[string]*frugal.Method
}

func NewFStoreClient(t frugal.FTransport, p *frugal.FProtocolFactory, middleware ...frugal.ServiceMiddleware) *FStoreClient {
	t.SetRegistry(frugal.NewFClientRegistry())
	methods := make(map[string]*frugal.Method)
	client := &FStoreClient{
		transport:       t,
		protocolFactory: p,
		oprot:           p.GetProtocol(t),
		methods:         methods,
	}
	methods["buyAlbum"] = frugal.NewMethod(client, client.buyAlbum, "buyAlbum", middleware)
	methods["enterAlbumGiveaway"] = frugal.NewMethod(client, client.enterAlbumGiveaway, "enterAlbumGiveaway", middleware)
	return client
}

// Do Not Use. To be called only by generated code.
func (f *FStoreClient) GetWriteMutex() *sync.Mutex {
	return &f.mu
}

func (f *FStoreClient) BuyAlbum(ctx *frugal.FContext, asin string, acct string) (r *Album, err error) {
	ret := f.methods["buyAlbum"].Invoke([]interface{}{ctx, asin, acct})
	if len(ret) != 2 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 2", len(ret)))
	}
	r = ret[0].(*Album)
	if ret[1] != nil {
		err = ret[1].(error)
	}
	return r, err
}

func (f *FStoreClient) buyAlbum(ctx *frugal.FContext, asin string, acct string) (r *Album, err error) {
	errorC := make(chan error, 1)
	resultC := make(chan *Album, 1)
	if err = f.transport.Register(ctx, f.recvBuyAlbumHandler(ctx, resultC, errorC)); err != nil {
		return
	}
	defer f.transport.Unregister(ctx)
	f.GetWriteMutex().Lock()
	if err = f.oprot.WriteRequestHeader(ctx); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.WriteMessageBegin("buyAlbum", thrift.CALL, 0); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	args := StoreBuyAlbumArgs{
		ASIN: asin,
		Acct: acct,
	}
	if err = args.Write(f.oprot); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.WriteMessageEnd(); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.Flush(); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	f.GetWriteMutex().Unlock()

	select {
	case err = <-errorC:
	case r = <-resultC:
	case <-time.After(ctx.Timeout()):
		err = frugal.ErrTimeout
	case <-f.transport.Closed():
		err = frugal.ErrTransportClosed
	}
	return
}

func (f *FStoreClient) recvBuyAlbumHandler(ctx *frugal.FContext, resultC chan<- *Album, errorC chan<- error) frugal.FAsyncCallback {
	return func(tr thrift.TTransport) error {
		iprot := f.protocolFactory.GetProtocol(tr)
		if err := iprot.ReadResponseHeader(ctx); err != nil {
			errorC <- err
			return err
		}
		method, mTypeId, _, err := iprot.ReadMessageBegin()
		if err != nil {
			errorC <- err
			return err
		}
		if method != "buyAlbum" {
			err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "buyAlbum failed: wrong method name")
			errorC <- err
			return err
		}
		if mTypeId == thrift.EXCEPTION {
			error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
			var error1 thrift.TApplicationException
			error1, err = error0.Read(iprot)
			if err != nil {
				errorC <- err
				return err
			}
			if err = iprot.ReadMessageEnd(); err != nil {
				errorC <- err
				return err
			}
			if error1.TypeId() == frugal.RESPONSE_TOO_LARGE {
				err = thrift.NewTTransportException(frugal.RESPONSE_TOO_LARGE, "response too large for transport")
				errorC <- err
				return nil
			}
			err = error1
			errorC <- err
			return err
		}
		if mTypeId != thrift.REPLY {
			err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "buyAlbum failed: invalid message type")
			errorC <- err
			return err
		}
		result := StoreBuyAlbumResult{}
		if err = result.Read(iprot); err != nil {
			errorC <- err
			return err
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			errorC <- err
			return err
		}
		if result.Error != nil {
			errorC <- result.Error
			return nil
		}
		resultC <- result.GetSuccess()
		return nil
	}
}

func (f *FStoreClient) EnterAlbumGiveaway(ctx *frugal.FContext, email string, name string) (r bool, err error) {
	ret := f.methods["enterAlbumGiveaway"].Invoke([]interface{}{ctx, email, name})
	if len(ret) != 2 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 2", len(ret)))
	}
	r = ret[0].(bool)
	if ret[1] != nil {
		err = ret[1].(error)
	}
	return r, err
}

func (f *FStoreClient) enterAlbumGiveaway(ctx *frugal.FContext, email string, name string) (r bool, err error) {
	errorC := make(chan error, 1)
	resultC := make(chan bool, 1)
	if err = f.transport.Register(ctx, f.recvEnterAlbumGiveawayHandler(ctx, resultC, errorC)); err != nil {
		return
	}
	defer f.transport.Unregister(ctx)
	f.GetWriteMutex().Lock()
	if err = f.oprot.WriteRequestHeader(ctx); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.WriteMessageBegin("enterAlbumGiveaway", thrift.CALL, 0); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	args := StoreEnterAlbumGiveawayArgs{
		Email: email,
		Name:  name,
	}
	if err = args.Write(f.oprot); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.WriteMessageEnd(); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	if err = f.oprot.Flush(); err != nil {
		f.GetWriteMutex().Unlock()
		return
	}
	f.GetWriteMutex().Unlock()

	select {
	case err = <-errorC:
	case r = <-resultC:
	case <-time.After(ctx.Timeout()):
		err = frugal.ErrTimeout
	case <-f.transport.Closed():
		err = frugal.ErrTransportClosed
	}
	return
}

func (f *FStoreClient) recvEnterAlbumGiveawayHandler(ctx *frugal.FContext, resultC chan<- bool, errorC chan<- error) frugal.FAsyncCallback {
	return func(tr thrift.TTransport) error {
		iprot := f.protocolFactory.GetProtocol(tr)
		if err := iprot.ReadResponseHeader(ctx); err != nil {
			errorC <- err
			return err
		}
		method, mTypeId, _, err := iprot.ReadMessageBegin()
		if err != nil {
			errorC <- err
			return err
		}
		if method != "enterAlbumGiveaway" {
			err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "enterAlbumGiveaway failed: wrong method name")
			errorC <- err
			return err
		}
		if mTypeId == thrift.EXCEPTION {
			error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
			var error1 thrift.TApplicationException
			error1, err = error0.Read(iprot)
			if err != nil {
				errorC <- err
				return err
			}
			if err = iprot.ReadMessageEnd(); err != nil {
				errorC <- err
				return err
			}
			if error1.TypeId() == frugal.RESPONSE_TOO_LARGE {
				err = thrift.NewTTransportException(frugal.RESPONSE_TOO_LARGE, "response too large for transport")
				errorC <- err
				return nil
			}
			err = error1
			errorC <- err
			return err
		}
		if mTypeId != thrift.REPLY {
			err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "enterAlbumGiveaway failed: invalid message type")
			errorC <- err
			return err
		}
		result := StoreEnterAlbumGiveawayResult{}
		if err = result.Read(iprot); err != nil {
			errorC <- err
			return err
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			errorC <- err
			return err
		}
		resultC <- result.GetSuccess()
		return nil
	}
}

type FStoreProcessor struct {
	*frugal.FBaseProcessor
}

func NewFStoreProcessor(handler FStore, middleware ...frugal.ServiceMiddleware) *FStoreProcessor {
	p := &FStoreProcessor{frugal.NewFBaseProcessor()}
	p.AddToProcessorMap("buyAlbum", &storeFBuyAlbum{handler: frugal.NewMethod(handler, handler.BuyAlbum, "BuyAlbum", middleware), writeMu: p.GetWriteMutex()})
	p.AddToProcessorMap("enterAlbumGiveaway", &storeFEnterAlbumGiveaway{handler: frugal.NewMethod(handler, handler.EnterAlbumGiveaway, "EnterAlbumGiveaway", middleware), writeMu: p.GetWriteMutex()})
	return p
}

type storeFBuyAlbum struct {
	handler *frugal.Method
	writeMu *sync.Mutex
}

func (p *storeFBuyAlbum) Process(ctx *frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := StoreBuyAlbumArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		p.writeMu.Lock()
		storeWriteApplicationError(ctx, oprot, thrift.PROTOCOL_ERROR, "buyAlbum", err.Error())
		p.writeMu.Unlock()
		return err
	}

	iprot.ReadMessageEnd()
	result := StoreBuyAlbumResult{}
	var err2 error
	var retval *Album
	ret := p.handler.Invoke([]interface{}{ctx, args.ASIN, args.Acct})
	if len(ret) != 2 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 2", len(ret)))
	}
	retval = ret[0].(*Album)
	if ret[1] != nil {
		err2 = ret[1].(error)
	}
	if err2 != nil {
		switch v := err2.(type) {
		case *PurchasingError:
			result.Error = v
		default:
			p.writeMu.Lock()
			storeWriteApplicationError(ctx, oprot, thrift.INTERNAL_ERROR, "buyAlbum", "Internal error processing buyAlbum: "+err2.Error())
			p.writeMu.Unlock()
			return err2
		}
	} else {
		result.Success = retval
	}
	p.writeMu.Lock()
	defer p.writeMu.Unlock()
	if err2 = oprot.WriteResponseHeader(ctx); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "buyAlbum", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageBegin("buyAlbum", thrift.REPLY, 0); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "buyAlbum", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "buyAlbum", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "buyAlbum", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "buyAlbum", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	return err
}

type storeFEnterAlbumGiveaway struct {
	handler *frugal.Method
	writeMu *sync.Mutex
}

func (p *storeFEnterAlbumGiveaway) Process(ctx *frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := StoreEnterAlbumGiveawayArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		p.writeMu.Lock()
		storeWriteApplicationError(ctx, oprot, thrift.PROTOCOL_ERROR, "enterAlbumGiveaway", err.Error())
		p.writeMu.Unlock()
		return err
	}

	iprot.ReadMessageEnd()
	result := StoreEnterAlbumGiveawayResult{}
	var err2 error
	var retval bool
	ret := p.handler.Invoke([]interface{}{ctx, args.Email, args.Name})
	if len(ret) != 2 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 2", len(ret)))
	}
	retval = ret[0].(bool)
	if ret[1] != nil {
		err2 = ret[1].(error)
	}
	if err2 != nil {
		p.writeMu.Lock()
		storeWriteApplicationError(ctx, oprot, thrift.INTERNAL_ERROR, "enterAlbumGiveaway", "Internal error processing enterAlbumGiveaway: "+err2.Error())
		p.writeMu.Unlock()
		return err2
	} else {
		result.Success = &retval
	}
	p.writeMu.Lock()
	defer p.writeMu.Unlock()
	if err2 = oprot.WriteResponseHeader(ctx); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "enterAlbumGiveaway", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageBegin("enterAlbumGiveaway", thrift.REPLY, 0); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "enterAlbumGiveaway", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "enterAlbumGiveaway", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "enterAlbumGiveaway", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			storeWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "enterAlbumGiveaway", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	return err
}

func storeWriteApplicationError(ctx *frugal.FContext, oprot *frugal.FProtocol, type_ int32, method, message string) {
	x := thrift.NewTApplicationException(type_, message)
	oprot.WriteResponseHeader(ctx)
	oprot.WriteMessageBegin(method, thrift.EXCEPTION, 0)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
}