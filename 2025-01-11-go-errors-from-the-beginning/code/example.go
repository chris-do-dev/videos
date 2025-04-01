package main

// Attempt to read the contents of a user-specified file
// If the file is successfully read, print the contents
// Else if an error is encoutnered,
//    If the file does not exist, fall back to a default and continue
//        If fallback filename specified, use that instead

//func main() {
//	content, err := os.ReadFile("./file.txt")
//	if err != nil {
//		fmt.Println("falling back to default file")
//
//		content, _ = os.ReadFile("./default.txt")
//	}
//	fmt.Print(string(content))
//}
