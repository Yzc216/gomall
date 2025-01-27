// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Phone, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint32
			v, offset, err = fastpb.ReadUint32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Role = append(x.Role, v)
			return offset, err
		})
	return offset, err
}

func (x *RegisterReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterReq[number], err)
}

func (x *RegisterReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserInfo = &v
	return offset, nil
}

func (x *RegisterReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PasswordConfirm, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterResp[number], err)
}

func (x *RegisterResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *RegisterResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint32
			v, offset, err = fastpb.ReadUint32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Role = append(x.Role, v)
			return offset, err
		})
	return offset, err
}

func (x *LoginReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginReq[number], err)
}

func (x *LoginReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LoginInfo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResp[number], err)
}

func (x *LoginResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *LoginResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint32
			v, offset, err = fastpb.ReadUint32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Role = append(x.Role, v)
			return offset, err
		})
	return offset, err
}

func (x *ResetPasswordReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ResetPasswordReq[number], err)
}

func (x *ResetPasswordReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *ResetPasswordReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ResetPasswordReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.NewPassword, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ResetPasswordResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ResetPasswordResp[number], err)
}

func (x *ResetPasswordResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsReset, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *SetUserRoleReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetUserRoleReq[number], err)
}

func (x *SetUserRoleReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *SetUserRoleReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint32
			v, offset, err = fastpb.ReadUint32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.NewRole = append(x.NewRole, v)
			return offset, err
		})
	return offset, err
}

func (x *SetUserRoleResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetUserRoleResp[number], err)
}

func (x *SetUserRoleResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsSet, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *UpdateUserInfoReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserInfoReq[number], err)
}

func (x *UpdateUserInfoReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserInfo = &v
	return offset, nil
}

func (x *UpdateUserInfoReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *UpdateUserInfoResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserInfoResp[number], err)
}

func (x *UpdateUserInfoResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsUpdated, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetUserInfoReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoReq[number], err)
}

func (x *GetUserInfoReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *GetUserInfoResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoResp[number], err)
}

func (x *GetUserInfoResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserInfo = &v
	return offset, nil
}

func (x *GetUserInfoListReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoListReq[number], err)
}

func (x *GetUserInfoListReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetUserInfoListReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetUserInfoListReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v uint64
			v, offset, err = fastpb.ReadUint64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.UserIds = append(x.UserIds, v)
			return offset, err
		})
	return offset, err
}

func (x *GetUserInfoListResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoListResp[number], err)
}

func (x *GetUserInfoListResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserInfos = append(x.UserInfos, &v)
	return offset, nil
}

func (x *GetUserInfoListResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetUserInfoListResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetUserInfoListResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BanUserReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BanUserReq[number], err)
}

func (x *BanUserReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadUint64(buf, _type)
	return offset, err
}

func (x *BanUserResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BanUserResp[number], err)
}

func (x *BanUserResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsBan, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetAvatar())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if x.Phone == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPhone())
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetEmail())
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if len(x.Role) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 7, len(x.GetRole()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint32(buf[offset:], numTagOrKey, x.GetRole()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *RegisterReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RegisterReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserInfo())
	return offset
}

func (x *RegisterReq) fastWriteField3(buf []byte) (offset int) {
	if x.PasswordConfirm == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPasswordConfirm())
	return offset
}

func (x *RegisterResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RegisterResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *RegisterResp) fastWriteField2(buf []byte) (offset int) {
	if len(x.Role) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetRole()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint32(buf[offset:], numTagOrKey, x.GetRole()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *LoginReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *LoginReq) fastWriteField1(buf []byte) (offset int) {
	if x.LoginInfo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLoginInfo())
	return offset
}

func (x *LoginReq) fastWriteField4(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPassword())
	return offset
}

