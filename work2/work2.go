package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main()  {
	file,err:=os.Open("top.txt")  //以只读方式打开文件
	if err!=nil {
		fmt.Println("failed to open the file,err:",err)
		return
	}
	defer file.Close()
    CpuRegexp,err:=regexp.Compile(`.\d\d\.\d.id\b`)   //找到cpu未利用率
    if err!=nil{
    	fmt.Println("failed to compile the cpuregexp,err:",err)
    	return
	}
    CpuDataRegexp,err:=regexp.Compile(`\d{0,3}\.\d`)   //找到未利用率的数据
    if err!=nil {
    	fmt.Println("failed to complie the cpuDataRegexp,err:",err)
    	return
	}
    MemoryStr,err:=regexp.Compile(`MiB.Mem(.*)\n`)   //找到Mib Mem行
    if err!=nil {
    	fmt.Println("failed to compile the MemoryStr,err:",err)
    	return
	}
	MemoryRegexp,err:=regexp.Compile(`\d{0,4}\.\d.used\b`)  //找到内存使用量
	if err!=nil {
		fmt.Println("failed to complie the MemoryRegexp,err:",err)
		return
	}
	MemoryDataRegexp,err:=regexp.Compile(`\d{0,4}\.\d`)  //找到内存使用量的数据
	if err!=nil {
		fmt.Println("failed to complie the MemoryDataRegexp,err:",err)
		return
	}  //多个判断能够一眼看出来问题在哪
	var(
		CpuUseMax=0.0
		CpuUseMin=100.0
		CpuUseSum=0.0
		MemoryUseMax=0.0
		MemoryUseMin=2000.0
		MemoryUseSum=0.0
		CpuUse [] float64
		MemoryUse [] float64
	)  //定义一些变量
	//开始读取文件
	reader:=bufio.NewReader(file)
	for {
		r,err:=reader.ReadString('\n')
		if err==io.EOF{
			fmt.Println("finish read")
			break   //必须先判断是否读完 不然永远不会出现err==io.EOF
		}
		if err!=nil{
			fmt.Println("failed to read the file,err:",err)
			return
		}

		CpuStr:=CpuRegexp.FindString(r)
		if CpuStr!=""{
			CpuDataStr:=CpuDataRegexp.FindString(CpuStr)//使用正则表达式匹配cpu信息和cpu为利用率
			CpuData,err:=strconv.ParseFloat(CpuDataStr,64)    //通过strconv包将string转为float64
			if err!=nil {
				fmt.Println("failed to parsefloat,err:",err)
				return
			}
			CpuUseData:=100.0-CpuData   //利用率
			CpuUse=append(CpuUse,CpuUseData)   //填入数组 比较和加入
			if CpuUseData>CpuUseMax{
				CpuUseMax=CpuUseData
			}
			if CpuUseData<CpuUseMin{
				CpuUseMin=CpuUseData
			}
			CpuUseSum+=CpuUseData
		}
		MemoryLine:=MemoryStr.FindString(r)   //使用正则表达式找到MiB Mem行
		if MemoryLine!=""{
			MemoryInfo := MemoryRegexp.FindString(MemoryLine)   //找到.used信息
			MemoryDataStr := MemoryDataRegexp.FindString(MemoryInfo)  //找到数据
			MemoryData, err := strconv.ParseFloat(MemoryDataStr, 64)  //字符串转浮点数
			if err != nil {
				fmt.Println("failed to parsefloat,err:", err)
				return
			}
			MemoryUse = append(MemoryUse, MemoryData)
			if MemoryData> MemoryUseMax {  //比较
				MemoryUseMax= MemoryData
			}
			if MemoryData < MemoryUseMin {
				MemoryUseMin = MemoryData
			}
			MemoryUseSum += MemoryData
		}
	}
	fmt.Println(CpuUse)
	fmt.Printf("cpuusemax:%f%%,cpuusemin:%f%%,cpuuseavg:%f%%\n",CpuUseMax,CpuUseMin,CpuUseSum/float64(len(CpuUse)))
	fmt.Println(MemoryUse)
	fmt.Printf("memoryusemax:%f,memoryusemin:%f,memoryuseavg:%f\n",MemoryUseMax,MemoryUseMin,MemoryUseSum/float64(len(MemoryUse)))
}
