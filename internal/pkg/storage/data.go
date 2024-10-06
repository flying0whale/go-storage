package storage

import "fmt"

type DataType int32

const (
  undefined DataType = -1
  Number    DataType = 0
  String    DataType = 1
)

type Data struct {
  Type DataType
  NumData int
  StrData string
}

func PrintType(dataType DataType) {
  switch dataType {
    case Number:
      fmt.Println("Number")
    case String:
      fmt.Println("String")
    default:
      fmt.Println("undefined")
  }
}
