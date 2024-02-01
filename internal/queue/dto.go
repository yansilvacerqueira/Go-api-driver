package queue

import (
  "encoding/json"
)

type QueueDto struct {
  FileName string `json:"file_name"`
  Path string `json:"path"`
  ID int `json:"id"`
}

func (q *QueueDto) Marshal() ([]byte, error) {
  return json.Marshal(q)
}

func (q *QueueDto) Unmarshal(data []byte) error {
  return json.Unmarshal(data, q)
}
