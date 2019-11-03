package main
import"fmt"
func factorial(n int) int{
	var k int =1
	for i:=1;i<=n;i++{
		k*=i
	}
	return k
}

func group() int{
	var arr1 =[4]int{1,2,3,4}
	for i:=0;i<4;i++{
		for k:=0;k<4;k++{
			if k!=i {
			}else{
				continue
			}
			for j:=0;j<4;j++{
				if (j!=i)&&(j!=k) {
				}else{
					continue
				}
				for h:=0;h<4 ;h++  {
					if (h!=i)&&(h!=j)&&(h!=k) {
						fmt.Printf("%d%d%d%d ",arr1[i],arr1[k],arr1[j],arr1[h])
					}else{
						continue
					}
				}
			}
		}
	}
	return 0
}

func main(){
	fmt.Printf("总数： %d\n",factorial(4))
	fmt.Printf("分别是：\n")
	group()
}