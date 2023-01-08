package id

import (
	"strconv"
)

type Uid int64

func (id Uid) To() int64 {
	return int64(id)
}

func To(id int64) Uid {
	return Uid(id)
}

func (id Uid) ToString() string {
	return strconv.FormatInt(int64(id), 10)
}

func FromInt64(v int64) Uid {
	return Uid(v)
}

func FromString(s string) (Uid, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1, err
	}
	return Uid(n), nil
}
