package main

import (
  "fmt"
  "log"
  "net/http"
  "webapi/models"
  "encoding/json"
  "github.com/gorilla/mux"
  "io/ioutil"
)

var ar []models.Article

func homePage(w http.ResponseWriter, r *http.Request){
//func homePage(){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    //fmt.Println("Endpoint Hit: homePage 1234")

    
    

    //fmt.Println(r.Host)
    // fmt.Println(ar[0].Content)

    // ar[0].Content = "1234"
    // fmt.Println(ar[0].Content)
    // fmt.Println(article)

}

func getArticles(w http.ResponseWriter, r *http.Request){
  fmt.Println("Endpoint Hit: returnAllArticles")
  json.NewEncoder(w).Encode(ar)
}

func initArticles(){
  article1 := models.Article{Content: "content1", Id: "1", Title: "title1"}
  article2 := models.Article{Content: "content2", Id: "2", Title: "title2"}
  ar = append(ar, article1)
  ar = append(ar, article2)
  //fmt.Println(ar[0].Content)
}

func getSingleArticles(w http.ResponseWriter, r *http.Request){
  var single_article *models.Article
 // vars := mux.Vars(r)
  //key := vars["id"]

  // for _, i := range ar {
  //   if i.Id == key {
  //     single_article = i
  //   }
  // }

  single_article = findArticle(mux.Vars(r)["id"])


  //fmt.Println(single_article)

  json.NewEncoder(w).Encode(single_article)  
}

func findArticle(id string) *models.Article {
  //var single *models.Article
  
  //fmt.Println(id)
  //fmt.Println(ar)
  var index int

  for inx, i := range ar {
    if i.Id == id {
      //fmt.Println(i.Id)
      //fmt.Println(id)
      //fmt.Println(i)
      //fmt.Println(&i)
      //single = &(i)
      //fmt.Println(single)
      index = inx
    }
  }


  //fmt.Println(ar[0])

  //fmt.Println(single)

  return &(ar[index])
}

func updateSingleArticles(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  key := vars["id"]

  article := findArticle(key)
  fmt.Println(article.Content)

  var newA models.Article
  resB, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(resB, &newA)

  article.Content = newA.Content
  article.Title = newA.Title

  fmt.Println(ar)
  json.NewEncoder(w).Encode(ar)
}

func deleteSingleArticles(w http.ResponseWriter, r *http.Request){
  id := mux.Vars(r)["id"]
  var index int 
  for inx, i := range ar {
    if (i.Id == id) {
      index = inx
    }
  }

  ar := append(ar[:index], ar[index + 1:]...)
  json.NewEncoder(w).Encode(ar)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)
  var article models.Article 
  json.Unmarshal(reqBody, &article)
  ar = append(ar, article)
  json.NewEncoder(w).Encode(ar)
}

func main(){
  initArticles()

  // var c []models.Article
  // a := models.Article{Content: "content123", Id: "123", Title: "title1234"}
  // c = append(c, a)

  // fmt.Println(a.Content)


  // b := &(c[0])
  // b.Content = "updated"
  // fmt.Println(b.Content)
  // fmt.Println(a.Content)
  // fmt.Println(c)


  router := mux.NewRouter().StrictSlash(true)




  router.HandleFunc("/", homePage)
  router.HandleFunc("/articles", getArticles)
  router.HandleFunc("/articles/{id}", getSingleArticles)
  router.HandleFunc("/articles/{id}/update", updateSingleArticles).Methods("POST")
  router.HandleFunc("/articles/{id}/delete", deleteSingleArticles)
  router.HandleFunc("/article", createNewArticle).Methods("POST")
  log.Fatal(http.ListenAndServe(":3001", router))

  
}



