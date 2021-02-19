package main

import (
	"fmt"
	"time"
)

type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}
type videoDetails struct {
	Author
	title string
	releaseDate string
	views int
	barrages int
	likes int
	isLike bool
	favorites int
	isFavorate bool
	coins int
	isCoin bool
	iscoin bool
	shares int
}
func (v *videoDetails) like(){
	if !v.isLike {
		v.likes++
		fmt.Println("点赞成功!")
	}else{
		v.likes--
		fmt.Println("取消点赞!")
	}
	v.isLike = !v.isLike
}
func (v *videoDetails) favorite(){
	if !v.isFavorate {
		v.favorites++
		fmt.Println("收藏成功!")
	}else{
		v.favorites--
		fmt.Println("取消收藏!")
	}
	v.isFavorate = !v.isFavorate
}
func (v *videoDetails) coin(d int){
	if !v.iscoin && d<=2 && d>=1 {
		v.coins += d
		v.iscoin = !v.iscoin
		fmt.Println("成功投了", d,"个币")
	}else{
		fmt.Println("你已投币")
	}
}
func (v *videoDetails) pressLike(){
	v.like()
	v.coin(2)
	v.favorite()
}
type Operations interface {
	like()
	favorite()
	coin()
	pressLike()
}
func publishVideo(a Author, title string) videoDetails{
	return videoDetails{
		Author:      a,
		title:       title,
		releaseDate: time.Now().Format("2006-01-02 15:04:05"),
		views:       0,
		barrages:    0,
		likes:       0,
		isLike:      false,
		favorites:   0,
		coins:       0,
		shares:      0,
	}

}
func main(){
	a1 := Author{
		Name: "张三",
		VIP: false,
		Icon: "无",
		Signature: "",
		Focus: 42,
	}
	str := "标题"
	pv := publishVideo(a1,str) //创建一个视频
	fmt.Println(pv)
	pv.like()
	pv.like()
	pv.coin(1)
	pv.coin(2)
	pv.favorite()
	pv.favorite()
	pv.favorite()
	fmt.Println(pv.coins)


}
