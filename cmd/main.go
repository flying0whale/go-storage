package main
import (
  "fmt"
  "homework/internal/pkg/storage"
)

func main() {
  var st = storage.NewStorage()
  st.Set("key", "value")
  st.Set("num", "23452")

  var value = st.Get("key")
  storage.PrintType(value.Type)
  switch value.Type {
    case storage.Number:
      fmt.Println(value.NumData)
    case storage.String:
      fmt.Println(value.StrData)
    default:
      fmt.Println("no value")
  }


  value = st.Get("num")
  storage.PrintType(value.Type)
}
