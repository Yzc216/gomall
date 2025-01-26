package util

import "github.com/sony/sonyflake"

func GenID() uint64 {
	// 创建 Sonyflake 实例
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}

	// 生成唯一 ID
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}
