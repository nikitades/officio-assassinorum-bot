package handlers

var noDead = "No dead people detected!"

func Dead() *string {
	return &noDead
}
