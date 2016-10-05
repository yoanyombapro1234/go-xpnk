package disqus

import (
	"fmt"
	"net/url"
)

// Get the details of a Disqus comment thread by thread id. 
// Gets /threads/details.json?api_key=app_api_key&thread=ListPostsResponse.Content.Thread
func (api *Api) GetThreadDetails(params url.Values) (res *DisqusThreadDetails, err error) {
	res = new(DisqusThreadDetails)
	err = api.get(fmt.Sprintf("/threads/details.json"), params, res)
	return
}
