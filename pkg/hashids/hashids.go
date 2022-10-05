package hashids

import (
	"errors"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"github.com/speps/go-hashids/v2"
)

var (
	ErrHashIDType = errors.New("hashid 类型错误")
)

type HashIDType int

const (
	TypeUser HashIDType = iota
	TypeText
	TypeNamespace
	TypeSpeaker
)

var data *hashids.HashIDData

func Init() {
	data = &hashids.HashIDData{
		MinLength: 4,
		Salt:      conf.SystemConfig.HashIDSalt,
		Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	}
}

// HashEncode 对给定数据计算HashID
func HashEncode(v []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = conf.SystemConfig.HashIDSalt

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	id, err := h.Encode(v)
	if err != nil {
		return "", err
	}
	return id, nil
}

// HashDecode 对给定数据计算原始数据
func HashDecode(raw string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = conf.SystemConfig.HashIDSalt

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return []int{}, err
	}

	return h.DecodeWithError(raw)

}

// HashIDEncode 编码 HashID
func HashIDEncode(id uint, t HashIDType) string {
	str, err := HashEncode([]int{int(id), int(t)})
	if err != nil {
		log.Log().Errorf("HashIDEncode error, %s", err)
	}

	return str
}

// HashIDDecode 解码 HashID
func HashIDDecode(s string, t HashIDType) (uint, error) {
	r, _ := HashDecode(s)
	if len(r) != 2 || r[1] != int(t) {
		return 0, ErrHashIDType
	}

	return uint(r[0]), nil
}
