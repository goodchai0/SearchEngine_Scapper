package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"ac":  "https://www.google.ac/search?q=",
	"ad":  "https://www.google.ad/search?q=",
	"ae":  "https://www.google.ae/search?q=",
	"af":  "https://www.google.com.af/search?q=",
	"ag":  "https://www.google.com.ag/search?q=",
	"ai":  "https://www.google.com.ai/search?q=",
	"al":  "https://www.google.al/search?q=",
	"am":  "https://www.google.am/search?q=",
	"ao":  "https://www.google.co.ao/search?q=",
	"ar":  "https://www.google.com.ar/search?q=",
	"as":  "https://www.google.as/search?q=",
	"at":  "https://www.google.at/search?q=",
	"au":  "https://www.google.com.au/search?q=",
	"az":  "https://www.google.az/search?q=",
	"ba":  "https://www.google.ba/search?q=",
	"bd":  "https://www.google.com.bd/search?q=",
	"be":  "https://www.google.be/search?q=",
	"bf":  "https://www.google.bf/search?q=",
	"bg":  "https://www.google.bg/search?q=",
	"bh":  "https://www.google.com.bh/search?q=",
	"bi":  "https://www.google.bi/search?q=",
	"bj":  "https://www.google.bj/search?q=",
	"bn":  "https://www.google.com.bn/search?q=",
	"bo":  "https://www.google.com.bo/search?q=",
	"br":  "https://www.google.com.br/search?q=",
	"bs":  "https://www.google.bs/search?q=",
	"bt":  "https://www.google.bt/search?q=",
	"bw":  "https://www.google.co.bw/search?q=",
	"by":  "https://www.google.by/search?q=",
	"bz":  "https://www.google.com.bz/search?q=",
	"ca":  "https://www.google.ca/search?q=",
	"kh":  "https://www.google.com.kh/search?q=",
	"cc":  "https://www.google.cc/search?q=",
	"cd":  "https://www.google.cd/search?q=",
	"cf":  "https://www.google.cf/search?q=",
	"cat": "https://www.google.cat/search?q=",
	"cg":  "https://www.google.cg/search?q=",
	"ch":  "https://www.google.ch/search?q=",
	"ci":  "https://www.google.ci/search?q=",
	"ck":  "https://www.google.co.ck/search?q=",
	"cl":  "https://www.google.cl/search?q=",
	"cm":  "https://www.google.cm/search?q=",
	"co":  "https://www.google.com.co/search?q=",
	"cr":  "https://www.google.co.cr/search?q=",
	"cu":  "https://www.google.com.cu/search?q=",
	"cv":  "https://www.google.cv/search?q=",
	"cy":  "https://www.google.com.cy/search?q=",
	"cz":  "https://www.google.cz/search?q=",
	"de":  "https://www.google.de/search?q=",
	"dj":  "https://www.google.dj/search?q=",
	"dk":  "https://www.google.dk/search?q=",
	"dm":  "https://www.google.dm/search?q=",
	"do":  "https://www.google.com.do/search?q=",
	"dz":  "https://www.google.dz/search?q=",
	"ec":  "https://www.google.com.ec/search?q=",
	"ee":  "https://www.google.ee/search?q=",
	"eg":  "https://www.google.com.eg/search?q=",
	"es":  "https://www.google.es/search?q=",
	"et":  "https://www.google.com.et/search?q=",
	"fi":  "https://www.google.fi/search?q=",
	"fj":  "https://www.google.com.fj/search?q=",
	"fm":  "https://www.google.fm/search?q=",
	"fr":  "https://www.google.fr/search?q=",
	"ga":  "https://www.google.ga/search?q=",
	"ge":  "https://www.google.ge/search?q=",
	"gf":  "https://www.google.gf/search?q=",
	"gg":  "https://www.google.gg/search?q=",
	"gh":  "https://www.google.com.gh/search?q=",
	"gi":  "https://www.google.com.gi/search?q=",
	"gl":  "https://www.google.gl/search?q=",
	"gm":  "https://www.google.gm/search?q=",
	"gp":  "https://www.google.gp/search?q=",
	"gr":  "https://www.google.gr/search?q=",
	"gt":  "https://www.google.com.gt/search?q=",
	"gy":  "https://www.google.gy/search?q=",
	"hk":  "https://www.google.com.hk/search?q=",
	"hn":  "https://www.google.hn/search?q=",
	"hr":  "https://www.google.hr/search?q=",
	"ht":  "https://www.google.ht/search?q=",
	"hu":  "https://www.google.hu/search?q=",
	"id":  "https://www.google.co.id/search?q=",
	"iq":  "https://www.google.iq/search?q=",
	"ie":  "https://www.google.ie/search?q=",
	"il":  "https://www.google.co.il/search?q=",
	"im":  "https://www.google.im/search?q=",
	"in":  "https://www.google.co.in/search?q=",
	"io":  "https://www.google.io/search?q=",
	"is":  "https://www.google.is/search?q=",
	"it":  "https://www.google.it/search?q=",
	"je":  "https://www.google.je/search?q=",
	"jm":  "https://www.google.com.jm/search?q=",
	"jo":  "https://www.google.jo/search?q=",
	"jp":  "https://www.google.co.jp/search?q=",
	"ke":  "https://www.google.co.ke/search?q=",
	"ki":  "https://www.google.ki/search?q=",
	"kg":  "https://www.google.kg/search?q=",
	"kr":  "https://www.google.co.kr/search?q=",
	"kw":  "https://www.google.com.kw/search?q=",
	"kz":  "https://www.google.kz/search?q=",
	"la":  "https://www.google.la/search?q=",
	"lb":  "https://www.google.com.lb/search?q=",
	"lc":  "https://www.google.com.lc/search?q=",
	"li":  "https://www.google.li/search?q=",
	"lk":  "https://www.google.lk/search?q=",
	"ls":  "https://www.google.co.ls/search?q=",
	"lt":  "https://www.google.lt/search?q=",
	"lu":  "https://www.google.lu/search?q=",
	"lv":  "https://www.google.lv/search?q=",
	"ly":  "https://www.google.com.ly/search?q=",
	"ma":  "https://www.google.co.ma/search?q=",
	"md":  "https://www.google.md/search?q=",
	"me":  "https://www.google.me/search?q=",
	"mg":  "https://www.google.mg/search?q=",
	"mk":  "https://www.google.mk/search?q=",
	"ml":  "https://www.google.ml/search?q=",
	"mm":  "https://www.google.com.mm/search?q=",
	"mn":  "https://www.google.mn/search?q=",
	"ms":  "https://www.google.ms/search?q=",
	"mt":  "https://www.google.com.mt/search?q=",
	"mu":  "https://www.google.mu/search?q=",
}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int) ([]string, error) {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func googleResultParsing(response *http.Response, rank int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	sel := doc.Find("div.g")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3.r")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err

}

func getScrapeClient(proxyString interface{}) *http.Client {

	switch v := proxyString.(type) {

	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	default:
		return &http.Client{}
	}
}

func GoogleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
	for _, page := range googlePages {
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := googleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {
	baseClient := getScrapeClient(proxyString)
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())
	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("SCraper recieved a non-200 status code suggesting a ban")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	res, err := GoogleScrape("amit kewot", "com", "en", nil, 1, 30, 10)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
