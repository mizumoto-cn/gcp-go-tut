package constexpr

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var instance *ops

type BqOperate struct {
	name    string
	apiName string
	id      int
}

func (b *BqOperate) GetName() string {
	return b.name
}

func (b *BqOperate) GetApiName() string {
	return b.apiName
}

func (b *BqOperate) GetId() int {
	return b.id
}

type ops struct {
	bs []BqOperate
}

func (o *ops) Get(i interface{}) BqOperate {
	switch i.(type) {
	case string:
		for _, v := range o.bs {
			if v.GetApiName() == i || v.GetName() == i {
				return v
			}
		}
	case int:
		for _, v := range o.bs {
			if v.GetId() == i {
				return v
			}
		}
	}
	return BqOperate{"not found", "not found", -1}
}

func GetOps() *ops {
	return instance
}

func init() {
	// read bqauditmetadata.json
	file, err := os.Open("bqauditmetadata.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var m map[string]string
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}
	b := []BqOperate{}
	// set the constants from 0 to 22
	i := 0
	for k, v := range m {
		b = append(b, BqOperate{k, v, i})
		i++
	}
	instance = &ops{b}
}
