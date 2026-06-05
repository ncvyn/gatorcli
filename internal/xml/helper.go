package xml

import "html"

func (f *RSSFeed) unescape() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for i, item := range f.Channel.Item {
		f.Channel.Item[i].Title = html.UnescapeString(item.Title)
		f.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}
