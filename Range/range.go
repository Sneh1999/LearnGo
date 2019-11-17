pacakge main

import "fmt"

func main() {

	nums := []int{2,3,4}
	sum:= 0
	// here it returns the index  and the value at that index
	// sum the value at the index , ignoring the index
	for _, num := range nums {
		sum +=num
	}

	fmt.Println("sum:", sum);

	for i, num := range  nums {
		if num==3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	//  for key , value it prints both 
	for k ,v := range  kvs {
		fmt.Println("%s -> %s\n", k ,v)

	}
	for k := range kvs {
		fmt.Println("key", k)
	}

	


}