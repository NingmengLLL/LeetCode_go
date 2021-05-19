package main

import "fmt"

func main() {

	//var a []int
	//a = append(a, 3, 8, 7, 2, 6)
	//
	//sort.Ints(a)
	//fmt.Printf("%#v\n", a)
	//
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//fmt.Printf("%#v", a)

	//rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	//var scoreMap = make(map[string]int, 200)
	//
	//for i := 0; i < 100; i++ {
	//	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	//	value := rand.Intn(100)          //生成0~99的随机整数
	//	scoreMap[key] = value
	//}
	////取出map中的所有key存入切片keys
	//var keys = make([]string, 0, 200)
	//for key := range scoreMap {
	//	keys = append(keys, key)
	//}
	////对切片进行排序
	//sort.Strings(keys)
	////按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
