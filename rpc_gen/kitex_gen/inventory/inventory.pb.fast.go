// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package inventory

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *QueryStockReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryStockReq[number], err)
}

func (x *QueryStockReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint64
			v, offset, err = fastpb.ReadUint64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.SkuId = append(x.SkuId, v)
			return offset, err
		})
	return offset, err
}

func (x *QueryStockResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryStockResp[number], err)
}

func (x *QueryStockResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	if x.CurrentStock == nil {
		x.CurrentStock = make(map[uint64]uint32)
	}
	var key uint64
	var value uint32
	offset, err = fastpb.ReadMapEntry(buf, _type,
		func(buf []byte, _type int8) (offset int, err error) {
			key, offset, err = fastpb.ReadUint64(buf, _type)
			return offset, err
		},
		func(buf []byte, _type int8) (offset int, err error) {
			value, offset, err = fastpb.ReadUint32(buf, _type)
			return offset, err
		})
	if err != nil {
		return offset, err
	}
	x.CurrentStock[key] = value
	return offset, nil
}

func (x *InventoryReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InventoryReq[number], err)
}

func (x *InventoryReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OrderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *InventoryReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v InventoryReq_Item
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Items = append(x.Items, &v)
	return offset, nil
}

func (x *InventoryReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Force, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *InventoryResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InventoryResp[number], err)
}

func (x *InventoryResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *InventoryResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v InventoryResp_Error
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Errors = append(x.Errors, &v)
	return offset, nil
}

func (x *ProductCreatedEvent) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductCreatedEvent[number], err)
}

func (x *ProductCreatedEvent) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ProductCreatedEvent) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SkuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductCreatedEvent) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.InitialStock, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ProductCreatedEvent) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Operator, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductUpdateEvent) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductUpdateEvent[number], err)
}

func (x *ProductUpdateEvent) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ProductUpdateEvent) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SkuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductUpdateEvent) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UpdatedStock, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ProductUpdateEvent) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Operator, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductDeleteEvent) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductDeleteEvent[number], err)
}

func (x *ProductDeleteEvent) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ProductDeleteEvent) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SkuName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductDeleteEvent) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Operator, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *InventoryReq_Item) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InventoryReq_Item[number], err)
}

func (x *InventoryReq_Item) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *InventoryReq_Item) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Quantity, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *InventoryResp_Error) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_InventoryResp_Error[number], err)
}

func (x *InventoryResp_Error) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SkuId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *InventoryResp_Error) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *QueryStockReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *QueryStockReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.SkuId) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetSkuId()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, x.GetSkuId()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *QueryStockResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *QueryStockResp) fastWriteField3(buf []byte) (offset int) {
	if x.CurrentStock == nil {
		return offset
	}
	for k, v := range x.GetCurrentStock() {
		offset += fastpb.WriteMapEntry(buf[offset:], 3,
			func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
				offset := 0
				offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, k)
				offset += fastpb.WriteUint32(buf[offset:], numIdxOrVal, v)
				return offset
			})
	}
	return offset
}

func (x *InventoryReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *InventoryReq) fastWriteField1(buf []byte) (offset int) {
	if x.OrderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOrderId())
	return offset
}

func (x *InventoryReq) fastWriteField2(buf []byte) (offset int) {
	if x.Items == nil {
		return offset
	}
	for i := range x.GetItems() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetItems()[i])
	}
	return offset
}

func (x *InventoryReq) fastWriteField3(buf []byte) (offset int) {
	if !x.Force {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 3, x.GetForce())
	return offset
}

func (x *InventoryResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *InventoryResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *InventoryResp) fastWriteField2(buf []byte) (offset int) {
	if x.Errors == nil {
		return offset
	}
	for i := range x.GetErrors() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetErrors()[i])
	}
	return offset
}

func (x *ProductCreatedEvent) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ProductCreatedEvent) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *ProductCreatedEvent) fastWriteField2(buf []byte) (offset int) {
	if x.SkuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSkuName())
	return offset
}

func (x *ProductCreatedEvent) fastWriteField3(buf []byte) (offset int) {
	if x.InitialStock == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetInitialStock())
	return offset
}

func (x *ProductCreatedEvent) fastWriteField4(buf []byte) (offset int) {
	if x.Operator == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetOperator())
	return offset
}

func (x *ProductUpdateEvent) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ProductUpdateEvent) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *ProductUpdateEvent) fastWriteField2(buf []byte) (offset int) {
	if x.SkuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSkuName())
	return offset
}

func (x *ProductUpdateEvent) fastWriteField3(buf []byte) (offset int) {
	if x.UpdatedStock == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 3, x.GetUpdatedStock())
	return offset
}

func (x *ProductUpdateEvent) fastWriteField4(buf []byte) (offset int) {
	if x.Operator == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetOperator())
	return offset
}

func (x *ProductDeleteEvent) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ProductDeleteEvent) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *ProductDeleteEvent) fastWriteField2(buf []byte) (offset int) {
	if x.SkuName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSkuName())
	return offset
}

