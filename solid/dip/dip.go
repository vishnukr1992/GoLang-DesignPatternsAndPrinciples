package solid

type SocialMediaPost struct {
	Posts []Post
}

// An abstraction
type ISocialMediaPostAnalyzer interface {
	GetPostByMediaType(mType SocialMedia) []string
}

func (smp *SocialMediaPost) GetPostByMediaType(mType SocialMedia) []string {
	fbPosts := make([]string, 0)
	for _, post := range smp.Posts {
		if post.From == Facebook {
			fbPosts = append(fbPosts, post.PostText)
		}
	}
	return fbPosts
}

// High level dependency that having the abstraction now
type SocialMediaPostAnalyzer struct {
	Analyzer ISocialMediaPostAnalyzer
}

// GetPostByMediaType helped to abstract the low level dependency's logic, So we do not need to change this business logic
// even if the low level data structure changed
func (smpa *SocialMediaPostAnalyzer) GetPostsInFacebook() []string {
	return smpa.Analyzer.GetPostByMediaType(Facebook)
}
