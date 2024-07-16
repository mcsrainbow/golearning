package jsontest

// 定义 BasicInfo 结构体
type BasicInfo struct {
	Name string `json:"name"` // 字符串类型的 Name, JSON 字段名为 "name"
	Age  int    `json:"age"`  // 整数类型的 Age, JSON 字段名为 "age"
}

// 定义 JobInfo 结构体
type JobInfo struct {
	Skills []string `json:"skills"` // 字符串切片类型的 Skills, JSON 字段名为 "skills"
}

// 定义 Employee 结构体
type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"` // BasicInfo 类型的 BasicInfo, JSON 字段名为 "basic_info"
	JobInfo   JobInfo   `json:"job_info"`   // JobInfo 类型的 JobInfo, JSON 字段名为 "job_info"
}
