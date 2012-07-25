package main

import (
  "net/http"
  "html/template"
  "github.com/kellegous/pork"
)

type jobs struct {
  tmpl     *template.Template
  articles []*Article
}

func (j *jobs) GetArticles() []*Article {
  return j.articles
}

func (j *jobs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func InitJobs(r pork.Router, tmpl *template.Template) (ArticleProvider, error) {
  return &jobs{
    tmpl: tmpl,
    articles: []*Article{
      &Article{
        Title:   "Lotus Development",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{1992, 0, 0},
      },
      &Article{
        Title:   "Pixel Technologies",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{1993, 0, 0},
      },
      &Article{
        Title:   "Heuristic Park",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{1995, 0, 0},
      },
      &Article{
        Title:   "Holistic Design",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{1997, 0, 0},
      },
      &Article{
        Title:   "AppForge",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{2000, 0, 0},
      },
      &Article{
        Title:   "Innuvo",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{2002, 0, 0},
      },
      &Article{
        Title:   "Google",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{2005, 7, 31},
      },
      &Article{
        Title:   "Monetology",
        Icon:    "img/icon-job.png",
        Date:    SimpleDate{2012, 3, 5},
      },
    },
  }, nil
}