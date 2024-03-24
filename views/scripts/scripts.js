
function selectInput(listItem) {
	const locationInput = document.getElementById("location");
	const locationResults = document.getElementById("location_results");
	const value = listItem.dataset.value;

	locationInput.value = value;
	locationResults.innerHTML = "";
}
