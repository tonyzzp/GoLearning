package main

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func download(ch chan string, wg *sync.WaitGroup) {
	me := os.Mkdir("pics", os.ModePerm)
	fmt.Println(me)
	for u := range ch {
		bytes := md5.Sum([]byte(u))
		key := hex.EncodeToString(bytes[:])
		resp, _ := http.Get(u)
		if resp != nil {
			f, _ := os.OpenFile("pics/"+key, os.O_RDWR, os.ModePerm)
			if f == nil {
				f, _ = os.Create("pics/" + key)
			}
			count, e := io.Copy(f, resp.Body)
			fmt.Println(count, e)
			f.Close()
			resp.Body.Close()
		}
	}
	fmt.Println("...")
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	resp, _ := http.Get("http://www.coolapk.com")
	bytes, _ := ioutil.ReadAll(resp.Body)
	content := string(bytes)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	pics := list.New()
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, _ := selection.Attr("src")
		if strings.HasPrefix(src, "http") {
			pics.PushBack(src)
		} else {
			pics.PushBack("http://www.coolapk.com" + src)
		}
	})

	wg := sync.WaitGroup{}
	ch := make(chan string, 100)
	wg.Add(1)
	go download(ch, &wg)
	for e := pics.Front(); e != nil; e = e.Next() {
		s := e.Value.(string)
		ch <- s
	}
	close(ch)
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println(end - start)
}
