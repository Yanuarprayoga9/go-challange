package main
import ("fmt")

func modifya (a *string ) string {
 return *a; 
}
func main() {
 var a string = "a";

 var modifyavar = modifya(&a);
 fmt.Println(a,modifyavar)
}

