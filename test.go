package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TestService struct {
}

/**
Author:charlie
Description:减少if-else重构
 */
/*
func Handle() error {
	var err error
	if Operation1(){
		if Operation2(){
			if Operation3(){
				if Operation4(){
					// do
				}else{
					err = OPERATION4FAILED
				}
			}else{
				err = OPERATION3FAILED
			}
		}else{
			err = OPERATION2FAILED
		}
	}else{
		err = OPERATION1FAILED
	}
	return err
}
*/

const (
	OPERATION1FAILED = "OPERATION1FAILED"
	OPERATION2FAILED = "OPERATION2FAILED"
	OPERATION3FAILED = "OPERATION3FAILED"
	OPERATION4FAILED = "OPERATION4FAILED"
)

func (s *TestService) Handle() error {
	if !Operation1() {
		return fmt.Errorf("%v", OPERATION1FAILED)
	}
	if !Operation2() {
		return fmt.Errorf("%v", OPERATION2FAILED)
	}
	if !Operation3() {
		return fmt.Errorf("%v", OPERATION3FAILED)
	}
	if !Operation4() {
		return fmt.Errorf("%v", OPERATION4FAILED)
	}
	//do
	return nil
}

//测试方法
func Operation1() bool {
	return false
}
func Operation2() bool {
	return false
}
func Operation3() bool {
	return false
}
func Operation4() bool {
	return false
}

/**
Auhtor:charlie
Description:求出指定和的加数,去重复数据
 */
func (s *TestService)Sum(target int,nums []int) map[int]int {
	var result = make(map[int]int)
	var resp = make(map[int]int)
	for _, value := range nums {
		_, ok := result[value]
		if ok {
			//如果有值
			resp[value] =  target - value
		} else {
			//存值
			result[target-value] = value
		}
	}
	return resp
}

/**
Auhtor:charlie
Description:堵塞主线程
*/
func (s *TestService)PrintGoroutine()  {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

/**
Author:charlie
Description:将ip转换为整数
 */
func (s *TestService)IpConvertInt(ip string) int {
	//先分成四段
	arrays := strings.Split(ip,".")
	var resp int
	for key, value := range arrays {
		array,err := strconv.Atoi(value)
		if err != nil {
			_ = fmt.Errorf("int类型转换为string类型失败 %v",err)
		}
		resp |= array << 8*key
	}
	return resp
}

/**
Author:charlie
Description:打印5个随机数
 */
func (s *TestService)PrintChannelData()  {
	wg := sync.WaitGroup{}
	wg.Add(1)
	resp := make(chan int,5) //5个缓冲容量
	go func() {
		for i := 0; i < 5; i++ {
			rand.New(rand.NewSource(time.Now().UnixNano())) //速度很快，用纳秒生成种子
			resp <- rand.Intn(100)
		}
		//关闭通道
		close(resp)
	}()
	go func() {
		defer wg.Done()
		for temp := range resp {
			fmt.Println(temp)
		}
	}()
	wg.Wait()
}