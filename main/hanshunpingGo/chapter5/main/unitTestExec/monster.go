package unitTestExec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const filePath = "d:/monsters.txt"

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (monster *Monster) Store() (err error) {
	// 将monster序列化成[]byte
	monsterStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("monster json marshal err=", err)
		return
	}

	// 写入文件
	err = ioutil.WriteFile(filePath, monsterStr, 0666)
	if err != nil {
		fmt.Println("monster write to file err=", err)
		return
	}

	fmt.Println("monster json marshal finished...")
	return
}

func (monster *Monster) ReStore() (err error) {
	// 先读取文件中的内容
	monsterStr, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("monster read from file err=", err)
		return
	}

	// 把[]byte反序列化成monster
	err = json.Unmarshal(monsterStr, &monster)
	if err != nil {
		fmt.Println("monster json unmarshal err=", err)
		return
	}

	fmt.Println("monster json unmarshal finished...")
	return
}
