//Package logf comment
// This file war generated by tars2go 1.1
// Generated from LogF.tars
package logf

import (
	"context"
	"fmt"
	m "tars/model"
	"tars/protocol/codec"
	"tars/protocol/res/requestf"
	"tars/util/current"
	"tars/util/tools"
)

//Log struct
type Log struct {
	s m.Servant
}

//Logger is the proxy function for the method defined in the tars file, with the context
func (_obj *Log) Logger(App string, Server string, File string, Format string, Buffer []string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(App, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(Server, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(File, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(Format, 4)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(Buffer)), 0)
	if err != nil {
		return err
	}
	for _, v := range Buffer {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "logger", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoggerWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Log) LoggerWithContext(ctx context.Context, App string, Server string, File string, Format string, Buffer []string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(App, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(Server, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(File, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(Format, 4)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(Buffer)), 0)
	if err != nil {
		return err
	}
	for _, v := range Buffer {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "logger", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoggerbyInfo is the proxy function for the method defined in the tars file, with the context
func (_obj *Log) LoggerbyInfo(Info *LogInfo, Buffer []string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Info.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(Buffer)), 0)
	if err != nil {
		return err
	}
	for _, v := range Buffer {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "loggerbyInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoggerbyInfoWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *Log) LoggerbyInfoWithContext(ctx context.Context, Info *LogInfo, Buffer []string, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Info.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(Buffer)), 0)
	if err != nil {
		return err
	}
	for _, v := range Buffer {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "loggerbyInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return nil
}

//SetServant sets servant for the service.
func (_obj *Log) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *Log) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

type _impLog interface {
	Logger(App string, Server string, File string, Format string, Buffer []string) (err error)
	LoggerbyInfo(Info *LogInfo, Buffer []string) (err error)
}
type _impLogWithContext interface {
	Logger(ctx context.Context, App string, Server string, File string, Format string, Buffer []string) (err error)
	LoggerbyInfo(ctx context.Context, Info *LogInfo, Buffer []string) (err error)
}

func logger(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var App string
	err = _is.Read_string(&App, 1, true)
	if err != nil {
		return err
	}
	var Server string
	err = _is.Read_string(&Server, 2, true)
	if err != nil {
		return err
	}
	var File string
	err = _is.Read_string(&File, 3, true)
	if err != nil {
		return err
	}
	var Format string
	err = _is.Read_string(&Format, 4, true)
	if err != nil {
		return err
	}
	var Buffer []string
	err, _, ty = _is.SkipToNoCheck(5, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		Buffer = make([]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_string(&Buffer[i0], 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}
	if withContext == false {
		_imp := _val.(_impLog)
		err = _imp.Logger(App, Server, File, Format, Buffer)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impLogWithContext)
		err = _imp.Logger(ctx, App, Server, File, Format, Buffer)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func loggerbyInfo(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Info LogInfo
	err = Info.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}
	var Buffer []string
	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		Buffer = make([]string, length, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = _is.Read_string(&Buffer[i1], 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}
	if withContext == false {
		_imp := _val.(_impLog)
		err = _imp.LoggerbyInfo(&Info, Buffer)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impLogWithContext)
		err = _imp.LoggerbyInfo(ctx, &Info, Buffer)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *Log) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "logger":
		err := logger(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "loggerbyInfo":
		err := loggerbyInfo(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
