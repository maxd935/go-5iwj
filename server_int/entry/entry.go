package entry

// func EntryHandler(w http.ResponseWriter, req *http.Request) {

// 	fmt.Println("EntryHandler")
// 	fmt.Println(myEntries)

// 	switch req.Method {
// 	case http.MethodGet:
// 		if len(myEntries) > 0 {
// 			for i, mess := range myEntries {
// 				fmt.Fprintf(w, "index: %s\n", i)
// 				fmt.Fprintf(w, "mess: %s\n", mess)
// 				fmt.Fprintf(w, "\n")
// 			}
// 		} else {
// 			fmt.Fprintf(w, "No Message")
// 		}
// 	case http.MethodPost:
// 		if err := req.ParseForm(); err != nil {
// 			fmt.Println("Something went bad")
// 			fmt.Fprintln(w, "Something went bad")
// 			return
// 		}
// 		for key, value := range req.PostForm {
// 			fmt.Println(key, "=>", value[0])
// 			myEntries[key] = value[0]
// 		}
// 		fmt.Println(myEntries)

// 		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
// 	}
// }
