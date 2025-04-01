package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tsitsishvili/social/internal/store"
	"log"
	"math/rand"
)

var Usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user: ", err)
		}
	}

	_ = tx.Commit()

	posts := generatePosts(100, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
		}
	}

	comments := generateComments(100, posts, users)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
		}
	}

	log.Println("Seeding completed")
}

func generateComments(count int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, count)

	for i := 0; i < count; i++ {
		post := posts[rand.Intn(len(posts))]
		user := users[rand.Intn(len(users))]

		comments[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: fmt.Sprintf("This is the comment %d", i),
		}
	}

	return comments
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)

	for i := 0; i < count; i++ {
		users[i] = &store.User{
			Username: Usernames[i%len(Usernames)] + "_" + fmt.Sprintf("%d", i),
			Email:    Usernames[i%len(Usernames)] + "_" + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)

	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   fmt.Sprintf("Post %d", i),
			Content: fmt.Sprintf("This is the body of post %d", i),
			Tags:    []string{"tag1", "tag2", "tag3"},
		}
	}

	return posts
}
