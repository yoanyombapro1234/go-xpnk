package disqus

import (
	"fmt"
	"net/url"
)

// Get the most recent Disqus comments posted by a user. 
// Gets /users/listPosts.json?api_key=app_api_key&user:username=DisqusUsername OR access_token=DisqusUserAccessToken
func (api *Api) GetUserRecentComments(params url.Values) (res *ListPostsResponse, err error) {
	res = new(ListPostsResponse)
	err = api.get(fmt.Sprintf("/users/listPosts.json"), params, res)
	return
}