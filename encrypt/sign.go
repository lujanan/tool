/**
自定义签名
*/
package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
)

// GenerateSign 连接字符串，MD5加密
// suKey 加密后缀
func GenerateSign(strMap map[string]string, suKey string) string {
	strList := strMapSort(strMap)
	if strList == nil {
		return ""
	}

	signStr := ""
	for _, value := range strList {
		signStr += value.Key + value.Val
	}
	if suKey != "" {
		signStr += suKey
	}

	h := md5.New()
	io.WriteString(h, signStr)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// CheckSign 校验加密串
func CheckSign(strMap map[string]string, sign, suKey string) bool {
	thisSign := GenerateSign(strMap, suKey)
	if thisSign == sign {
		return true
	}
	return false
}

// keyval 自定义排序
type keyval struct {
	Key string
	Val string
}

type strMapList []keyval

func (list strMapList) Len() int {
	return len(list)
}

func (list strMapList) Less(i, j int) bool {
	return list[i].Key < list[j].Key
}

func (list strMapList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// strMapSort 字符串类型字典顺序排序
func strMapSort(strMap map[string]string) []keyval {
	if len(strMap) <= 0 {
		return nil
	}

	list := strMapList{}
	for key, val := range strMap {
		list = append(list, keyval{Key: key, Val: val})
	}
	sort.Sort(list)
	return list
}
