package profiling

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 创建请求字符串
func createRequest() string {
	payload := make([]int, 100, 100) // 创建一个大小为 100 的整型切片
	for i := 0; i < 100; i++ {
		payload[i] = i // 将 0 到 99 的值赋值给切片
	}
	req := Request{"demo_transaction", payload} // 创建请求对象
	v, err := json.Marshal(&req)                // 序列化请求对象为 JSON
	if err != nil {
		panic(err) // 如果发生错误, 则抛出异常
	}
	return string(v) // 返回 JSON 字符串
}

// 处理请求
func processRequest(reqs []string) []string {
	reps := []string{} // 创建一个空的字符串切片, 用于存储响应
	for _, req := range reqs {
		reqObj := &Request{}              // 创建一个空的请求对象
		reqObj.UnmarshalJSON([]byte(req)) // 反序列化 JSON 字符串到请求对象
		//	json.Unmarshal([]byte(req), reqObj)

		var buf strings.Builder            // 创建一个字符串构建器
		for _, e := range reqObj.PayLoad { // 遍历请求负载
			buf.WriteString(strconv.Itoa(e)) // 将整型转换为字符串并追加到构建器
			buf.WriteString(",")             // 追加逗号
		}
		repObj := &Response{reqObj.TransactionID, buf.String()} // 创建响应对象
		repJson, err := repObj.MarshalJSON()                    // 序列化响应对象为 JSON
		//repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err) // 如果发生错误, 则抛出异常
		}
		reps = append(reps, string(repJson)) // 将响应 JSON 字符串追加到响应切片
	}
	return reps // 返回响应切片
}

// 旧版处理请求(未优化)
func processRequestOld(reqs []string) []string {
	reps := []string{} // 创建一个空的字符串切片, 用于存储响应
	for _, req := range reqs {
		reqObj := &Request{}                // 创建一个空的请求对象
		json.Unmarshal([]byte(req), reqObj) // 反序列化 JSON 字符串到请求对象
		ret := ""                           // 创建一个空字符串用于存储负载数据
		for _, e := range reqObj.PayLoad {  // 遍历请求负载
			ret += strconv.Itoa(e) + "," // 将整型转换为字符串并追加逗号
		}
		repObj := &Response{reqObj.TransactionID, ret} // 创建响应对象
		repJson, err := json.Marshal(&repObj)          // 序列化响应对象为 JSON
		if err != nil {
			panic(err) // 如果发生错误, 则抛出异常
		}
		reps = append(reps, string(repJson)) // 将响应 JSON 字符串追加到响应切片
	}
	return reps // 返回响应切片
}
