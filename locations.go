package main

type LocationAreaResponse struct{
Count int 
Next *string
Previous *string
Results	[]LocationArea
}

type LocationArea struct {
	Name string `json:"name"`
	URL string `json:"url"`
}