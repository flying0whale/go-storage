package storage

import "strconv"

type Storage struct {
  data map[string]Data
}

func NewStorage() Storage {
  return Storage {
    data: make(map[string]Data),
  }
}

func (st Storage) Set (key, val string) {
  var value = Data {
    Type: String,
    StrData: val,
  }

  if num, err := strconv.Atoi(val); err == nil {
    value.Type =  Number
    value.NumData = num
  }

  st.data[key] = value
}

func (st Storage) Get (key string) *Data {
  if val, ok := st.data[key]; ok {
    return &val
  }

  return nil
}

func Hello () {

}

func (st Storage) Type (key string) DataType {
  if val, ok := st.data[key]; ok {
    return val.Type
  }

  return undefined
}
