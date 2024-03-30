
function selectInput(listItem) {
	const locationInput = document.getElementById("location");
	const locationResults = document.getElementById("location_results");
	const value = listItem.dataset.value;

	locationInput.value = value;
	locationResults.innerHTML = "";
}
// Search Bar
const resultsEle = document.getElementById("results");
const input = document.getElementById("location");

input.addEventListener("input", () => {
	resultsEle.style.display = "block";
})

input.addEventListener("blur", () => {
	setTimeout(() => {
		resultsEle.style.display = "none";
	}, 200);
})
// Add Listing location input field
const locationInput = document.getElementById("location");
const locationResultsEle = document.getElementById("location_element");

locationInput.addEventListener("input", () => {
	locationResultsEle.style.display = "block";
})

input.addEventListener("blur", () => {
	setTimeout(() => {
		locationResultsEle.style.display = "none";
	}, 200);
})


