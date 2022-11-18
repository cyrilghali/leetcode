package main
import "fmt"

func isValid(s string) bool {
    var queue []string;
    for _, char := range s {
        switch string(char) {
            case "(":
                queue = append(queue, ")")
            case "[":
                queue = append(queue, "]")
            case "{":
                queue = append(queue, "}")
            default:
                if len(queue)== 0 || string(char) != queue[len(queue)-1]{
                    return false
            }
            queue = queue[:len(queue)-1]
        }
    }
    return len(queue) == 0;    
}

func main(){
    s := "(]"
    fmt.Println(isValid(s))
}
