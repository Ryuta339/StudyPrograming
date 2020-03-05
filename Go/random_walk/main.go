package main

import (
	"fmt"
	"math/rand"
)

func addLine(list map[string][]string, line []string) {
	if list[line[0]] == nil {
		list[line[0]] = make([]string, 0)
	}
	for i := 1; i < len(line); i++ {
		vp := line[i-1]
		v := line[i]
		if list[v] == nil {
			list[v] = make([]string, 0)
		}
		list[vp] = append(list[vp], v)
		list[v] = append(list[v], vp)
	}
}

func main() {

	田園都市線 := []string{
		"渋谷",
		"池尻大橋",
		"三軒茶屋",
		"駒澤大学",
		"桜新町",
		"用賀",
		"二子玉川",
		"溝の口",
		"鷺沼",
		"たまプラーザ",
		"あざみ野",
		"青葉台",
		"長津田",
		"つくし野",
		"すずかけ台",
		"南町田グランベリーパーク",
		"つきみ野",
		"中央林間",
	}

	東横線 := []string{
		"渋谷",
		"中目黒",
		"学芸大学",
		"自由が丘",
		"田園調布",
		"多摩川",
		"武蔵小杉",
		"日吉",
		"綱島",
		"菊名",
		"横浜",
	}

	目黒線 := []string{
		"目黒",
		"武蔵小山",
		"大岡山",
		"田園調布",
	}

	大井町線 := []string{
		"二子玉川",
		"自由が丘",
		"大岡山",
		"旗の台",
		"大井町",
	}

	池上線 := []string{
		"五反田",
		"旗の台",
		"蒲田",
	}

	多摩川線 := []string{
		"蒲田",
		"多摩川",
	}


	list := make(map[string][]string, 100)

	addLine(list, 田園都市線)
	addLine(list, 東横線)
	addLine(list, 目黒線)
	addLine(list, 大井町線)
	addLine(list, 池上線)
	addLine(list, 多摩川線)

	/*
	for k, v := range (list) {
		fmt.Printf("%s => %v\n", k, v)
	}
	*/

	count := int (1)
	station := "大岡山"
	for station != "すずかけ台" {
		fmt.Printf ("%5d, %v\n", count, station);
		l := len (list[station])
		n := rand.Intn (l)
		station = list[station][n]
		count ++
	}
	fmt.Printf ("%5d, %v\n",count,station);

}