func (x *LoginResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *LoginResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *LoginResp) fastWriteField2(buf []byte) (offset int) {
	if len(x.Role) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetRole()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint32(buf[offset:], numTagOrKey, x.GetRole()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *ResetPasswordReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ResetPasswordReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *ResetPasswordReq) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *ResetPasswordReq) fastWriteField3(buf []byte) (offset int) {
	if x.NewPassword == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetNewPassword())
	return offset
}

func (x *ResetPasswordResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ResetPasswordResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsReset {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsReset())
	return offset
}

func (x *SetUserRoleReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SetUserRoleReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *SetUserRoleReq) fastWriteField2(buf []byte) (offset int) {
	if len(x.NewRole) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 2, len(x.GetNewRole()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint32(buf[offset:], numTagOrKey, x.GetNewRole()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *SetUserRoleResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SetUserRoleResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsSet {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsSet())
	return offset
}

func (x *UpdateUserInfoReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpdateUserInfoReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserInfo())
	return offset
}

func (x *UpdateUserInfoReq) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *UpdateUserInfoResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateUserInfoResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsUpdated {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsUpdated())
	return offset
}

func (x *GetUserInfoReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserInfoReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetUserInfoResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserInfoResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserInfo())
	return offset
}

func (x *GetUserInfoListReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetUserInfoListReq) fastWriteField1(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetPage())
	return offset
}

func (x *GetUserInfoListReq) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetPageSize())
	return offset
}

func (x *GetUserInfoListReq) fastWriteField3(buf []byte) (offset int) {
	if len(x.UserIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 3, len(x.GetUserIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteUint64(buf[offset:], numTagOrKey, x.GetUserIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *GetUserInfoListResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetUserInfoListResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserInfos == nil {
		return offset
	}
	for i := range x.GetUserInfos() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserInfos()[i])
	}
	return offset
}

func (x *GetUserInfoListResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetUserInfoListResp) fastWriteField3(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetPage())
	return offset
}

func (x *GetUserInfoListResp) fastWriteField4(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetPageSize())
	return offset
}

func (x *BanUserReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *BanUserReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteUint64(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *BanUserResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *BanUserResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsBan {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsBan())
	return offset
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetAvatar())
	return n
}

func (x *User) sizeField5() (n int) {
	if x.Phone == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPhone())
	return n
}

func (x *User) sizeField6() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetEmail())
	return n
}

func (x *User) sizeField7() (n int) {
	if len(x.Role) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(7, len(x.GetRole()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint32(numTagOrKey, x.GetRole()[numIdxOrVal])
			return n
		})
	return n
}

func (x *RegisterReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	return n
}

func (x *RegisterReq) sizeField1() (n int) {
	if x.UserInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUserInfo())
	return n
}

func (x *RegisterReq) sizeField3() (n int) {
	if x.PasswordConfirm == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPasswordConfirm())
	return n
}

func (x *RegisterResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RegisterResp) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *RegisterResp) sizeField2() (n int) {
	if len(x.Role) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetRole()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint32(numTagOrKey, x.GetRole()[numIdxOrVal])
			return n
		})
	return n
}

func (x *LoginReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField4()
	return n
}

func (x *LoginReq) sizeField1() (n int) {
	if x.LoginInfo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetLoginInfo())
	return n
}

func (x *LoginReq) sizeField4() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPassword())
	return n
}

func (x *LoginResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *LoginResp) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *LoginResp) sizeField2() (n int) {
	if len(x.Role) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetRole()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint32(numTagOrKey, x.GetRole()[numIdxOrVal])
			return n
		})
	return n
}

func (x *ResetPasswordReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ResetPasswordReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *ResetPasswordReq) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *ResetPasswordReq) sizeField3() (n int) {
	if x.NewPassword == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetNewPassword())
	return n
}

func (x *ResetPasswordResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ResetPasswordResp) sizeField1() (n int) {
	if !x.IsReset {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsReset())
	return n
}

