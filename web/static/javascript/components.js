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
    var dropZone = document.getElementById('beam-upload');
    
    // If there are no children, create the board
    if (dropZone.children.length === 0) {
        CreateBoard(dropZone);
    }

    return dropZone;
}