//Package propertyf comment
// This file war generated by tars2go 1.1
// Generated from PropertyF.tars
package propertyf

import (
	"fmt"
	"tars/protocol/codec"
)

//StatPropInfo strcut implement
type StatPropInfo struct {
	Policy string `json:"policy"`
	Value  string `json:"value"`
}

func (st *StatPropInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *StatPropInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Policy, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Value, 1, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *StatPropInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatPropInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *StatPropInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Policy, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Value, 1)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *StatPropInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
