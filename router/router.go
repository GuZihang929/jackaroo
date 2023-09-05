package router

import (
	"fmt"
	"sync"
	"time"
	_60 "xiangxiang/jackaroo/common/app/360"
	"xiangxiang/jackaroo/common/app/Alibaba"
	"xiangxiang/jackaroo/common/app/Baidu"
	"xiangxiang/jackaroo/common/app/Meituan"
	"xiangxiang/jackaroo/common/app/Tencent"
	"xiangxiang/jackaroo/common/app/Wangyi"
	"xiangxiang/jackaroo/common/app/Weiruan"
	Bilibili "xiangxiang/jackaroo/common/app/bilibili"
	"xiangxiang/jackaroo/common/app/huawei"
	"xiangxiang/jackaroo/common/app/jingdong"
	"xiangxiang/jackaroo/common/app/lilisi"
	"xiangxiang/jackaroo/common/app/mihoyo"
	"xiangxiang/jackaroo/common/app/xiaohongshu"
	"xiangxiang/jackaroo/common/app/xiaomi"
	"xiangxiang/jackaroo/common/app/zijie"
)

//func New() *gin.Engine {
//
//}

// 并发爬取信息
func Hello() {
	var wg sync.WaitGroup
	wg.Add(15)
	go func() {
		defer wg.Done()
		for {
			if pan, err := zijie.Header(""); pan != false {
				fmt.Println("zijie爬完了")
				break
			} else {
				fmt.Println("zijie重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := xiaomi.Header(""); pan != false {
				fmt.Println("xiaomi爬完了")
				break
			} else {
				fmt.Println("xiaomi重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := xiaohongshu.Header(""); pan != false {
				fmt.Println("xiaohongshu爬完了")
				break
			} else {
				fmt.Println("xiaohongshu重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := lilisi.Header(""); pan != false {
				fmt.Println("lilisi爬完了")
				break
			} else {
				fmt.Println("lilisi重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := huawei.Header(""); pan != false {
				fmt.Println("huawei爬完了")
				break
			} else {
				fmt.Println("huawei重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Alibaba.Header(""); pan != false {
				fmt.Println("alibaba爬完了")
				break
			} else {
				fmt.Println("alibaba重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Baidu.Header(""); pan != false {
				fmt.Println("百度爬完了")
				break
			} else {
				fmt.Println("百度重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Bilibili.Header(""); pan != false {
				fmt.Println("Bilibili爬完了")
				break
			} else {
				fmt.Println("bilibili重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := jingdong.Header(""); pan != false {
				fmt.Println("jingdong爬完了")
				break
			} else {
				fmt.Println("京东重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Meituan.Header(""); pan != false {
				fmt.Println("美团爬完了")
				break
			} else {
				fmt.Println("美团重试", err)
				time.Sleep(3 * time.Second)
				continue
			}

		}
	}()
	//公司网络屏蔽了米哈游 所以需要换网络才可以进去并爬到
	go func() {
		defer wg.Done()
		for {
			if pan, err := mihoyo.Header(""); pan != false {
				fmt.Println("米哈游爬完了")
				break
			} else {
				fmt.Println("米哈游重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Tencent.Header(""); pan != false {
				fmt.Println("腾讯爬完了")
				break
			} else {
				fmt.Println("腾讯重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Wangyi.Header(""); pan != false {
				fmt.Println("网易爬完了")
				break
			} else {
				fmt.Println("网易重试", err)
				time.Sleep(3 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := Weiruan.Header(""); pan != false {
				fmt.Println("微软爬完了")
				break
			} else {
				fmt.Println("微软重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			if pan, err := _60.Header(""); pan != false {
				fmt.Println("360爬完了")
				break
			} else {
				fmt.Println("360重试", err)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	wg.Wait()
	fmt.Println("所有网站都爬取完毕了")
}
