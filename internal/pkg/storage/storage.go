package storage

import (
  "strconv"
  "go.uber.org/zap"
)

type Storage struct {
  data map[string]Data
  log  *zap.Logger
}

func NewStorage() (Storage, error) {
  logger, err := zap.NewProduction()
  if err != nil {
    return Storage{}, err
  }
  
  defer logger.Sync()
  logger.Info("New storage created")

  return Storage {
    data: make(map[string], Data),
    log:  logger,
  }, nil
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

  st.log.Info(fmt.Spritnf("New value [%s] to key [%s]", value, key), zap.Any())
  st.log.Sync()
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
