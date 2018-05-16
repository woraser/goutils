package helper

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSubstr(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	str := "党的领导是中国特色社会主义最本质的特征松树番茄,谁喜欢吃西红柿"
	fmt.Println(Substr(str, 2, 6))  //领导是中
	fmt.Println(Substr(str, 2, 60)) //领导是中国特色社会主义最本质的特征松树番茄,谁喜欢吃西
	fmt.Println(Substr(str, 20, 6)) //""
}

func TestRandomStr(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	f := runtime.FuncForPC(pc)
	fmt.Printf("\n\n\n------%s--------\n", f.Name())
	fmt.Println(RandomStr(5))  //hHnZV
	fmt.Println(RandomStr(10)) //3X4gPDCu2y
}

func TestBase62(t *testing.T) {
	var i int64 = 349879
	b62 := Base62Encode(i)
	fmt.Println(b62)
	fmt.Println(Base62Decode(b62))
}

func TestTraditionalToSimplified(t *testing.T) {
	tra := "無錫，簡稱“錫”，古稱新吳、梁溪、金匱，江蘇省地級市，被譽為“太湖明珠”。無錫位於江蘇省南部，地處長江三角洲平原、江南腹地，太湖流域。北倚長江，南濱太湖，東接蘇州，西連常州，構成蘇錫常都市圈 [1]  ，是長江經濟帶、長江三角洲城市群的重要城市，也是中央軍委無錫聯勤保障中心駐地。京杭大運河從無錫穿過，作為中國大運河的壹段，入選世界遺產名錄。條:1:条,偽:2:伪,廬:3:庐,聶:4:聂,緻:5:致,檔:6:档,棲:7:栖,啟:8:启,墳:9:坟,漿:10:浆,黴:11:霉,贓:12:赃,ａｂｃａ@￥@#%#ｓｄ🎈🎉ｆ我E２３４３４５んエォサ６３＃＄％＾＄＆％＾（＆我"
	fmt.Println("tra:\t", tra)
	sim := TraditionalToSimplified(tra)
	fmt.Println("sim:\t", sim)
}
