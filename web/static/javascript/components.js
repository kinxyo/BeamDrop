/* Toast */
export function AddToast(message, ParentElement) {
	var ToastMessage = document.createElement("p");
	ToastMessage.textContent = message;
	ToastMessage.classList.add("appear");

	ParentElement.appendChild(ToastMessage);

	setTimeout(function () {
		ToastMessage.classList.remove("appear");
		ToastMessage.classList.add("disappear");
	}, 3000);

	// Removing the paragraph from the DOM when the transition ends
	ToastMessage.addEventListener("transitionend", function () {
		if (ToastMessage.classList.contains("disappear")) {
			ToastMessage.parentNode.removeChild(ToastMessage);
		}
	});
}

/* BeamDrop Board */
export function CreateBoard(dropZone) {
    
    // Making it visible
    dropZone.style.opacity = 1;

    // Header
    var header = document.createElement('h1');
    header.textContent = 'BeamDrop';
    dropZone.appendChild(header);

    // Button
    var button = document.createElement('button');
    button.textContent = 'drop';
    dropZone.appendChild(button);

    // Remove Noise
    var tag = document.querySelector('label');
    tag.style.opacity = 0;
}

export function GetBoard() {
    var dropZone = document.getElementById('beam-aim');
    
    // If there are no children, create the board
    if (dropZone.children.length === 0) {
        CreateBoard(dropZone);
    }

    return dropZone;
}