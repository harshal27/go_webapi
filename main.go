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
  fmt.Fprintf(w, "Welcome to the HomePage!")
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
}

func getSingleArticles(w http.ResponseWriter, r *http.Request){
  var single_article *models.Article
  var er error
  single_article, er = findArticle(mux.Vars(r)["id"])

  if (er != nil){
    fmt.Println(er)
    resp,_ := json.Marshal(er)
    w.Write(resp)  
    return
  }

  json.NewEncoder(w).Encode(single_article)  
}

func findArticle(id string) (*models.Article, error) {
  index := -1
  for inx, i := range ar {
    if i.Id == id {
      index = inx
    }
  }
  if (index == -1){
    return nil, models.ApiError{Err: "ARTICLE NOT FOUND", Code: 400, Metadata: "meta"}
  }
  return &(ar[index]), nil
}

func updateSingleArticles(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  key := vars["id"]

  article, er := findArticle(key)

  if (er != nil){
    fmt.Println(er)
    resp,_ := json.Marshal(er)
    w.Write(resp)
    return
  }

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

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", homePage)
  router.HandleFunc("/articles", getArticles)
  router.HandleFunc("/articles/{id}", getSingleArticles)
  router.HandleFunc("/articles/{id}/update", updateSingleArticles).Methods("POST")
  router.HandleFunc("/articles/{id}/delete", deleteSingleArticles)
  router.HandleFunc("/article", createNewArticle).Methods("POST")

  log.Fatal(http.ListenAndServe(":3001", router))
}



