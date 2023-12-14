package lib

type NavLink struct {
	Name string
	Link string
}

type Page struct {
	Nav  []NavLink	
}

func Navbar() []NavLink {
	return []NavLink{
		{
			Name: "/",
			Link: "/",
		},
	}
}
				
