var result1 = ""

try {
	result1 = dosomething()
catch (e)
	log("something went wrong")
}

try {
	dosomethingelse(result1)
catch (e)
	log("something else went wrong")
}
