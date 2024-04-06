package solid

// This principle is saying that the high level business logic should not depending on
// low level dependencies, all low level dependencies should handles in low level itself
// So if any changes happened in low level dependency then high level logics will not fail

// for an example lets take a social media post analysis.

type PostAnalyzer struct {
}

// here we failed the DIP, because analyzer is high level feature of MediaPost, if media data structure changed then we will have to
//change the analyzer also, this can be avoided by abstracting this function here
func (pa *PostAnalyzer) GetPostsInFacebook(posts []Post) []string {
	fbPosts := make([]string, 0)
	for _, post := range posts {
		if post.From == Facebook {
			fbPosts = append(fbPosts, post.PostText)
		}
	}
	return fbPosts
}
