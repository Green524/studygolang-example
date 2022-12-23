package json

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}
type Movie1 struct {
}

var movies = []Movie{
	{Title: "三体奈飞版", Year: 2022, Color: true, Actors: []string{"cpq", "cpq1", "cpq2"}},
	{Title: "三体动画版", Year: 2022, Color: true, Actors: []string{"一画开天垃圾"}},
	{Title: "三体企鹅版", Year: 2022, Actors: []string{"qq"}},
}

func test() {

}
