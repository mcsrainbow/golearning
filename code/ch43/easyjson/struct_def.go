package jsontest

type BasicInfo struct {
	Name string `json:"name"` // 字段 "Name" , 类型 string , 对应 JSON 键 "name"
	Age  int    `json:"age"`  // 字段 "Age" , 类型 int , 对应 JSON 键 "age"
}

type JobInfo struct {
	Skills []string `json:"skills"` // 字段 "Skills" , 类型 []string , 对应 JSON 键 "skills"
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"` // 字段 "BasicInfo" , 类型 BasicInfo , 对应 JSON 键 "basic_info"
	JobInfo   JobInfo   `json:"job_info"`   // 字段 "JobInfo" , 类型 JobInfo , 对应 JSON 键 "job_info"
}
