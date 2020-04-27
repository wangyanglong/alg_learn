package main

import (
    "fmt"
)

func quickSort(array []int, left,right int){
    if left >= right{
        return
    }

    pivot := array[left]
    i,j:=left,right
    for i !=j{
        for array[j] >= pivot && i<j{
            j --
        }
        for pivot >= array[i] && i<j{
            i++
        }

        if i <j{
            tmp:= array[i]
            array[i] = array[j]
            array[j] = tmp
        }
        fmt.Printf("i:%v,j:%v\n",i,j)
    }

    array[left] = array[i]
    array[i] = pivot

    quickSort(array,0,i-1)
    quickSort(array,i+1,right)
}

func main(){
    array := []int{1,3,4,9,8,2,5,7,6,10}
    quickSort(array,0,len(array)-1)

    fmt.Printf("array:%+v",array)
}
