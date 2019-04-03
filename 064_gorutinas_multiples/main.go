package main

import (
	"fmt"
	"net/http"
	"time"
)

type RSS struct {
	Nombre string
	Url    string
}

func (rss *RSS) leer() {
	t1 := time.Now()
	resp, err := http.Get(rss.Url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	t2 := time.Since(t1).Seconds()
	fmt.Println(rss.Nombre, "lectura completada en", t2, "segundos")
}

func main() {
	listaRSS := []*RSS{
		&RSS{Nombre: "ESPN", Url: "http://www.espn.com/espn/rss/news"},
		&RSS{Nombre: "NY Times", Url: "http://rss.nytimes.com/services/xml/rss/nyt/World.xml"},
		&RSS{Nombre: "Wall Street Journal", Url: "https://feeds.a.dj.com/rss/RSSWorldNews.xml"},
	}
	for _, rss := range listaRSS {
		go rss.leer()
	}
	time.Sleep(time.Second * 6)
}
