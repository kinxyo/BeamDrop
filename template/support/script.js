/* IMPORTS */
import { appendFilenames } from "./helper.js";

/* Main Function */
document.addEventListener("DOMContentLoaded", function () {
	var form = document.querySelector("#drop");
	var fileInput = document.querySelector("#beam");
	var dataTransfer = new DataTransfer();
	var filenames = [];

	form.addEventListener("drop", function (e) {
		e.preventDefault();
		for (var i = 0; i < e.dataTransfer.files.length; i++) {
			dataTransfer.items.add(e.dataTransfer.files[i]);
			filenames.push(e.dataTransfer.files[i].name);
		}
		appendFilenames(filenames);
		fileInput.files = dataTransfer.files;
	});

	form.addEventListener("dragover", function (e) {
		e.preventDefault();
	});

	form.addEventListener("submit", function (e) {
		e.preventDefault();
		var formData = new FormData(e.target);
		fetch("/beam", {
			method: "POST",
			body: formData,
		})
			.then((response) => {
				console.log(response);
				return response.json();
			})
			.then((data) => {
				console.log(data);
				let board = document.getElementById("beam-upload");
				board.textContent = "Generated link: ";

				let a = document.createElement("a");
				a.href = "http://localhost:3000/drop/" + data.link;
				a.textContent = data.link;

				board.appendChild(a);
			})
			.catch((error) => {
				console.error(error);
				document.getElementById("beam-upload").textContent =
					"File upload failed";
			});
	});
});