func (x *ProductDeleteEvent) fastWriteField4(buf []byte) (offset int) {
	if x.Operator == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetOperator())
	return offset
}

func (x *InventoryReq_Item) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *InventoryReq_Item) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *InventoryReq_Item) fastWriteField2(buf []byte) (offset int) {
	if x.Quantity == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetQuantity())
	return offset
}

func (x *InventoryResp_Error) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *InventoryResp_Error) fastWriteField1(buf []byte) (offset int) {
	if x.SkuId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetSkuId())
	return offset
}

func (x *InventoryResp_Error) fastWriteField2(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *QueryStockReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *QueryStockReq) sizeField1() (n int) {
	if len(x.SkuId) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetSkuId()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint64(numTagOrKey, x.GetSkuId()[numIdxOrVal])
			return n
		})
	return n
}

func (x *QueryStockResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField3()
	return n
}

func (x *QueryStockResp) sizeField3() (n int) {
	if x.CurrentStock == nil {
		return n
	}
	for k, v := range x.GetCurrentStock() {
		n += fastpb.SizeMapEntry(3,
			func(numTagOrKey, numIdxOrVal int32) int {
				n := 0
				n += fastpb.SizeUint64(numTagOrKey, k)
				n += fastpb.SizeUint32(numIdxOrVal, v)
				return n
			})
	}
	return n
}

func (x *InventoryReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *InventoryReq) sizeField1() (n int) {
	if x.OrderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOrderId())
	return n
}

func (x *InventoryReq) sizeField2() (n int) {
	if x.Items == nil {
		return n
	}
	for i := range x.GetItems() {
		n += fastpb.SizeMessage(2, x.GetItems()[i])
	}
	return n
}

func (x *InventoryReq) sizeField3() (n int) {
	if !x.Force {
		return n
	}
	n += fastpb.SizeBool(3, x.GetForce())
	return n
}

func (x *InventoryResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *InventoryResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *InventoryResp) sizeField2() (n int) {
	if x.Errors == nil {
		return n
	}
	for i := range x.GetErrors() {
		n += fastpb.SizeMessage(2, x.GetErrors()[i])
	}
	return n
}

func (x *ProductCreatedEvent) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ProductCreatedEvent) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *ProductCreatedEvent) sizeField2() (n int) {
	if x.SkuName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSkuName())
	return n
}

func (x *ProductCreatedEvent) sizeField3() (n int) {
	if x.InitialStock == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetInitialStock())
	return n
}

func (x *ProductCreatedEvent) sizeField4() (n int) {
	if x.Operator == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetOperator())
	return n
}

func (x *ProductUpdateEvent) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ProductUpdateEvent) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *ProductUpdateEvent) sizeField2() (n int) {
	if x.SkuName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSkuName())
	return n
}

func (x *ProductUpdateEvent) sizeField3() (n int) {
	if x.UpdatedStock == 0 {
		return n
	}
	n += fastpb.SizeUint32(3, x.GetUpdatedStock())
	return n
}

func (x *ProductUpdateEvent) sizeField4() (n int) {
	if x.Operator == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetOperator())
	return n
}

func (x *ProductDeleteEvent) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField4()
	return n
}

func (x *ProductDeleteEvent) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *ProductDeleteEvent) sizeField2() (n int) {
	if x.SkuName == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSkuName())
	return n
}

func (x *ProductDeleteEvent) sizeField4() (n int) {
	if x.Operator == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetOperator())
	return n
}

func (x *InventoryReq_Item) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *InventoryReq_Item) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *InventoryReq_Item) sizeField2() (n int) {
	if x.Quantity == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetQuantity())
	return n
}

func (x *InventoryResp_Error) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *InventoryResp_Error) sizeField1() (n int) {
	if x.SkuId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetSkuId())
	return n
}

func (x *InventoryResp_Error) sizeField2() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMessage())
	return n
}

var fieldIDToName_QueryStockReq = map[int32]string{
	1: "SkuId",
}

var fieldIDToName_QueryStockResp = map[int32]string{
	3: "CurrentStock",
}

var fieldIDToName_InventoryReq = map[int32]string{
	1: "OrderId",
	2: "Items",
	3: "Force",
}

var fieldIDToName_InventoryResp = map[int32]string{
	1: "Success",
	2: "Errors",
}

var fieldIDToName_ProductCreatedEvent = map[int32]string{
	1: "SkuId",
	2: "SkuName",
	3: "InitialStock",
	4: "Operator",
}

var fieldIDToName_ProductUpdateEvent = map[int32]string{
	1: "SkuId",
	2: "SkuName",
	3: "UpdatedStock",
	4: "Operator",
}

var fieldIDToName_ProductDeleteEvent = map[int32]string{
	1: "SkuId",
	2: "SkuName",
	4: "Operator",
}

var fieldIDToName_InventoryReq_Item = map[int32]string{
	1: "SkuId",
	2: "Quantity",
}

var fieldIDToName_InventoryResp_Error = map[int32]string{
	1: "SkuId",
	2: "Message",
}
