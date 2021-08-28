package helper

import (
	"../service"
	"../LineThrift"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

type mentionMsg struct {
	MENTIONEES []struct {
		S string `json:"S"`
		E string `json:"E"`
		M string `json:"M"`
	} `json:"MENTIONEES"`
}

func MaxRevision (a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func InArray(arr []string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func InArray_int64(arr []int64, data int64) bool {
   for _, a := range arr {
      if a == data {
         return true
      }
   }
   return false
}

func InMap(dict map[string]bool, key string) bool {
    _, ok := dict[key]
    return ok
}

func IndexOf(data []string, element string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}

func Remove(s []string, element string) []string {
	i := IndexOf(s, element)
	s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func IsBanned(from string) bool {
	if InArray(service.Banned, from) == true {
		return true
	}
	return false
}

func GetMidFromMentionees(data string) []string{
	var midmen []string
	var midbefore []string
	res := mentionMsg{}
	json.Unmarshal([]byte(data), &res)
	for _, v := range res.MENTIONEES {
		if InArray(midbefore, v.M) == false {
			midbefore = append(midbefore, v.M)
			midmen = append(midmen, v.M)
		} 
	}

	return midmen
}

func Log(optype LineThrift.OpType, logtype string, str string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a:=time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0"+strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0"+strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0"+strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := yyyy+"-"+MM+"-"+dd+" "+hhconv+":"+mmconv+":"+ssconv
	fmt.Println("["+times+"]["+optype.String()+"]["+logtype+"]"+str)
}