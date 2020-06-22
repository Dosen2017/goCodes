package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	// "strconv"
)

var db orm.Ormer

type MovieInfo struct {
	Id int64
	Movie_id int64
	Movie_name string
	Movie_pic string
	Movie_director string
	Movie_writer string
	Movie_country string
	Movie_language string
	Movie_main_character string
	Movie_type string
	Movie_on_time string
	Movie_span string
	Movie_grade string
	_Create_time string
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo)(int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}

func GetMovieDirector(movieHtml string) string {

/*
	const (
		cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	)
	 
	compile := regexp.MustCompile(cityListReg)
	 
	submatch := compile.FindAllSubmatch(contents, -1)
	 
	for _, m := range submatch {
		fmt.Println("url:" , string(m[1]), "city:", string(m[2]))
	}
*/

	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">([^<]+)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}


	directors := ""
	for _, d := range result {
		directors += d[1] + ","
	}

	return string(strings.Trim(directors,","))
}

func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	mainCharacters := ""
	for _, d := range result {
		mainCharacters += d[1] + ","
	}

	return string(strings.Trim(mainCharacters,","))
}

func GetMovieScore(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}

	// score, _ := strconv.ParseFloat(result[0][1], 64)
	return result[0][1]
}

func GetMovieType(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:genre">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	// score, _ := strconv.ParseFloat(result[0][1], 64)
	return result[0][1]
}

func GetMovieReleaseDate(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:initialReleaseDate" content=".*?">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	// score, _ := strconv.ParseFloat(result[0][1], 64)
	return result[0][1]
}

func GetMovieRuntime(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}

	reg := regexp.MustCompile(`<span property="v:runtime" content=".*?">(.*)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	// score, _ := strconv.ParseFloat(result[0][1], 64)
	return result[0][1]
}

func GetMovieUrls(movieHtml string) []string {
	// if movieHtml == "" {
	// 	return ""
	// }
	var movieSet []string
	reg := regexp.MustCompile(`<a href="(https://movie.douban.com/.*)" class="" >`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return movieSet
	}
	
	for _, d := range result {
		movieSet = append(movieSet, d[1])
	}

	return movieSet
}


