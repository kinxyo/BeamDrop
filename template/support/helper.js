export function appendFilenames(filenames) {
	let dropZone = GetBoard();

	// Removing the existing list if it exists
	let outdated = dropZone.querySelector("ul");
	if (outdated) {
		dropZone.removeChild(outdated);
	}

	var Beam = document.createElement("ul");
	for (var i = 0; i < filenames.length; i++) {
		var file = document.createElement("li");
		file.textContent = filenames[i];
		Beam.appendChild(file);
	}

	var button = dropZone.querySelector("button");
	dropZone.insertBefore(Beam, button);
}

function CreateBoard(dropZone) {
    
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

function GetBoard() {
    var dropZone = document.getElementById('beam-upload');
    
    // If there are no children, create the board
    if (dropZone.children.length === 0) {
        CreateBoard(dropZone);
    }

    return dropZone;
}
