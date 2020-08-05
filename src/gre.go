package main

import (
	"./common"
	"./model"
	"./util"
	"bufio"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//初始化struct
	words := model.Word{}
	files := model.File{}

	//循环遍历文件夹
	for {
		//每次都新建数组，避免循环导致的数组大小循环增加
		fileArr := make([]model.File, 0)
		dir, e := ioutil.ReadDir(common.FileDir)
		if e != nil {
			fmt.Printf("open dir failed,error = {%s},请检查文件路径是否正确\n", e)
			return
		}
		fmt.Println("📖️小刘😔还要😔继续😔背单词😔!!!😔")
		for i, f := range dir {
			if f.IsDir() {
				fmt.Println("此文件夹内存在目录,请删除目录，并保证文件都是已.xlsx结尾的excel文件")
			}
			fmt.Println("[", i, "]", f.Name(), "🉐️")
			//给file 赋值
			files.FileName = f.Name()
			files.FileId = i
			fileArr = append(fileArr, files)
		}
		//打印数组内容
		//fmt.Printf("%v\n", fileArr)

		//获取键盘输入的数字
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Printf("请选择要复习的文件:")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			return
		}

		switch input {
		case input:
			i, err := ReplaceN(input)
			if err != nil {
				fmt.Println("ReplaceN : 类型转换异常,请输入有效的文件序号!", err)
				break
			}

			iArr := make([]int, 0)
			for _, f := range fileArr {
				iArr = append(iArr, f.FileId)
			}

			b, _ := util.Contain(i, iArr)
			if !b {
				fmt.Println("宁选择的单词本不存在，请选择正确的单词本")
				break
			}
			fileName := fileArr[i].FileName
			fmt.Println("宁正在复习", fileName)
			//复习主方法
			Review(words, common.FileDir+"/"+fileName)
		}
	}
}

/**
@Description 去掉 input 中的 \n
@param input:控制台输入的数字 eg:  1\n  2\n
*/
func ReplaceN(input string) (int, error) {
	re := regexp.MustCompile("\\n")
	newStr := re.ReplaceAllString(input, "")
	i, err := strconv.Atoi(newStr)
	return i, err
}

/**
@Description 背诵主逻辑
@param words:单词对象
@param excelFileName:excel文件
*/
func Review(words model.Word, excelFileName string) {
	//初始化单词数组
	wordArr := make([]model.Word, 0)
	//打开文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	//遍历
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			//给每个word赋值
			for i, cell := range row.Cells {
				text := cell.String()
				if 0 == i {
					//excel sheet 第一列：单词名称
					words.Name = strings.Replace(text, " ", "", -1)
				} else if 1 == i {
					//excel sheet 第二列：单词释义
					words.Explain = text
				}
			}
			//每个单词的
			words.Id = i
			//逐个添加到切片中
			wordArr = append(wordArr, words)
		}
	}
	//
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Printf("请选择要复习的单元:")

		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There were errors reading, exiting program.")
			return
		}
		n, err := ReplaceN(input)
		if err != nil {
			fmt.Println("ReplaceN : 类型转换异常，请输入有效的单元", err)
			break
		}
		fmt.Println("🀀🀄︎🀁🀂🀃🀅🀆🀇🀈🀉🀊🀋🀌🀍🀎🀏🀐🀑🀒🀓🀔🀕🀖🀗🀘🀙🀚🀛🀜🀝🀞🀟🀠🀡🀢🀣🀤🀥🀦🀧🀨🀩")
		fmt.Printf("🦌️ 您现在正在复习单元 [%v] 🦌\n", n)
		fmt.Printf("🦌 请选择背诵频率, 单位[秒/个] 🦌️")

		//手动设置背诵频率
		i, err := sleepTime(inputReader)
		if err != nil {
			println("设置背诵频率发生异常,请输入[1-99999...]之间的整数 ")
			break
		}

		fmt.Printf("🦌️ 宁的背诵频率为，[%v 秒/个] 🦌️\n", i)
		fmt.Println("🀀🀄︎🀁🀂🀃🀅🀆🀇🀈🀉🀊🀋🀌🀍🀎🀏🀐🀑🀒🀓🀔🀕🀖🀗🀘🀙🀚🀛🀜🀝🀞🀟🀠🀡🀢🀣🀤🀥🀦🀧🀨🀩")

		length := len(wordArr)
		//设置假的单元数，让用户有一种有很多单元需要学习的错觉，有循序渐进学习的满足感，其实选哪个都是随机选100个🌶🐔
		switch input {
		case "1\n":
			wordArr = getRandomWords(i, length, wordArr)
		case "2\n":
			wordArr = getRandomWords(i, length, wordArr)
		case "3\n":
			wordArr = getRandomWords(i, length, wordArr)
		case "4\n":
			wordArr = getRandomWords(i, length, wordArr)
		case "5\n":
			wordArr = getRandomWords(i, length, wordArr)
		case "6\n":
			wordArr = getRandomWords(i, length, wordArr)
		}
		//如果length = 0 背完
		if 0 == len(wordArr) {
			fmt.Println("恭喜宁背完了!!🉑🉑 🀀🀄︎🀁🀂🀃🀅🀆🀇🀈🀉🀊🀋🀌🀍🀎🀏🀐🀑🀒🀓🀔🀕🀖🀗🀘🀙🀚🀛🀜🀝🀞🀟🀠🀡🀢🀣🀤🀥🀦🀧🀨🀩 🉑🉑")
			break
		}

	}
}

