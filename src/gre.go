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
	"time"
)

func main() {
	//初始化struct
	words := model.Word{}
	files := model.File{}

	fileArr := make([]model.File, 0)

	//循环遍历文件夹

	for {
		dir, e := ioutil.ReadDir(common.FileDir)
		if e != nil {
			fmt.Println("open dir failed", e)
		}
		fmt.Println("📖️小刘😔还要😔继续😔背单词😔!!!😔")
		for i, f := range dir {
			if f.IsDir() {
				fmt.Println("文件夹下文件是目录，请改为.xslx格式")
			}
			fmt.Println("【", i, "】", f.Name(), "🉐️")
			//给file 赋值
			files.File_name = f.Name()
			files.File_id = i
			fileArr = append(fileArr, files)
		}

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
				fmt.Println("ReplaceN : 类型转换异常", err)
			}

			iArr := make([]int, 0)
			for _, f := range fileArr {
				iArr = append(iArr, f.File_id)
			}

			b, err := util.Contain(i, iArr)
			if err != nil {
				fmt.Println("")
			}
			if !b {
				fmt.Println("宁选择的单词本不存在，请选择正确的单词本")
				break
			}
			fileName := fileArr[i].File_name
			fmt.Println("宁正在复习", fileName)
			Review(words, common.FileDir+"/"+fileName)
		}
	}
}

/**
去掉 input 中的 \n
*/
func ReplaceN(input string) (int, error) {
	re := regexp.MustCompile("\\n")
	newStr := re.ReplaceAllString(input, "")
	i, err := strconv.Atoi(newStr)
	return i, err
}

/**
背诵主逻辑
*/
func Review(words model.Word, excelFileName string) {
	//初始化集合
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
					words.Name = text
				} else if 1 == i {
					words.Explain = text
				}
			}
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
			fmt.Println("ReplaceN是 : 类型转换异常", err)
		}
		fmt.Println("--------------------")
		fmt.Printf("🦌️ 您现在正在复习单元 [%v],", n)
		fmt.Printf("请选择背诵频率, 单位[秒/个]  🦌️")
		//手动设置背诵频率
		i, done := SleepTime(err, inputReader)
		if done {
			return
		}

		fmt.Printf("🦌️ 宁的背诵频率为，[%v 秒/个]  🦌️\n", i)
		fmt.Println("--------------------")

		length := len(wordArr)

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
			fmt.Println("恭喜宁背完了!!")
			break
		}

	}
}

/**
设置单词背诵间隔
*/
func SleepTime(err error, inputReader *bufio.Reader) (int, bool) {
	sleep, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return 0, true
	}
	i, err := ReplaceN(sleep)
	if err != nil {
		fmt.Println("ReplaceN : 类型转换异常", err)
	}
	return i, false
}

/**
  获取随机单词逻辑
*/
func getRandomWords(sleepTime int, length int, w []model.Word) (newWord []model.Word) {

	sub := length

	//每次循环100个单词，如果最后的单词不够100个则用剩余的单词
	res := 100
	if 100 > length {
		res = length
	}
	for i := 0; i < res; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))

		x := r.Intn(sub)
		wordLen := len(w[x].Name)
		space := 0
		if 20 > wordLen {
			space = 20 - wordLen
		}

		fmt.Printf("[%s]", w[x].Name)
		for i := 0; i < space; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("[%s]\n\n", w[x].Explain)

		//删除已经背过的单词
		w = append(w[:x], w[x+1:]...)
		//控制长度避免数组越界，因为单词少一个，切片长度需要和单词数量相对应也要少一个
		sub = sub - 1
		//背一个单词睡3秒钟,除非手动控制
		time.Sleep(time.Second * time.Duration(sleepTime))
		////手动控制输入，输入为回车才可继续执行循环
		//inputReader := bufio.NewReader(os.Stdin)
		//input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		//
		//if err != nil {
		//	fmt.Println("There were errors reading, exiting program.")
		//	return
		//}

	}

	//背100个，总数就减100个
	length = length - 100
	if length < 0 {
		length = 0
	}
	fmt.Printf("剩余需要复习的单词数量 = 【%v】", length)
	return w
}
