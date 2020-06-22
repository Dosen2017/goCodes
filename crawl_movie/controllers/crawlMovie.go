package controllers

import (
	"github.com/astaxie/beego"
	"crawl_movie/models"
	"github.com/astaxie/beego/httplib"
	// "fmt"
	"time"
	"log"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {
	
	var movieInfo models.MovieInfo

	models.ConnectRedis("127.0.0.1:6379")

	sUrl := "https://movie.douban.com/subject/32659890/"
	models.PutinQueue(sUrl)

	
	i := int64(1)
	for {
		length := models.GetQueueLength()
		log.Println("-----", length)
		if length == 0 {
			break
		}

		sUrl = models.PopfromQueue()

		if models.IsVisit(sUrl) {
			continue
		}

		rsp := httplib.Get(sUrl)
		sMovieHtml, err := rsp.String()
		if err != nil {
			panic(err)
		}

		movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
		log.Println("+++++", movieInfo.Movie_name)
		//记录电影信息
		if movieInfo.Movie_name != "" {

			// if models.IsInsert(movieInfo.Movie_name) {
			// 	continue
			// }

			movieInfo.Id = i 
			// movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
			movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_grade = models.GetMovieScore(sMovieHtml)
			movieInfo.Movie_type = models.GetMovieType(sMovieHtml)
			movieInfo.Movie_on_time = models.GetMovieReleaseDate(sMovieHtml)
			movieInfo.Movie_span = models.GetMovieRuntime(sMovieHtml)

			models.AddMovie(&movieInfo)

			// models.AddToDYSet(movieInfo.Movie_name)

			i++
		}

		//提取该页面的所有连接
		urls := models.GetMovieUrls(sMovieHtml)
		for _, url := range urls {
			models.PutinQueue(url)

			c.Ctx.WriteString(url + "\n")
		}

		//sUrl，应当记录到，访问Set中
		models.AddToSet(sUrl)

		time.Sleep(time.Second)
	}

	c.Ctx.WriteString("end of crawl!")

	// movieInfo := models.MovieInfo{}
	// movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
	// movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
	// movieInfo.Movie_main_character = models.GetMovieMainCharacters(sMovieHtml)
	// movieInfo.Movie_grade = models.GetMovieScore(sMovieHtml)
	// movieInfo.Movie_type = models.GetMovieType(sMovieHtml)
	// movieInfo.Movie_on_time = models.GetMovieReleaseDate(sMovieHtml)
	// movieInfo.Movie_span = models.GetMovieRuntime(sMovieHtml)

	// id, _ := models.AddMovie(&movieInfo)

	// c.Ctx.WriteString(fmt.Sprintf("%v", id))


	// c.Ctx.WriteString(fmt.Sprintf("%v", models.GetMovieUrls(sMovieHtml)))

	
	// urls := models.GetMovieUrls(sMovieHtml)
	// for _, url := range urls {
	// 	models.PutinQueue(url)
	// }

	// c.Ctx.WriteString(fmt.Sprintf("%v", urls))
	// c.Ctx.WriteString("导演：" + models.GetMovieDirector(sMovieHtml) + "\n")

	// c.Ctx.WriteString("电影名：" + models.GetMovieName(sMovieHtml) + "\n")

	// c.Ctx.WriteString("主演：" + models.GetMovieMainCharacters(sMovieHtml) + "\n")

	// c.Ctx.WriteString("评分：" + models.GetMovieScore(sMovieHtml) + "\n")

	// c.Ctx.WriteString("类型：" + models.GetMovieType(sMovieHtml) + "\n")

	// c.Ctx.WriteString("上映日期：" + models.GetMovieReleaseDate(sMovieHtml) + "\n")

	// c.Ctx.WriteString("片长：" + models.GetMovieRuntime(sMovieHtml) + "\n")
}
