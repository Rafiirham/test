package models

type Author struct {
	Author_id int    `json:"author_id"`
	Name      string `json:"name"`
}

type Authors struct {
	Author []Author `json:"authors"`
}
