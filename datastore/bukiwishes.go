package datastore

import (
    "fmt"
    "github.com/deini/bukilist"
    "strings"
    "time"
)

type bukiwishesStore struct{ *Datastore }

// func (s *postsStore) Get(id int) (*thesrc.Post, error) {
//     var posts []*thesrc.Post
//     if err := s.dbh.Select(&posts, `SELECT * FROM post WHERE id=$1;`, id); err != nil {
//         return nil, err
//     }
//     if len(posts) == 0 {
//         return nil, thesrc.ErrPostNotFound
//     }
//     return posts[0], nil
// }

// func (s *postsStore) List(opt *thesrc.PostListOptions) ([]*thesrc.Post, error) {

//     sql := `SELECT * FROM post`

//     var conds []string
//     if opt.CodeOnly {
//         conds = append(conds, "classification LIKE 'CODE%'")
//     }
//     if len(conds) > 0 {
//         sql += " WHERE (" + strings.Join(conds, ") AND (") + ")"
//     }

//     sql += " ORDER BY submittedat DESC LIMIT $1 OFFSET $2;"

//     var posts []*thesrc.Post
//     err := s.dbh.Select(&posts, sql, opt.PerPageOrDefault(), opt.Offset())
//     if err != nil {
//         return nil, err
//     }
//     return posts, nil
// }
func (s *bukiwishesStore) GetBukiwishes([]*bukilist.Bukiwish, error) {
    var bukiwishes []*bukilist.Bukiwish
    query := Db.Find(&bukiwishes)

    if query.Error != nil {
        return nil, query.Error
    }

    return bukiwishes, nil
}

// func (s *postsStore) Submit(post *thesrc.Post) (bool, error) {
//     retries := 3
//     var wantRetry bool

// retry:
//     retries--
//     wantRetry = false
//     if retries == 0 {
//         return false, fmt.Errorf("failed to submit post with URL %q after retrying", post.LinkURL)
//     }

//     var created bool
//     err := transact(s.dbh, func(tx modl.SqlExecutor) error {
//         var existing []*thesrc.Post
//         if err := tx.Select(&existing, `SELECT * FROM post WHERE linkurl=$1 LIMIT 1;`, post.LinkURL); err != nil {
//             return err
//         }
//         if len(existing) > 0 {
//             *post = *existing[0]
//             return nil
//         }

//         if err := tx.Insert(post); err != nil {
//             if strings.Contains(err.Error(), `violates unique constraint "post_linkurl"`) {
//                 time.Sleep(time.Duration(rand.Intn(75)) * time.Millisecond)
//                 wantRetry = true
//                 return err
//             }
//             return err
//         }

//         created = true
//         return nil
//     })
//     if wantRetry {
//         goto retry
//     }
//     return created, err
// }