func (x *SetUserRoleReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SetUserRoleReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *SetUserRoleReq) sizeField2() (n int) {
	if len(x.NewRole) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(2, len(x.GetNewRole()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint32(numTagOrKey, x.GetNewRole()[numIdxOrVal])
			return n
		})
	return n
}

func (x *SetUserRoleResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SetUserRoleResp) sizeField1() (n int) {
	if !x.IsSet {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsSet())
	return n
}

func (x *UpdateUserInfoReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpdateUserInfoReq) sizeField1() (n int) {
	if x.UserInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUserInfo())
	return n
}

func (x *UpdateUserInfoReq) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(2, x.GetUserId())
	return n
}

func (x *UpdateUserInfoResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateUserInfoResp) sizeField1() (n int) {
	if !x.IsUpdated {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsUpdated())
	return n
}

func (x *GetUserInfoReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserInfoReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *GetUserInfoResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserInfoResp) sizeField1() (n int) {
	if x.UserInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUserInfo())
	return n
}

func (x *GetUserInfoListReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetUserInfoListReq) sizeField1() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetPage())
	return n
}

func (x *GetUserInfoListReq) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetPageSize())
	return n
}

func (x *GetUserInfoListReq) sizeField3() (n int) {
	if len(x.UserIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(3, len(x.GetUserIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeUint64(numTagOrKey, x.GetUserIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *GetUserInfoListResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetUserInfoListResp) sizeField1() (n int) {
	if x.UserInfos == nil {
		return n
	}
	for i := range x.GetUserInfos() {
		n += fastpb.SizeMessage(1, x.GetUserInfos()[i])
	}
	return n
}

func (x *GetUserInfoListResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetTotal())
	return n
}

func (x *GetUserInfoListResp) sizeField3() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetPage())
	return n
}

func (x *GetUserInfoListResp) sizeField4() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetPageSize())
	return n
}

func (x *BanUserReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *BanUserReq) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeUint64(1, x.GetUserId())
	return n
}

func (x *BanUserResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *BanUserResp) sizeField1() (n int) {
	if !x.IsBan {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsBan())
	return n
}

var fieldIDToName_User = map[int32]string{
	1: "Username",
	2: "Password",
	4: "Avatar",
	5: "Phone",
	6: "Email",
	7: "Role",
}

var fieldIDToName_RegisterReq = map[int32]string{
	1: "UserInfo",
	3: "PasswordConfirm",
}

var fieldIDToName_RegisterResp = map[int32]string{
	1: "UserId",
	2: "Role",
}

var fieldIDToName_LoginReq = map[int32]string{
	1: "LoginInfo",
	4: "Password",
}

var fieldIDToName_LoginResp = map[int32]string{
	1: "UserId",
	2: "Role",
}

var fieldIDToName_ResetPasswordReq = map[int32]string{
	1: "UserId",
	2: "Password",
	3: "NewPassword",
}

var fieldIDToName_ResetPasswordResp = map[int32]string{
	1: "IsReset",
}

var fieldIDToName_SetUserRoleReq = map[int32]string{
	1: "UserId",
	2: "NewRole",
}

var fieldIDToName_SetUserRoleResp = map[int32]string{
	1: "IsSet",
}

var fieldIDToName_UpdateUserInfoReq = map[int32]string{
	1: "UserInfo",
	2: "UserId",
}

var fieldIDToName_UpdateUserInfoResp = map[int32]string{
	1: "IsUpdated",
}

var fieldIDToName_GetUserInfoReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetUserInfoResp = map[int32]string{
	1: "UserInfo",
}

var fieldIDToName_GetUserInfoListReq = map[int32]string{
	1: "Page",
	2: "PageSize",
	3: "UserIds",
}

var fieldIDToName_GetUserInfoListResp = map[int32]string{
	1: "UserInfos",
	2: "Total",
	3: "Page",
	4: "PageSize",
}

var fieldIDToName_BanUserReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_BanUserResp = map[int32]string{
	1: "IsBan",
}
