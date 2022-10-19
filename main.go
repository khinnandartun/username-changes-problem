package main

import "fmt"

func possibleChanges(usernames []string) []string {
	var result []string
	for _, name := range usernames {
		var yn string = "NO"
		for i := 0; i < len(name)-1; i++ {
			if name[i] > name[i+1] {
				yn = "YES"
				break
			}
		}
		result = append(result, yn)
	}
	return result

}
func main() {
	fmt.Println(possibleChanges([]string{"foo", "bar", "baz"}))
}
