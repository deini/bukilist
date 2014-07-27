package thesrc

import (
    "errors"
    "net/http"
    "strconv"
    "time"

    "github.com/sourcegraph/deini/router"
)

type Bukiwish struct {
    Id          int64
    Name        string `sql:"not null;unique;type:varchar(100)"`
    Description string
    Rate        int
    Worth       bool
    Location    string
    Geolocation string
    CreatedAt   time.Time
}

type BukiwishesService interface {
    // Get a bukiwish
    // Get(id int) (*Bukiwish, error)

    // Get all bukiwishes
    GetAll() ([]*Bukiwish, error)

    // Submit a post. If this post's link URL has never been submitted, post.ID
    // will be a new ID, and created will be true. If it has been submitted
    // before, post.ID will be the ID of the previous post, and created will be
    // false.
    // NewBukiwish(bukiwish *Bukiwish) (created bool, err error)
}

// var (
//     ErrBukiwishNotFound = errors.New("Bukiwish not found")
// )

// type bukiwishesService struct{ client *Client }

func GetAll() ([]*Bukiwish, error) {
    var bukiwishes []*Bukiwish
    _, err = s.client.Do(req, &bukiwishes)
    if err != nil {
        return nil, err
    }

    return bukiwishes, nil
}

// func (s *postsService) Submit(post *Post) (bool, error) {
//     url, err := s.client.url(router.SubmitPost, nil, nil)
//     if err != nil {
//         return false, err
//     }

//     req, err := s.client.NewRequest("POST", url.String(), post)
//     if err != nil {
//         return false, err
//     }

//     resp, err := s.client.Do(req, &post)
//     if err != nil {
//         return false, err
//     }

//     return resp.StatusCode == http.StatusCreated, nil
// }