/**
设置单词背诵间隔
*/
func sleepTime(inputReader *bufio.Reader) (int, error) {
	sleep, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return 0, err
	}
	i, err := ReplaceN(sleep)
	if err != nil {
		fmt.Println("ReplaceN : 类型转换异常", err)
		return 0, err
	}
	return i, err
}

/**
  获取随机单词逻辑
*/
func getRandomWords(sleepTime int, length int, w []model.Word) []model.Word {

	sub := length

	//每次循环100个单词，如果最后的单词不够100个则用剩余的单词
	res := 100
	if 100 > length {
		res = length
	}
	for i := 0; i < res; i++ {
		//随机因子,基于时间戳，每次都不一样 x就是随机的数字
		r := rand.New(rand.NewSource(time.Now().Unix()))
		x := r.Intn(sub)

		idLen := len(strconv.Itoa(i))
		//仅仅为了前端展示需要，表示序号和单词之间的空格数
		idSpace := 0
		if 5 > idLen {
			idSpace = 5 - idLen
		}
		fmt.Printf("[%v]", i)
		for i := 0; i < idSpace; i++ {
			fmt.Printf(" ")
		}

		/*idLen := len(string(w[x].Id))
		//仅仅为了前端展示需要，表示序号和单词之间的空格数
		idSpace := 0
		if 5 > idLen {
			idSpace = 5 - idLen
		}
		fmt.Printf("[%v]", w[x].Id)
		for i := 0; i < idSpace; i++ {
			fmt.Printf(" ")
		}*/

		wordLen := len(w[x].Name)
		//仅仅为了前端展示需要，表示单词和释义之间的空格数
		space := 0
		if 20 > wordLen {
			space = 20 - wordLen
		}

		fmt.Printf("[%s]", w[x].Name)
		for i := 0; i < space; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("[%s]\n\n", w[x].Explain)

		//删除已经背过的单词：删除数组某个元素:a = append(a[:i], a[i+1:]...)
		w = append(w[:x], w[x+1:]...)
		//控制长度避免数组越界，因为单词少一个，切片长度需要和单词数量相对应也要少一个
		sub = sub - 1
		//背一个单词睡3秒钟,除非手动控制
		time.Sleep(time.Second * time.Duration(sleepTime))
	}

	//背100个，总数就减100个
	length = length - 100
	if length < 0 {
		length = 0
	}
	fmt.Printf("剩余需要复习的单词数量 = 【%v】", length)
	//返回数组中剩余的单词,此时单词数组的个数等于 len(w) - 100 |  length == length - 100
	return w
}
